// Kronos 5G Emulator console. Plain JS, no build step.
// Every operation is a command line sent to POST /api/v1/commands, as typed
// on the CUI; events arrive as Server-Sent Events from /api/v1/events.
"use strict";

const MAX_EVENTS = 4000; // events kept in the page (sequence view and log)
// PAGES maps page names (also the URL hash) to their titles.
const PAGES = {
  dashboard: "Dashboard",
  ues: "Subscribers",
  ops: "Operations",
  signals: "Signals",
  sequence: "Sequence",
  scenario: "Scenario",
  logs: "Logs",
  license: "License",
};
// KINDS are the event kinds, as offered by the log filters.
const KINDS = ["message", "state", "info", "snap", "command", "link", "warn", "error"];

// S holds the state of the page.
const S = {
  page: "dashboard",
  status: null,
  started: null,
  commands: [],
  signals: [],
  scenario: null,
  target: "", // IMSI that commands go to; "" is every subscriber
  ues: [],
  uesTotal: 0,
  ueState: "",
  ueQuery: "",
  ueDetail: null,
  events: [], // recent events, oldest first
  lastSeq: 0,
  lost: 0,
  seq: { src: "", notes: true, follow: true, selected: 0 },
  log: { kinds: new Set(KINDS), src: "", q: "", follow: true },
  history: [],
  histPos: 0,
  streaming: false,
};

// ---------- small helpers (no DOM) ----------

// pad left-pads n with zeros to w digits.
function pad(n, w) {
  return String(n).padStart(w, "0");
}

// timeOf renders an event time as hh:mm:ss.mmm (local time).
function timeOf(t) {
  const d = new Date(t);
  if (isNaN(d)) return "";
  return `${pad(d.getHours(), 2)}:${pad(d.getMinutes(), 2)}:${pad(d.getSeconds(), 2)}.${pad(d.getMilliseconds(), 3)}`;
}

// dateTimeOf renders a time as yyyy-mm-dd hh:mm:ss (local time).
function dateTimeOf(t) {
  const d = new Date(t);
  if (isNaN(d)) return "";
  return `${d.getFullYear()}-${pad(d.getMonth() + 1, 2)}-${pad(d.getDate(), 2)} ${timeOf(d).slice(0, 8)}`;
}

// hexDump renders hex as offset, octets and printable characters, 16 a line.
function hexDump(hex) {
  const bytes = (hex || "").match(/../g) || [];
  const lines = [];
  for (let i = 0; i < bytes.length; i += 16) {
    const row = bytes.slice(i, i + 16);
    let hx = "";
    let asc = "";
    for (let j = 0; j < 16; j++) {
      hx += (j < row.length ? row[j] : "  ") + (j === 7 ? "  " : " ");
      if (j < row.length) {
        const c = parseInt(row[j], 16);
        asc += c >= 0x20 && c < 0x7f ? String.fromCharCode(c) : ".";
      }
    }
    lines.push(`${pad(i.toString(16), 4)}  ${hx} ${asc}`);
  }
  return lines.join("\n");
}

// commandName returns the command of a line ("" for @file or no command).
function commandName(line) {
  const m = /^\s*([A-Za-z_]\w*)\s*(\(|$)/.exec(line);
  return m ? m[1] : "";
}

// firstArg returns the first argument of a call, unquoted.
function firstArg(line) {
  const m = /\(\s*"?([^",)]*)"?/.exec(line);
  return m ? m[1].trim() : "";
}

// withTarget puts target(IMSI). in front of a command that subscribers run,
// as the CUI user would type it. Lines that already have a
// target, batches and commands of the console or the node stay as they are.
function withTarget(line, target, commands, signals) {
  const t = line.trim();
  if (!target || t.startsWith("@") || /^target\s*\(/.test(t)) return t;
  const name = commandName(t);
  const c = commands.find((x) => x.name === name);
  if (!c || c.scope.includes("console")) return t;
  let ue = c.scope.includes("ue");
  if (c.scope.includes("signal")) {
    const s = signals.find((x) => x.name === firstArg(t));
    ue = !!s && s.scope === "UE";
  }
  if (c.scope.includes("var") || c.scope.includes("assign")) {
    // A public variable is a timer of the node, or one value for all: no target.
    const name = firstArg(t).split("=")[0].trim();
    const v = ((S.scenario && S.scenario.vars) || []).find((x) => x.name === name);
    ue = !v || !v.public;
  }
  return ue ? `target(${target}).${t}` : t;
}

// buildLine makes a command line from form values; it returns "" and an
// error message when a required argument is missing. Arguments are
// positional: an empty optional one before a given one takes its default.
function buildLine(c, values) {
  let last = -1;
  values.forEach((v, i) => {
    if ((v || "").trim()) last = i;
  });
  const parts = [];
  for (let i = 0; i < c.args.length; i++) {
    const a = c.args[i];
    let v = (values[i] || "").trim();
    if (!v) {
      if (!a.optional && !a.variadic) return ["", `${c.name}: enter ${a.name}`];
      if (i > last) break;
      if (!a.default) return ["", `${c.name}: enter ${a.name} (it cannot be left out when a later argument is given)`];
      v = a.default;
    }
    if (a.kind === "format") parts.push(JSON.stringify(v));
    else if (a.variadic) parts.push(...v.split(",").map((x) => x.trim()).filter(Boolean));
    else parts.push(v);
  }
  return [`${c.name}(${parts.join(", ")})`, ""];
}

// messageLabel is NGAP[:NAS] of a message event.
function messageLabel(ev) {
  return ev.nas ? `${ev.ngap}:${ev.nas}` : ev.ngap || "?";
}

// eventText is the one-line text of any event.
function eventText(ev) {
  if (ev.kind === "message") {
    const arrow = ev.dir === "tx" ? "->" : "<-";
    const extra = ev.dir === "tx" ? ev.signal && `(${ev.signal})` : ev.handler && `[${ev.handler}]`;
    return `${arrow} ${messageLabel(ev)}${extra ? "  " + extra : ""}`;
  }
  let text = ev.text || "";
  if (ev.ngap && !text.includes(ev.ngap)) text = `${ev.ngap}: ${text}`;
  if (ev.handler && ev.kind !== "state") text += `  [${ev.handler}]`;
  return text;
}

// ---------- DOM helpers ----------

// nodes turns children (nested arrays, strings, null for nothing) into nodes.
function nodes(kids) {
  return kids
    .flat(Infinity)
    .filter((k) => k != null && k !== false)
    .map((k) => (k instanceof Node ? k : document.createTextNode(String(k))));
}

// h creates an element with props (class, text, on* handlers, dataset, attributes) and children.
function h(tag, props, ...kids) {
  const el = document.createElement(tag);
  for (const [k, v] of Object.entries(props || {})) {
    if (v == null || v === false) continue;
    if (k === "class") el.className = v;
    else if (k === "text") el.textContent = v;
    else if (k.startsWith("on")) el.addEventListener(k.slice(2), v);
    else if (k === "dataset") Object.assign(el.dataset, v);
    else el.setAttribute(k, v === true ? "" : v);
  }
  el.append(...nodes(kids));
  return el;
}

// fill replaces the children of el.
function fill(el, ...kids) {
  el.replaceChildren(...nodes(kids));
}

// $ returns the element with the given id.
const $ = (id) => document.getElementById(id);

// badge makes a small status label; cls picks its color.
function badge(text, cls) {
  return h("span", { class: "badge" + (cls ? " " + cls : "") }, text);
}

// notice shows a message in the notice area; an empty text clears it.
function notice(text, cls) {
  const n = $("notice");
  n.textContent = text || "";
  n.className = cls || "";
}

// ---------- API ----------

const api = {
  // get reads JSON from the API and throws on an error status.
  async get(path) {
    const r = await fetch("/api/v1" + path, { cache: "no-store" });
    const body = await r.json().catch(() => ({}));
    if (!r.ok) throw new Error(body.error || `${r.status} ${r.statusText}`);
    return body;
  },
  // send makes a JSON request (the editor's POST and PUT).
  async send(method, path, body) {
    const r = await fetch("/api/v1" + path, {
      method,
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });
    const b = await r.json().catch(() => ({}));
    if (!r.ok) throw new Error(b.error || `${r.status} ${r.statusText}`);
    return b;
  },
  // run posts a command line to /api/v1/commands.
  async run(line) {
    const r = await fetch("/api/v1/commands", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ line }),
    });
    const body = await r.json().catch(() => ({}));
    if (!r.ok) throw new Error(body.error || `${r.status} ${r.statusText}`);
    return body;
  },
};

// ---------- commands ----------

// out appends a line to the command output, keeping the last 400 lines.
function out(text, cls) {
  const o = $("cmdout");
  o.append(h("div", { class: cls || "out" }, text));
  while (o.childNodes.length > 400) o.firstChild.remove();
  o.scrollTop = o.scrollHeight;
}

// run sends a command line (with the target put in front when it applies)
// and shows the reply under the command bar.
async function run(line) {
  const eff = withTarget(line, S.target, S.commands, S.signals);
  if (!eff) return null;
  $("cmdout").hidden = false;
  out("gnodeb> " + eff, "cmd");
  try {
    const r = await api.run(eff);
    if (r.output) for (const l of r.output.split("\n")) out(l);
    if (r.error) out("error: " + r.error, "err");
    soon();
    return r;
  } catch (e) {
    out("error: " + e.message, "err");
    return null;
  }
}

