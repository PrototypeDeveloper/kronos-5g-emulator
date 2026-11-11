# kronos-5g-emulator

A 5G gNodeB pseudo-call (emulator). It connects to a 5G core over N2 (NGAP/SCTP) and N3 (GTP-U) and drives subscribers (UEs) through a scenario.

[日本語](README_ja.md) | Manual: [docs/manual_en.md](docs/manual_en.md)

## Build

Ubuntu 24.04 or later is assumed (Go 1.22 or later is needed).

```sh
sudo apt update
sudo apt install -y git make golang-go

git clone <URL of this repository> kronos-5g-emulator
cd kronos-5g-emulator
make build GOPROXY=https://proxy.golang.org,direct
```

- The build makes only `./gnodeb` at the repository root (one static binary, the web console included).
- The first build sets `GOPROXY` to fetch the dependencies; after that, `make build` alone works (the Makefile builds offline).
- On a distribution whose `golang-go` is older than 1.22, use the Go of [go.dev](https://go.dev/dl/).

## Start

1. Make SCTP available in the kernel (once).

   ```sh
   sudo modprobe sctp
   ```

   If the module is not found, run `sudo apt install -y linux-modules-extra-$(uname -r)` and try again.

2. Edit `config/config.yaml` for your network: the AMF address (`gnodeb.ngap.dst`), the local addresses (`gnodeb.ngap.src`, `gnodeb.gtpu.src`), the PLMN, the IMSI and keys of the subscribers, …

3. Check the files, then start.

   ```sh
   ./gnodeb check -c config/config.yaml
   ./gnodeb -c config/config.yaml
   ```

4. At the `gnodeb>` prompt, `linkup` connects to the AMF and the scenario starts. `help` lists the commands; Ctrl-D quits.

With `gnodeb.web.src` set, the web console is at `http://<address of gnodeb.web.src>:<port>/` as well.

## License

Apache License 2.0 ([LICENSE](LICENSE), [NOTICE](NOTICE))
