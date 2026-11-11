# kronos-5g-emulator

Go で書いた 5G の gNodeB 疑似呼です。`gnodeb` は 5G コアに N2（NGAP/SCTP）と N3（GTP-U）でつながり、シナリオに沿って多数の UE（IMSI が連番の加入者）を動かします。操作は、端末（CUI）と、同じバイナリに入っている Web コンソールのどちらからでもできます。メッセージと画面の表示は英語です。

[English](README.md)

## 機能

- NG Setup、Registration（5G AKA、NAS セキュリティ）、PDU セッションの確立と解放、UE Context Release、Service Request、Paging、Deregistration
- キャプチャしたテンプレート（`signal.txt`）から NGAP メッセージを作り、動的な項目（ID、鍵、PLMN、TEID など）を自動で補正する
- シナリオ（`scenario.txt`）: 状態、entry、受信したメッセージごとのハンドラ
- GTP-U の U-plane: `ping` と、DN の iperf3 サーバと試験する iperf3 互換の UDP クライアント（`iperf`）
- UE ごとの TUN デバイス（`ue.tun`。任意）。ホストの `iperf3` や `curl` などが、カーネルの TCP/IP で PDU セッションを使える。これを通る内蔵の iperf3 TCP クライアント `iperftcp` もある（シナリオからも使える）
- Web コンソール: ダッシュボード、加入者、信号のボタン、デコード付きのシーケンス図、ログ、信号の編集（`signal.txt` の信号をデコードして値と構造を変え、更新や新しい名前での追加をする。`reload` でファイルの信号を送るようになる）
- バッチ（`@ファイル`、`--exec`）と、全メッセージの JSONL トレース

## 必要なもの

- SCTP をカーネルで使える Linux（`sctp` モジュール。カーネル 3.17 以降）
- ビルドには Go 1.22 以降（Ubuntu 24.04 の `golang-go` パッケージで足りる）

## ビルド

```sh
make build    # ./gnodeb（静的リンクのバイナリ）
make test
```

Makefile はオフラインでビルドします（`GOPROXY=off`）。Go のモジュールキャッシュに依存モジュールがまだ無い環境では、最初に一度だけ次を実行します。

```sh
make GOPROXY=https://proxy.golang.org,direct
```

## 使い方

`config/config.yaml`（AMF・UPF のアドレス、PLMN、加入者の鍵など）を編集します。そのあと、ファイルを検証してから起動します。

```sh
./gnodeb check -c config/config.yaml   # config.yaml・signal.txt・scenario.txt を検証する
./gnodeb -c config/config.yaml         # 起動して、gnodeb> にコマンドを入力する
```

コマンドの例:

```
gnodeb> linkup                                   # AMF に接続する。以降はシナリオが動く
gnodeb> ping(10.10.10.20, 5)
gnodeb> iperf(10.10.10.20, 5201, 100M, 10s, dl)
gnodeb> target(208930000000001).send(AN_UEContextReleaseRequest)
gnodeb> help
```

- **Web コンソール:** config の `gnodeb.web.src` で待ち受けます。
- **ほかの起動のしかた:** `--no-cui` は CUI なしで常駐し、`--exec ファイル` はバッチを実行して終了します。
- **出力先:** ログとトレースは、カレントディレクトリの `log/` に出ます。

### TUN デバイス（`ue.tun`）

config の `ue:` に `tun: true` を書くと、PDU セッションを持つ UE ごとに、UE のアドレスを持つ TUN デバイス（先頭の加入者は `kronos-ue1`）をホストに作ります。あわせて、そのアドレスから送る通信をデバイスに流す送信元ルールも作ります。作成には `CAP_NET_ADMIN` が要ります。

```sh
make setcap     # ビルドしてから sudo setcap cap_net_admin+ep ./gnodeb を実行する（ビルドのたびに必要）
```

セッションがある間に、UE のアドレスに bind してプログラムを動かします。

```sh
iperf3 -c 10.10.10.20 -B 10.60.0.1 -t 10        # TCP の上り
iperf3 -c 10.10.10.20 -B 10.60.0.1 -t 10 -R     # TCP の下り
```

コマンド `iperftcp` を使うと、同じ TCP の試験を gnodeb から実行できます。シナリオからも使えます。

```
gnodeb> iperftcp(10.10.10.20, 5201, 0, 10s, ul)  # 帯域 0 は TCP の流せるだけ
```

デバイスは、セッションがなくなったとき（解放や登録解除）、`linkdown`、gnodeb の終了で消えます。

## ライセンス

Copyright 2026 Kronos Systems。Apache License 2.0 で公開しています。[LICENSE](LICENSE) と [NOTICE](NOTICE) を参照してください。