// ---------- target ----------

// setTarget makes imsi the target of commands ("" for every subscriber) and updates the views that show it.
function setTarget(imsi) {
  S.target = imsi || "";
  S.ueDetail = null;
  renderTarget();
  if (S.page === "ues") renderUEs();
  if (S.page === "scenario") renderScenario();
  loadDetail();
}

// targetChip renders the current target, with a button to go back to all subscribers.
function targetChip() {
  if (!S.target) return h("span", { class: "chip all", title: "Pick a subscriber on the Subscribers page to make it the only target" }, "Target: all subscribers");
  return h(
    "span",
    { class: "chip", title: "Commands get target(" + S.target + "). in front" },
    "Target: " + S.target,
    h("button", { type: "button", title: "Back to all subscribers", onclick: () => setTarget("") }, "×"),
  );
}

// renderTarget shows the target in the command bar and the side bar.
function renderTarget() {
  fill($("cmdTarget"), targetChip());
  const side = $("sideTarget");
  fill(side, targetChip());
  if (S.ueDetail) side.append(h("small", null, `state ${S.ueDetail.state} · ${S.ueDetail.rm} · ${S.ueDetail.cm}`));
}

// ---------- pages ----------

// showPage switches to a page, keeps it in the URL hash and refreshes its contents.
function showPage(name) {
  if (!PAGES[name]) name = "dashboard";
  S.page = name;
  for (const b of document.querySelectorAll("nav button")) b.classList.toggle("active", b.dataset.page === name);
  for (const p of Object.keys(PAGES)) $("page-" + p).hidden = p !== name;
  $("content").classList.toggle("fill", $("page-" + name).classList.contains("fill"));
  $("title").textContent = PAGES[name];
  if (location.hash !== "#" + name) history.replaceState(null, "", "#" + name);
  if (name === "dashboard") renderDashboard();
  if (name === "ues") loadUEs();
  if (name === "scenario") renderScenario();
  if (name === "signals") loadSignalFile();
  if (name === "license") renderLicense();
  if (name === "sequence") scrollEnd($("ladder"), S.seq.follow);
  if (name === "logs") scrollEnd($("loglist"), S.log.follow);
}

// scrollEnd scrolls el to the bottom when follow is on.
function scrollEnd(el, follow) {
  if (el && follow) el.scrollTop = el.scrollHeight;
}

// ---------- side bar and dashboard ----------

// setupBadge shows the NG Setup result of the node.
function setupBadge(n) {
  if (n.setup === "ok") return badge("OK");
  if (n.setup === "failed") return badge("failed", "err");
  return badge(n.connected ? "pending" : "—", "off");
}

// renderSide updates the side bar and the link buttons from the status.
function renderSide() {
  const st = S.status;
  if (!st) return;
  const n = st.node;
  $("sideSctp").replaceWith(Object.assign(badge(n.connected ? "up" : "down", n.connected ? "" : "off"), { id: "sideSctp" }));
  $("sideSetup").replaceWith(Object.assign(setupBadge(n), { id: "sideSetup" }));
  $("version").textContent = `gnodeb ${st.version || ""} · ${st.config || ""}`;
  $("linkup").disabled = n.connected;
  $("linkdown").disabled = !n.connected;
  const peer = $("ladderPeer");
  if (peer) peer.textContent = peerName();
  renderTarget();
}

// setStream shows whether the event stream is live, with the count of dropped events.
function setStream(on) {
  S.streaming = on;
  const b = on ? badge(S.lost ? `live (${S.lost} dropped)` : "live", S.lost ? "warn" : "") : badge("down (reconnecting)", "warn");
  $("sideStream").replaceWith(Object.assign(b, { id: "sideStream" }));
}

// card makes a dashboard card: a title with an optional badge, a big value and sub lines.
function card(title, extra, big, ...subs) {
  return h(
    "div",
    { class: "card" },
    h("h3", null, title, extra),
    h("div", { class: "big" }, big),
    subs.filter(Boolean).map((s) => h("div", { class: "sub" }, s)),
  );
}

// stateNames lists the scenario states in file order, then any other state in the status.
function stateNames() {
  const names = [];
  if (S.scenario) for (const b of S.scenario.blocks) if (b.kind === "state") names.push(b.name);
  if (S.status) for (const k of Object.keys(S.status.states)) if (!names.includes(k)) names.push(k);
  return names;
}

// renderDashboard builds the dashboard from the latest status.
function renderDashboard() {
  const page = $("page-dashboard");
  const st = S.status;
  if (!st) {
    fill(page, h("p", { class: "muted" }, "Loading…"));
    return;
  }
  const n = st.node;
  const c = st.counters;
  // Cards: SCTP link, NG Setup, subscribers and message counters.
  const cards = h(
    "div",
    { class: "cards" },
    card(
      "SCTP (N2)",
      badge(n.connected ? "up" : "down", n.connected ? "" : "off"),
      n.connected ? "connected" : "not connected",
      n.connected ? `${n.local} → ${n.remote}` : "connect with linkup",
      n.connected ? `streams out ${n.outStreams} · in ${n.inStreams}` : "",
    ),
    card("NG Setup", setupBadge(n), n.peerName || (n.setup === "failed" ? "failed" : "—"), n.cause, ...(n.peerInfo || [])),
    card(
      "Subscribers",
      null,
      String(st.subscribers),
      `registered ${st.rm.REGISTERED || 0} · connected ${st.cm.CONNECTED || 0} · idle ${st.cm.IDLE || 0}`,
      `N3 tunnels ${st.sessions}`,
    ),
    card(
      "Messages",
      c.errors ? badge(`errors ${c.errors}`, "err") : null,
      `${c.tx} / ${c.rx}`,
      "sent / received",
      `warnings ${c.warnings} · unhandled ${c.unhandled}` + (c.dropped ? ` · dropped from the display ${c.dropped}` : ""),
    ),
  );

  // Subscribers per state as bars.
  const max = Math.max(1, st.subscribers);
  const bars = h("div", { class: "bars" });
  for (const name of stateNames()) {
    const v = st.states[name] || 0;
    bars.append(
      h(
        "div",
        { class: "bar" },
        h("span", { class: "mono" }, name),
        h("div", { class: "track" }, h("div", { class: "fill", style: `width:${(100 * v) / max}%` })),
        h("span", { class: "num mono" }, String(v)),
      ),
    );
  }
  // NGAP messages sent and received, by name.
  const msgs = Object.entries(c.messages || {}).sort((a, b) => a[0].localeCompare(b[0]));
  const table = h(
    "table",
    null,
    h("thead", null, h("tr", null, h("th", null, "NGAP"), h("th", { class: "num" }, "Sent"), h("th", { class: "num" }, "Received"))),
    h(
      "tbody",
      null,
      msgs.map(([name, m]) =>
        h("tr", null, h("td", { class: "mono" }, name), h("td", { class: "num" }, String(m.tx)), h("td", { class: "num" }, String(m.rx))),
      ),
    ),
  );
  // Runtime facts.
  const started = new Date(st.started);
  const up = Math.max(0, Math.floor((Date.now() - started) / 1000));
  const facts = h(
    "dl",
    { class: "facts" },
    [
      ["Config file", st.config],
      ["Role", st.role],
      ["snap", st.snap ? "on" : "off (base.snap: false)"],
      ["Started", dateTimeOf(started)],
      ["Uptime", `${Math.floor(up / 3600)} h ${Math.floor((up % 3600) / 60)} min ${up % 60} s`],
      ["Version", st.version],
    ].map(([k, v]) => [h("dt", null, k), h("dd", null, v || "—")]),
  );
  fill(page,
    cards,
    h(
      "div",
      { class: "grid2" },
      h("div", { class: "panel" }, h("h2", null, "Subscribers by state"), stateNames().length ? bars : h("p", { class: "empty" }, "No states")),
      h("div", { class: "panel" }, h("h2", null, "NGAP messages"), msgs.length ? table : h("p", { class: "empty" }, "Nothing sent or received yet")),
    ),
    h("div", { class: "panel" }, h("h2", null, "Runtime"), facts),
  );
}

// ---------- subscribers ----------

// buildUEsPage builds the Subscribers page: filters, the table and the detail panel.
function buildUEsPage() {
  const sel = h(
    "select",
    { id: "ueState", onchange: (e) => ((S.ueState = e.target.value), loadUEs()) },
    h("option", { value: "" }, "All states"),
    stateNames().map((s) => h("option", { value: s }, s)),
  );
  const q = h("input", {
    id: "ueQuery",
    type: "search",
    placeholder: "Filter by IMSI",
    oninput: (e) => ((S.ueQuery = e.target.value.trim()), renderUEs()),
  });
  fill($("page-ues"),
    h(
      "div",
      { class: "toolbar" },
      sel,
      q,
      h("span", { id: "ueCount", class: "muted grow" }),
      h("button", { type: "button", class: "secondary", onclick: () => setTarget("") }, "Target all subscribers"),
    ),
    h(
      "div",
      { class: "ue-layout" },
      h("div", { class: "panel table-wrap" }, h("table", { id: "ueTable" })),
      h("div", { class: "panel", id: "ueDetail" }, h("p", { class: "muted" }, "Pick a subscriber to see its details here. The picked subscriber becomes the target of commands.")),
    ),
  );
}

