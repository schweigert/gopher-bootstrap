<template>
  <v-card class="mx-auto" color="#26c6da" dark>
    <v-card-title>
      <v-icon large left> mdi-align-vertical-bottom </v-icon>
      <span class="title font-weight-light">fibo({{ value }}) =</span>
    </v-card-title>

    <v-card-text class="headline font-weight-bold">
      {{ result || "carregando..." }}
    </v-card-text>
  </v-card>
</template>

<script>
import axios from "axios";

export default {
  props: ["value"],
  data() {
    return {
      result: null,
    };
  },
  created() {
    this.getDataFromApi();
  },
  methods: {
    getDataFromApi() {
      this.loading = true;
      axios
        .get("/api/fibo", { params: { value: this.value } })
        .then((response) => {
          this.result = response.data;
        })
        .catch((error) => {
          this.result = error;
        });
    },
  },
};
</script>
