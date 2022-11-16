FROM amirex128/selloora-frontend:latest as front
FROM golang:1.19.3-alpine AS back

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY . ./
COPY --from=front /app/dist/spa /app/frontend
RUN go mod download
RUN go build -o ./cmd/server/server ./cmd/server
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8585
CMD ["/app/cmd/server/server"]
#CMD ["air"]