// loadUEs fetches the subscriber list (filtered by state on the server) and renders it.
async function loadUEs() {
  try {
    const q = "/ues?limit=1000" + (S.ueState ? "&state=" + encodeURIComponent(S.ueState) : "");
    const r = await api.get(q);
    S.ues = r.items;
    S.uesTotal = r.total;
    renderUEs();
  } catch (e) {
    notice("Cannot read the subscriber list: " + e.message);
  }
}

// renderUEs fills the subscriber table, filtered by IMSI; a click on a row picks or unpicks the target.
function renderUEs() {
  const table = $("ueTable");
  if (!table) return;
  const items = S.ueQuery ? S.ues.filter((u) => u.imsi.includes(S.ueQuery)) : S.ues;
  $("ueCount").textContent = `${S.uesTotal} subscriber${S.uesTotal === 1 ? "" : "s"}` + (items.length !== S.uesTotal ? ` (${items.length} shown)` : "");
  const head = h(
    "thead",
    null,
    h("tr", null, ["IMSI", "State", "RM", "CM", "5G-TMSI", "UE IP", "RAN / AMF ID"].map((t) => h("th", null, t))),
  );
  const body = h("tbody");
  for (const u of items) {
    const ips = (u.sessions || []).filter((s) => s.ueIp).map((s) => s.ueIp + (s.active ? "" : " (N3 down)"));
    body.append(
      h(
        "tr",
        { class: "clickable" + (u.imsi === S.target ? " selected" : ""), onclick: () => setTarget(u.imsi === S.target ? "" : u.imsi) },
        h("td", { class: "mono" }, u.imsi),
        h("td", { class: "mono" }, u.state),
        h("td", null, badge(u.rm === "REGISTERED" ? "registered" : "deregistered", u.rm === "REGISTERED" ? "" : "off")),
        h("td", null, badge(u.cm === "CONNECTED" ? "connected" : "idle", u.cm === "CONNECTED" ? "info" : "off")),
        h("td", { class: "mono", title: u.guti ? "5G-GUTI " + u.guti : null }, u.guti ? u.guti.slice(-8) : "—"),
        h("td", { class: "mono" }, ips.join(", ") || "—"),
        h("td", { class: "mono" }, `${u.ranUeNgapId ?? "—"} / ${u.amfUeNgapId ?? "—"}`),
      ),
    );
  }
  fill(table, head, body);
}

// loadDetail fetches the details of the target subscriber and renders them.
async function loadDetail() {
  if (!S.target) {
    S.ueDetail = null;
    renderDetail();
    return;
  }
  try {
    S.ueDetail = await api.get("/ues/" + encodeURIComponent(S.target));
  } catch (e) {
    S.ueDetail = null;
  }
  renderDetail();
  renderTarget();
}

// sessionTable renders the PDU sessions of a subscriber.
function sessionTable(ss) {
  const hex = (v) => (v ? "0x" + v.toString(16) : "—");
  const head = h("thead", null, h("tr", null, ["ID", "UE IP", "UPF", "UL TEID", "DL TEID", "QFI", "N3"].map((t) => h("th", null, t))));
  const rows = ss.map((s) =>
    h(
      "tr",
      null,
      h("td", null, String(s.id)),
      h("td", { class: "mono" }, s.ueIp || "—"),
      h("td", { class: "mono" }, s.upf || "—"),
      h("td", { class: "mono" }, hex(s.ulTeid)),
      h("td", { class: "mono" }, hex(s.dlTeid)),
      h("td", null, (s.qfi || []).join(", ") || "—"),
      h("td", null, badge(s.active ? "up" : "down", s.active ? "" : "off")),
    ),
  );
  return h("div", { class: "table-scroll" }, h("table", null, head, h("tbody", null, rows)));
}

// renderDetail shows the target subscriber's state, NAS security, PDU sessions and variables.
function renderDetail() {
  const box = $("ueDetail");
  if (!box) return;
  const u = S.ueDetail;
  if (!u) {
    fill(box, h("p", { class: "muted" }, "Pick a subscriber to see its details here. The picked subscriber becomes the target of commands."));
    return;
  }
  const kv = (pairs) => h("dl", { class: "kv" }, pairs.map(([k, v]) => [h("dt", null, k), h("dd", { class: "mono" }, v ?? "—")]));
  const sec = u.security;
  const parts = [
    h("h2", null, "Subscriber " + u.imsi),
    kv([
      ["Scenario state", u.state],
      ["RM / CM", `${u.rm} / ${u.cm}`],
      ["5G-GUTI", u.guti],
      ["RAN-UE-NGAP-ID", u.ranUeNgapId],
      ["AMF-UE-NGAP-ID", u.amfUeNgapId],
      ["NAS security", sec ? `${sec.integrity} / ${sec.ciphering} (ngKSI ${sec.ngKsi})` : "none"],
      ["NAS COUNT", sec ? `UL ${sec.ulCount} · DL ${sec.dlCount}` : "—"],
      ["Last received", u.lastMsg],
    ]),
  ];
  const ss = u.sessions || [];
  parts.push(h("h2", null, "PDU sessions"));
  parts.push(ss.length ? sessionTable(ss) : h("p", { class: "empty" }, "None"));
  const vars = Object.entries(u.vars || {});
  if (vars.length) {
    parts.push(h("h2", null, "Private variables"));
    parts.push(kv(vars.sort((a, b) => a[0].localeCompare(b[0]))));
  }
  fill(box, parts);
}

// ---------- operations ----------

// selectOf makes a select from [value, label] pairs, with a "(none)" entry when optional.
function selectOf(options, optional) {
  return h(
    "select",
    null,
    optional ? h("option", { value: "" }, "(none)") : null,
    options.map(([v, label]) => h("option", { value: v }, label)),
  );
}

// placeholderFor gives the hint shown in an empty argument input.
function placeholderFor(cmd, arg) {
  const p = {
    "ping:dst": "10.10.10.20",
    "ping:count": "1 (0: until stop)",
    "iperf:dst": "10.10.10.20",
    "iperf:rate": "1M (100M, 1G, …)",
    "iperftcp:dst": "10.10.10.20",
    "iperftcp:rate": "0 (no limit; 100M, 1G, …)",
    "wait:time": "3s",
    "set:var = value": "a = a + 5",
    "sync_send:dst": "127.0.0.1:12345",
  }[`${cmd}:${arg.name}`];
  if (p) return p;
  if (arg.default) return arg.default;
  return arg.variadic ? "values (comma-separated)" : "value";
}

// commandForm makes the form of one command, with an input per argument; submitting runs it.
function commandForm(c) {
  const inputs = c.args.map((a) => {
    let el;
    if (a.kind === "signal") el = selectOf(S.signals.map((s) => [s.name, `${s.name} (${s.scope})`]), a.optional);
    else if (a.kind === "state") el = selectOf(stateNames().map((s) => [s, s]), a.optional);
    else if (a.kind === "command") el = selectOf(S.commands.map((x) => [x.name, x.name]), a.optional);
    else if (a.kind === "choice") el = selectOf((a.choices || []).map((x) => [x, x]), a.optional);
    else el = h("input", { type: "text", placeholder: a.kind === "format" ? "format (e.g. state %s)" : placeholderFor(c.name, a), size: 16 });
    return { a, el };
  });
  // What the command runs on, by the first scope word it has.
  const scopeText = [
    ["console", "console"],
    ["signal", "depends on the signal"],
    ["target", "the selected subscriber, or without one the node and all subscribers"],
    ["var", "node for a public variable, subscribers for a private one"],
    ["assign", "once for a public variable, subscribers for a private one"],
  ].find(([k]) => c.scope.includes(k));
  const scope = scopeText ? scopeText[1] : "subscribers";
  return h(
    "form",
    {
      class: "cmd-form",
      onsubmit: (e) => {
        e.preventDefault();
        const [line, err] = buildLine(
          c,
          inputs.map((i) => i.el.value),
        );
        if (err) {
          $("cmdout").hidden = false;
          out("error: " + err, "err");
          return;
        }
        run(line);
      },
    },
    h("div", { class: "name" }, c.name),
    h(
      "div",
      { class: "args" },
      inputs.map(({ a, el }) => h("label", null, a.name + (a.optional ? " (optional)" : ""), el)),
      c.args.length ? null : h("span", { class: "muted" }, "no arguments"),
      h("div", { class: "help" }, `${c.help} (runs on: ${scope})`),
    ),
    h("button", { type: "submit" }, "Run"),
  );
}

