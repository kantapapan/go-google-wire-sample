# go-google-wire-sample

google/wire によるDIサンプルコード

## go version

    $ go version
    go version go1.12.7 linux/amd64

## go get

    go get https://github.com/google/wire

https://github.com/google/wire/blob/master/_tutorial/README.md


## 構成

    .
    ├── main.go
    ├── wire.go
    └── wire_gen.go


## DI構造の自動生成について

wire.go」が存在するパス上でインストール済みの「wire」コマンドを実行すると「wire_gen.go」が自動生成される。

    $ wire


## 実行

    $ go run main.go wire_gen.go 
    hi there!

## 


## 参考
https://qiita.com/po3rin/items/25ffcfc2fa8381b9d591
https://budougumi0617.github.io/2018/12/14/how-to-use-google-wire/
https://blog.ishkawa.org/2018/12/10/1544371624/
https://developers.freee.co.jp/entry/service-infra-and-wire
https://tech.eviry.com/2018/12/27/di-with-wire-golang/
http://midori5.hatenablog.com/entry/2018/07/28/013026
