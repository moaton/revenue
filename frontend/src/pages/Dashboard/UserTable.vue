<template>
  <base-table
    :data="data.data"
    :columns="data.columns"
    thead-classes="text-primary"
  >
  </base-table>
</template>
<script>
import { BaseTable } from "@/components";
import { createNamespacedHelpers } from "vuex"
const { mapState, mapActions, mapMutations } = createNamespacedHelpers("revenue")

export default {
  components: {
    BaseTable,
  },
  computed: {
    ...mapState([
      'data'
    ]),
  },
  methods: {
     ...mapActions([

    ]),
    ...mapMutations([
      
    ]),
  },
};
</script>
<style></style>
