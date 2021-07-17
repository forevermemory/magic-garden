<template>
    <div>
        <!-- <IndexBaseHeader /> -->
        {{message}}
    </div>
</template>

<script>
// import {listGames,userAddGames} from '@/api/request.js'
// import IndexBaseHeader from '@/components/headers/IndexBaseHeader.vue'
// import VueEvent from '@/event/index.js'

export default {
    name: 'AddGameResult',
    data() {
        return {
            message:'',
        }
    },
    components: {
        // IndexBaseHeader,
    },
    created(){
        let user = this.$userinfo.get()
        this.message = user.addGameResultMessage
    },
    mounted(){
        
    },
    methods: {
       
    }
}
</script>

<style >
/* // <style lang="scss"> */



</style>