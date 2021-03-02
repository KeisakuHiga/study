# Cookie と CORS
## 前提
- 技術スタック
  - フロントエンド：Nuxt.js →　http://localhost:3000
  - バックエンド：Node.js →　http://localhost:4000
  - [universal-cookie](https://www.npmjs.com/package/universal-cookie)を使う

## やりたかったこと
- サーバーで発行されたcredential情報をcookieとして保存すること
- cookieに保存されているcredential情報をサーバーに送付して、それを元に認証機能を実装すること

## はまったこと
- クライアントからのリクエストに対して、サーバー側からのレスポンスに`Set-Cookie`ヘッダーを付与することは出来て、ブラウザのdeveloper toolでも`Set-Cookie`ヘッダーがレスポンスにあることを確認出来ているのに、クライアントでcookieがセットされない。

## 原因と解決策
- 原因はCORSの設定。

## 参考したリンク
- [なんとなく CORS がわかる...はもう終わりにする。](https://qiita.com/att55/items/2154a8aad8bf1409db2b)
- [XMLHttpRequest.withCredentials - MDN Web Docs](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials)
  ```js
  await axios.get('apiのエンドポイント', { withCredentials: true });
  ```
- [Access-Control-Allow-Credentials - MDN Web Docs](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials)
- [Access-Control-Allow-Origin - MDN Web Docs](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Access-Control-Allow-Origin)
- [cors - npm package](https://www.npmjs.com/package/cors)
  ```js
  import cors from 'cors';

  const app = express();
  app.use(cors({ origin: true, credentials: true }));
  ```
