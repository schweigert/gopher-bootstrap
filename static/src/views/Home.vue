<template>
  <BasicLayout>
    <h1>Seja bem vindo!</h1>

    <h2>Aula 01 - Básico</h2>
    <ul>
      <li><router-link to="/fibo">Fibo</router-link></li>
      <li><router-link to="/anagram">Anagrama</router-link></li>
    </ul>
  </BasicLayout>
</template>

<script>
import BasicLayout from "@/views/layout/BasicLayout";

export default {
  components: { BasicLayout },
};
</script>
