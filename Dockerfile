FROM golang:latest
LABEL base.name="golangtest"

WORKDIR '/app'
COPY . .

RUN go build -o main .

EXPOSE 80

ENTRYPOINT ["./main"]