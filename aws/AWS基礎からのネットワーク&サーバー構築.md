# AWS 基礎からのネットワーク&サーバー構築

## 1. ネットワークを構築する / 実験用の VPC を作成する

1. IP アドレス

   1. パブリック IP アドレス  
      インターネットに接続する際に用いる IP アドレスのこと
   1. プライベート IP アドレス
      インターネットで使われない IP アドレスのこと  
      |IP アドレスの範囲|
      |---|
      |10.0.0.0 ~ 10.255.255.255|
      |172.16.0.0 ~ 172.31.255.255|
      |192.168.0.0 ~ 192.168.255.255|
   1. IP アドレス範囲と表記方法  
      「ネットワーク部」と「ホスト部」があるが、「ホスト部」に割り当てる IP アドレスの範囲は **「２の n 乗個で区切る」** ルールがある。よく使われるのは「256 個(後ろから 8 ビット分)」「65536 個(後ろから 16 ビット分)」。

1. CIDR 表記とサブネットマスク表記  
   「ネットワーク部」と「ホスト部」がどこで区切られているかを簡単に表す。

   1. 普通の表記  
      192.168.0.0 ~ 192.168.255.255 (16 ビットがネットワーク部)  
      192.168.1.0 ~ 192.168.1.255 (24 ビットがネットワーク部)

   1. CIDR 表記  
      192.168.0.0/16 (16 ビットがネットワーク部)  
      192.168.1.0/24 (24 ビットがネットワーク部)

   1. サブネットマスク表記  
      192.168.0.0/255.255.0.0 (16 ビットがネットワーク部)  
      192.168.1.0/255.255.255.0 (24 ビットがネットワーク部)

1. VPC を作る

   1. AWS management console から VPC を 選択
   1. リージョン northeast を選択
   1. VPC の作成を押下
   1. ネームタグに名前をつける
   1. CIDR block 「10.0.0.0/16」で作成する

1. サブネット

   1. 割り当てられた CIDER ブロックを、そのまま使わずに、小さな CIDER ブロックに分割して利用する
   1. 上記の VPC は 「/16」でネットワークを区切った（10.0.0.0 ~ 10.0.255.255、つまり 65,536 個の IP アドレスがある）けれど、今回、サブネットは「/24」で区切る、つまり、256 アドレス \* 256 個サブネット
   1. サブネットに分ける目的は、障害発生時に備えた「物理的な隔離」と「セキュリティ上の理由（サブネットでセキュリティの強弱をつけたい場合など）」があげられる

1. パブリックサブネットを作る

   1. サブネットに作成する
   1. 名前をつける（パブリックサブネット、とか）
   1. 分割したい VPC を選択する
   1. CIDR ブロックを設定する（10.0.1.0/24）← 8 bit がホスト部
   1. 作成押下

1. インターネットゲートウェイの作成

   1. 作成した VPC とインターネットが繋がるためのポイントが「インターネットゲートウェイ」
   1. インターネットゲートウェイが無いと、そもそも外界と繋がれないから大事な存在

1. ルーターとルートテーブル

   1. ルーターは沢山あって、受け取ったパケットの中にある「宛先 IP アドレス」確認して、転送先ネットワークを決める
   1. 「宛先 IP アドレス」に最短で到達させるためにはどのネットワークに飛ばせばいいかを考える
   1. その時に使うのが「ルートテーブル」で、そこにはどのネットワークに流すかの設定が書いてある
   1. ルートテーブルにネットワーク情報が無いと、うまいこと転送できないから、事前に設定する必要あり
      1. EGP (Exterior Gateway Protocol)  
         ISP（インターネットサービスプロバイダー）などの「ある程度お大きなネットワーク」は、そのネットワークを管理する「AS 番号（Autonomous System 番号）」を持っていて、それをやりとりして、「どのネットワークの先に、どのネットワークが接続されているか」を、大まかにやりとり把握する。
         1. BGP (Border Gateway Protocol)
      1. IGP (Interior Gateway Protocol)  
         EGP 内部のルーター同志で、ルートテーブルの情報をやりとりする。
         1. OSPF (Open Shortest Path First)
         1. RIP (Routing Information Protocol)

1. デフォルトゲートウェイをインターネットに向けて設定する
   1. ルートテーブルの作成
      1. ネームタグで名前をつける（パブリックルートテーブル、とか）
      1. 作成済み VPC を選択する
      1. 作成ボタン押下
   1. ルートテーブルをサブネットに割り当てる
      1. 作成済みルートテーブルを選択
      1. サブネットの関連付けから編集
      1. 割り当てたいサブネットにチェックを付けて保存
   1. デフォルトゲートウェイをインターネットゲートウェイに設定する
      1. 作成済みルートテーブルを選択
      1. ルートから編集
      1. 送信先に「0.0.0.0/0」を入力、ターゲットに「igw-xxxx」を選択して保存
      1. （確認１）サブネットから、「パブリックサブネット」の「ルートテーブル」を確認する
      1. （確認２）「0.0.0.0/0」が、インターネットゲートウェイに設定されていることを確認する

