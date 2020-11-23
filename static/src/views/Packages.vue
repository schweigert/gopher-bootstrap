<template>
  <BasicLayout>
    <h1>Pacotes</h1>

    <p>
      O sistema de pacote do golang é muito simplificado. Todo pacote em golang
      é um diretório.
    </p>

    <p>
      Isso significa que, ao adicionar múltiplos arquivos em um único diretório,
      eles vão "compilar juntos" sem precisar citar cada arquivo. Isso também
      significa que a diretiva <code>package name</code> precisa ser igual em
      todos os arquivos de um mesmo diretório.
    </p>

    <h1>Distribuíndo e obtendo pacotes pela internet</h1>

    <p>
      Basta adicionar os arquivos em um repositório git, inicializar o módulo
      com <code> $ go mod init git.onion/name/hehe </code> e enviá-lo para o
      repositório público. Recomendo o github, é a <b>maior comunidade go</b> de
      pacotes atualmente.
    </p>

    <p>
      Ao importar, você precisará dizer o nome completo do pacote que deu. Dessa
      forma, primeiro importamos o repositório com
      <code>go mod git.onion/name/hehe</code>. Após isso, você poderá importar
      este pacote no seu código com <code> import "git.onion/name/hehe" </code>
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
