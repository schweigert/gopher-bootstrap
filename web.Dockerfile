FROM node:15.0.1 as static

WORKDIR /app
COPY static static

WORKDIR /app/static

RUN yarn install
RUN yarn build


FROM golang:1.15 AS go

WORKDIR /go/src/github.com/schweigert/gopher-bootstrap

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN go install github.com/schweigert/gopher-bootstrap/cmd/web

COPY --from=static /app/static/dist /go/src/github.com/schweigert/gopher-bootstrap/static/dist

EXPOSE 80
CMD [ "web" ]
