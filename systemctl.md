# systemctl
## 参考にさせて頂いたサイト
[【systemctl】の使い方 〜オプション一覧,自動起動,list-units/list-unit-filesの見方, reloadとrestartの違い, runningとexitedの違い〜](https://milestone-of-se.nesuke.com/sv-basic/linux-basic/systemctl/#:~:text=%E3%81%AE%E3%81%9F%E3%82%81%E3%81%AB-,systemctl%20%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%A8%E3%81%AF,%E3%82%92%E3%82%B3%E3%83%B3%E3%83%88%E3%83%AD%E3%83%BC%E3%83%AB%E3%81%99%E3%82%8B%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%81%A7%E3%81%99%E3%80%82a)
## コマンド一覧
- 基本

  `$ systemctl [命令] [サービス名]`

- 開始

  `$ systemctl start [サービス名]`
  
- 停止

  `$ systemctl stop [サービス名]`

- 再起動

  `$ systemctl restart [サービス名]`

- 再読み込み

  `$ systemctl reload [サービス名]`

- 特定サービスの起動状態表示

  `$ systemctl status [サービス名]`

- 全てのサービスの起動状態表示

  `$ systemctl list-units -t service`

- サーバーが起動したタイミングでサービスも自動起動させる

  `$ systemctl enable [サービス名]`

- サーバーが起動したタイミングでサービスも自動起動させない

  `$ systemctl disable [サービス名]`

- 特定のサービスが自動起動するか否かを確認するコマンド

  `$ systemctl is-enabled [サービス名]`

- サービスが自動起動するか否かを確認するコマンド（chkconfig --listに代わるコマンド）

  `$ systemctl list-unit-files -t service`
