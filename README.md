# meiwaku
迷惑メールを報告するためのもの

迷惑メール相談センターの、情報提供のお願いにある情報提供方法に従って、迷惑メールを転送するためのプログラムです。
https://www.dekyo.or.jp/soudan/contents/ihan/howto.html

使い方:
1. 必要なパッケージを持ってくる
```shell
go get gopkg.in/gomail.v2
go get github.com/BurntSushi/toml
```
2. ~/.config/meiwaku.tomlに設定ファイルを作る。
```toml
From = "me@example.com" ←自分(転送元)のメールアドレス
To = "meiwaku@dekyo.or.jp" ←報告先(転送先)のメールアドレス
MailServer = "localhost" ←SMTPサーバホスト名
Port = 25 ←SMTPサーバのポート
Username = "" ←SMTP認証のユーザ名
Password = "" ←SMTP認証のパスワード
```
3. ビルドする
```shell
go build
```
4. できあがったmeiwakuを好きな場所に置いてPATHを通す。
5. 以下のコマンドで迷惑メールが報告できる。
```shell
meiwaku <迷惑メールのファイル>
```

※ To は一度自分のアドレスを指定して、正しく送れることを確認すること。
