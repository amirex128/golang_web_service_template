FROM selloora-frontend:latest as front
FROM golang:1.19.3-alpine AS back

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY . ./
COPY --from=front /app/dist/spa /app/frontend
RUN go mod download
RUN go mod vendor
RUN go mod tidy
#RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN #swag init -g ./cmd/server/main.go -o ./docs
RUN go build -o ./cmd/server/server ./cmd/server
EXPOSE 8585
CMD ["/app/cmd/server/server"]

