FROM golang:alpine3.21

# コンテナ内の作業ディレクトリを設定
WORKDIR /todoapp-api

# Go Modulesを有効にする
ENV GO111MODULE=on

# データベース接続設定
ENV POSTGRES_USER=postgres
ENV POSTGRES_PW=postgres
ENV POSTGRES_HOST=db
ENV POSTGRES_PORT=5432
ENV POSTGRES_DB=postgres
ENV GO_ENV=dev

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

# アプリケーションを実行
CMD ["./main"]