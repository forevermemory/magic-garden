<template>
  <div class="index">
        {{error}}
        
        <IndexHeader />

         
        广场

    </div>
</template>

<script>
// import {login,listGames} from '@/api/request.js'
import IndexHeader from '@/components/headers/IndexHeader.vue'

export default {
    data() {
        return {
            userinfo:{},
            username:'未知用户',
            error:'',
            page:{
                pageSize:100000,
                pageNo:1,
            },
            games:[],
        }
    },
    components: {
        IndexHeader,
        // Header
    },
    created(){
        
    },
    mounted(){
       
    },
    methods: {
        
        
    }
}
</script>

<style >
/* // <style lang="scss"> */

.module-title {
    margin: 8px 0 0;
    padding: 0 5px;
    height: 20px;
    line-height: 20px;
    border-bottom: 4px solid #9FC6EC;
    color: #000;
    font-weight: bold;
}



</style>