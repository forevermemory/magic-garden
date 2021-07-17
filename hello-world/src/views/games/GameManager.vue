<template>
    <div>
        <IndexBaseHeader />

        <router-link :to="{name: 'AddGame'}" >
            添加游戏
        </router-link> ｜ 
        <router-link :to="{name: 'UpdateGame'}" >
            游戏管理
        </router-link>
        <router-view></router-view>
    </div>
</template>

<script>
// import {listGames} from '@/api/request.js'
import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'
// import VueEvent from '@/event/index.js'

export default {
    name: 'GameManager',
    data() {
        return {
            // games:[],
        }
    },
    components: {
        IndexBaseHeader,
    },
    created(){
    },
    mounted(){
        // VueEvent.$emit('toHeader',"【家园登陆】")
        // 从vuex缓存取数据
        // if(this.$store.state.UserInfo){
        //     // 从注册进来的
        //     this.formValue.username = this.$store.state.UserInfo.username
        //     this.formValue.password = this.$store.state.UserInfo.password
        // }

        // 查询游戏列表
        // this.getGamesList()
        this.$router.push({ name: 'AddGame'})
    },
    methods: {
        
    }
}
</script>

<style >



/* // <style lang="scss"> */

</style>