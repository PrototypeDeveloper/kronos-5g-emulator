# kronos-5g-emulator 使い方マニュアル

[English](manual_en.md)

gNodeB 疑似呼 `gnodeb` の使い方をまとめます。ビルドと最初の起動は [README_ja.md](../README_ja.md) を見てください。

## 目次

1. [概要](#1-概要)
2. [起動とオプション](#2-起動とオプション)
3. [config.yaml](#3-configyaml)
4. [signal.txt（送る信号のテンプレート）](#4-signaltxt送る信号のテンプレート)
5. [scenario.txt（シナリオ）](#5-scenariotxtシナリオ)
6. [コマンド](#6-コマンド)
7. [CUI とバッチ](#7-cui-とバッチ)
8. [Web コンソール](#8-web-コンソール)
9. [U-plane（ping・iperf・TUN）](#9-u-planepingiperftun)
10. [マッピング（自動補正）と ON/OFF](#10-マッピング自動補正と-onoff)
11. [疑似呼間通信（sync_send / sync_receive）](#11-疑似呼間通信sync_send--sync_receive)
12. [ログとトレース](#12-ログとトレース)
13. [制限事項](#13-制限事項)

## 1. 概要

`gnodeb` は 5G の gNodeB として 5G コア（AMF・UPF）に接続し、多数の加入者（UE）を動かす試験ツールです。UE は端末として模擬せず、gNodeB が加入者ごとの NAS（登録、認証、セキュリティ、PDU セッション）まで処理します。

入力は `config/` にある 3 つのファイルです。

| ファイル | 内容 |
|---|---|
| `config.yaml` | 自局と 5G コアのアドレス、PLMN、加入者（IMSI・鍵）、PDU セッションなど |
| `signal.txt` | 送る信号のテンプレート（NGAP の PDU を 16 進で書いたもの）。送るときに ID や鍵などを自動で補正する |
| `scenario.txt` | 信号を受けたときに何をするか（状態とハンドラ）。加入者ごとに動く |

操作は、端末の CUI（`gnodeb>` のプロンプト）と、同じバイナリに入っている Web コンソールのどちらからでもできます。メッセージと画面の表示は英語です。

基本の流れは次のとおりです。

1. `linkup` で AMF に SCTP で接続します。`node` ブロックの `entry` が動き、NG Setup Request を送ります。
2. NG Setup が成功すると、全加入者が IMSI の順に起動し、`start` 状態の `entry` が動きます（例: Registration Request を送る）。
3. 以後は、受信した信号ごとにシナリオのハンドラが動きます。CUI や Web コンソールからコマンドを打って、信号を送ったり状態を変えたりもできます。

## 2. 起動とオプション

```
gnodeb [-c config.yaml] [--no-cui] [--no-web]   疑似呼を起動する（linkup で N2 に接続する）
gnodeb [-c config.yaml] --exec FILE             FILE のコマンドを順に実行して終了する（バッチ）
gnodeb check [-c config.yaml] [-v]              config・signal・scenario を検証する
gnodeb version                                  版を表示する
```

| オプション | 意味 |
|---|---|
| `-c FILE` | config.yaml のパス（既定は `config/config.yaml`）。scenario.txt と signal.txt は同じディレクトリから読む |
| `-v` | `check` で、信号とシナリオをどう解釈したかも表示する |
| `--no-cui` | CUI なしで常駐する（Ctrl-C か SIGTERM で終了）。Web コンソールが無いときは、起動と同時に linkup する |
| `--no-web` | Web コンソールを開かない |
| `--exec FILE` | バッチ。1 行ずつ実行し、失敗した行で止まる（終了コード 1） |

`gnodeb check` は、エラー（起動できない）・警告・情報を行と列つきで表示します。シナリオや信号を書き換えたら、まず `check` を通してください。起動時にも同じ検証をし、エラーがあれば起動しません。

## 3. config.yaml

数字だけの値（IMSI、MNC、SD、TAC など）は、ほかの YAML ツールで数値に変わらないよう、引用符で囲むことをお勧めします（例 `sd: "010203"`。`gnodeb` 自身は文字列として読みます）。

### base

| キー | 意味 | 既定 |
|---|---|---|
| `snap` | `true` のときだけ `snap()` を実行する（負荷試験では `false` にして表示の負荷を無くす） | false |
| `debug` | ログをデバッグの詳しさにする | false |
| `scenario` / `signal` | シナリオと信号のファイル名（config.yaml からの相対パス） | `scenario.txt` / `signal.txt` |
| `role` | `gnb` だけ（AMF の役は未実装） | gnb |

### ue（加入者）

| キー | 意味 |
|---|---|
| `imsi` | 先頭の加入者の IMSI（15 桁） |
| `subscription` | 加入者の数。IMSI は `imsi` から連番になる（既定 1） |
| `imeisv` | IMEISV（16 桁。任意）。Security Mode Complete と Identity Response に使う |
| `suci.protectionscheme` | SUCI の保護方式。`0`（null スキーム）だけに対応 |
| `suci.routingindicator` | ルーティング指示子（4 桁まで） |
| `credential.key` / `credential.op` | 鍵 K と OP または OPc（32 桁の 16 進） |
| `credential.type` | `OP` か `OPC` |
| `credential.amf` | Milenage の AMF（4 桁の 16 進。先頭のビット（separation bit）が 1 のもの） |
| `nssai` | Requested NSSAI（`sst`、`sd` のリスト。任意） |
| `session` | PDU セッションのリスト: `id`（1〜15）、`type`（`IPv4` / `IPv6` / `IPv4v6`）、`dnn`、`snssai` |
| `integrity` / `ciphering` | UE が対応を申告する NIA0〜3 / NEA0〜3（既定は 0〜2）。3（ZUC）は未実装 |
| `tun` | `true` で、PDU セッションを持つ UE ごとにホストに TUN デバイスを作る（9 章） |

`ue.type` と `ue.apn` は古い書き方で、`session` の ID 1 として読みます。

### gnodeb（自局）

| キー | 意味 |
|---|---|
| `name` | NG Setup の RANNodeName（省くとテンプレートの値） |
| `gnodebid` / `gnodebidlength` | gNB ID（16 進）と長さ（22〜32 ビット。既定 32） |
| `nci` | NR セル ID（9 桁の 16 進。既定は gNB ID の後ろを 0 で埋めたもの） |
| `plmn.mcc` / `plmn.mnc` | PLMN |
| `snssai` | NG Setup で申告するスライス（`sst`、`sd` のリスト） |
| `tac` | TAC（10 進。16 進なら `0x` を付ける） |
| `ngap.src` / `ngap.dst` | N2 の自局（任意。`port: 0` は任意のポート）と AMF のアドレス・ポート（38412） |
| `gtpu.src` | N3 の自局のアドレス・ポート（2152） |
| `gtpu.dst` | 書くと、5G コアが通知した UPF のアドレスの代わりにこれを使う（任意） |
| `web.src` | Web コンソールの待ち受け（無ければ開かない） |
| `http.src` | 疑似呼間通信の待ち受けと送信元（11 章。無ければ使えない） |

## 4. signal.txt（送る信号のテンプレート）

1 つの信号は、名前と NGAP の PDU（16 進）で書きます。名前は自由で、`send(名前)` で使います。`#` から行末はコメントです。

```
NGSetupRequest {
  00150035000004001b00090002f839500000000100524008028035672d72616e...
}

InitialUEMeesageRegistrationRequest [nas=keep] {
  000f40480000050055000200010026001a197e004179000d0102f839...
}
```

名前の後ろの `[...]` は送り方です。

| 送り方 | 意味 |
|---|---|
| （省略）`auto` | NGAP と NAS の動的な項目を補正し、NAS を保護して送る |
| `nas=keep` | NGAP は補正するが、NAS の本体はテンプレートのまま（保護はする） |
| `raw` | 何も変えずにテンプレートのバイト列を送る |

- テンプレートは、実機のキャプチャ（Wireshark で NGAP の PDU をコピーしたもの）や、Web コンソールの Signals 画面で作れます。
- 保護された NAS を含むテンプレートは、中身が読めること（平文か NEA0）が必要です。
- UE 関連のメッセージ（NGAP の UE ID を持つもの）は加入者が送り、それ以外はノードが送ります。
- 補正する項目は 10 章を見てください。補正しない項目はテンプレートの値のまま送ります。

## 5. scenario.txt（シナリオ）

### 5.1 例

```
var {
	public  count = 0       # 疑似呼全体で 1 つ
	private t     = 5s      # 加入者ごと
}

node {
	entry {
		send(NGSetupRequest)                # SCTP の接続時
	}
	receive(NGSetupResponse) {
		snap("NG Setup OK")
	}
}

any {                                       # どの状態でも
	receive(UEContextReleaseCommand) {
		send(UEContextReleaseComplete)
	}
}

start {                                     # 加入者の最初の状態
	entry {
		control(10)                         # 全加入者で 1 秒に 10 人まで
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

### 5.2 ブロック

| ブロック | 意味 |
|---|---|
| `var { … }` | 変数の宣言（5.5） |
| `node { … }` | ノード（gNB 自身）のハンドラ。UE 関連でない信号（NG Setup の応答、NG Reset など）を受ける |
| `any { … }` | 加入者のハンドラで、どの状態でも有効。状態ブロックに無いときに探す |
| `状態名 { … }` | 加入者の状態。`start` が最初の状態。名前は自由 |

ブロックの中に書けるもの:

| 項目 | いつ動くか |
|---|---|
| `entry { … }` | `node`: SCTP が接続したとき。状態: その状態に入ったとき（`start` は NG Setup の成功後）。`any` には書けない |
| `receive(NGAP名) { … }` / `receive(NGAP名:NAS名) { … }` | その信号を受けたとき。NAS 名は NGAP が運ぶ NAS の名前（5GSM なら内側の名前でも一致） |
| `expire(変数) { … }` | `timer(変数)` で始めたタイマーが満了したとき |
| `sync_receive(SYNC\|HO) { … }` | 他の疑似呼から通信を受けたとき（11 章） |

### 5.3 ハンドラの探し方

- 加入者: いまの状態のブロック → `any` の順に探します。同じブロックでは `NGAP:NAS` の一致が `NGAP` だけの一致より先です。
- ノード: `node` ブロックだけを探します。Paging は 5G-S-TMSI で加入者に渡し、該当者がいなければ `node` で受けます。
- 見つからなければ「unhandled」の警告イベントになります（組み込みの処理は済んでいます）。
- 名前の大文字・小文字は区別しません（`DownLinkNasTransport` は `DownlinkNASTransport` に一致します）。
- 認証、セキュリティ、セッションの情報の取り込みなどの組み込み処理は、ハンドラより先に行います。

### 5.4 文

ハンドラと entry の中には、コマンド（6 章）と `if` を書きます。

```
if ($ue == 1 && count > 10) {
	send(X)
} else if (mode == 2) {
	snap("mode 2")
} else {
	set(mode = 1)
}
```

- `move(状態)` はその場で状態を変え、ハンドラが終わってから遷移先の `entry` を実行します（同じ状態への move でも実行します）。entry の中の move による連鎖は 16 回までです。
- `#` と `//` から行末はコメントです（文の後ろにも書けます。日本語も可）。
- 予約語（状態名に使えない）: `var` `public` `private` `node` `any` `receive` `entry` `expire` `sync_receive` `if` `else` `timeout` `target`

### 5.5 変数と式

```
var {
	public  a = 1         # int
	public  rate = 0.5    # float
	private s = "text"    # string
	private t = 5s        # duration（100ms、1m30s など）
	private n = -1
}
```

- `public` は疑似呼全体で 1 つの値で、全加入者とノードから見えます。`private` は加入者ごとの値で、ノードからは使えません。
- `private` は SCTP の切断（linkdown）で初期値に戻ります。
- 型は初期値で決まり、変わりません。代入は `set(変数 = 式)` です。int の値は float の変数に入れられます。
- 式の演算子は C と同じ優先順位です: `||` `&&` `==` `!=` `<` `<=` `>` `>=` `+` `-` `*` `/` `%` と単項の `-` `!`、括弧。
  - int どうしは int（`/` は切り捨て）。int と float を混ぜると float。
  - `+` の片方が文字列なら、もう片方を文字にしてつなぎます（`"ue" + $ue` → `ue1`）。
  - duration どうしの足し引き、duration と数の掛け割りができます。
  - 比較の結果は true / false で、`if` の条件と `&&` `||` `!` に使います（変数には入れられません）。
  - 0 で割ると実行時のエラーになります。
- 組み込みの値:

| 名前 | 値 |
|---|---|
| `$ue` | 加入者の番号（1 から） |
| `$imsi` | 加入者の IMSI |
| `$state` | 加入者のいまの状態 |
| `$msg` | 最後に受けたトリガ（`NGAP:NAS`） |
| `$time` | 現在時刻 |

実行時にコマンドや条件が失敗すると、エラーのイベントを出してその文だけをやめ、次の文に進みます。

### 5.6 タイマー

duration 型の変数がタイマーになります。

- `timer(t)` で始め、満了すると `expire(t) { … }` が動きます。探す順は、満了した時点の状態のブロック → `any` です。
- 動いているタイマーに `timer(t)` を実行すると、最初からやり直します。
- `stoptimer(t)` で止めます。状態の遷移では止まらず、linkdown で止まります。
- `private` の変数は加入者のタイマー、`public` の変数はノードのタイマーです（ノードのタイマーは `node` の `expire`）。

### 5.7 呼量の制御（control）

`control(n)` は、その文を通る回数を、全加入者の合計で 1 秒に n 回までにします。

- 通る時刻を 1/n 秒おきに割り当て、加入者は自分の番まで待ちます（`control(5)` は 0.2 秒おき）。
- 手順をループさせるシナリオで、`send` の前に置いて負荷量（呼/秒）を決めます。
- `control(0)` は制限なしです。n には変数や式も書けます（`control(rate)` にして、`set(rate = 20)` で負荷を変える）。
- 待つ間は、その加入者の受信の処理も待つので、手順の頭に置いてください。
- linkdown で待ちは打ち切られ、その本体の残りは実行しません。

## 6. コマンド

| コマンド | 使える場所 | 内容 |
|---|---|---|
| `send(信号)` | シナリオ・CUI/GUI | signal.txt の信号を補正して送る。UE の信号は加入者、ノードの信号はノードから |
| `move(状態)` | シナリオ・CUI/GUI | 状態を変えて entry を実行する（CUI/GUI からは強制遷移） |
| `snap(書式, 値…)` | シナリオ・CUI/GUI | Printf 形式で表示する（`base.snap: true` のときだけ）。例 `snap("%s %d", $imsi, a)` |
| `set(変数 = 式)` | シナリオ・CUI/GUI | 変数に代入する。CUI/GUI では、public は 1 回、private は加入者ごとに実行する |
| `timer(変数)` / `stoptimer(変数)` | シナリオ・CUI/GUI | タイマーを始める / 止める |
| `control(n)` | シナリオ | 1 秒に n 回までに抑える（5.7） |
| `ping(宛先[, 回数[, 間隔[, サイズ]]])` | シナリオ・CUI/GUI | PDU セッション経由の ICMP Echo（回数 0 は stop まで） |
| `iperf(宛先[, ポート[, 帯域[, 時間[, ul\|dl[, 長さ]]]]])` | シナリオ・CUI/GUI | DN の iperf3 サーバと UDP の試験 |
| `iperftcp(…)` | シナリオ・CUI/GUI | 同じく TCP の試験（TUN 経由。9 章） |
| `stop([ping\|iperf])` | シナリオ・CUI/GUI | 動いている ping / iperf / iperftcp を止める |
| `mapon([項目])` / `mapoff([項目])` / `mapget([項目])` | シナリオ・CUI/GUI | マッピングの ON/OFF と確認（10 章） |
| `sync_send(SYNC\|HO, 宛先)` | シナリオ・CUI/GUI | 他の疑似呼に通信を送る（11 章） |
| `linkup` / `linkdown` | CUI/GUI | SCTP の接続と切断。切断すると加入者は初期状態に戻る |
| `reload` | CUI/GUI | signal.txt を読み直す（エラーがあれば何も変えない） |
| `wait(時間)` | CUI/GUI | 指定時間だけ待つ（バッチ用。1 時間まで） |
| `help([コマンド])` | CUI/GUI | コマンドの一覧と使い方 |

- **対象の指定:** CUI/GUI では、加入者に対するコマンドの前に `target(IMSI).` を付けると、その加入者だけに実行します（例 `target(208930000000001).send(AN_UEContextReleaseRequest)`）。付けなければ全加入者に実行します。
- **書き方の省略:** 引数の無いコマンドは括弧を省けます（`linkup`）。IPv4 アドレスと `アドレス:ポート` は引用符なしで書けます（`ping(10.0.0.1, 3)`）。
- **帯域の書き方:** `100M`、`1G`、`500k` のように書きます。小数は引用符で囲みます（`"1.5G"`）。

## 7. CUI とバッチ

- `gnodeb>` のプロンプトにコマンドを入力します。Ctrl-D で終了します。受信や状態の変化は、イベントとして随時表示されます。
- `@ファイル` で、ファイルに書いたコマンドを 1 行ずつ順に実行します（バッチ）。失敗した行で止まります。`--exec ファイル` で起動すると、バッチを実行して終了します（成功 0、失敗 1）。
- バッチの中では、応答を待つのに `wait(3s)` を使います。`#` と `//` から行末はコメントで、`@` は入れ子にできます。

```
# 登録 → ping → 解放 → Service Request → ping → 登録解除
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

## 8. Web コンソール

`gnodeb.web.src` を設定すると、ブラウザから `http://<アドレス>:<ポート>/` で使えます。認証や TLS は無いので、信頼できるネットワークで使ってください。

| 画面 | 内容 |
|---|---|
| Dashboard | SCTP と NG Setup の状態、加入者の数（登録・接続）、警告などの集計 |
| Subscribers | 加入者の一覧と詳細（状態、NAS の COUNT、PDU セッション、変数）。選ぶと、コマンドの対象（target）になる |
| Operations | 信号のボタン、コマンドの入力フォーム、バッチ |
| Signals | signal.txt の信号をデコードして表示・編集する。値の変更、任意 IE やリストの要素の追加・削除、別名での追加、新しい信号の作成（NGAP の種類を選んで作る）。補正される項目には「auto」の印と、切り替える項目名（`[mapping: plmn]` など）が付く。保存は signal.txt を書き換え、`reload` で送る信号に反映する |
| Sequence | 送受信のシーケンス図。メッセージを選ぶとデコードした中身を表示する |
| Scenario | 変数の値と、ブロック・ハンドラの一覧（加入者の数つき） |
| Logs | イベントの一覧 |
| License | ライセンス |

画面の下のコマンド欄には、CUI と同じコマンドを入力できます。

## 9. U-plane（ping・iperf・TUN）

PDU セッションが確立すると、gnodeb は UE の IP アドレスで GTP-U のトンネルを使えるようになります。

- `ping(宛先, 回数, 間隔, サイズ)`: ICMP Echo を送り、損失と RTT（min/avg/max/mdev）を表示します。DN 側から UE のアドレスへの ping にも応答します。
- `iperf(宛先, ポート, 帯域, 時間, ul|dl, 長さ)`: DN で動いている `iperf3 -s` を相手に UDP の試験をします（iperf3 互換。`dl` は iperf3 の `-R`）。

### TUN デバイス（ue.tun）

`ue.tun: true` にすると、PDU セッションを持つ UE ごとに、UE のアドレスを持つ TUN デバイス（先頭の加入者は `kronos-ue1`）と、そのアドレスから送る通信をデバイスに流す送信元ルールをホストに作ります。ホストの `iperf3` や `curl` などが、カーネルの TCP/IP で PDU セッションを使えます。

作成には `CAP_NET_ADMIN` が必要です。ビルドのたびに次を実行してください（`setcap` は `libcap2-bin` パッケージにあります）。

```sh
make setcap     # ビルドしてから sudo setcap cap_net_admin+ep ./gnodeb を実行する
```

```sh
iperf3 -c 10.0.0.1 -B 10.60.0.1 -t 10        # UE のアドレスに bind して TCP の上り
iperf3 -c 10.0.0.1 -B 10.60.0.1 -t 10 -R     # 下り
```

`iperftcp(宛先, ポート, 帯域, 時間, ul|dl, 長さ)` は、同じ TCP の試験を gnodeb の中から実行します（シナリオからも使えます。帯域 0 は TCP の流せるだけ）。デバイスは、セッションがなくなったとき、linkdown、終了で消えます。

## 10. マッピング（自動補正）と ON/OFF

送るとき、テンプレートの動的な項目を、config と加入者の状態に合わせて補正します（マッピング）。補正は次の項目ごとに ON/OFF できます。起動時はすべて ON です。

| 項目 | 補正する内容 |
|---|---|
| `plmn` | PLMN（NG Setup / RAN Configuration Update の gNB ID と TA、ULI、Cell Traffic Trace のセル） |
| `tac` | TAC（TA のリスト、ULI） |
| `nci` | NR セル ID（ULI、Cell Traffic Trace） |
| `gnodebid` | gNB ID |
| `ranname` | RANNodeName（`gnodeb.name`） |
| `snssai` | S-NSSAI（NG Setup のスライス、Requested NSSAI、ULNASTransport） |
| `timestamp` | ULI の TimeStamp（送信時刻） |
| `ngapid` | RAN-UE-NGAP-ID と AMF-UE-NGAP-ID（Path Switch の Source AMF ID、NG Reset のリストを含む） |
| `identity` | SUCI・5G-GUTI・5G-S-TMSI、Identity Response（求められた種類: SUCI・5G-GUTI・IMEI・5G-S-TMSI・IMEISV） |
| `ngksi` | ngKSI |
| `res` | Authentication Response の RES*、Authentication Failure（synch failure）の AUTS |
| `seccap` | UE Security Capability（Path Switch Request を含む） |
| `imeisv` | Security Mode Complete の IMEISV |
| `session` | PDU セッション ID・種別・DNN・PTI（ULNASTransport と 5GSM） |
| `status` | PDU Session Status、Uplink Data Status、Allowed PDU Session Status |
| `tunnel` | NGAP の PDU セッションのリスト（Setup・Release・Modify・Notify・Handover・Path Switch の応答、DL の TEID とアドレス、QoS フロー） |
| `container` | NAS message container（作り直しと暗号化） |

```
gnodeb> mapoff(plmn)                             # ノードと全加入者で PLMN をテンプレートのまま送る
gnodeb> target(208930000000001).mapoff(ngapid)   # その加入者だけ
gnodeb> mapget                                   # ON でない項目の一覧（mixed は一部だけ OFF）
gnodeb> mapon                                    # 全項目を ON に戻す
```

- 項目を省くと全項目が対象です。全項目を OFF にすると、テンプレートとバイト単位で同じものを送ります（NAS の保護は除く）。
- シナリオでは、そのブロックの主体（`node` ならノード、それ以外はその加入者）に効きます。
- 切り替えた状態は linkdown では戻りません。

## 11. 疑似呼間通信（sync_send / sync_receive）

疑似呼どうしが JSON の HTTP/REST で通信します。ハンドオーバーで、Source の加入者の情報を Target に渡すのに使います。両方の config に `gnodeb.http.src` が必要です（送信元のアドレスと待ち受けを兼ねます）。

```
# Source
connected {
	receive(…) { sync_send(HO, 10.0.0.31:12345) }
}

# Target（同じ IMSI の加入者で受ける）
start {
	sync_receive(HO)   { move(handover) }
	sync_receive(SYNC) { snap("sync from the source") }
}
```

| 種類 | 内容 |
|---|---|
| `SYNC` | 単純な同期。送り元と IMSI だけを運ぶ |
| `HO` | ハンドオーバー。加入者のマッピングの情報（登録状態、5G-GUTI、ngKSI、Kamf、NAS セキュリティの鍵と COUNT、PDU セッション、Source の AMF-UE-NGAP-ID）を運ぶ |

- 受け側は同じ IMSI の加入者で受けます（ノードから送った SYNC はノードで受けます）。探す順は receive と同じです。
- HO を受けた Target の加入者は、その情報で以後の NAS を保護し、マッピングします。
- Target の AMF からの HandoverRequest は、HO を受けた加入者に振り分けます（PDU セッションの UL TEID で照合）。HandoverRequestAcknowledge には、割り当てた DL TEID を入れます。Xn の場合は PathSwitchRequest を送ります。
- 送った側（Source）は何もしません。その後はシナリオの制御どおりです。

HTTP の形:

```
POST http://<宛先>/kronos/v1/sync
Content-Type: application/json

{"kind": "HO", "from": "10.0.0.30:12345", "imsi": "208930000000001",
 "handover": {"registered": true, "guti": "f202f839…", "ngKSI": 1, "kamf": "…",
   "security": {"integrity": 2, "ciphering": 2, "knasInt": "…", "knasEnc": "…", "uplinkCount": 4, "downlinkCount": 3},
   "sessions": [{"id": 1, "ueIp": "10.60.0.7", "upf": "10.0.0.40", "ulTeid": 1, "qfi": [1]}],
   "amfUeNgapId": 100, "nextPTI": 2}}
```

応答は `{"accepted": true}`（202）か `{"accepted": false, "error": "…"}` です（400 は形式の誤り、404 はその IMSI の加入者がいない、503 はノード宛てで SCTP が未接続）。認証も TLS も無く、鍵が平文で流れるので、試験用のネットワークで使ってください。

## 12. ログとトレース

カレントディレクトリの `log/` に出力します。

| ファイル | 内容 |
|---|---|
| `log/gnodeb.txt` | アプリケーションのログ |
| `log/trace-YYYYMMDD-HHMMSS.jsonl` | 全イベント（送受信した PDU の 16 進、状態の変化、snap、警告など）を 1 行 1 件の JSON で |

## 13. 制限事項

- SUCI は null スキームだけです（Profile A/B は未実装）。
- 128-NEA3 / 128-NIA3（ZUC）は未実装です。
- EAP-AKA' の認証には対応していません（5G AKA のみ）。
- AMF の役（反転動作）は未実装です。
- I-RNTI による Paging（RRC_INACTIVE）には対応していません。
- libnas がデコードできない NAS メッセージは、名前の判定だけ行います。
- Rel-17/18 の MBS・AIoT などのメッセージは、NGAP ID と ULI の補正だけ行います。