// renderOps builds the Operations page: signal buttons, command forms, batch and template viewer.
function renderOps() {
  // Signal buttons, grouped by NGAP message.
  const groups = new Map();
  for (const s of S.signals) {
    const k = s.ngap || "(undecodable signals)";
    if (!groups.has(k)) groups.set(k, []);
    groups.get(k).push(s);
  }
  const sigPanel = h(
    "div",
    { class: "panel" },
    h("h2", null, "Send a signal"),
    h(
      "p",
      { class: "muted" },
      "UE signals are sent by the target subscriber (all subscribers when none is picked). Node signals (dark buttons) are sent by the node. A button runs send(signal name).",
    ),
    [...groups.entries()].map(([ngap, list]) =>
      h(
        "div",
        { class: "sig-group" },
        h("h3", null, ngap),
        h(
          "div",
          { class: "sig-buttons" },
          list.map((s) =>
            h(
              "button",
              { type: "button", class: s.scope === "UE" ? "" : "node", title: `${s.label} (${s.scope}, ${s.mode}, signal.txt line ${s.line})`, onclick: () => run(`send(${s.name})`) },
              s.name,
            ),
          ),
        ),
      ),
    ),
  );
  // A form for each command that can be run interactively.
  const cmds = S.commands.filter((c) => c.where.includes("interactive"));
  const cmdPanel = h("div", { class: "panel" }, h("h2", null, "Commands"), cmds.map(commandForm));
  // Batch: run an @file on the host of gnodeb.
  const batch = h("input", { type: "text", placeholder: "e.g. cmds.txt (an absolute path, or relative to where gnodeb was started)", class: "grow" });
  const batchPanel = h(
    "form",
    {
      class: "panel",
      onsubmit: (e) => {
        e.preventDefault();
        if (batch.value.trim()) run("@" + batch.value.trim());
      },
    },
    h("h2", null, "Batch (@file)"),
    h("p", { class: "muted" }, "Runs a file on the host of gnodeb one line at a time, stopping at the first line that fails."),
    h("div", { class: "toolbar" }, batch, h("button", { type: "submit" }, "Run")),
  );
  // Template viewer: a loaded signal decoded and as a hex dump.
  const tplSel = selectOf(S.signals.map((s) => [s.name, `${s.name} (${s.label})`]), false);
  const tplBox = h("div");
  const tplPanel = h(
    "div",
    { class: "panel" },
    h("h2", null, "Template contents"),
    h(
      "div",
      { class: "toolbar" },
      tplSel,
      h(
        "button",
        {
          type: "button",
          class: "secondary",
          onclick: async () => {
            fill(tplBox, h("p", { class: "muted" }, "Loading…"));
            try {
              const r = await api.get("/signals/" + encodeURIComponent(tplSel.value));
              fill(tplBox,
                r.error ? h("p", { class: "k-error" }, r.error) : null,
                r.tree ? h("div", { class: "tree" }, treeNode(r.tree, 0)) : null,
                h("pre", { class: "hexdump" }, hexDump(r.signal.hex)),
              );
            } catch (e) {
              fill(tplBox, h("p", { class: "k-error" }, e.message));
            }
          },
        },
        "Show",
      ),
    ),
    h("p", { class: "muted" }, "When sending, the dynamic fields among these values (IDs, keys, PLMN, …) are corrected automatically."),
    tplBox,
  );
  fill($("page-ops"), sigPanel, cmdPanel, batchPanel, tplPanel);
}

// ---------- signal editor ----------
// signal.txt on disk, one edit at a time through POST /signaledit (the
// server keeps no state: the page holds the octets and the undo history).
// Save writes the file; reload makes gnodeb send the file's signals.

const SE = {
  file: null, // GET /signalfile
  name: "", // the signal edited ("" for a new one)
  isNew: false,
  mode: "auto",
  hex: "",
  base: "", // the file's octets: changes are shown against them
  baseMode: "auto",
  tree: null,
  diags: [],
  history: [], // {hex, mode} before each edit
  error: "",
  closed: new Set(), // paths of the nodes folded
  newName: null, // the name typed for "Save as new" (null: not asked)
  newType: "", // the NGAP message picked for a new signal
  newFilter: "", // the filter of the NGAP message list
  types: null, // GET /signaltypes: NGAP, 5GMM and 5GSM messages the editor makes
  busy: false,
};

// signalDirty tells whether the edited signal differs from what signal.txt has.
function signalDirty() {
  return SE.isNew ? SE.hex !== "" : SE.hex !== SE.base || SE.mode !== SE.baseMode;
}

// loadSignalFile reads signal.txt and renders the editor.
async function loadSignalFile() {
  try {
    SE.file = await api.get("/signalfile");
    if (!SE.types) SE.types = await api.get("/signaltypes");
  } catch (e) {
    SE.file = null;
    SE.error = "Cannot read signal.txt: " + e.message;
  }
  renderSignals();
}

// signalEdit asks the server for the tree of hex, after op when there is one.
// On success the previous octets go to the undo history (with keep).
async function signalEdit(hex, mode, op, keep) {
  if (SE.busy) return false;
  SE.busy = true;
  try {
    const r = await api.send("POST", "/signaledit", { name: SE.name, hex, mode, op: op || undefined });
    if (keep) SE.history.push({ hex: SE.hex, mode: SE.mode });
    SE.hex = r.hex;
    SE.mode = mode;
    SE.tree = r.tree;
    SE.diags = r.diags || [];
    SE.error = "";
    return true;
  } catch (e) {
    SE.error = e.message;
    return false;
  } finally {
    SE.busy = false;
    renderSignals();
  }
}

// signalOp applies one edit to the current octets, keeping the undo history.
function signalOp(op) {
  return signalEdit(SE.hex, SE.mode, op, true);
}

// openSignal loads a signal of signal.txt into the editor (asking before discarding changes).
async function openSignal(entry) {
  if (signalDirty() && !confirm(`Discard the unsaved changes to ${SE.name || "the new signal"}?`)) return;
  Object.assign(SE, { name: entry.name, isNew: false, mode: entry.mode, base: entry.hex, baseMode: entry.mode, history: [], closed: new Set(), newName: null, error: "" });
  SE.hex = entry.hex;
  await signalEdit(entry.hex, entry.mode, null, false);
}

// newSignal starts an empty new signal (asking before discarding changes).
function newSignal() {
  if (signalDirty() && !confirm(`Discard the unsaved changes to ${SE.name || "the new signal"}?`)) return;
  Object.assign(SE, { name: "", isNew: true, mode: "auto", hex: "", base: "", baseMode: "auto", tree: null, diags: [], history: [], closed: new Set(), newName: "", newType: "", newFilter: "", error: "" });
  renderSignals();
}

// saveSignal writes the edited signal back into signal.txt.
async function saveSignal() {
  try {
    SE.file = await api.send("PUT", "/signalfile/" + encodeURIComponent(SE.name), { hex: SE.hex, mode: SE.mode });
    SE.base = SE.hex;
    SE.baseMode = SE.mode;
    SE.error = "";
    notice(`Saved ${SE.name} to signal.txt. Run reload to send it.`);
  } catch (e) {
    SE.error = "Cannot save: " + e.message;
  }
  renderSignals();
}

// saveSignalAsNew adds the edited octets to signal.txt under the typed name.
async function saveSignalAsNew() {
  const name = (SE.newName || "").trim();
  if (!name) {
    SE.error = "Type a name for the new signal (letters, digits and _).";
    renderSignals();
    return;
  }
  try {
    SE.file = await api.send("POST", "/signalfile", { name, hex: SE.hex, mode: SE.mode });
    Object.assign(SE, { name, isNew: false, base: SE.hex, baseMode: SE.mode, newName: null, error: "" });
    notice(`Added ${name} to signal.txt. Run reload to send it.`);
  } catch (e) {
    SE.error = "Cannot add: " + e.message;
  }
  renderSignals();
}

// reloadSignals runs reload so gnodeb sends the file's signals, then reads the signal lists again.
async function reloadSignals() {
  const r = await run("reload");
  if (!r || r.error) return;
  notice("gnodeb sends the signals of signal.txt now.");
  SE.error = "";
  try {
    S.signals = await api.get("/signals");
    renderOps();
  } catch (e) {
    /* the operations page keeps its list */
  }
  await loadSignalFile();
}

// rangeText describes a node's size or value range, e.g. "1..8 octets (extensible)".
function rangeText(n) {
  if (n.min == null && n.max == null) return "";
  const unit = { int: "", octets: " octets", bits: " bits", string: " characters", list: " items" }[n.k] ?? "";
  const r = n.min === n.max ? `${n.min}` : `${n.min ?? "…"}..${n.max ?? "…"}`;
  return r + unit + (n.ext ? " (extensible)" : "");
}

// editInput makes a value input that calls onCommit when changed; Escape restores the value.
function editInput(n, cls, value, onCommit) {
  const v = value ?? "";
  const inp = h("input", {
    class: "ev " + cls,
    value: v,
    size: Math.max(6, Math.min(64, v.length + 2)),
    spellcheck: "false",
    disabled: SE.busy,
    title: [rangeText(n), n.note].filter(Boolean).join(" · ") || null,
  });
  inp.addEventListener("change", () => {
    if (inp.value !== v) onCommit(inp.value.trim());
  });
  inp.addEventListener("keydown", (e) => {
    if (e.key === "Escape") {
      inp.value = v;
      inp.blur();
    }
  });
  return inp;
}

