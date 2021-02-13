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
   $ nslookup <DNS>
    # IPアドレスが返却される
   ```
   ```console
   $ nslookup <IPアドレス>
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
   $ . ~/.nvm/nvm.sh
   $ nvm install node
   $ node -e "console.log('Running Node.js ' + process.version)" # Running Node.js VERSION ← のように表示されたら成功！
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
      $ telnet <ec2インスタンスのIP> 8080

        Trying 18.179.53.166...
        Connected to <ec2インスタンスのIP>.
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
      $ telnet <ec2インスタンスのIP> 8080

        Trying 18.179.53.166...
        Connected to <ec2インスタンスのIP>.
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

インターネットから隠された（接続出来ない）サブネットを構築する！セキュリティを高められるよ。そこに DB サーバーを立てる。パブリックサブネットは「10.0.1.0/24」だったけど、プライベートサブネットは「10.0.2.0/24」にする。

1. プライベートサブネットを作る  
   作り方は基本的にパブリックサブネットを作るときと一緒。「CIDR ブロック」に「10.0.2.0/24」を設定する。ルートテーブルはデフォルトの状態で、「自身のネットワーク（送信先：10.0.0.0/16）」に対してのルーティングがされているように設定する。なぜならインターネットに接続しないから。

1. DB サーバーを構築する  
   EC2 タブで Web サーバーを構築したように、DB サーバーを構築する。サブネットは 1.で作成したプライベートサブネットを選択して、自動割り当てパブリック IP は「無効化」とする。またプライベート IP は「10.0.2.10」に設定。セキュリティグループ名は「DB-SG」とし、タイプには SSH に加えて、「MYSQL/Aurora」を選択、送信元を 「任意の場所」に設定。

1. `ping`コマンドで疎通確認！  
   `ping`コマンドはサーバー間での疎通を確認するときに使うコマンドで、「ICMP（Internet Control Message Protocol）」を用いている。
   1. この ICMP が通るように「DB-SG」の設定を編集（「全ての ICMP」を許可するように設定）する。
   1. Web サーバーに接続
      ```console
      $ ssh -i my-key ec2-user@<WebサーバーのパブリックIP>
      ```
   1. 接続後、DB サーバーに疎通確認！
      ```console
      $ ping 10.0.2.10（<DBサーバーのプライベートIP>）
      ```

## 6. NAT を構築する

1.  NAT(Network Address Translation)の仕組み

    1. NAT は IP アドレスを変換する装置（パブリック IP アドレス ⇄ プライベート IP アドレス）
    1. プライベートサブネット側からインターネットへアクセスを可能にする！
    1. インターネットからプライベートサブネット側にはプライベート IP アドレスは公開されていない
    1. NAT と似たものに NAPT (Network Address and Port Translation)がある

1.  NAT インスタンス  
    NAT ソフトウェアがあらかじめインストールされた AMI から起動した EC2 インスタンス

1.  NAT ゲートウェイ(今回はこっちを使ってみる)  
    NAT 専用に構成された仮想的なコンポーネントで、配置するサブネットを設定するだけで構成できるから簡単！

1.  パブリックサブネットとプライベートサブネットを NAT でつなげる

    1.  VPC メニューから「NAT ゲートウェイ」を選択して、「NAT ゲートウェイの作成」を押下
    1.  サブネットはパブリックサブネットで、Elastic IP を割り当てる
    1.  プライベートサブネットからインターネットに対して通信するとき、パケットが NAT ゲートウェイの方向に流れるようにルートテーブルを更新する

        1.  NAT ゲートウェイ ID を確認する
        1.  プライベートサブネットのデフォルトゲートウェイを NAT ゲートウェイに向ける設定にする
        1.  疎通確認するよ  
            プライベートサブネットに立てた DB サーバーに踏み台サーバーを利用して接続して、`curl`コマンドでインターネットと疎通できているか確認してみよう。
            ```console
            $ curl https://www.yahoo.co.jp
            ```
            HTML を取得出来ていれば成功！！
        1.  NAT ゲートウェイを使う必要がないなら削除しよう！課金されるぞ！

