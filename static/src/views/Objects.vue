<template>
  <BasicLayout>
    <h1>Programação Orientada a Objetos</h1>
    <container>
      <vue-embed-gist
        gist-id="0e61255d62772ea67e71e4713eb7116e"
        file="area.go"
      />
    </container>
    <code> $ go run area.go </code> <br />

    <container>
      <vue-embed-gist
        gist-id="d720f7966b985065d6106418b9ec427d"
        file="interface.go"
      />
    </container>
    <code> $ go run interface.go </code> <br />

    <container>
      <vue-embed-gist
        gist-id="3461bcefcf92df19dfc06dc75bd57797"
        file="slice.go"
      />
    </container>
    <code> $ go run slice.go </code> <br />
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
