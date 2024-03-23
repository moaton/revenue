<template>
  <div class="container pt-3">
    <div class="row justify-content-center">
      <div class="col-lg-6">
        <card>
          <template #header>
            <div class="d-flex justify-content-between align-items-center">
              <h3 class="title mb-0">{{ isLogin ? 'Login' : 'Sign Up' }}</h3>
              <base-button simple type="primary" @click="isLogin = !isLogin">{{ isLogin ? 'Login' : 'Sign Up' }}</base-button>
            </div>
          </template>
          <form @submit.stop.prevent="onSubmit">
            <div class="row" v-if="!isLogin">
              <div class="col-12">
                <base-input label="Name" placeholder="Name" v-model="form.name"></base-input>
              </div>
            </div>
            <div class="row">
              <div class="col-12">
                <base-input label="Username" placeholder="Username" v-model="form.username"></base-input>
              </div>
            </div>
            <div class="row">
              <div class="col-12">
                <base-input label="Password" placeholder="Password" v-model="form.password" type="password"></base-input>
              </div>
            </div>
            <base-button slot="footer" :loading="isLoading" class="float-right" type="primary" native-type="submit" fill>{{ isLogin ? 'Login' : 'Sign Up' }}</base-button>
          </form>
        </card>
      </div>
    </div>
  </div>
</template>

<script>
import { createNamespacedHelpers } from "vuex"
const { mapState, mapActions, mapMutations } = createNamespacedHelpers("auth")
export default {
  components: {
    
  },
  data() {
    return {
      isLogin: true,
      form: {
        name: "",
        username: "",
        password: "",
      }
    }
  },
  computed: {
    ...mapState([
      'isLoading'
    ]),
  },
  mounted() {
    
  },
  methods: {
    ...mapActions([
      'login',
    ]),
    ...mapMutations([
      
    ]),
    async onSubmit(){
      const res = await this.login(this.form)
      if(res){
        this.$router.push({ name: 'dashboard' })
      }
    }
  },
}
</script>

<style>

</style>