## 2. サーバーを構築する

1. EC2 インスタンスを作成

   1. リージョンの確認
   1. インスタンス作成開始
      1. AMI の選択
      1. インスタンスタイプの選択
      1. インスタンス詳細情報の設定
         1. ネットワークに「作成済み VPC」を選択
         1. サブネットに作成済み「パブリックサブネット」を選択
         1. 自動お割り当てパブリック IP は「有効化」
         1. ネットワークインターフェイスのオプライマリ IP に「10.0.1.10」を入力
         1. 「ストレージの追加」へ
      1. ストレージ
         1. デフォルトのままで OK。仮想ハードディスク（EBS：Elastic Block Store）が設定される。
         1. 「Add Tags」へ
      1. Add Tags
         1. キーに「Name」、値に「Web サーバー」と入力
         1. 「セキュリティグループの設定」へ
      1. セキュリティグループの設定
         1. セキュリティグループ名に「WEB-SG」と入力
         1. 「確認と作成」へ
      1. 確認と作成
         1. 「作成」押下
         1. キーペアを作成＆ダウンロード
         1. インスタンスを作成し起動する
         1. 「インスタンス表示」へ
      1. インスタンス表示
         1. status が running とか起動中とかになるのを確認する
         1. status check が「2/2 チェックに合格しました。」となるのを確認する

1. SSH で EC2 インスタンスに接続する
   ```console
   $ ssh -i my-key.pem ec2-user@xx.xx.xx.xx (EC2インスタンスのパブリックIP)
   ```
   もし pem ファイルのパーミッションで怒られたら、、、
   ```console
   $ chmod 400 my-key.pem
   ```
   最初の質問には`yes [enter]`で答えると、ウェルカムメッセージが表示され、接続完了。
1. パブリック IP アドレスを固定化する
   1. Elastic IP メニューから、「新しいアドレスの割り当て」押下し、アドレスを割り当てする
   1. 確保したアドレスをインスタンスに割り当てる（アクションからインスタンスを選択）
   1. 「関連付け」押下
1. ファイアウォールで接続制限
   1. AWS では「セキュリティグループ」がファイアウォールの役割を担う
   1. デフォルトでは「SSH・ポート 22」に対して、全ての通信（0.0.0.0/0）を許可している
   1. このままではそれ以外の通信ははじかれる為、必要に応じてセキュリティグループの設定を調整する必要がある。
   1. 「インバウンドルール」は外から中（インスタンス）にアクセスしようとする通信に対する制限
   1. 「アウトバウンドルール」は中（インスタンス）から外にアクセスしようとする通信に対する制限

## 3. Web サーバーソフトをインストール

1. Apache HTTP Server を インストールする

   1. SSH で AWS のインスタンスに接続

      ```console
      $ ssh -i my-key.pem ec2-user@xx.xx.xx.xx(パブリックサブネットに立てたサーバーのパブリックDNS)
      ```

   1. install Apach

      ```console
      $ sudo yum -y install httpd # -y means that user will not check the instalation detail and install it instantly
      ```

   1. start Apach

      ```console
      $ sudo service httpd start
      ```

   1. configure Apach will be started when the server starts

      ```console
      $ sudo chkconfig httpd on
      ```

   1. check if the configuration has been set properly

      ```console
      $ sudo chkconfig --list httpd
      ```

   1. check processes

      ```console
      $ ps -ax | grep httpd
      ```

   1. check network status

      ```console
      $ sudo lsof -i -n -P
      ```

   1. check DNS server
      ```
      $ nslookup ec2-xx-xx-xx-xx.ap-northeast-1.compute.amazonaws.com(パブリックサブネットに立てたサーバーのパブリックDNS)
      ```

1. ファイアウォールの設定  
   今の状態では http 通信はできない。なぜならポート８０番がファイアウォールで拒否されているから！  
   セキュリティグループの設定を編集してポート８０番への通信を許可するようにしよう。

   1. セキュリティーグループ・WEB-SG の設定画面へ
   1. インバウンドルールを選択して、タイプ：「カスタムルール」、ポート範囲：「８０」、送信元：「0.0.0.0/0」（全てのホスト）で保存。
   1. Web サーバーのパブリック IP アドレスをブラウザから叩いてみて疎通確認する。（http://xx.xx.xx.xx）
   1. Apach のウェルカムページが表示されれば成功！

