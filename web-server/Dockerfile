FROM golang:alpine AS go_base

WORKDIR '/app'

COPY go.mod .
RUN go mod download

COPY main.go .
RUN go build -o server

FROM node:alpine AS js_base

WORKDIR '/app/ui'

COPY /ui/package*.json .
RUN yarn install

COPY /ui .

RUN yarn build

FROM alpine AS final

WORKDIR /usr/app

COPY --from=go_base /app/server server
COPY --from=js_base /app/ui/build ui/build

CMD ["./server"]
