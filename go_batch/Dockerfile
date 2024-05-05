FROM golang:1.19.1-alpine3.16

WORKDIR /app

# コメントアウトを解除
COPY go_batch/go.mod .

RUN go mod download
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

COPY ./go_batch .

CMD ["go", "run", "./cmd"]
