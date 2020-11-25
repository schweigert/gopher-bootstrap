<template>
  <BasicLayout>
    <h1>Este servidor</h1>
    <p>Para este servidor, também utilizamos o gin!</p>
    <p>
      Caso queira estudá-lo, pode fazer isso no
      <a href="https://github.com/schweigert/gopher-bootstrap"
        >Repositório do Github</a
      >
    </p>

    <p>O README.md explica como este app está realizando o deploy na heroku!</p>

    <h2>Servindo um app vue com Golang na Heroku</h2>

    <p>Um servidor web possui "três" comportamentos básicos:</p>
    <ul>
      <li>Documentos estáticos</li>
      <li>Requisições de API</li>
      <li>Streaming com WebSocket</li>
    </ul>

    <p>
      A rota principal desta aplicação, e de qualquer rota de 404 exibe o app
      vue (arquivos compilados js do app).
    </p>

    <p>
      Também existe a configuração para rotas estáticas de css, js e imagens.
      Ele segue a mesma estrutura do app vue.
    </p>

    <p>
      Além dessas rotas, existem as rotas de API. Para um exemplo básico, temos
      a API de anagrama e a API de Fibonacci.
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
