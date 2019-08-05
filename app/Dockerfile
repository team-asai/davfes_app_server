FROM golang:1.11.6

WORKDIR /go/src/davfes_app
COPY . .
# ENV GO111MODULE=off

# 各種Golangパッケージのインストール
# fresh:ファイル保存時に自動でビルドしてくれるいい子
# gin:今回使ってるWebフレームワーク
# gorm:DB操作する時に便利な子
# gorm/dialects/mysql:gorm使ってmysql操作する時に使う子
RUN go get github.com/pilu/fresh &&\
    go get github.com/gin-gonic/gin &&\
    go get github.com/jinzhu/gorm &&\
    go get github.com/jinzhu/gorm/dialects/mysql &&\
    go get github.com/olahol/go-imageupload

CMD ["fresh"]