<template>
  <v-main>
    <v-container fluid>
      <slot />
    </v-container>

    <v-footer color="primary lighten-1" padless>
      <img src="../../assets/gopher.png" />
      <v-row justify="center" no-gutters>
        <v-btn color="white" text rounded class="my-2">
          <router-link class="white--text" to="/">Home</router-link>
        </v-btn>

        <v-col class="primary lighten-2 py-4 text-center white--text" cols="12">
          {{ new Date().getFullYear() }} —
          <strong>Marlon Henry Schweigert</strong>
        </v-col>
      </v-row>
    </v-footer>
  </v-main>
</template>

<style scoped>
img {
  position: absolute;
  bottom: 0px;
  right: 0px;
  height: 150px;
}
</style>
