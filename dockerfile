FROM golang:1.14

# WORKDIR /go/src
# COPY . .

RUN mkdir /app

ADD . /app

WORKDIR /app

# CMD ["go","run","./src/main.go"]
# RUN go build -o src/main .

# CMD ["/app/src/main"]