// editValue renders the editor of a node's value.
function editValue(n) {
  const set = (value, extra) => signalOp({ kind: "set", path: n.p, value, ...extra });
  const hint = (t) => (t ? h("small", { class: "eh" }, t) : null);
  switch (n.k) {
    case "int":
      if (n.enum && n.enum.length) return enumValue(n, set);
    // fall through: an integer without names
    case "string":
    case "octets":
    case "plmn":
    case "ip":
      // A fixed size is said by the size itself; a range is shown.
      return [editInput(n, "ev-" + n.k, n.v, (v) => set(v)), hint(n.h), hint(n.k === "octets" && n.min === n.max ? "" : rangeText(n))];
    case "bits": {
      let bits = n.bits;
      const val = editInput(n, "ev-bits", n.v, (v) => set(v, { bits }));
      const len = editInput(n, "ev-len", String(n.bits), (v) => {
        bits = Number(v);
        set(val.value.trim(), { bits });
      });
      return [val, len, hint(n.min != null && n.min === n.max ? "bits" : "bits " + rangeText(n))];
    }
    case "bool": {
      const sel = h("select", { class: "ev", disabled: SE.busy, onchange: (e) => set(e.target.value) }, ["true", "false"].map((b) => h("option", { value: b, selected: n.v === b }, b)));
      return sel;
    }
    case "choice":
      // A CHOICE, or an IE container entry (its IE, with the IE ID and criticality as the hint).
      return [
        h(
          "select",
          { class: "ev echoice", disabled: SE.busy, title: "The alternative of the CHOICE (the IE of an IE container entry)", onchange: (e) => signalOp({ kind: "choose", path: n.p, value: e.target.value }) },
          n.alt ? null : h("option", { value: "", selected: true }, "pick one"),
          n.alts.map((a) => h("option", { value: a, selected: a === n.alt }, a)),
        ),
        hint(n.h),
      ];
    case "list":
      return hint([n.h, rangeText(n)].filter(Boolean).join(", "));
    case "container": {
      const box = h("span", { class: "ehex", hidden: true }, editInput(n, "ev-octets", n.v, (v) => signalOp({ kind: "hex", path: n.p, value: v })));
      // A NAS-PDU or NAS container: make a NAS message of a type there.
      const kinds = SE.types && (n.nas === "5gmm" ? SE.types.gmm : n.nas === "5gsm" ? SE.types.gsm : null);
      const has = n.c && n.c.length > 0;
      const nasSel = kinds
        ? h(
            "select",
            { class: "ev echoice", disabled: SE.busy, title: "Make a new NAS message of this type here (its fields get their smallest values)", onchange: (e) => e.target.value && signalOp({ kind: "nas", path: n.p, value: e.target.value }) },
            h("option", { value: "" }, has ? "replace with a NAS message…" : "pick a NAS message"),
            kinds.map((k) => h("option", { value: k }, k)),
          )
        : null;
      return [
        hint(n.h),
        nasSel,
        h("button", { type: "button", class: "small secondary", title: "Replace the octets of this nested message as hex", onclick: () => (box.hidden = !box.hidden) }, "hex"),
        box,
      ];
    }
    case "ext":
      return h("span", { class: "tv" }, n.v);
  }
  return hint(n.h);
}

// enumValue renders an ENUMERATED value: its names in a select, and a box
// for a number the type does not name (for tests of invalid values).
function enumValue(n, set) {
  const named = n.enum.some((o) => String(o.v) === n.v);
  const box = h("span", { class: "ehex", hidden: named }, editInput(n, "ev-int", n.v, (v) => set(v)));
  const sel = h(
    "select",
    { class: "ev echoice", disabled: SE.busy, title: rangeText(n), onchange: (e) => e.target.value !== "" && set(e.target.value) },
    named ? null : h("option", { value: "", selected: true }, `${n.v} (no name)`),
    n.enum.map((o) => h("option", { value: String(o.v), selected: String(o.v) === n.v }, `${o.n} (${o.v})${o.ext ? " ext" : ""}`)),
  );
  return [sel, h("button", { type: "button", class: "small secondary", title: "Enter a number instead", onclick: () => (box.hidden = !box.hidden) }, "#"), box];
}

// editActions renders what can be done to a node's structure.
function editActions(n) {
  if (SE.busy) return [];
  const acts = [];
  if (n.add && n.add.length) {
    acts.push(
      h(
        "select",
        { class: "ev eadd", title: "Add an optional field", onchange: (e) => e.target.value && signalOp({ kind: "add", path: n.p, value: e.target.value }) },
        h("option", { value: "" }, "+ field"),
        n.add.map((a) => h("option", { value: a }, a)),
      ),
    );
  }
  if (n.k === "list") {
    const full = n.max != null && !n.ext && n.c && n.c.length >= n.max;
    acts.push(h("button", { type: "button", class: "small secondary", disabled: full, title: "Add an item at the end", onclick: () => signalOp({ kind: "add", path: n.p }) }, "+ item"));
  }
  if (n.del) acts.push(h("button", { type: "button", class: "small secondary edel", title: "Remove " + n.n, onclick: () => signalOp({ kind: "remove", path: n.del }) }, "×"));
  return acts;
}

// editNode renders a node of the editor tree with its value, actions and (foldable) children.
function editNode(n) {
  const kids = n.c || [];
  const folded = SE.closed.has(n.p);
  const caret = kids.length
    ? h("button", { type: "button", class: "caret", title: folded ? "Open" : "Fold", onclick: () => (folded ? SE.closed.delete(n.p) : SE.closed.add(n.p), renderSignals()) }, folded ? "▸" : "▾")
    : h("span", { class: "caret" });
  const row = h(
    "div",
    { class: "erow" + (n.auto ? " auto" : "") },
    caret,
    h("span", { class: "tn", title: n.p || "(the message)" }, n.n),
    editValue(n),
    n.auto ? h("span", { class: "badge off", title: (n.note ? n.note + ". " : "") + "Mapped: replaced when the signal is sent (send mode auto or nas=keep)" }, "auto") : null,
    editActions(n),
    n.note && !n.auto ? h("small", { class: "enote" }, n.note) : null,
  );
  if (!kids.length || folded) return h("div", { class: "enode" }, row);
  return h("div", { class: "enode" }, row, h("div", { class: "ekids" }, kids.map(editNode)));
}

// hexDiff is hexDump with the octets that differ from base marked.
function hexDiff(hex, base) {
  const bytes = (hex || "").match(/../g) || [];
  const old = (base || "").match(/../g) || [];
  const pre = h("pre", { class: "hexdump" });
  for (let i = 0; i < bytes.length; i += 16) {
    pre.append(pad(i.toString(16), 4) + "  ");
    let asc = "";
    for (let j = 0; j < 16; j++) {
      const k = i + j;
      if (k < bytes.length) {
        pre.append(bytes[k] !== old[k] ? h("span", { class: "chg" }, bytes[k]) : bytes[k]);
        const c = parseInt(bytes[k], 16);
        asc += c >= 0x20 && c < 0x7f ? String.fromCharCode(c) : ".";
      } else pre.append("  ");
      pre.append(j === 7 ? "  " : " ");
    }
    pre.append(" " + asc + "\n");
  }
  return pre;
}

// renderSignals builds the Signals page: the signal list, the field editor and the octets panel.
// newSignalForm asks for the name and the NGAP message of a new signal: a
// filter narrows the list, the select picks from it.
function newSignalForm() {
  const name = h("input", { value: SE.newName || "", placeholder: "letters, digits and _", spellcheck: "false" });
  name.addEventListener("input", () => (SE.newName = name.value));
  const ngap = (SE.types && SE.types.ngap) || [];
  const filter = h("input", { value: SE.newFilter, placeholder: "filter, e.g. NAS or Setup", spellcheck: "false" });
  const pick = h("select", { size: 8, class: "ngap-pick" });
  const fillPick = () => {
    const q = SE.newFilter.toLowerCase();
    const shown = ngap.filter((t) => t.toLowerCase().includes(q));
    pick.replaceChildren(...shown.map((t) => h("option", { value: t, selected: t === SE.newType }, t)));
  };
  fillPick();
  filter.addEventListener("input", () => {
    SE.newFilter = filter.value;
    fillPick();
  });
  pick.addEventListener("change", () => (SE.newType = pick.value));
  const fail = (text) => ((SE.error = text), renderSignals());
  const create = () => {
    const n = (SE.newName || "").trim();
    if (!/^[A-Za-z0-9_]+$/.test(n)) return fail("Give the new signal a name of letters, digits and _.");
    if (SE.file && SE.file.signals.some((e) => e.name === n)) return fail(`signal.txt already has ${n}: pick another name.`);
    if (!ngap.includes(SE.newType)) return fail("Pick the NGAP message from the list.");
    SE.newName = n;
    signalEdit("", SE.mode, { kind: "new", value: SE.newType }, false);
  };
  pick.addEventListener("dblclick", create);
  return h(
    "div",
    { class: "newsig" },
    h("h3", null, "New signal"),
    h("div", { class: "toolbar" }, h("label", null, "name", name)),
    h("div", { class: "toolbar" }, h("label", null, "NGAP message", filter), h("small", null, `${ngap.length} messages`)),
    pick,
    h("div", { class: "toolbar" }, h("button", { type: "button", disabled: SE.busy, onclick: create }, "Create"), h("button", { type: "button", class: "secondary", onclick: discardSignal }, "Cancel")),
    h(
      "p",
      { class: "muted" },
      "The mandatory fields get their smallest values. Edit them, add the optional ones you need, pick the NAS message of a NAS-PDU from its list, then Add. Pasting the octets on the right works too.",
    ),
  );
}

