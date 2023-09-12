FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /bin/app

EXPOSE 3000

CMD ["/bin/app"]