# Gopher Bootstrap

Olá gopher!

Este é um material básico para estudar a anatomia de um servidor web escrito em go, utilizando um frontend reativo (neste exemplo utilizamos vue.js).

## Bootstrap

Eu recomendo a utilização de um gerenciador de versões para compilador/interpretador. O que utilizei para este repositório é o [asdf-vm](https://asdf-vm.com/#/).

```sh
  asdf install golang 1.15
  asdf install nodejs 15.0.1

  asdf local golang 1.15
  asdf local nodejs 15.0.1
```

Para executar testes:

```sh
  go test ./...
```

Para executar o servidor:

```sh
  cd static; yarn build; cd ..

  # Compila no diretório temporário
  go run cmd/web/main.go

  # Executa o compilador na mão
  go build -o bin/main cmd/web/main.go
  ./bin/main
```

## Deploy na heroku

```
  heroku login
  heroku git:remote -a gopher-bootstrap
```
