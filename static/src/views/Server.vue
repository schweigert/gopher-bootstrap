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
