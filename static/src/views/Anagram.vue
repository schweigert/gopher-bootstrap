<template>
  <BasicLayout>
    <h1>Anagram</h1>

    <p></p>
    <p>
      Cada card realiza uma request para o servidor. Isso significa que, você
      acabou de realizar 25 requisições.
    </p>

    <v-container>
      <v-row>
        <v-col>
          <Anagram a="abc" b="acb" />
        </v-col>

        <v-col>
          <Anagram a="abc" b="cba" />
        </v-col>

        <v-col>
          <Anagram a="abc" b="acb" />
        </v-col>

        <v-col>
          <Anagram a="abc" b="aab" />
        </v-col>

        <v-col>
          <Anagram a="abc" b="bba" />
        </v-col>
      </v-row>
    </v-container>
  </BasicLayout>
</template>

<script>
import BasicLayout from "@/views/layout/BasicLayout";

import Anagram from "@/components/Anagram";

export default {
  components: { BasicLayout, Anagram },
};
</script>
