# Linux / Docker / Kubernetes コマンドのチートシート
## Linux
- 現在のディレクトリを表示
  ```bash
  $ pwd
  ```
- 現在のディレクト内ファイルを表示
  ```bash
  $ ls [-a -l]
  ```
- ディレクトリ間を移動
  ```bash
  $ cd 
  ```
- ファイルの中身を表示
  ```bash
  $ cat <file name>
  ```
- 新しいディレクトリ作成
  ```bash
  $ mkdir <new directory name>
  ```
- 新しいファイルを作成する
  ```bash
  $ touch <new file name>
  ```
- ファイル書き込み（上書き）
  ```bash
  $ echo > "contents"
  ```
- ファイル書き込み（追記）
  ```bash
  $ echo >> "contents"
  ```
- ファイルやディレクトリを検索
  ```bash
  $ find <検索の始まり（例えばルートから探すときは `.`）> -type <ディレクト検索： `d`、ファイル検索：`f`> -name <検索したい文字（例えば`html`）>
  ```
## Docker
- 新しいイメージの取得
  ```bash
  $ docker pull <image name:tag name>
  ```
- イメージ一覧の取得
  ```bash
  $ docker images
  ```
- イメージの削除
  ```bash
  $ docker rmi <image name>
  ```
- イメージの履歴を表示
  ```bash
  $ docker history <image name>
  ```
- コンテナ一覧の取得
  ```bash
  $ docker ps -a
  ```
- コンテナ情報の取得
  ```bash
  $ docker inspect <container name>
  ```
- コンテナの作成と起動
  ```bash
  $ docker run <image name>
  ```
  - オプション
    - (公開ポートの指定) -p
    - (ボリュームマウント) -v
    - (コンテナ名を指定) --name 
    - (環境変数を設定) -e 
    - (コンテナ停止後コンテナ自動削除) --rm 
    - (バックグラウンドで起動) -d 
  ```bash
  $ docker run -p <host port>:<container port> -v <path to source>:<path to mount the source> --name <container name> -e NEW_ENV=<env variable> --rm -d <image name>
  ```
- コンテナのログ表示
  ```bash
  $ docker logs <container name>
  ```
- 停止中のコンテナの起動
  ```bash
  $ docker run <container name>
  ```
- 起動中のコンテナの停止
  ```bash
  $ docker stop <container name>
  ```
- 起動中のコンテナに入る（tty テレタイプライターで繋ぐ）
  ```bash
  $ docker exec -it <container name> <シェル名、例えば `sh`>
  ```
- コンテナの削除
  ```bash
  $ docker rm <container name>
  ```
- 起動中のコンテナから新しいイメージ作成
  ```bash
  $ docker commit <container name> <new image name>
  ```
- Dockerfileから新しいイメージ作成
  ```bash
  $ docker build --tag <new image name> <path to Dockerfile>
  ```
- イメージのタグ付け
  ```bash
  $ docker tag <image name> <username>/<repo name>
  ```
- Docker Hubにログイン
  ```bash
  $ docker login
  ```
- DockerHubにイメージをアップロード
  ```bash
  $ docker push <username>/<image name>
  ```
- docker-composeコマンド
  - コンテナを起動する
    ```bash
    $ docker compose -f <docker compose用yamlファイル名> up [--scale <containerName>=<numOfContainer>]
    ```
  - コンテナを停止して削除する
    ```bash
    $ docker compose -f <docker compose用yamlファイル名> down
    ```
- ネットワーク
  - ネットワーク一覧を取得
    ```bash
    $ docker network ls
    ```
  - ネットワークを作成（--driverオプション付き）
    ```bash
    $ docker network create --driver bridge <network name>
    ```
- ネットワークを削除
    ```bash
    $ docker network rm <network name>
    ```
- Swarmを有効にする
  - Swarmの状態を取得
    ```bash
    $ docker info | grep -i swarm
    ```
## Swarm
  - Swarmを有効にする
    ```bash
    $ docker swarm init
    ```
  - Node一覧を取得する
    ```bash
    $ docker node ls
    ```
  - serviceの一覧を取得する
    ```bash
    $ docker service ls
    ```
  - serviceを作成する
    ```bash
    $ docker service create --name <service name> <image name>
    ```
  - serviceをスケールアップ
    ```bash
    $ docker service scale <service name>=<num>
    ```
  - serviceをローリングアップデート
    ```bash
    $ docker service update --image <new image:tag> <service name>
    ```
  - serviceの起動中のコンテナ一覧を取得
    ```bash
    $ docker service ps <service name>
    ```
## Kubernetes
- minikube/kubectlのインストール
  ```bash
  $ brew install minikube
  $ brew install kubectl
  ```
- minikube起動・ステータス確認・停止・削除
  ```bash
  $ minikube start
  $ minikube status
  $ minikube stop
  $ minikube delete
  ```
- Podの起動
  ```bash
  $ kubectl run --image <imageName:tag> --restart Never <newPodName> // like 'docker run --image <imageName:tag> --name <newContainerName> <imageName>'
  ```
- Podの一覧取得
  ```bash
  $ kubectl get pods // like 'docker ps'
  ```
- Pod内のログ表示（Log）
  ```bash
  $ kubectl logs <podName> // like 'docker logs <containerName>'
  ```
- Podのメタデータ取得
  ```bash
  $ kubectl describe pod <podName> // like 'docker inspect <containerName>
  ```
- Podのなかにシェルで入る
  ```bash
  $ kubectl exec -it <podName> sh // like 'docker exec -it <containerName> sh
  ```
- Podを削除
  ```bash
  $ kubectl delete pod <podName> // like 'docker rm <containerName>'
  ```
- コンテナの環境変数の指定（`--env KEY=value`）
  ```bash
  $ kubectl run --env KEY=value --image <imageName:tag> --restart Never <newPodName> // like 'docker run --env KEY=value -d --name <newContainerName> <imageName>'
  ```
- コンテナのポート設定（`--port <portNumber>`）
  ```bash
  $ kubectl run --port <portNumber> --image <imageName:tag> --restart Never <newPodName>
  ```
- Service
  - Serviceの一覧取得
    ```bash
    $ kubectl get service
    ```
  - ClusterIPタイプのServiceを作る
    ```bash
    $ kubectl expose pod <targetPodName> --type ClusterIP --port <portNumber> --name <newServiceName>
    // 疎通確認のため、クラスター内にcurlコンテナを作ってClusterIPサービス経由でPodにアクセスしてみる。
    $ kubectl run --restart Never --image curlimages/curl:7.68.0 -it --rm curl sh
    $ # curl <newServiceName>:<portNumber>
    ```
  - NodePortタイプのServiceを作る
    ```bash
    $ kubectl expose pod <targetPodName> --type NodePort --port <portNumber> --name <newServiceName>
    ```
  - 
    ```bash
    $ 
    ```
  - 
    ```bash
    $ 
    ```
  - 
    ```bash
    $ 
    ```
  - 
    ```bash
    $ 
    ```
- 
  ```bash
  $ 
  ```
- 
  ```bash
  $ 
  ```
