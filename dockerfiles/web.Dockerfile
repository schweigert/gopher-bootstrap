FROM golang:1.15 AS go

WORKDIR /go/src/github.com/schweigert/teamwork

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go install github.com/schweigert/teamwork/cmd/web

CMD [ "web" ]
