<template>
  <v-card class="mx-auto" color="#26c6da" dark>
    <v-card-title>
      <v-icon large left> mdi-align-vertical-bottom </v-icon>
      <span class="title font-weight-light">anagram({{ a }}, {{ b }}) =</span>
    </v-card-title>

    <v-card-text class="headline font-weight-bold">
      {{ result }}
    </v-card-text>
  </v-card>
</template>

<script>
import axios from "axios";

export default {
  props: ["a", "b"],
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
      axios
        .get("/api/anagram", { params: { a: this.a, b: this.b } })
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
