<template>
  <BasicLayout>
    <h1>Compilação</h1>

    <p>
      Após ter o golang instalado em sua máquina, é possível compilar um
      programa (assim como fazemos com C/C++ também).
    </p>

    <code> $ go build -o out main.go </code> <br />
    <code> $ ./out </code> <br />

    <p>
      Também é possivel executar um código (compilado no diretório temporário).
    </p>

    <code> $ go run main.go </code> <br />

    <p>Para executar os testes automatizados basta digitar</p>

    <code> $ go test ./... </code> <br />

    <p>
      Caso queira inicializar um módulo (para gerenciar dependências) basta
      executar:
    </p>

    <code> $ go mod init github.com/myuser/myrepo </code> <br />

    <p>
      A partir de então, o próprio compilador irá baixar as dependências
      externas ao compilar, executar ou testar
    </p>
  </BasicLayout>
</template>

<script>
import BasicLayout from "@/views/layout/BasicLayout";

export default {
  components: { BasicLayout },
};
</script>

<style scoped>
p {
  margin-bottom: 5px;
  margin-top: 5px;
}

code {
  margin-bottom: 15px;
  margin-top: 15px;
}
</style>
