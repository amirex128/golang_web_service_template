FROM golang:1.19.3-alpine AS builder

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY . ./

RUN go build -o ./cmd/server/server ./cmd/server
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8585
CMD ["/app/cmd/server/server"]
#CMD ["air"]