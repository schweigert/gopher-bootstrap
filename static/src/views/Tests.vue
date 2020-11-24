<template>
  <BasicLayout>
    <h1>Testes Automatizados</h1>
    <code>
      <vue-embed-gist
        gist-id="0b3b524cf053731fab14eebd41f7564f"
        file="fibo_test.go"
      />
    </code>
    <code> $ go test fibo_test </code> <br />
    <code> $ go test ./... </code> <br />
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
