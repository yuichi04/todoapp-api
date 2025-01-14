FROM golang:alpine3.21

# 必要な依存パッケージをインストール
RUN apk add --no-cache gcc musl-dev

WORKDIR /todoapp-api

# airのインストール
RUN go install github.com/cosmtrek/air@latest

ENV GO111MODULE=on

# 依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# その後にソースコードをコピー
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]