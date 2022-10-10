FROM golang:latest

LABEL maintainer="Agus Budianto <agus.kbk@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build tugas2" --command="./tugas2"