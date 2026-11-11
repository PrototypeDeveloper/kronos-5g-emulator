# kronos-5g-emulator

5G の gNodeB 疑似呼です。5G コアに N2（NGAP/SCTP）と N3（GTP-U）で接続し、シナリオに沿って加入者（UE）を動かします。

[English](README.md) ・ 使い方は [docs/manual_ja.md](docs/manual_ja.md)

## ビルド

Ubuntu 24.04 以降を想定しています（Go 1.22 以降が必要です）。

```sh
sudo apt update
sudo apt install -y git make golang-go

git clone <このリポジトリの URL> kronos-5g-emulator
cd kronos-5g-emulator
make build GOPROXY=https://proxy.golang.org,direct
```

- できるのはリポジトリ直下の `./gnodeb` だけです（静的リンクの 1 ファイル。Web コンソールも入っています）。
- 依存モジュールの取得のため、初回は `GOPROXY` を付けます。2 回目以降は `make build` だけでビルドできます（Makefile はオフラインでビルドします）。
- `golang-go` が 1.22 より古いディストリビューションでは、[go.dev](https://go.dev/dl/) の Go を使ってください。

## 起動

1. SCTP をカーネルで使えるようにします（一度だけ）。

   ```sh
   sudo modprobe sctp
   ```

   モジュールが見つからない場合は、`sudo apt install -y linux-modules-extra-$(uname -r)` を実行してから、もう一度試します。

2. `config/config.yaml` を環境に合わせて編集します（AMF のアドレス `gnodeb.ngap.dst`、自局のアドレス `gnodeb.ngap.src` / `gnodeb.gtpu.src`、PLMN、加入者の IMSI と鍵など）。

3. 設定を検証してから起動します。

   ```sh
   ./gnodeb check -c config/config.yaml
   ./gnodeb -c config/config.yaml
   ```

4. `gnodeb>` のプロンプトで `linkup` と入力すると AMF に接続し、シナリオが動き始めます。`help` でコマンドの一覧、Ctrl-D で終了です。

`gnodeb.web.src` を設定していれば、ブラウザで `http://<gnodeb.web.src のアドレス>:<ポート>/` を開くと Web コンソールからも操作できます。

## ライセンス

Apache License 2.0（[LICENSE](LICENSE)、[NOTICE](NOTICE)）