// discardSignal drops the new signal, or the unsaved changes, and goes back to no signal.
function discardSignal() {
  if (signalDirty() && !confirm(`Discard ${SE.isNew ? "the new signal" : "the unsaved changes to " + SE.name}?`)) return;
  Object.assign(SE, { name: "", isNew: false, mode: "auto", hex: "", base: "", baseMode: "auto", tree: null, diags: [], history: [], closed: new Set(), newName: null, newType: "", newFilter: "", error: "" });
  renderSignals();
}

function renderSignals() {
  const page = $("page-signals");
  if (!page) return;
  const f = SE.file;
  // Left: the signals of signal.txt.
  const list = h(
    "div",
    { class: "sig-list" },
    h("div", { class: "toolbar" }, h("button", { type: "button", onclick: newSignal, title: "Make a new signal from hex octets" }, "New"), h("button", { type: "button", class: "secondary", onclick: reloadSignals, title: "Read signal.txt again and send its signals (the reload command)" }, "Reload")),
    f ? h("small", { class: "mono", title: f.path }, f.path.split("/").pop()) : null,
    f && f.pending ? h("p", { class: "pending" }, "signal.txt has changes that gnodeb does not send yet. Reload sends them.") : null,
    f
      ? f.signals.map((e) =>
          h(
            "button",
            { type: "button", class: "item" + (e.name === SE.name && !SE.isNew ? " selected" : ""), title: `line ${e.line}`, onclick: () => openSignal(e) },
            h("span", null, e.name, e.mode !== "auto" ? h("span", { class: "badge info" }, e.mode) : null, e.loaded ? null : h("span", { class: "badge warn", title: "Saved but not reloaded" }, "not sent yet")),
            h("small", null, e.label),
          ),
        )
      : null,
  );

  // Middle: the editor header (mode, undo, save) and the field tree.
  const modeSel = h(
    "select",
    { disabled: SE.busy || (!SE.tree && !SE.isNew), title: "The send mode: auto fills in the dynamic fields, nas=keep keeps the NAS, raw sends the octets as they are", onchange: (e) => (SE.hex ? signalEdit(SE.hex, e.target.value, null, true) : ((SE.mode = e.target.value), renderSignals())) },
    ["auto", "nas=keep", "raw"].map((m) => h("option", { value: m, selected: m === SE.mode }, m)),
  );
  const dirty = signalDirty();
  const nameBox =
    SE.newName != null
      ? h(
          "span",
          { class: "toolbar-inline" },
          (() => {
            const i = h("input", { value: SE.newName, placeholder: "new signal name", spellcheck: "false" });
            i.addEventListener("input", () => (SE.newName = i.value));
            i.addEventListener("keydown", (e) => e.key === "Enter" && saveSignalAsNew());
            return i;
          })(),
          h("button", { type: "button", disabled: !SE.hex || SE.busy, onclick: saveSignalAsNew }, "Add"),
          SE.isNew ? null : h("button", { type: "button", class: "secondary", onclick: () => ((SE.newName = null), (SE.error = ""), renderSignals()) }, "Cancel"),
        )
      : null;
  const head = h(
    "div",
    { class: "toolbar" },
    h("strong", { class: "sig-title" }, SE.isNew ? (SE.tree && SE.newName ? `New: ${SE.newName}` : "New signal") : SE.name || "Pick a signal"),
    dirty ? badge("unsaved", "warn") : null,
    h("label", null, "mode", modeSel),
    h("span", { class: "grow" }),
    h("button", { type: "button", class: "secondary", disabled: !SE.history.length || SE.busy, onclick: () => { const p = SE.history.pop(); signalEdit(p.hex, p.mode, null, false); }, title: "Undo the last edit" }, "Undo"),
    SE.isNew ? null : h("button", { type: "button", class: "secondary", disabled: !dirty || SE.busy, onclick: () => signalEdit(SE.base, SE.baseMode, null, true), title: "Back to what signal.txt has" }, "Revert"),
    SE.isNew ? h("button", { type: "button", class: "secondary", disabled: SE.busy, onclick: discardSignal, title: "Drop the new signal" }, "Discard") : null,
    SE.isNew ? null : h("button", { type: "button", disabled: !dirty || SE.busy || !SE.name, onclick: saveSignal, title: "Write this signal into signal.txt" }, "Save"),
    SE.isNew || SE.newName != null ? null : h("button", { type: "button", class: "secondary", disabled: !SE.tree || SE.busy, onclick: () => ((SE.newName = ""), renderSignals()), title: "Add a copy under a new name" }, "Save as new"),
  );
  const editor = h(
    "div",
    { class: "sig-editor" },
    head,
    nameBox ? h("div", { class: "toolbar" }, nameBox) : null,
    SE.error ? h("p", { class: "k-error" }, SE.error) : null,
    SE.tree
      ? h("div", { class: "etree" }, editNode(SE.tree))
      : SE.isNew
        ? newSignalForm()
        : h("p", { class: "muted" }, "Pick a signal on the left to see and edit its fields, or make a new one with New."),
    h("p", { class: "muted" }, "Fields marked auto are mapped: gnodeb replaces them when it sends the signal (send mode auto, and the NGAP fields in nas=keep), so a value set here goes out as it is only with raw. Lengths follow the content."),
  );

  // Right: the octets with changes marked, the checks and a box to replace the hex.
  const raw = h("textarea", { class: "ev rawhex", rows: 4, spellcheck: "false", placeholder: "hex octets of the whole signal" }, SE.hex);
  const side = h(
    "div",
    { class: "sig-side" },
    h("h3", null, "Octets", SE.hex ? h("small", null, ` ${SE.hex.length / 2} (signal.txt: ${SE.base.length / 2})`) : null),
    SE.hex ? hexDiff(SE.hex, SE.base) : h("p", { class: "muted" }, "—"),
    h("h3", null, "Checks"),
    SE.diags.length
      ? h("ul", { class: "diags" }, SE.diags.map((d) => h("li", null, badge(d.sev, d.sev === "error" ? "err" : d.sev === "warning" ? "warn" : "info"), " ", d.msg)))
      : h("p", { class: "muted" }, SE.tree ? "No problems." : "—"),
    h("h3", null, "Replace with hex"),
    raw,
    h("div", { class: "toolbar" }, h("button", { type: "button", class: "secondary", disabled: SE.busy, onclick: () => signalEdit(SE.hex, SE.mode, { kind: "hex", path: "", value: raw.value }, true) }, "Apply")),
  );
  fill(page, h("div", { class: "sig-layout" }, list, editor, side));
}

// ---------- decoded messages ----------

// treeNode renders a decoded message as nested details, the first four levels open.
function treeNode(n, depth) {
  const name = h("span", { class: "tn" }, n.n);
  const value = n.v ? h("span", { class: "tv" }, n.v) : null;
  if (!n.c || !n.c.length) return h("div", { class: "leaf" }, name, value);
  const d = h("details", depth < 4 ? { open: true } : null, h("summary", null, name, value));
  d.append(h("div", { class: "kids" }, n.c.map((k) => treeNode(k, depth + 1))));
  return d;
}

// ---------- sequence ----------

// buildSequencePage builds the Sequence page: filters, the ladder and the message detail.
function buildSequencePage() {
  const src = h("select", {
    id: "seqSrc",
    onchange: (e) => {
      S.seq.src = e.target.value;
      renderLadder();
    },
  });
  fillSources(src, S.seq.src);
  const notes = h("input", {
    type: "checkbox",
    checked: S.seq.notes,
    onchange: (e) => {
      S.seq.notes = e.target.checked;
      renderLadder();
    },
  });
  const follow = h("input", {
    type: "checkbox",
    checked: S.seq.follow,
    onchange: (e) => {
      S.seq.follow = e.target.checked;
      scrollEnd($("ladder"), S.seq.follow);
    },
  });
  fill($("page-sequence"),
    h(
      "div",
      { class: "seq-layout" },
      h(
        "div",
        { class: "seq-main" },
        h(
          "div",
          { class: "toolbar" },
          h("label", null, "Source", src),
          h("label", null, notes, "Show other events too"),
          h("label", null, follow, "Follow"),
          h("span", { class: "grow" }),
          h("button", { type: "button", class: "secondary small", onclick: clearEvents }, "Clear"),
        ),
        h("div", { class: "ladder", id: "ladder" }),
      ),
      h("div", { class: "seq-detail", id: "seqDetail" }, h("p", { class: "muted" }, "Pick an arrow to see the decoded message and its hex dump.")),
    ),
  );
  renderLadder();
}

// fillSources lists node, web and the subscribers seen in the events.
function fillSources(sel, current) {
  const seen = new Set();
  for (const e of S.events) seen.add(e.src);
  if (current) seen.add(current);
  const list = [...seen].filter((x) => x).sort((a, b) => (a === "node" ? -1 : b === "node" ? 1 : a.localeCompare(b)));
  fill(sel, h("option", { value: "" }, "All"), list.map((s) => h("option", { value: s, selected: s === current }, s)));
}

