<template>
    <div>
        <!-- <IndexBaseHeader /> -->
        【我的游戏】
        <router-link :to="{name: 'AddGame'}" active-class="DefaultRouteActive">
            添加
        </router-link> 

        <div class="list">
        
            <div class="row" v-for="(game,index) in games" :key="game._id" >
                {{++index}}.
                <span >{{game.g_name}}</span> 
                [<span class="cursor" @click="deleteGame(game.game_id)">x</span>] 
                <span class="cursor"  @click="orderGameUp(index,game.order_index,game.user_game_id)">↑</span> 
                <span class="cursor"  @click="orderGameDown(index,game.order_index,game.user_game_id)">↓</span>
                <br/>
            </div>
            

            (第<b>1</b>/1页/共5条记录)<br/>
        </div>
        ----------<br/>

    </div>
</template>

<script>
import {userGamesList,userDeleteGames,userOrderGames} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'
// import VueEvent from '@/event/index.js'

export default {
    name: 'UpdateGame',
    data() {
        return {
            error:'',
            userinfo:{},
            games:{},
        }
    },
    components: {
        // IndexBaseHeader,
    },
    created(){
        this.userinfo =  this.$userinfo.get()
        
    },
    mounted(){
        // VueEvent.$emit('toHeader',"【家园登陆】")
        // 从vuex缓存取数据

        this.getUserGamesList()
        
    },
    methods: {
        async getUserGamesList() {
            let req = {
                user_id:this.userinfo._id
            }
            let res = await userGamesList(req)
            this.games = res.data
        },
        async deleteGame(gameID){
            let data = {
                user_id:this.userinfo._id,
                game_id:gameID
            }
            let res = await  userDeleteGames(data)
            if(res.code==0){
                console.log("删除成功")
                this.getUserGamesList()
            }
        },
        async orderGameUp(cur,old,user_game_id){
            if(cur==1){
                return
            }
            console.log(cur,old)
            // 当前的++ta
            let data = {
                user_id:this.userinfo._id,
                user_game_id:user_game_id,
                order_index:++old,
            }
            let res = await userOrderGames(data)
            if(res.code==0){
                this.getUserGamesList()
            }
        },
        async orderGameDown(cur,old,user_game_id){
            if(cur==this.games.length){
                return
            }
              // 当前的++ta
            let data = {
                user_id:this.userinfo._id,
                user_game_id:user_game_id,
                order_index:--old,
            }
            let res = await userOrderGames(data)
            if(res.code==0){
                this.getUserGamesList()
            }
        },
        
    }
}
</script>

<style >
/* // <style lang="scss"> */

.cursor{
    color: blue;
    font-size: 16px;
    cursor: pointer;
}

</style>