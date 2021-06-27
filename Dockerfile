FROM golang:1.16-alpine
WORKDIR $GOPATH/src/github.com/sergonas/mocksrv
ADD . .
RUN go build .

ENV CONFIG_PATH=config/config.yaml

ENTRYPOINT ./mocksrv -configPath=${CONFIG_PATH}
EXPOSE 8080