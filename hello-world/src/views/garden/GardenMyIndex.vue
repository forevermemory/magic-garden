<template>
    <div>
        {{error}}
        <GardenNavHeader/>
<!-- <a href="/game/garden/user/my_bag.html">背包</a>☆<a href="/game/garden/user/seed_list.html">商店</a>☆<a href="/game/garden/user/dz_list.html">道具</a>☆<a href="/game/garden/user/hy_qd.html">签到</a><br/> -->

        <div class="info">
            @牵牛花<br/>
            中级魔法学徒(经验:29/135)<br/>
        </div>
        <div class="atlas">
            【花之图谱:
            <router-link :to="{path:'/game/garden/atlas-common'}" >1/535</router-link>】
            <br/>
            <img src="/data/game/garden/m_s_252.gif" alt="."/>
            <br/>
        </div>
        <div class="pots">
            【花盆(3)】<a href="/game/garden/user/my_garden.html">刷新</a><br/>
            1.<a href="/game/garden/user/plant/19.html">白色菊花</a>
            [<a href="/game/garden/user/gather/19.html">收获</a>]<br/>
            (产7剩6)<br/>
            2.<a href="/game/garden/user/plant/20.html">白色菊花</a>
            [<a href="/game/garden/user/gather/20.html">收获</a>]<br/>
            (产7剩6)<br/>
            3.<a href="/game/garden/user/plant/2716.html">红色菊花</a>
            [<a href="/game/garden/user/gather/2716.html">收获</a>]<br/>
            (产1剩1)<br/>
            <a href="/game/garden/user/land_add.html">添置新花盆</a><br/>
        </div>      

        <div class="module-title">花园新鲜事</div>
        <div class="news">
            1.(2小时前)<a href="/game/garden/garden.html?gid=35798849">@yet</a>种出了银芽柳花朵<br/>
            2.(22小时前)<a href="/game/garden/garden.html?gid=35803380"><font color="#ff0000">麋鹿迷露</font></a>来花园摘走了1朵白色菊花。<br/>
            3.(22小时前)<a href="/game/garden/garden.html?gid=35803380"><font color="#ff0000">麋鹿迷露</font></a>来花园摘走了1朵白色菊花。<br/>
            <a href="/game/garden/user/message_list.html">查看更多>></a><br/>
            <br/>
        </div>

        <div class="device">
            【花园设施】<br/>
            花园名:niyongwei的花园 <a href="/game/garden/user/garden_edit.html?act=1">[更改]</a><br/>
            公告板:劳动光荣，偷窃可耻！ <a href="/game/garden/user/garden_edit.html?act=2">[更改]</a><br/>
            <br/>
            开通蓝钻会员才可以使用一键功能！<br/>
            <br/>
        </div>
        <div class="tools">

            <router-link :to="{path:'/game/garden/backpack'}">背包</router-link>☆
            <router-link :to="{path:'/game/garden/shop'}">商店</router-link>☆
            <router-link :to="{path:'/game/garden/props'}">道具</router-link>☆
            <router-link :to="{path:'/game/garden/signin'}">签到</router-link>☆
            <router-link :to="{path:'/game/garden/help'}">帮助</router-link>
            <br/>
            <router-link :to="{path:'/game/garden/shop-sterious'}">神秘商店</router-link>☆
            <router-link :to="{path:'/game/garden/market'}">花种交易市场</router-link>
            <br/>

        </div>
        <br/>
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
import GardenNavHeader from '@/components/headers/GardenNavHeader.vue'

export default {
    name: 'GardenMyIndex',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
        GardenNavHeader,
    },
    created(){
    },
    mounted(){
        
        // 查询游戏列表
        // this.getGamesList()
        // this.$router.push({ name: 'AddGame'})

    },
    methods: {
        
    }
}
</script>

<style >



/* // <style lang="scss"> */

</style>