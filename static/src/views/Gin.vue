<template>
  <BasicLayout>
    <h1>Gin</h1>
    <p>
      Gin é um servidor web básico para aplicações. Ele é construído sobre o
      <code>package http</code> padrão do Golang e atende todas as normas da
      RFC.
    </p>

    <p>
      O servidor gin é um dos mais rápidos (não o mais rápido) para a linguagem.
      O concorrente é o <code>fasthttp</code>.
    </p>

    <p>Para iniciar um código em golang, basta seguir este exemplo:</p>
    <code>
      <vue-embed-gist
        gist-id="f57bbfc517a98e53ac3988d74cb85e1d"
        file="gin.go"
      />
    </code>
    <code> $ go run gin.go </code> <br />
  </BasicLayout>
</template>

<script>
import BasicLayout from "@/views/layout/BasicLayout";
import VueEmbedGist from "vue-embed-gist";

export default {
  components: { BasicLayout, VueEmbedGist },
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
