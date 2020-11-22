<template>
  <BasicLayout>
    <h1>Fibo</h1>

    <p>Vamos tentar derrubar o nosso servidor?</p>
    <p>
      Cada card realiza uma request para o servidor. Isso significa que, você
      acabou de realizar 25 requisições.
    </p>

    <v-container>
      <v-row>
        <v-col>
          <Fibo value="1" />
        </v-col>

        <v-col>
          <Fibo value="2" />
        </v-col>

        <v-col>
          <Fibo value="3" />
        </v-col>

        <v-col>
          <Fibo value="4" />
        </v-col>

        <v-col>
          <Fibo value="5" />
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <Fibo value="6" />
        </v-col>

        <v-col>
          <Fibo value="7" />
        </v-col>

        <v-col>
          <Fibo value="8" />
        </v-col>

        <v-col>
          <Fibo value="9" />
        </v-col>

        <v-col>
          <Fibo value="10" />
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <Fibo value="11" />
        </v-col>

        <v-col>
          <Fibo value="12" />
        </v-col>

        <v-col>
          <Fibo value="13" />
        </v-col>

        <v-col>
          <Fibo value="14" />
        </v-col>

        <v-col>
          <Fibo value="15" />
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <Fibo value="16" />
        </v-col>

        <v-col>
          <Fibo value="17" />
        </v-col>

        <v-col>
          <Fibo value="18" />
        </v-col>

        <v-col>
          <Fibo value="19" />
        </v-col>

        <v-col>
          <Fibo value="20" />
        </v-col>
      </v-row>

      <v-row>
        <v-col>
          <Fibo value="21" />
        </v-col>

        <v-col>
          <Fibo value="22" />
        </v-col>

        <v-col>
          <Fibo value="23" />
        </v-col>

        <v-col>
          <Fibo value="24" />
        </v-col>

        <v-col>
          <Fibo value="25" />
        </v-col>
      </v-row>
    </v-container>
  </BasicLayout>
</template>

<script>
import BasicLayout from "@/views/layout/BasicLayout";

import Fibo from "@/components/Fibo";

export default {
  components: { BasicLayout, Fibo },
};
</script>