1.  デフォルトゲートウェイとかその辺の知識

    参考：[デフォルトゲートウェイとは](https://livra.geolocation.co.jp/iplearning/654/)

    > 同一ネットワーク内に存在しない外部ネットワークと通信する必要がある際に、どのゲートウェイ（ルータ）を使用したらよいかをあらかじめ指定しておきます。 ネットワーク内部のホストから見ると、「デフォルト（既定）」のゲートウェイさえ知っておけば、外部ネットワークとの通信はゲートウェイが取り持ってくれるわけです。

    参考：[Wiki デフォルトゲートウェイ](https://ja.wikipedia.org/wiki/%E3%83%87%E3%83%95%E3%82%A9%E3%83%AB%E3%83%88%E3%82%B2%E3%83%BC%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A4)

    > コンピュータネットワークにおいてデフォルトゲートウェイ（default gateway）とは、内部ネットワークと外部ネットワークを接続するためのノードである。IP ネットワークにおいて、経路が分からない IP アドレス宛のパケットは、デフォルトゲートウェイ又はデフォルトルートに転送される。一般的に、デフォルトゲートウェイはルーターである。

    参考：[デフォルトルート](https://ja.wikipedia.org/wiki/%E3%83%87%E3%83%95%E3%82%A9%E3%83%AB%E3%83%88%E3%83%AB%E3%83%BC%E3%83%88)

    > 宛先のコンピュータが同じスイッチングハブに接続されている場合など、経路が明確な場合、コンピュータはその経路情報を使用して通信を行うため、デフォルトルートの設定は無視される。しかし、それ以外の場合は、全てのパケットはデフォルトルートの設定に従って、デフォルトゲートウェイに転送される。

    参考：[0.0.0.0](https://ja.wikipedia.org/wiki/0.0.0.0)

    > IPv4 において、0.0.0.0 は全ビットが 0 の IP アドレスであり、無効、不明、または適用外の対象を指定するために使用されるルーティング不可のメタアドレスである。このアドレスには、いくつかの特別な意味が割り当てられている。

## 7. DB を用いたブログシステムの構築

1. DB サーバーに MySQL をインストールして、WordPress からデータを保存できるようにデータベースを作成する

   1. MySQL のインストール  
      Web サーバーを踏み台にして DB サーバーに接続して[AWS EC2 AmazonLinux2 MySQL をインストールする](https://qiita.com/miriwo/items/eb09c065ee9bb7e8fe06)
   1. MySQL の起動と初期設定
      ```console
      $ sudo service mysqld start
      ```
   1. MySQL の管理者の初期パスワードを確認  
      `/var/log/mysqld.log`に temporary password が記述されてる。以下のような箇所に`xyz`っていう風に初期パスワードが記載されている

      ```console
      $ sudo less /var/log/mysqld.log

      ****************　省略　**************************
      〜〜〜 A temporary password is generated for root@localhost: xyz
      ****************　省略　**************************
      ***
      ```

   1. MySQL の管理者パスワードを設定  
      上記で確認したパスワードで mysql に接続

      ```console
      $[ec2-user@ip-10-0-2-10 ~]$ mysql -u root -p
         Enter password:
         Welcome to the MySQL monitor.  Commands end with ; or \g.
         Your MySQL connection id is 12
         Server version: 8.0.23

         Copyright (c) 2000, 2021, Oracle and/or its affiliates.

         Oracle is a registered trademark of Oracle Corporation and/or its
         affiliates. Other names may be trademarks of their respective
         owners.

         Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

         mysql>
      ```

      接続できたら、初期パスワードを変更しとく

      ```console
         mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '大文字、小文字、特殊文字(アスタリスク等)を含む新しいrootユーザのパスワード';
      ```

      じゃないと怒られる

      ```console
         mysql> なんかのsqlコマンド

         ERROR 1820 (HY000): You must reset your password using ALTER USER statement before executing this statement.
      ```

   1. WordPress 用のデータベースを作成

      1. （root の初期パスワード変更後）データベースを作成する。DB 名は「wordpress_db」とする。
         ```console
         mysql> CREATE DATABASE wordpress_db DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
         ```
         ※ 文字セット（charset_name）と照合順序（collation_name）指定はこの辺を参考にして  
         → [MySQL:文字セットと照合順序)](https://www.dbonline.jp/mysql/ini/index5.html)
      1. ユーザーを作成する。ユーザー名は「wordpress_user」とする。
         ```console
         mysql> CREATE USER wordpress_user IDENTIFIED BY '大文字、小文字、特殊文字(アスタリスク等)を含む新しいwordpress_userの為のパスワード';
         ```
         ※ ユーザーの作成はこの辺を参考にして  
         → [MySQL:ユーザーを作成する（CREATE USER 文）)](https://www.dbonline.jp/mysql/user/index1.html)
      1. 「wordpress_user」に「wordpress_db」に対する全てのアクセス権を与える。
         ```console
         mysql> GRANT ALL ON wordpress_db.* TO wordpress_user;
         ```
         ※ アクセス権限周りについてはこの辺を参考にして  
         → [MySQL:ユーザーに権限を設定する(GRANT 文)](https://www.dbonline.jp/mysql/user/index6.html)
      1. 権限設定を反映させるコマンドを実行
         ```console
         mysql> flush privileges;
         ```
      1. wordpress_user が登録されたかの確認
         ```console
         mysql> SELECT user, host FROM mysql.user;
         ```
         wordpress_user の host が%となっていれば、全てのホストから接続可能になっていることになる。
      1. DB サーバーが起動したタイミングで MySQL も起動するように設定しておく

         ```console
         mysql> exit
         $ sudo systemctl enable  mysqld
         $ sudo systemct is-enabled mysqld

            # enabled ってなってたら設定できてる
         ```

         ※ `systemctl`のコマンドはこの辺を参考に  
         → [systemctl コマンド](https://qiita.com/sinsengumi/items/24d726ec6c761fc75cc9)

1. Web サーバーに WordPress をインストールする

1. WordPress の初期設定をして、1.のデータベースを使うように構成する

## 8. TCP / IP による通信の仕組みを理解する
