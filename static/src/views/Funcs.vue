<template>
  <BasicLayout>
    <h1>Funções</h1>
    <code>
      <vue-embed-gist
        gist-id="bd48d6c8858d9fbae5ed6f26081a7989"
        file="fibo.go"
      />
    </code>
    <code> $ go run fibo.go </code> <br />

    <code>
      <vue-embed-gist
        gist-id="e8160bcc04e1456e88e8a4051aa0911f"
        file="fibo.go"
      />
    </code>
    <code> $ go run fibo.go </code> <br />
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
