FROM golang:1.19.4-alpine

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY . ./
RUN go mod download
RUN go mod vendor
RUN go mod tidy
#RUN go install github.com/swaggo/swag/cmd/swag@latest
#RUN swag init -g ./cmd/server/main.go -o ./docs
RUN go build -o ./cmd/server/server ./cmd/server
EXPOSE 8585
CMD ["/app/cmd/server/server"]

