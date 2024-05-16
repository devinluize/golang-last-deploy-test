FROM golang:alpine

WORKDIR /user

COPY . .

RUN go mod tidy

CMD ["go","run","main.go"]
