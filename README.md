# kronos-5g-emulator

A 5G gNodeB emulator written in Go. `gnodeb` connects to a 5G core over N2
(NGAP/SCTP) and N3 (GTP-U) and runs many emulated UEs (subscribers with
consecutive IMSIs) from a scenario. You operate it from a terminal or from a web
console built into the same binary.

[Japanese](README_ja.md)

## Features

- NG Setup, Registration (5G AKA, NAS security), PDU session establishment and
  release, UE context release, Service Request, Paging and Deregistration
- NGAP messages built from captured templates (`signal.txt`): the dynamic
  fields (IDs, keys, PLMN, TEIDs …) are filled in automatically
- Scenarios (`scenario.txt`): states, entry blocks and handlers for received
  messages
- User plane over GTP-U: `ping`, and `iperf`, an iperf3-compatible UDP client
  for an iperf3 server on the data network
- Optional TUN device per UE (`ue.tun`), so host programs such as `iperf3`
  or `curl` can use a PDU session through the kernel's TCP/IP, and
  `iperftcp`, a built-in iperf3 TCP client that runs through it (from a
  scenario too)
- Web console: dashboard, subscribers, signal buttons, a sequence diagram with
  decoded messages, logs, and a signal editor (decode a signal of
  `signal.txt`, change its values and structure, save it or add a new one;
  `reload` sends the file's signals)
- Batches (`@file`, `--exec`) and a JSONL trace of every message

## Requirements

- Linux with SCTP in the kernel (the `sctp` module, kernel 3.17 or later)
- Go 1.22 or later to build (the `golang-go` package of Ubuntu 24.04 is enough)

## Build

```sh
make build    # ./gnodeb, a static binary
make test
```

The Makefile builds offline (`GOPROXY=off`). If the Go module cache on your
machine does not have the dependencies yet, run this once:

```sh
make GOPROXY=https://proxy.golang.org,direct
```

## Run

Edit `config/config.yaml` (AMF and UPF addresses, PLMN, subscriber keys …).
Then check the files and start:

```sh
./gnodeb check -c config/config.yaml   # validate config.yaml, signal.txt and scenario.txt
./gnodeb -c config/config.yaml         # start, then type commands at the gnodeb> prompt
```

Example commands:

```
gnodeb> linkup                                   # connect to the AMF; the scenario runs from there
gnodeb> ping(10.10.10.20, 5)
gnodeb> iperf(10.10.10.20, 5201, 100M, 10s, dl)
gnodeb> target(208930000000001).send(AN_UEContextReleaseRequest)
gnodeb> help
```

- **Web console:** it listens on `gnodeb.web.src` of the config.
- **Other modes:** `--no-cui` runs headless, and `--exec FILE` runs a batch and
  exits.
- **Output:** logs and traces go to `log/` under the current directory.

### TUN devices (`ue.tun`)

With `tun: true` under `ue:` in the config, each UE with a PDU session gets a
TUN device on the host (`kronos-ue1` for the first subscriber) that holds the UE
address, and a source rule that sends what the host sends from that address
through it. Creating them needs `CAP_NET_ADMIN`:

```sh
make setcap     # build, then sudo setcap cap_net_admin+ep ./gnodeb (needed after each build)
```

Then, while the session is up, run programs bound to the UE address:

```sh
iperf3 -c 10.10.10.20 -B 10.60.0.1 -t 10        # TCP, uplink
iperf3 -c 10.10.10.20 -B 10.60.0.1 -t 10 -R     # TCP, downlink
```

The command `iperftcp` runs the same TCP test from gnodeb, so a scenario can
start it too:

```
gnodeb> iperftcp(10.10.10.20, 5201, 0, 10s, ul)  # rate 0: as fast as TCP goes
```

The device goes away with the session (a release or a deregistration), at
`linkdown` and when gnodeb exits.

## License

Copyright 2026 Kronos Systems. Licensed under the Apache License 2.0: see
[LICENSE](LICENSE) and [NOTICE](NOTICE).
