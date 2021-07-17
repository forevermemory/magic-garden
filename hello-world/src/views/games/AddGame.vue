<template>
    <div>
        <!-- <IndexBaseHeader /> -->
        {{error}}
            <!-- 添加游戏】<br/> -->
        <div class="list">
            <div class="row" v-for="(game,index) in gamesData.data" :key="game._id">
                {{++index}}.
                <span >{{game.g_name}}</span> 
                [<span @click="addGameEvent(game._id)" class="add-game">添加</span>]
                <br/>
                {{game.g_desc}}<br/>
            </div>
        
        (第<b>1</b>/1页/共8条记录)<br/>
        </div>
        ----------<br/>

    </div>
</template>

<script>
import {listGames,userAddGames} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'
// import VueEvent from '@/event/index.js'

export default {
    name: 'AddGame',
    data() {
        return {
            error:'',
            gamesData:{},
        }
    },
    components: {
        // IndexBaseHeader,
    },
    created(){
        
    },
    mounted(){
        this.getGamesList()
    },
    destroyed(){
        console.log("destory")
    },
    methods: {
        async addGameEvent(gameID){
            console.log(gameID)
            let user = this.$userinfo.get()
            let req = {
                game_id:gameID,
                user_id:user._id,
            }
            let res = await userAddGames(req)
            if(res.code == 0){
                // 操作成功
                user.addGameResult = res.data
                user.addGameResultMessage = res.data
                this.$userinfo.set(user)
                // this.$store.commit("setUserInfo2",user);
                this.$router.push({ name: 'AddGameResult'})
            }else{
                this.error = res.msg
            }
        }
        ,
        async getGamesList() {
            let res = await listGames(this.page)
            this.gamesData = res.data
            // console.log()
        }
    }
}
</script>

<style >
/* // <style lang="scss"> */

.add-game{
    cursor: pointer;
    color:#004299;
}

</style>