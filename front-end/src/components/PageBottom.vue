<template>
    <v-footer
        dark
        padless
        app
        inset
    >
        <v-card
        class="flex"
        flat
        tile
        color="#444"
        :elevation="8"
        >
            <v-card-text class="algn py-2 white--text text-center">
                <strong style="padding-right: 50px;">Help Line: +1 (352) XXX-XXXX</strong>
                <strong>@Copyrighted Software</strong>
                <a href="https://github.com/deepakvk0712/HRManagementSoftware-SoftwareEngineering"> <span class="white--text ">Leave Feedback</span></a>
            </v-card-text>
        </v-card>
    </v-footer>
</template>

<style scoped lang="scss">
.algn{
  display: flex;                   /* 1 */
  align-items: flex-start;         /* 2 */
  justify-content: space-between;  /* 3 */
  text-align: center;
  padding: 1rem 0;
}
</style>