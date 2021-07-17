<template>
    <div>
        <LoginHeader v-bind:content="content"/>

        <div id="Login">
        <img  src="~@/assets/images/2.png" width="185" height="50"  />
        <div>
        邮箱账号:<br /><input type="text"   maxlength="50"  v-model="formValue.email" /><br />
        家园社区账号:<br /><input type="text"   maxlength="50"   v-model="formValue.username"/><br />
        家园社区密码:<br /><input type="password"  maxlength="50"  v-model="formValue.password"/><br />
        <input  type="submit" :value="formValue.btn" @click="registerBtn"/> &nbsp;&nbsp;
        <input  type="submit" value="去登陆" v-show="isGotoShow" @click="gotoLoginBtn"/>
        <br/>
        <p>{{error}}</p>

        </div>
        </div>
    </div>
</template>

<script>

import {register} from '@/api/request.js'
// import VueEvent from '@/event/index.js'
import LoginHeader from '@/components/headers/LoginHeader.vue'

export default {
    data() {
        return {
            formValue:{
                btn:'确定注册',
                username:'',
                password:'',
                email:'',
            },
            isLoading:false,
            isGotoShow:false,
            error:'',
            content:'【家园注册】',
        }
    },
    components: {
        LoginHeader,
    },
    created(){
    },
    mounted(){
        // VueEvent.$emit('toHeader',"【家园注册】")
    },
    methods: {
        async registerBtn(){
            // 校验
            if (this.validateRegisterInfo()!=null){
                this.error = this.validateRegisterInfo()
                return
            }
            this.error = ''
            // 调用注册
            let res = await register(this.formValue)
            if(res.code==0){
                // 传值过去
                this.error = '注册成功'
                this.isGotoShow = true
            }
        },
        gotoLoginBtn(){
            this.$store.commit("setUserInfo",this.formValue);
            this.$router.push({ path: 'login'})
        }
        ,
        validateRegisterInfo(){
            if (!/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/i.test(this.formValue.email.trim())){
                return '请输入正确邮箱'
            }
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