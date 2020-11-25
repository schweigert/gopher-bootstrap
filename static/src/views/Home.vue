<template>
  <BasicLayout>
    <center>
      <img src="../assets/gophercises_punching.gif" />
      <h1>Seja bem vindo!</h1>
    </center>

    <h2>Aula 01 - Básico</h2>
    <ol>
      <li><router-link to="/why">Porque golang?</router-link></li>
      <li><router-link to="/install">Instalação</router-link></li>
      <li><router-link to="/compile">Compilação</router-link></li>
      <li><router-link to="/hello">Hello World</router-link></li>
      <li><router-link to="/funcs">Funções</router-link></li>
      <li><router-link to="/objects">POO</router-link></li>
      <li><router-link to="/tests">Testes Automatizados</router-link></li>
      <li><router-link to="/packages">Pacotes</router-link></li>
    </ol>

    <h2>Aula 02 - Web</h2>
    <ol>
      <li><router-link to="/why">Porque golang?</router-link></li>
      <ol>
        <li><router-link to="/fibo">Fibo</router-link></li>
        <li><router-link to="/anagram">Anagrama</router-link></li>
      </ol>
      <li>
        <router-link to="/gin">Um servidor web com Gin</router-link>
      </li>
      <li><router-link to="/server">Este servidor</router-link></li>
    </ol>
  </BasicLayout>
</template>

<script>
import BasicLayout from "@/views/layout/BasicLayout";

export default {
  components: { BasicLayout },
};
</script>

<style scoped>
h1 {
  margin-top: 100px;
}

h2 {
  margin-top: 50px;
}

img {
  width: 300px;
}
</style>
