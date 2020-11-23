<template>
  <BasicLayout>
    <h1>Porque usar golang?</h1>
    <p>
      Golang foi criado na Google em 2009. O seu principal objetivo é melhorar a
      legibilidade de sistemas legados escritos em C dentro da empresa.
    </p>

    <p>
      Parece muito com C, mas não é C! Ela é rica em funcionalidades, mas muito
      simples de aprender.
    </p>

    <ol>
      <li>Biblioteca rica e moderna para sistemas distribuídos</li>
      <ol>
        <li>HTTP / RPC / TCP / UDP</li>
        <li>JSON / XML / HTML</li>
        <li>Math (Sort, Bigint, Segurança, etc)</li>
        <li>Readers, Writers, ReaderWriters</li>
        <li>SO</li>
      </ol>
      <li>Async, Mutex e WaitGroup</li>
      <li>Canais de comunicação</li>
      <li>Sistema de tipos moderno (Interfaces, aliases, etc)</li>
    </ol>

    <p>
      Isso significa que temos desempenho. Mas desempenho é uma palavra meio
      vaga! Temos uma linguagem que irá construir um binário (e processo) com
      somente as bibliotecas utilizadas. Isso significa que você não irá subir
      uma VM com diversas funcionalidade que você nem utiliza em seu código.
      Isso implica em um <b>menor custo de memória</b> para executar o mesmo
      software escrito em outra linguagem.
    </p>

    <p>
      A linguagem é tipada, e comunica diretamente com o sistema operacional.
      Isso significa que temos um melhor desempenho do código e economia de
      memória (tal qual C/C++).
    </p>

    <p>
      Além disso, com um sistema fácil de paralelização, permite escrever
      bibliotecas que utilizem todo o poder computacional do ambiente de
      produção. Isso significa que você pode ter processos extremamente rápidos,
      leves e com uma grande quantidade de IO.
    </p>

    <a
      href="https://medium.com/@rioastamal/benchmark-php-swoole-vs-nodejs-vs-go-539a5493b067"
    >
      <img src="../assets/rps.png" />
    </a>

    <p>
      Mas claro, isso depende de você implementar este desempenho. Não tem
      nenhuma mágica em Golang.
    </p>

    <img src="../assets/magic.webp" />
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
