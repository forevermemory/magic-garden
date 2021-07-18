<template>
    <div>
        <IndexBaseHeader />
        <br>
        {{error}}
        <br>
        <br>
        <Back/>
        <!-- <router-view></router-view> -->
    </div>
</template>

<script>
import {getGameByfullPath} from '@/api/request.js'
import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'
import Back from '@/components/Back.vue'
// import VueEvent from '@/event/index.js'

export default {
    name: 'EnterGame',
    data() {
        return {
            // games:[],
            error:'',
            game:{},
        }
    },
    components: {
        IndexBaseHeader,
        Back,
    },
    beforeRouteEnter (to, from, next) {
        // 在渲染该组件的对应路由被 confirm 前调用
        // 不！能！获取组件实例 `this`
        // 因为当钩子执行前，组件实例还没被创建

        // if(to.query.game.g_state!=1){
        //     // 还没有开通
        //     // this.error = `[${to.query.game.g_name}]暂未开通,敬请期待`
        // }
        // return

        next()

        
    },
    beforeRouteUpdate (to, from, next) {
        // 在当前路由改变，但是该组件被复用时调用
        // 举例来说，对于一个带有动态参数的路径 /foo/:id，在 /foo/1 和 /foo/2 之间跳转的时候，
        // 由于会渲染同样的 Foo 组件，因此组件实例会被复用。而这个钩子就会在这个情况下被调用。
        // 可以访问组件实例 `this`
        next()
    },
    beforeRouteLeave (to, from, next) {
        // 导航离开该组件的对应路由时调用
        // 可以访问组件实例 `this`
        next()
    },
    created(){
    },
    mounted(){
        
        // console.log(this.$router.history.current)
        let path = this.$router.history.current.query.g_url
        if(!path){
            this.error = '请正确操作!'
            return
        }
        this.error = ''
        this.handleGameStatus2(path)


    },
    methods: {
        async handleGameStatus2(url){
            let res = await getGameByfullPath({g_url:url})
            if(res.code==0 && res.data.total>0){
                if(res.data.data[0].g_state!=1){
                    // console.log(res.data)
                    this.error = ` [${res.data.data[0].g_name}] 暂未开通,敬请期待!`
                    return
                }
                // 
                console.log('enter:',res.data.data[0].g_url)
                this.$router.push({ path: res.data.data[0].g_url})
            }
        },
        // async handleGameStatus(gameID){
        //     let res = await getGameByID({id:gameID})
        //     if(res.code==0){
        //         if(res.data.g_state!=1){
        //             console.log(res.data)
        //             this.error = ` [${res.data.g_name}] 暂未开通,敬请期待!`
        //             return
        //         }
        //         // 
        //         this.$router.push({ path: res.data.g_url})
        //     }
        // }
    }
}
</script>

<style >



/* // <style lang="scss"> */

</style>