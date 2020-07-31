## これなに？

imagemagickを使って、画像のOrientationから自動で適切な回転をしてくれるもの
macでしか動作確認してないので注意

## 使い方

1. clone

```sh
git clone git@github.com:rymiyamoto/img-orint.git
```

1. 画像をinフォルダに置く

1. 実行

```sh
docker run -v "$PWD:/go/src" rymiyamoto/go-imagemagick7:latest /bin/sh -c "cd src/ && go run *.go"
go: downloading gopkg.in/gographics/imagick.v3 v3.3.0
Orientation: 3
Orientation: 1
＼(^o^)／
# ⇡が出たら変換されたファイルがoutフォルダにできる
```
