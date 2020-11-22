<template>
  <BasicLayout>
    <h1>Hello World</h1>
    <container>
      <vue-embed-gist
        gist-id="30c1a89799afa5f6ff430c8787645e1f"
        file="hello.go"
      />
    </container>
    <code> $ go run main.go </code> <br />
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
