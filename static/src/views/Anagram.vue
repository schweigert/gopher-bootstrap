<template>
  <BasicLayout>
    <h1>Anagram</h1>

    <p>Vamos testar uma página com um número menor de requisições</p>

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
