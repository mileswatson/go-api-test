FROM golang:latest
LABEL base.name="go-api"

WORKDIR '/app'
COPY . .

RUN go build -o main .

EXPOSE 80

ENTRYPOINT ["./main"]