// ladderMatch tells whether an event goes on the ladder under the current filters.
function ladderMatch(ev) {
  if (S.seq.src && ev.src !== S.seq.src) return false;
  return ev.kind === "message" || S.seq.notes;
}

// ladderRow renders an event as an arrow (messages) or a note line (other events).
function ladderRow(ev) {
  const time = h("span", { class: "ltime" }, timeOf(ev.time));
  const src = h("span", { class: "lsrc", title: ev.src }, ev.src);
  if (ev.kind === "message") {
    const sub = ev.dir === "tx" ? ev.signal : ev.handler && `[${ev.handler}]`;
    return h(
      "div",
      {
        class: "lrow msg" + (ev.seq === S.seq.selected ? " selected" : ""),
        dataset: { seq: ev.seq },
        title: `#${ev.seq} ${ev.dir === "tx" ? "sent" : "received"} ${messageLabel(ev)} (${ev.len} octets)`,
        onclick: () => selectMessage(ev.seq),
      },
      time,
      src,
      h(
        "div",
        { class: "lane" },
        h(
          "div",
          { class: `arrow ${ev.dir}` + (ev.handler === "unhandled" ? " unhandled" : "") },
          h("span", { class: "lbl" }, messageLabel(ev), sub ? h("small", null, "  " + sub) : null),
        ),
      ),
    );
  }
  return h("div", { class: "lrow note" }, time, src, h("div", { class: "lane k-" + ev.kind, title: eventText(ev) }, h("span", null, eventText(ev))));
}

// peerName is the AMF name learned in NG Setup, or "AMF".
function peerName() {
  return S.status && S.status.node.peerName ? S.status.node.peerName : "AMF";
}

// ladderHead renders the header row of the ladder with the gNB and AMF ends.
function ladderHead() {
  return h(
    "div",
    { class: "lhead" },
    h("span", { class: "ltime" }, "Time"),
    h("span", { class: "lsrc" }, "Source"),
    h("div", { class: "ends" }, h("span", null, "gNB (gnodeb)"), h("span", { id: "ladderPeer" }, peerName())),
  );
}

// renderLadder redraws the whole ladder from the kept events.
function renderLadder() {
  const el = $("ladder");
  if (!el) return;
  const frag = document.createDocumentFragment();
  frag.append(ladderHead());
  for (const ev of S.events) if (ladderMatch(ev)) frag.append(ladderRow(ev));
  fill(el, frag);
  scrollEnd(el, S.seq.follow);
}

// ladderAppend adds new events to the ladder, drops old rows and follows the end when scrolled there.
function ladderAppend(batch) {
  const el = $("ladder");
  if (!el) return;
  const atEnd = el.scrollHeight - el.scrollTop - el.clientHeight < 60;
  const frag = document.createDocumentFragment();
  for (const ev of batch) if (ladderMatch(ev)) frag.append(ladderRow(ev));
  el.append(frag);
  while (el.childNodes.length > MAX_EVENTS + 1) el.childNodes[1].remove(); // keep the header
  if (S.seq.follow && atEnd) el.scrollTop = el.scrollHeight;
  const sel = $("seqSrc");
  if (sel && batch.some((e) => ![...sel.options].some((o) => o.value === e.src))) fillSources(sel, S.seq.src);
}

// selectMessage marks a message on the ladder and shows it decoded with its hex dump.
async function selectMessage(seq) {
  S.seq.selected = seq;
  for (const r of document.querySelectorAll(".lrow.selected")) r.classList.remove("selected");
  const row = document.querySelector(`.lrow[data-seq="${seq}"]`);
  if (row) row.classList.add("selected");
  const box = $("seqDetail");
  fill(box, h("p", { class: "muted" }, "Loading…"));
  try {
    const r = await api.get("/trace/" + seq);
    const ev = r.event;
    const kv = [
      ["Seq", `#${ev.seq}`],
      ["Time", dateTimeOf(ev.time) + "." + pad(new Date(ev.time).getMilliseconds(), 3)],
      ["Direction", ev.dir === "tx" ? "sent (gNB → AMF)" : "received (AMF → gNB)"],
      ["Subscriber", ev.src],
      ["Message", messageLabel(ev)],
      [ev.dir === "tx" ? "Signal" : "Handler", ev.dir === "tx" ? ev.signal : ev.handler],
      ["State", ev.state],
      ["Length / stream", `${ev.len} octets / ${ev.stream ?? 0}`],
    ];
    fill(box,
      h("h2", null, messageLabel(ev)),
      h("dl", { class: "kv" }, kv.map(([k, v]) => [h("dt", null, k), h("dd", { class: "mono" }, v || "—")])),
      r.error ? h("p", { class: "k-error" }, r.error) : null,
      r.tree
        ? [
            h(
              "div",
              { class: "toolbar" },
              h("button", { type: "button", class: "secondary small", onclick: () => box.querySelectorAll("details").forEach((d) => (d.open = true)) }, "Expand all"),
              h("button", { type: "button", class: "secondary small", onclick: () => box.querySelectorAll("details").forEach((d) => (d.open = false)) }, "Collapse all"),
            ),
            h("div", { class: "tree" }, treeNode(r.tree, 0)),
          ]
        : null,
      h("h2", null, "Hex dump (NGAP)"),
      h("pre", { class: "hexdump" }, hexDump(ev.hex)),
    );
  } catch (e) {
    fill(box, h("p", { class: "k-error" }, e.message));
  }
}

// ---------- scenario ----------

// stmtLine is one line of a scenario body; the lines inside an if and its
// closing braces carry no bullet, so the if reads as one statement.
function stmtLine(line) {
  return h("li", { class: /^[ }]/.test(line) ? "cont" : null }, line);
}

// renderScenario shows the scenario: its variables and its blocks with handlers and subscriber counts.
function renderScenario() {
  const page = $("page-scenario");
  const sc = S.scenario;
  if (!sc) {
    fill(page, h("p", { class: "muted" }, "Loading…"));
    return;
  }
  const counts = S.status ? S.status.states : {};
  const current = S.target && S.ueDetail ? S.ueDetail.state : "";
  const vars = h(
    "div",
    { class: "panel" },
    h("h2", null, "Variables"),
    sc.vars.length
      ? h(
          "table",
          null,
          h("thead", null, h("tr", null, ["Name", "Type", "Scope", "Value"].map((t) => h("th", null, t)))),
          h(
            "tbody",
            null,
            sc.vars.map((v) =>
              h(
                "tr",
                null,
                h("td", { class: "mono" }, v.name),
                h("td", null, v.type),
                h("td", null, v.public ? "public (shared)" : "private (per subscriber)"),
                h("td", { class: "mono" }, v.public ? v.value : `initial ${v.value}`),
              ),
            ),
          ),
        )
      : h("p", { class: "empty" }, "None"),
  );
  // A card for each block (node, any, state) with its entry calls and handlers.
  const blocks = h(
    "div",
    { class: "blocks" },
    sc.blocks.map((b) => {
      const n = b.kind === "state" ? counts[b.name] || 0 : 0;
      const cls = "block" + (n ? " busy" : "") + (b.kind === "state" && b.name === current ? " current" : "");
      const kind = { node: "node", any: "any (all states)", state: "state" }[b.kind];
      return h(
        "div",
        { class: cls },
        h(
          "header",
          null,
          h("span", null, badge(kind, b.kind === "state" ? "info" : "off"), " ", h("strong", null, b.name)),
          h("span", null, b.kind === "state" ? badge(`${n} UE${n === 1 ? "" : "s"}`, n ? "" : "off") : null, h("small", null, `  line ${b.line}`)),
        ),
        b.entry && b.entry.length ? h("div", { class: "hdl" }, h("div", { class: "trig" }, "entry"), h("ul", null, b.entry.map(stmtLine))) : null,
        (b.handlers || []).map((hd) =>
          h("div", { class: "hdl" }, h("div", { class: "trig", title: `line ${hd.line}` }, `${hd.kind || "receive"}(${hd.trigger})`), h("ul", null, hd.calls.map(stmtLine))),
        ),
      );
    }),
  );
  fill(page,
    h("p", { class: "muted" }, `${sc.path}. Highlighted states have subscribers in them; the state of the picked subscriber stands out.`),
    vars,
    blocks,
  );
}

// ---------- log ----------

// buildLogsPage builds the Logs page: kind, source and text filters and the log list.
// ---------- license ----------

// renderLicense shows the copyright and the license in short.
function renderLicense() {
  const version = S.status && S.status.version ? S.status.version : "";
  const apache = "https://www.apache.org/licenses/LICENSE-2.0";
  fill(
    $("page-license"),
    h(
      "div",
      { class: "panel" },
      h("h2", null, "kronos-5g-emulator", version ? h("small", null, " gnodeb " + version) : null),
      h("p", null, "Copyright 2026 Kronos Systems."),
      h("p", null, "Licensed under the Apache License, Version 2.0. The full text is in LICENSE at the top of the source tree."),
    ),
    h(
      "div",
      { class: "panel" },
      h("h2", null, "Apache License 2.0 in short"),
      h(
        "ul",
        { class: "license-points" },
        h("li", null, "You may use, copy, modify and distribute the software, in source or binary form, for any purpose, commercial ones included."),
        h("li", null, "Keep the copyright and license notices with what you distribute, and give the recipients a copy of the license."),
        h("li", null, "Mark the files you changed as changed."),
        h("li", null, "Contributors grant a patent license for their contributions; it ends for anyone who sues over patents in the software."),
        h("li", null, "No right to use the trademarks or names of the licensor is granted."),
        h("li", null, "The software is provided \"AS IS\", without warranties or conditions of any kind."),
      ),
      h("p", { class: "muted" }, "This summary is no substitute for the license: ", h("a", { href: apache, target: "_blank", rel: "noopener" }, apache), "."),
    ),
  );
}

