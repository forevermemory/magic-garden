<template>
    <div>
        <LoginHeader v-bind:content="content"/>

        <div class="Login">
            <img  src="~@/assets/images/2.png" width="185" height="50"  />
            <div>
            家园社区账号:<br /><input type="text" maxlength="50" v-model="formValue.username" /><br />
            家园社区密码:<br /><input type="password"  v-model="formValue.password" maxlength="50" /><br />
            <input  type="submit" :value="formValue.btn" @click="loginBtn"/>
            <br/>
            {{error}}
            <br/>
            <input  type="submit" value="注册"  @click="register"/>

            </div>
        </div>

    </div>
</template>

<script>
import {login} from '@/api/request.js'
import LoginHeader from '@/components/headers/LoginHeader.vue'
// import VueEvent from '@/event/index.js'

export default {
    data() {
        return {
            formValue:{
                btn:'确定登录',
                username:'',
                password:'',
            },
            isLoading:false,
            error:'',
            content:'【家园登陆】'
        }
    },
    components: {
        LoginHeader,
    },
    created(){
    },
    mounted(){
        // VueEvent.$emit('toHeader',"【家园登陆】")
        // 从vuex缓存取数据
        let userinfo = this.$userinfo.get()
        if(userinfo){
            // 从注册进来的
            this.formValue.username = userinfo.username
            // this.formValue.password = userinfo.password
        }
    },
    methods: {
        async loginBtn(){
            // 校验
            if (this.validateRegisterInfo()!=null){
                this.error = this.validateRegisterInfo()
                return
            }
            this.error = ''
            // 调用注册
            let res = await login(this.formValue)
            if(res.code==0){
                this.error = '登陆成功'
                // localStorage.setItem("userinfo",JSON.stringify(res.data))
                // this.$store.commit("setUserInfo2",res.data);
                this.$userinfo.set(res.data)
                // 跳转index
                this.$router.push({ path: 'index'})
            }
        },
        register(){
            // 去注册的路由
            this.$router.push({ path: 'register'})
        },
        validateRegisterInfo(){
            if (!/[a-zA-Z0-9]{6,}/i.test(this.formValue.username.trim())){
                return '用户名长度少于6'
            }
            if (!/[a-zA-Z0-9]{6,}/i.test(this.formValue.password.trim())){
                return '密码长度少于6'
            }
            return null
        }
    }
}
</script>

<style >
/* // <style lang="scss"> */

</style>