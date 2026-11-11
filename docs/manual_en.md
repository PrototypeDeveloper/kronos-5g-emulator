# kronos-5g-emulator manual

[日本語](manual_ja.md)

How to use the gNodeB pseudo-call `gnodeb`. For the build and the first start, see [README.md](../README.md).

## Contents

1. [Overview](#1-overview)
2. [Starting and options](#2-starting-and-options)
3. [config.yaml](#3-configyaml)
4. [signal.txt (templates of the signals to send)](#4-signaltxt-templates-of-the-signals-to-send)
5. [scenario.txt (the scenario)](#5-scenariotxt-the-scenario)
6. [Commands](#6-commands)
7. [CUI and batches](#7-cui-and-batches)
8. [Web console](#8-web-console)
9. [User plane (ping, iperf, TUN)](#9-user-plane-ping-iperf-tun)
10. [Mapping (auto-fix) and turning it on/off](#10-mapping-auto-fix-and-turning-it-onoff)
11. [Communication between pseudo-calls (sync_send / sync_receive)](#11-communication-between-pseudo-calls-sync_send--sync_receive)
12. [Logs and traces](#12-logs-and-traces)

## 1. Overview

`gnodeb` is a test tool that acts as a 5G gNodeB: it connects to a 5G core (AMF, UPF) and drives many subscribers (UEs). It does not simulate the UEs as devices; the gNodeB itself handles the NAS of each subscriber (registration, authentication, security, PDU sessions).

It reads three files in `config/`.

| File | Contents |
|---|---|
| `config.yaml` | the local and 5G core addresses, the PLMN, the subscribers (IMSI, keys), the PDU sessions … |
| `signal.txt` | templates of the signals to send (NGAP PDUs in hex). IDs, keys and other dynamic fields are fixed when they are sent |
| `scenario.txt` | what to do when a signal comes (states and handlers), run per subscriber |

It is driven from the terminal (CUI, the `gnodeb>` prompt) or from the web console built into the same binary. Messages and screens are in English.

The basic flow:

1. `linkup` connects to the AMF over SCTP. The `entry` of the `node` block runs and sends NG Setup Request.
2. When NG Setup succeeds, every subscriber starts, in IMSI order, and the `entry` of the `start` state runs (e.g. it sends Registration Request).
3. From then on, the scenario's handlers run for each signal received. Commands on the CUI or the web console send signals, change states and more.

## 2. Starting and options

```
gnodeb [-c config.yaml] [--no-cui] [--no-web]   start the pseudo-call (linkup connects N2)
gnodeb [-c config.yaml] --exec FILE             run the commands of FILE in order and exit (batch)
gnodeb check [-c config.yaml] [-v]              check config, signal and scenario
gnodeb version                                  print the version
```

| Option | Meaning |
|---|---|
| `-c FILE` | path of config.yaml (default `config/config.yaml`); scenario.txt and signal.txt are read from the same directory |
| `-v` | with `check`, also print what was understood from the signals and the scenario |
| `--no-cui` | run without the CUI (stop with Ctrl-C or SIGTERM); without a web console, linkup runs at once |
| `--no-web` | do not open the web console |
| `--exec FILE` | batch: run one line at a time, stop at a failing line (exit code 1) |

**Privileges:** the TUN device of each UE (chapter 9, for TCP iperf and more) needs `CAP_NET_ADMIN`. Run `./gnodeb` after `make setcap` (it builds, then gives the capability; needed after each build), or start it with `sudo ./gnodeb`. Without the capability it still starts, but makes no TUN devices and says so at the start. Started with sudo, it writes the files of `log/` as root: mind that when you start it again without sudo.

`gnodeb check` reports errors (it cannot start), warnings and notes with their line and column. Run it after each change to the scenario or the signals. The start runs the same checks and does not start on an error.

## 3. config.yaml

Quote values made of digits only (IMSI, MNC, SD, TAC …) so that other YAML tools do not turn them into numbers (e.g. `sd: "010203"`; `gnodeb` itself reads them as strings).

### base

| Key | Meaning | Default |
|---|---|---|
| `snap` | `snap()` runs only when `true` (set `false` in load tests to save the output) | false |
| `debug` | debug log level | false |
| `scenario` / `signal` | file names of the scenario and the signals (relative to config.yaml) | `scenario.txt` / `signal.txt` |
| `role` | `gnb` only (the AMF role is not implemented) | gnb |

### ue (the subscribers)

| Key | Meaning |
|---|---|
| `imsi` | IMSI of the first subscriber (15 digits) |
| `subscription` | number of subscribers; their IMSIs count up from `imsi` (default 1) |
| `imeisv` | IMEISV (16 digits, optional), for Security Mode Complete and Identity Response |
| `suci.protectionscheme` | SUCI protection scheme: `0` the null scheme (the MSIN in clear), `1` Profile A (X25519), `2` Profile B (secp256r1). With Profile A/B each Registration encrypts the MSIN with a new ephemeral key |
| `suci.keyid` | the home network public key identifier (0–255; Profile A/B) |
| `suci.key` | the home network public key in hex: 32 octets (64 digits) for Profile A, 33 compressed (66 digits) or 65 uncompressed octets (130 digits) for Profile B. It must pair with the private key of the 5G core (UDM) |
| `suci.routingindicator` | routing indicator (up to 4 digits) |
| `credential.key` / `credential.op` | key K and OP or OPc (32 hex digits) |
| `credential.type` | `OP` or `OPC` |
| `credential.amf` | the AMF field of Milenage (4 hex digits, with the top bit, the separation bit, set) |
| `nssai` | Requested NSSAI (a list of `sst`, `sd`; optional) |
| `session` | the PDU sessions: `id` (1–15), `type` (`IPv4` / `IPv6` / `IPv4v6`), `dnn`, `snssai` |
| `integrity` / `ciphering` | NIA0–3 / NEA0–3 the UE announces (default 0–2); 3 (ZUC) is not implemented |
`ue.type` and `ue.apn` are the old form, read as the session with ID 1. The former `ue.tun` is no longer needed (it is ignored, with a note).

### gnodeb (this gNB)

| Key | Meaning |
|---|---|
| `name` | RANNodeName of NG Setup (none keeps the template's) |
| `gnodebid` / `gnodebidlength` | gNB ID (hex) and its length (22–32 bits, default 32) |
| `nci` | NR cell identity (9 hex digits; default: the gNB ID followed by zero bits) |
| `plmn.mcc` / `plmn.mnc` | the PLMN |
| `snssai` | the slices announced in NG Setup (a list of `sst`, `sd`) |
| `tac` | the TAC (decimal; hex with `0x`) |
| `ngap.src` / `ngap.dst` | N2: the local end (optional; `port: 0` for any port) and the AMF's address and port (38412) |
| `gtpu.src` | N3: the local address and port (2152) |
| `gtpu.dst` | when set, used instead of the UPF address the 5G core names (optional) |
| `web.src` | where the web console listens (none: no web console) |
| `http.src` | where the pseudo-call link listens and sends from (chapter 11; none: no link) |

## 4. signal.txt (templates of the signals to send)

A signal is a name and an NGAP PDU in hex. The name is free; `send(name)` uses it. `#` starts a comment up to the end of the line.

```
NGSetupRequest {
  00150035000004001b00090002f839500000000100524008028035672d72616e...
}

InitialUEMeesageRegistrationRequest [nas=keep] {
  000f40480000050055000200010026001a197e004179000d0102f839...
}
```

The `[...]` after the name is the send mode.

| Mode | Meaning |
|---|---|
| (none) `auto` | fix the dynamic fields of NGAP and NAS, protect the NAS, send |
| `nas=keep` | fix NGAP but keep the NAS body as written (it is still protected) |
| `raw` | send the template's octets unchanged |

- Templates come from captures of real equipment (an NGAP PDU copied from Wireshark) or from the Signals page of the web console.
- A template with protected NAS must have a readable body (plain or NEA0).
- UE-associated messages (with NGAP UE IDs) are sent by the subscribers, the others by the node.
- Chapter 10 lists what is fixed. Everything else is sent as the template has it.

## 5. scenario.txt (the scenario)

### 5.1 Example

```
var {
	public  count = 0       # one for the whole pseudo-call
	private t     = 5s      # one per subscriber
}

node {
	entry {
		send(NGSetupRequest)                # when SCTP connects
	}
	receive(NGSetupResponse) {
		snap("NG Setup OK")
	}
}

any {                                       # in every state
	receive(UEContextReleaseCommand) {
		send(UEContextReleaseComplete)
	}
}

start {                                     # the first state of a subscriber
	entry {
		control(10)                         # at most 10 subscribers a second, all together
		send(InitialUEMeesageRegistrationRequest)
	}
	receive(DownLinkNasTransport:AuthenticationRequest) {
		send(UpLinkNASTransportAuthenticationResponse)
	}
	receive(DownLinkNasTransport:SecurityModeCommand) {
		send(UpLinkNASTransportSecurityModeComplete)
	}
	receive(InitialContextSetupRequest:RegistrationAccept) {
		send(InitialContextSetupResponse)
		send(UplinkNASTransportRegistrationComplete)
		set(count = count + 1)
		move(sessionestablishment)
	}
}

sessionestablishment {
	entry {
		send(UplinkNASTransportPDUSessionEstablishmentRequest)
		timer(t)
	}
	receive(PDUSessionResourceSetupRequest:PDUSessionEstablishmentAccept) {
		stoptimer(t)
		send(PDUSessionResourceSetupResponse)
	}
	expire(t) {
		snap("no answer in 5 s (%s)", $imsi)
	}
}
```

### 5.2 Blocks

| Block | Meaning |
|---|---|
| `var { … }` | variable declarations (5.5) |
| `node { … }` | handlers of the node (the gNB itself): signals that are not UE-associated (NG Setup answers, NG Reset …) |
| `any { … }` | handlers of a subscriber valid in every state, looked up when the state block has none |
| `state { … }` | a state of a subscriber; `start` is the first one. Any name |

What a block holds:

| Item | When it runs |
|---|---|
| `entry { … }` | `node`: when SCTP connects. A state: when the subscriber enters it (`start`: after NG Setup succeeds). Not in `any` |
| `receive(NGAP) { … }` / `receive(NGAP:NAS) { … }` | when that signal comes. NAS is the name of the NAS message the NGAP message carries (for 5GSM, the inner name matches too) |
| `expire(var) { … }` | when the timer started by `timer(var)` runs out |
| `sync_receive(SYNC\|HO) { … }` | when another pseudo-call sends a message (chapter 11) |

### 5.3 How handlers are found

- A subscriber: the block of its current state, then `any`. Within a block, an `NGAP:NAS` match comes before an `NGAP` one.
- The node: the `node` block only. Paging goes to the subscriber with its 5G-S-TMSI, and to `node` when there is none.
- With no handler, an "unhandled" warning event is given (the built-in processing has been done).
- Names ignore case (`DownLinkNasTransport` matches `DownlinkNASTransport`).
- The built-in processing (authentication, security, taking session information …) is done before the handler.

### 5.4 Statements

Handlers and entries hold commands (chapter 6) and `if`.

```
if ($ue == 1 && count > 10) {
	send(X)
} else if (mode == 2) {
	snap("mode 2")
} else {
	set(mode = 1)
}
```

- `move(state)` changes the state at once; the `entry` of the new state runs when the handler is done (also for a move to the same state). Moves from entries chain up to 16 times.
- `#` and `//` start a comment up to the end of the line (also after a statement; any language).
- Reserved words (cannot name a state): `var` `public` `private` `node` `any` `receive` `entry` `expire` `sync_receive` `if` `else` `timeout` `target`

### 5.5 Variables and expressions

```
var {
	public  a = 1         # int
	public  rate = 0.5    # float
	private s = "text"    # string
	private t = 5s        # duration (100ms, 1m30s …)
	private n = -1
}
```

- A `public` variable is one value for the whole pseudo-call, seen by every subscriber and the node. A `private` one has a value per subscriber; the node cannot use it.
- `private` variables go back to their initial values when SCTP goes down (linkdown).
- The type comes from the initial value and never changes. `set(var = expression)` assigns; an int value fits a float variable.
- Expressions have the operators of C with their precedence: `||` `&&` `==` `!=` `<` `<=` `>` `>=` `+` `-` `*` `/` `%`, unary `-` `!`, and parentheses.
  - int with int is int (`/` truncates); mixing int and float gives float.
  - `+` with a string on one side joins the text of both (`"ue" + $ue` → `ue1`).
  - durations add and subtract, and multiply or divide by a number.
  - a comparison gives true / false, for `if` and `&&` `||` `!` (variables cannot hold it).
  - a division by 0 is an error when it runs.
- Builtins:

| Name | Value |
|---|---|
| `$ue` | the number of the subscriber (from 1) |
| `$imsi` | the IMSI of the subscriber |
| `$state` | the current state of the subscriber |
| `$msg` | the last trigger received (`NGAP:NAS`) |
| `$time` | the current time |
| `$ngap.Name` / `$nas.Name` | an IE of the last signal received ([5.8](#58-ies-of-the-received-signal-nas--ngap)) |

When a command or a condition fails at run time, an error event is given, that statement is skipped and the next one runs.

### 5.6 Timers

A duration variable is a timer.

- `timer(t)` starts it; when it runs out, `expire(t) { … }` runs, looked up in the block of the state at that moment, then in `any`.
- `timer(t)` on a running timer starts it again from the beginning.
- `stoptimer(t)` stops it. A state change does not stop it; linkdown does.
- A `private` variable is a timer of the subscriber, a `public` one a timer of the node (`expire` in `node`).

### 5.7 Pacing (control)

`control(n)` lets at most n passes a second through that statement, counted over all subscribers.

- Turns are 1/n second apart, and a subscriber waits for its turn (`control(5)`: one every 0.2 s).
- In a scenario that loops through a procedure, put it before the `send` to set the load (calls per second).
- `control(0)` means no limit. n can be a variable or an expression (`control(rate)`, changed with `set(rate = 20)`).
- While it waits, the subscriber's received messages wait too: put it at the head of a procedure.
- linkdown cuts the wait short, and the rest of that body does not run.

### 5.8 IEs of the received signal ($nas / $ngap)

`$ngap.Name` and `$nas.Name` put the value of an IE of the last signal received into an expression: branch on what the core returned, or keep it in a variable.

```
receive(DownlinkNASTransport:RegistrationReject) {
	if ($nas.Cause == 3) {                     # 5GMM cause #3 Illegal UE
		snap("illegal ue, T3502 = %d", $nas.T3502Value)
	}
}
receive(UEContextReleaseCommand) {
	if ($ngap.Cause == "Nas" && $ngap.Cause.Nas == "deregister") {
		move(start)
	}
}
receive(PDUSessionResourceSetupRequest:PDUSessionEstablishmentAccept) {
	set(upf = $ngap.TransportLayerAddress)   # the UPF address ("10.10.30.40", ...)
	set(qfi = $ngap.QosFlowIdentifier)
}
```

- The names are those of the IEs in the signal editor of the web console (Signals). Case does not matter.
- A name after `.` is looked for inside the IE before it. The first name is looked for in the whole signal (its own IEs first, then deeper). The names in between can be left out, so `$ngap.TransportLayerAddress` names a deep IE directly. When several IEs have the name, the first one found is taken.
- `[index]` (from 0) picks a list item: `$ngap.PDUSessionResourceSetupListSUReq[1].PDUSessionID`.
- `$ngap` also looks inside the transfers of PDU sessions (PDUSessionResourceSetupRequestTransfer, ...).
- `$nas` reads the NAS messages of the signal in plain. When a UL/DL NAS Transport carries a 5GSM message, that 5GSM message is looked in first, then the IEs of the transport.
- The values:

| IE | value |
|---|---|
| an integer | int |
| an NGAP ENUMERATED | its name as a string (`"deregister"`, ...; the name in Signals) |
| a CHOICE | the name of the alternative it took (`$ngap.Cause` → `"RadioNetwork"`, ...); the value inside is `$ngap.Cause.RadioNetwork` |
| octets | hex as a string (a PLMN is `"02f839"`, a TEID `"00001001"`) |
| a TransportLayerAddress | the IP address as a string |
| another BIT STRING | an int up to 63 bits, longer ones hex as a string |

- A group of IEs (a GUAMI, ...) or a list itself has no value: name an IE inside it, or an item with `[index]`.
- An IE the signal does not have (an optional IE that is absent, an alternative of a CHOICE not taken, an index past the end) has no value: comparisons such as `==` and `<` are false whatever the other side (only `!=` is true), and arithmetic or `set` with it is an error at run time. snap shows it as `<nil>`.
- It is the last signal that subscriber (the node, in a `node` handler) received. In an `entry` or an `expire`, it is the last one received by then. Commands of the CUI/GUI cannot use it.

## 6. Commands

| Command | Where | What it does |
|---|---|---|
| `send(signal)` | scenario, CUI/GUI | send a signal of signal.txt with its fields fixed: UE signals from the subscriber, node signals from the node |
| `move(state)` | scenario, CUI/GUI | change the state and run its entry (forced from the CUI/GUI) |
| `snap(format, value…)` | scenario, CUI/GUI | print Printf style (only with `base.snap: true`), e.g. `snap("%s %d", $imsi, a)` |
| `set(var = expression)` | scenario, CUI/GUI | assign a variable; on the CUI/GUI a public one is set once, a private one per subscriber |
| `timer(var)` / `stoptimer(var)` | scenario, CUI/GUI | start / stop a timer |
| `control(n)` | scenario | at most n a second (5.7) |
| `ping(dst[, count[, interval[, size]]])` | scenario, CUI/GUI | ICMP echo through the PDU session (count 0: until stop) |
| `iperf(dst[, port[, rate[, time[, ul\|dl[, len]]]]])` | scenario, CUI/GUI | a UDP test with an iperf3 server on the DN |
| `iperftcp(…)` | scenario, CUI/GUI | the same with TCP (through TUN, chapter 9) |
| `stop([ping\|iperf])` | scenario, CUI/GUI | stop the running ping / iperf / iperftcp |
| `mapon([item])` / `mapoff([item])` / `mapget([item])` | scenario, CUI/GUI | turn the mapping on / off, show it (chapter 10) |
| `sync_send(SYNC\|HO, dst)` | scenario, CUI/GUI | send a message to another pseudo-call (chapter 11) |
| `linkup` / `linkdown` | CUI/GUI | connect / disconnect SCTP; on a disconnect the subscribers go back to their initial state |
| `reload` | CUI/GUI | read signal.txt again (nothing changes when it has an error) |
| `wait(time)` | CUI/GUI | wait (for batches; up to 1 hour) |
| `help([command])` | CUI/GUI | the commands and how to use them |

- **Target:** on the CUI/GUI, `target(IMSI).` before a command for subscribers runs it on that subscriber only (e.g. `target(208930000000001).send(AN_UEContextReleaseRequest)`); without it, on all subscribers.
- **Short forms:** a command without arguments may omit the parentheses (`linkup`). IPv4 addresses and `address:port` need no quotes (`ping(10.0.0.1, 3)`).
- **Rates:** written `100M`, `1G`, `500k`; a fraction goes in quotes (`"1.5G"`).

## 7. CUI and batches

- Type commands at the `gnodeb>` prompt; Ctrl-D quits. What is received and state changes show up as events.
- `@file` runs the commands of a file one line at a time (a batch) and stops at a failing line. Starting with `--exec file` runs a batch and exits (0 on success, 1 on a failure).
- In a batch, `wait(3s)` waits for answers. `#` and `//` start comments, and `@` can nest.

```
# registration → ping → release → Service Request → ping → deregistration
linkup
wait(3s)
ping(10.0.0.1, 3)
wait(4s)
target(208930000000001).send(AN_UEContextReleaseRequest)
wait(1s)
target(208930000000001).send(InitialUEMeesageServiceRequest)
wait(2s)
ping(10.0.0.1, 3)
wait(4s)
target(208930000000001).send(InitialUEMeesageDeregistrationRequest)
wait(1s)
```

## 8. Web console

With `gnodeb.web.src` set, open `http://<address>:<port>/` in a browser. It has no authentication or TLS: use it on a trusted network.

| Page | Contents |
|---|---|
| Dashboard | SCTP and NG Setup, the subscribers (registered, connected), counts of warnings … |
| Subscribers | the subscribers and their details (state, NAS COUNTs, PDU sessions, variables); the one picked is the target of commands |
| Operations | buttons for the signals, forms for the commands, batches |
| Signals | decode and edit the signals of signal.txt: change values, add and remove optional IEs and list items, add a copy under a new name, create a new signal from an NGAP message type. Fixed fields are marked "auto" with the item that turns their fixing on and off (`[mapping: plmn]` …). Saving rewrites signal.txt; `reload` sends the new signals |
| Sequence | the sequence diagram of what was sent and received; a message shows its decoded contents |
| Scenario | the variables and the blocks and handlers (with subscriber counts) |
| Logs | the events |
| License | the license |

The command line at the bottom takes the same commands as the CUI.

## 9. User plane (ping, iperf, TUN)

Once a PDU session is up, gnodeb uses the GTP-U tunnel with the UE's IP address.

- `ping(dst, count, interval, size)` sends ICMP echo requests and reports the loss and the RTT (min/avg/max/mdev). It also answers pings from the DN to the UE address.
- `iperf(dst, port, rate, time, ul|dl, len)` runs a UDP test with `iperf3 -s` on the DN (iperf3 compatible; `dl` is iperf3's `-R`).

### TUN devices

When gnodeb has `CAP_NET_ADMIN` (`./gnodeb` after `make setcap`, or `sudo ./gnodeb`; chapter 2), each UE with a PDU session gets a TUN device on the host holding its address (`kronos-ue1` for the first subscriber) and a source rule that sends the traffic from that address into the device. Nothing needs to be configured. Host programs such as `iperf3` and `curl` then use the PDU session through the kernel's TCP/IP.

```sh
make setcap     # build, then sudo setcap cap_net_admin+ep ./gnodeb (setcap is in libcap2-bin)
```

```sh
iperf3 -c 10.0.0.1 -B 10.60.0.1 -t 10        # bound to the UE address: TCP uplink
iperf3 -c 10.0.0.1 -B 10.60.0.1 -t 10 -R     # downlink
```

`iperftcp(dst, port, rate, time, ul|dl, len)` runs the same TCP test from inside gnodeb (from a scenario too; rate 0: as fast as TCP goes). The devices go away with their sessions, on linkdown and at exit.

## 10. Mapping (auto-fix) and turning it on/off

When a signal is sent, the dynamic fields of the template are fixed from the config and the state of the subscriber (the mapping). Each of these items can be turned on and off; all are on at the start.

| Item | What it fixes |
|---|---|
| `plmn` | the PLMN (gNB ID and TA of NG Setup / RAN Configuration Update, the ULI, the cell of Cell Traffic Trace) |
| `tac` | the TAC (TA list, ULI) |
| `nci` | the NR cell identity (ULI, Cell Traffic Trace) |
| `gnodebid` | the gNB ID |
| `ranname` | RANNodeName (`gnodeb.name`) |
| `snssai` | S-NSSAIs (slices of NG Setup, Requested NSSAI, ULNASTransport; NSSAA Complete takes the S-NSSAI of the command received) |
| `timestamp` | the TimeStamp of the ULI (the time of sending) |
| `ngapid` | RAN-UE-NGAP-ID and AMF-UE-NGAP-ID (also the source AMF ID of Path Switch and the lists of NG Reset) |
| `identity` | SUCI, 5G-GUTI, 5G-S-TMSI, and Identity Response (the type asked for: SUCI, 5G-GUTI, IMEI, 5G-S-TMSI, IMEISV) |
| `ngksi` | ngKSI (Registration, Service, Control Plane Service and De-registration Request) |
| `res` | RES* of Authentication Response, the AUTS of Authentication Failure (synch failure) |
| `seccap` | UE security capability (also in Path Switch Request) |
| `imeisv` | the IMEISV of Security Mode Complete |
| `session` | PDU session ID, type, DNN and PTI (ULNASTransport and 5GSM; a request gets a new PTI, Release Complete, Modification Complete, Modification Command Reject and PDU Session Authentication Complete the PTI of the command received) |
| `status` | PDU session status, uplink data status, allowed PDU session status (Registration, Service and Control Plane Service Request, Notification Response) |
| `tunnel` | the PDU session lists of NGAP (answers of setup, release, modify, notify, handover and path switch: DL TEID and address, QoS flows) |
| `container` | the NAS message container (rebuilt and ciphered) |

```
gnodeb> mapoff(plmn)                             # the node and every subscriber send the template's PLMN
gnodeb> target(208930000000001).mapoff(ngapid)   # that subscriber only
gnodeb> mapget                                   # the items that are not on (mixed: off for some only)
gnodeb> mapon                                    # every item back on
```

- Without an item, every item. With every item off, the octets sent are the template's (but the NAS protection).
- In a scenario, they apply to the subject of the block (the node in `node`, else that subscriber).
- They stay as set across linkdown.

From the NAS received, the subscriber takes the following (before the handler runs).

| NAS received | What is taken |
|---|---|
| Authentication Request | RES* and the keys (AUTN checked); the SQN for a synch failure |
| Security Mode Command | NAS protection starts; whether the IMEISV is asked for |
| Registration Accept / Configuration Update Command | the 5G-GUTI (paging is matched with it too); for Configuration Update Command, whether an acknowledgement or a registration is asked for is shown |
| Identity Request | the type of identity asked for (put into Identity Response) |
| Network Slice-Specific Authentication Command | the S-NSSAI (put into the complete) |
| PDU Session Establishment Accept / Reject | the UE's IP address and the session; a reject removes a session that never came up |
| PDU Session Release / Modification / Authentication Command | the PTI (put into the complete or command reject) |
| Deregistration Accept / Deregistration Request (from the network) / Registration Reject | de-registered; the PDU sessions are released locally |
| Authentication Reject, Service Reject (#3, #6, #7, #9) | besides, the 5G-GUTI and the NAS security context are deleted (Service Reject #10 only de-registers) |

## 11. Communication between pseudo-calls (sync_send / sync_receive)

Pseudo-calls talk to each other with JSON over HTTP/REST, e.g. to hand a subscriber over from the source of a handover to its target. Both configs need `gnodeb.http.src` (where they listen and send from).

```
# source
connected {
	receive(…) { sync_send(HO, 10.0.0.31:12345) }
}

# target (the subscriber with the same IMSI takes it)
start {
	sync_receive(HO)   { move(handover) }
	sync_receive(SYNC) { snap("sync from the source") }
}
```

| Kind | Contents |
|---|---|
| `SYNC` | a plain synchronization: only the sender and the IMSI |
| `HO` | a handover: the mapping information of the subscriber (registration state, 5G-GUTI, ngKSI, Kamf, NAS security keys and COUNTs, PDU sessions, the source's AMF-UE-NGAP-ID) |

- The subscriber with the same IMSI takes it (a SYNC sent from the node goes to the node), looked up as for receive.
- After a HO, the target's subscriber protects and maps its NAS with that information.
- The HandoverRequest from the target's AMF goes to the subscriber that took the HO (matched by the uplink TEIDs of its PDU sessions); HandoverRequestAcknowledge carries the DL TEIDs allocated. For an Xn handover, send PathSwitchRequest.
- The sender (the source) does nothing more; the scenario decides what follows.

On the wire:

```
POST http://<dst>/kronos/v1/sync
Content-Type: application/json

{"kind": "HO", "from": "10.0.0.30:12345", "imsi": "208930000000001",
 "handover": {"registered": true, "guti": "f202f839…", "ngKSI": 1, "kamf": "…",
   "security": {"integrity": 2, "ciphering": 2, "knasInt": "…", "knasEnc": "…", "uplinkCount": 4, "downlinkCount": 3},
   "sessions": [{"id": 1, "ueIp": "10.60.0.7", "upf": "10.0.0.40", "ulTeid": 1, "qfi": [1]}],
   "amfUeNgapId": 100, "nextPTI": 2}}
```

The answer is `{"accepted": true}` (202) or `{"accepted": false, "error": "…"}` (400: a malformed message, 404: no subscriber with that IMSI, 503: for the node while SCTP is down). There is no authentication or TLS and the keys travel in clear: use it on a test network.

## 12. Logs and traces

Written to `log/` in the current directory.

| File | Contents |
|---|---|
| `log/gnodeb.txt` | the application log |
| `log/trace-YYYYMMDD-HHMMSS.jsonl` | every event (PDUs sent and received in hex, state changes, snap, warnings …), one JSON object per line |
