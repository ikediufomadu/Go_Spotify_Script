FROM golang:1.21.5

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

LABEL author="ikedi"
