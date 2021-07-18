<template>
    <div>
        {{error}}

        <div class="market">
            <router-link :to="{path:'/game/garden/market'}">花种交易市场</router-link>
            |
            <router-link :to="{path:'/game/garden/market-onsale'}">我的橱窗</router-link>
        </div>
【出售花种】<br/>
 花种名称：梅花<br/>
 花种数量：1<br/>
 <form method='post' action='add.asp?muui=sok'>
<input type="hidden" name="did" value="3"/>
<input type="hidden" name="hname" value="梅花"/>
<input type="hidden" name="amount" value="1"/>

出售数量：<input type="text" name="shuliang" maxlength="2" size="2" value="1"/><br/>
币种：<select name="bi" value="1">
<option value="1">GB</option>
<option value="2">元宝</option>
</select><br/>
出售单价<input type="text" name="jia" maxlength="10" size="2" value="100"/><br/>
<input type="submit" value="确定出售" />
</form>
 
----------<br/>
<a href="/game/garden/user/my_garden.html">花园</a>><a href="/game/garden/shop/index.asp">市场</a>>出售花种<br/>


    </div>
</template>

<script>
// import {getGameByID} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'

export default {
    name: 'GardenToolsMarketSale2',
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