function buildLogsPage() {
  const kinds = KINDS.map((k) =>
    h(
      "label",
      null,
      h("input", {
        type: "checkbox",
        checked: S.log.kinds.has(k),
        onchange: (e) => {
          if (e.target.checked) S.log.kinds.add(k);
          else S.log.kinds.delete(k);
          renderLog();
        },
      }),
      k,
    ),
  );
  const src = h("input", {
    type: "search",
    placeholder: "Source (IMSI, node, web)",
    size: 22,
    oninput: (e) => {
      S.log.src = e.target.value.trim();
      renderLog();
    },
  });
  const q = h("input", {
    type: "search",
    placeholder: "Filter by text",
    size: 18,
    oninput: (e) => {
      S.log.q = e.target.value.trim().toLowerCase();
      renderLog();
    },
  });
  const follow = h("input", {
    type: "checkbox",
    checked: S.log.follow,
    onchange: (e) => {
      S.log.follow = e.target.checked;
      scrollEnd($("loglist"), S.log.follow);
    },
  });
  fill($("page-logs"),
    h(
      "div",
      { class: "logs-main" },
      h("div", { class: "toolbar" }, kinds, src, q, h("label", null, follow, "Follow"), h("span", { class: "grow" }), h("button", { type: "button", class: "secondary small", onclick: clearEvents }, "Clear")),
      h("div", { class: "loglist", id: "loglist" }),
    ),
  );
  renderLog();
}

// logMatch tells whether an event passes the log filters.
function logMatch(ev) {
  if (!S.log.kinds.has(ev.kind)) return false;
  if (S.log.src && !String(ev.src).includes(S.log.src)) return false;
  if (S.log.q && !eventText(ev).toLowerCase().includes(S.log.q)) return false;
  return true;
}

// logLine renders an event as a log line.
function logLine(ev) {
  const cls = { message: ev.dir === "tx" ? "out" : "", warn: "warn", error: "err", snap: "", state: "info", link: "info" }[ev.kind];
  return h(
    "div",
    { class: "logline" },
    h("span", { class: "t" }, timeOf(ev.time)),
    h("span", { class: "s", title: ev.src }, ev.src),
    h("span", null, badge(ev.kind, cls === undefined ? "off" : cls)),
    h("span", { class: "x k-" + ev.kind }, eventText(ev)),
  );
}

// renderLog redraws the whole log from the kept events.
function renderLog() {
  const el = $("loglist");
  if (!el) return;
  const frag = document.createDocumentFragment();
  for (const ev of S.events) if (logMatch(ev)) frag.append(logLine(ev));
  fill(el, frag);
  scrollEnd(el, S.log.follow);
}

// logAppend adds new events to the log, drops old lines and follows the end when scrolled there.
function logAppend(batch) {
  const el = $("loglist");
  if (!el) return;
  const atEnd = el.scrollHeight - el.scrollTop - el.clientHeight < 60;
  const frag = document.createDocumentFragment();
  for (const ev of batch) if (logMatch(ev)) frag.append(logLine(ev));
  el.append(frag);
  while (el.childNodes.length > MAX_EVENTS) el.firstChild.remove();
  if (S.log.follow && atEnd) el.scrollTop = el.scrollHeight;
}

// ---------- events ----------

// pending holds the events received but not drawn yet; flush draws them in a batch.
let pending = [];
let flushQueued = false;

// queueEvent buffers an event and schedules a flush (dropping the oldest when far behind).
function queueEvent(ev) {
  pending.push(ev);
  if (pending.length > MAX_EVENTS * 2) pending.splice(0, pending.length - MAX_EVENTS);
  if (!flushQueued) {
    flushQueued = true;
    requestAnimationFrame(flush);
    setTimeout(flush, 300); // also when the tab is hidden (no animation frames)
  }
}

// flush draws the buffered events, counting gaps in seq as dropped events.
function flush() {
  flushQueued = false;
  if (!pending.length) return;
  let batch = pending;
  pending = [];
  for (const ev of batch) {
    if (S.lastSeq && ev.seq > S.lastSeq + 1) S.lost += ev.seq - S.lastSeq - 1;
    if (ev.seq > S.lastSeq) S.lastSeq = ev.seq;
  }
  for (const ev of batch) S.events.push(ev);
  if (S.events.length > MAX_EVENTS) S.events.splice(0, S.events.length - MAX_EVENTS);
  if (batch.length > MAX_EVENTS) batch = batch.slice(-MAX_EVENTS);
  ladderAppend(batch);
  logAppend(batch);
  if (S.lost) setStream(true);
  if (batch.some((e) => e.kind !== "message" || e.handler === "unhandled")) soon();
}

// clearEvents empties the sequence view and the log.
function clearEvents() {
  S.events = [];
  renderLadder();
  renderLog();
}

// connect opens the event stream; a new start time in hello means gnodeb restarted.
function connect() {
  const es = new EventSource("/api/v1/events?tail=" + MAX_EVENTS);
  es.addEventListener("hello", (m) => {
    const hello = JSON.parse(m.data);
    if (S.started && S.started !== hello.started) {
      // gnodeb was restarted: its event numbers start again.
      S.events = [];
      S.lastSeq = 0;
      S.lost = 0;
      pending = [];
      renderLadder();
      renderLog();
      loadStatic().catch(() => {});
      notice("gnodeb restarted", "info");
    } else if (!S.started) {
      S.events = [];
      S.lastSeq = 0;
    }
    S.started = hello.started;
    setStream(true);
    if ($("notice").textContent.startsWith("Event stream")) notice("");
  });
  es.onmessage = (m) => queueEvent(JSON.parse(m.data));
  es.onerror = () => {
    setStream(false);
    notice("Event stream lost. Reconnecting…");
  };
}

// ---------- refresh ----------

let refreshTimer = null;

// soon refreshes the snapshots shortly, once for a burst of events.
function soon() {
  clearTimeout(refreshTimer);
  refreshTimer = setTimeout(refresh, 250);
}

// refresh reads the status again and redraws the views that depend on it.
async function refresh() {
  try {
    S.status = await api.get("/status");
  } catch (e) {
    notice("Cannot reach gnodeb: " + e.message);
    return;
  }
  if ($("notice").textContent.startsWith("Cannot reach gnodeb")) notice("");
  renderSide();
  if (S.page === "dashboard") renderDashboard();
  if (S.page === "license" && !$("page-license").querySelector("small")) renderLicense(); // the version arrives with the status
  if (S.page === "ues") await loadUEs();
  if (S.target) await loadDetail();
  if (S.page === "scenario") {
    try {
      S.scenario.vars = await api.get("/vars");
    } catch (e) {
      /* the next refresh tries again */
    }
    renderScenario();
  }
}

// loadStatic loads the commands, signals and scenario and builds the pages that use them.
async function loadStatic() {
  const [commands, signals, scenario] = await Promise.all([api.get("/commands"), api.get("/signals"), api.get("/scenario")]);
  S.commands = commands;
  S.signals = signals;
  S.scenario = scenario;
  renderOps();
  buildUEsPage();
  buildSequencePage();
  buildLogsPage();
  renderScenario();
}

// init wires up the navigation and the command bar, then loads the data, opens the event stream and polls the status.
function init() {
  for (const b of document.querySelectorAll("nav button")) b.addEventListener("click", () => showPage(b.dataset.page));
  $("linkup").addEventListener("click", () => run("linkup"));
  $("linkdown").addEventListener("click", () => {
    if (confirm("Disconnecting SCTP puts every subscriber back to the initial state. Disconnect?")) run("linkdown");
  });
  $("cmdToggle").addEventListener("click", () => ($("cmdout").hidden = !$("cmdout").hidden));
  // Command bar: Enter runs the line; Up and Down walk the history.
  const input = $("cmdline");
  $("cmdform").addEventListener("submit", (e) => {
    e.preventDefault();
    const line = input.value.trim();
    if (!line) return;
    if (S.history[S.history.length - 1] !== line) S.history.push(line);
    S.histPos = S.history.length;
    input.value = "";
    run(line);
  });
  input.addEventListener("keydown", (e) => {
    if (e.key !== "ArrowUp" && e.key !== "ArrowDown") return;
    e.preventDefault();
    S.histPos = Math.max(0, Math.min(S.history.length, S.histPos + (e.key === "ArrowUp" ? -1 : 1)));
    input.value = S.history[S.histPos] || "";
  });
  renderTarget();
  showPage(location.hash.slice(1));
  loadStatic()
    .catch((e) => notice("Cannot read the settings from gnodeb: " + e.message))
    .then(refresh)
    .then(connect);
  setInterval(() => {
    if (!document.hidden) refresh();
  }, 2000);
}

if (typeof document !== "undefined") init();
