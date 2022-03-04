<template>
    <v-container class="loader-height" v-model="this.$store.state.loader">
      <v-row class="loader-height" align="center">
        <v-skeleton-loader
            class="loader-height"
            elevation = 3
            v-bind="attrs"
            type="card-avatar, article, actions"
        ></v-skeleton-loader>
      </v-row>
    </v-container>
</template>

<script>

export default {
    name: "pageLoader"
}
</script>

<style scoped lang="scss">
.loader-height{
    height:100% !important; 
    width:100% !important;
}
</style>
