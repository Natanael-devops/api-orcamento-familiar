FROM golang:1.19.0-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN ls

RUN go build -o /server

CMD [ "/server" ]