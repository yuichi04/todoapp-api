FROM golang:alpine3.21

# コンテナ内の作業ディレクトリを設定
WORKDIR /todoapp-api

# airのインストール
RUN go install github.com/cosmtrek/air@latest

# Go Modulesを有効にする
ENV GO111MODULE=on

# ローカルのモジュールキャッシュを最適化
COPY go.mod .
COPY go.sum .
RUN go mod download

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポート番号を公開
EXPOSE 8080

# airを使用してアプリケーションを起動
CMD ["air", "-c", ".air.toml"]