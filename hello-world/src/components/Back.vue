

<template>
    <div class="back">
        <span @click="goBack" class="mycursor">[返回]</span>
    </div>
</template>

<script>
export default {
    name: 'Back',
    props: {
    },
    data() {
        return {
            
        }
    },
    mounted(){
    },
    destroyed(){
    },
    methods:{
        goBack(){
            this.$router.go(-1)
        },
        

    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
