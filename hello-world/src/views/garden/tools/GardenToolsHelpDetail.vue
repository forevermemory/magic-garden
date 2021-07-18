<template>
    <div>
        {{error}}
        【花园帮助】<br/>
<a href="/game/garden/user/my_garden.html">花园</a>><a href="/game/garden/help.html">帮助</a><br/>
怎样种出想要的颜色的花?<br/>
在商店中只能购买某一类花的种子,至于种出什么颜色的花就得看您的运气啦。据有经验的老花匠透露,每个人都有自己较容易种出的花色,也有较不容易种出的花色。要么勤奋的种花,要么就多去好友那看看,说不定正好可以互补有无。<br/>
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>><a href="/game/garden/help.html">帮助</a><br/>
        
    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsHelpDetail',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
    },
    created(){
    },
    mounted(){
        
        // this.$router.push({ name: 'AddGame'})

    },
    methods: {
        
    }
}
</script>

<style >



/* // <style lang="scss"> */

</style>