1. VPC 内の EC2 インスタンスへ DNS（Domain Name System）名 が自動で割り振られる設定  
   IP アドレスでの表記は覚えづらいから、DNS 名が便利。EC2 インスタンスへ DNS 名を割り当てる設定をしよう。

   1. VPC を選択
   1. 「アクション」から「DNS ホスト名の編集」を選択
   1. 「はい」を選択すると、設定完了
   1. 選択した VPC 内に立てた EC2 インスタンスを洗濯
   1. パブリック/プライベート DNS を確認すると、割り当てられているのが確認できるはず
   1. パブリック DNS を使ってブラウザから Web サーバー（EC2 インスタンス）にアクスしてみよう

1. DNS <-> IP アドレス を確認できる `nslookup`コマンド
   ```console
   $ nslookup DNS
    # IPアドレスが返却される
   ```
   ```console
   $ nslookup IPアドレス
    # DNSが返却される
   ```

## 4. HTTP の動きを確認する

`telnet`コマンドを使って HTTP リクエストを投げてみる実験をする

1. 実験用の EC2 インスタンスをパブリックサブネットに立てる
1. インバウンドルールでポート８０８０番への通信を許可する
1. 実験用サーバーに接続する
1. Node.js をインストールする  
   [AWS のチュートリアルからインストール方法がわかります！](https://docs.aws.amazon.com/ja_jp/sdk-for-javascript/v2/developer-guide/setting-up-node-on-ec2-instance.html)

   ```console
   $ cd ~ $ curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
   $ . ~/.nvm/nvm.sh $ nvm install node $ node -e "console.log('Running Node.js ' + process.version)" # Running Node.js VERSION ← のように表示されたら成功！
   ```

1. app.js を作成して以下のコードをはる

   ```console
   $ vi app.js # app.jsを作る
   ```

   ```js
   var http = require('http');
   http
   	.createServer(function (req, res) {
   		console.log('Hello World!');
   		var data = {
   			RequestHeader: req.headers,
   		};

   		if (req.method == 'GET') {
   			response(res, data);
   		} else {
   			req.on('data', function (body) {
   				data.RequestBody = body.toString();
   				req.on('end', function () {
   					response(res, data);
   				});
   			});
   		}
   	})
   	.listen(8080);

   function response(res, data) {
   	console.log('request is coming!!');
   	var json = JSON.stringify(data);
   	res.writeHead(200, {
   		'Content-Type': 'application/json',
   		'Content-Length': json.length,
   	});
   	res.end(json);
   }
   ```

   ```console
   $ cat app.js # ちゃんとファイル出来てるか確認。
   $ node app.js # nodeサーバーを起動する
   ```

1. ローカルに別のターミナルを開いて、`telnet`で通信してみる！

   1. まずは GET request

      ```console
      $ telnet ec2-18-179-53-166.ap-northeast-1.compute.amazonaws.com 8080

        Trying 18.179.53.166...
        Connected to ec2-18-179-53-166.ap-northeast-1.compute.amazonaws.com.
        Escape character is '^]'.
        GET / HTTP/1.1[enter] # 入力後enter
        User-Agent: OreOreGetRequest[enter][enter] # 入力後enter

        HTTP/1.1 200 OK  # 結果が返ってくる!!
        Content-Type: application/json
        Content-Length: 51
        Date: Wed, 10 Feb 2021 10:20:42 GMT
        Connection: keep-alive
        Keep-Alive: timeout=5

        {"RequestHeader":{"user-agent":"OreOreGetRequest"}}
      ```

   1. 続いて POST request

      ```console
      $ telnet ec2-18-179-53-166.ap-northeast-1.compute.amazonaws.com 8080

        Trying 18.179.53.166...
        Connected to ec2-18-179-53-166.ap-northeast-1.compute.amazonaws.com.
        Escape character is '^]'.
        POST / HTTP/1.1
        User-Agent: OreOrePostRequest
        Content-Length: 4

        test
        HTTP/1.1 200 OK
        Content-Type: application/json
        Content-Length: 94
        Date: Wed, 10 Feb 2021 10:24:29 GMT
        Connection: keep-alive
        Keep-Alive: timeout=5

        {"RequestHeader":{"user-agent":"OreOrePostRequest","content-length":"4"},"RequestBody":"test"}
      ```

1. 5.で app.js を起動していたが、そのターミナルには以下のようにログがあるはず
   ```console
   [ec2-user@ip-10-0-1-20 ~]$ node app.js
   Hello World!
   request is coming!!
   ```

## 5. プライベートサブネットを構築する

```console
$
```

```console
$
```

## 6. NAT を構築する

## 7. DB を用いたブログシステムの構築

## 8. TCP / IP による通信の仕組みを理解する
