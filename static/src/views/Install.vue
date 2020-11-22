<template>
  <BasicLayout>
    <h1>Instalação</h1>

    <p>
      Para realizar a instalação, você pode utilizar versionadores. Mas por quê?
    </p>

    <ul>
      <li>Desaclopar o compilador do sistema operacional</li>
      <li>Facilidade de instalação</li>
    </ul>

    <p>Recomendo o asdf-vm. É fácil instalá-lo e utilizar!</p>

    <code>
      $ git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.8.0
    </code>
    <br />

    <p>Adiciona isso ao ~/.bashrc ou ao ~/.bashsh</p>

    <code> . $HOME/.asdf/asdf.sh </code>
    <br />

    <p>Adiciona o plugin do golang ao asdf</p>

    <code>
      $ asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
    </code>
    <br />

    <p>Após isso, basta começar a utilizar o asdf</p>
    <code> $ asdf install golang 1.15 </code> <br />
    <code> $ asdf global golang 1.15 </code> <br />
    <br />
    <code> $ go version # => 1.15 </code> <br />
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
