FROM golang:1.12

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 8002

CMD [ "go", "run", "./main.go"]