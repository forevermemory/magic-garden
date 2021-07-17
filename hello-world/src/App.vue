<template>
  <div id="app">
    <!-- <Header v-bind:content="headerContent"/>  -->
    <router-view></router-view>
    <Footer/> 
  </div>
</template>

<script>
import Footer from './components/Footer.vue'
// import Header from './components/Header.vue'
// import VueEvent from '@/event/index.js'
// 

export default {
    name: 'App',
    components: {
        Footer,
        // Header
    },
    data() {
        return {
            headerContent:''
        }
    },
    mounted(){
        // VueEvent.$on('toHeader',(data)=>{this.headerContent = data})
    }
}
</script>

<style>
/* #app {
    font-size: 14px;
} */

/*重置*/
body,div,p,a,table,textarea,form,input,img,ul,li{margin:0;padding:0;}
body{font-size:14px;margin:5px;line-height:1.5;}

li{list-style:none;text-indent:0;}
img,a img{border:0;}
a{text-decoration:none;color:#004299;}
a:visited{ 
    text-decoration: none;
    color:#004299;
}
a:visited.router-link-active{
  color: red
}


/* a:visited {color:#800080;} */
em{color:#9B9B9B;font-size:12px;}
textarea{font-size:13px;width:98%;margin:0 0 5px 0;border:1px solid #9FC6EC;}

/*图标*/
.logo{padding:5px 0 0 10px;background:#fff;}

/*页头*/
.header{background:url("~@/assets/images/header.gif") repeat-x;height:22px;line-height:22px;}

/*页脚*/
.footer{padding:5px;background:#F5FFFA;border-top:1px solid #9FC6EC;}
.footer p{line-height:22px;color:#515151;}
.footer p a{color:#0051A4;}
.footer .version a{color:#004299;}

/*列表*/
.list{line-height:1.6;}
.list a{color:#0051A4;}
.list img{margin-right:2px;vertical-align:middle;}
.list .row{padding:3px;border-bottom:1px solid #E3E6EB;}
.list .row form{margin:0px;display: inline}

/*导航*/
.bar{height:25px;padding:0 5px;background:#71afe3 url("~@/assets/images/bar.gif") repeat-x;line-height:25px;color:#fff;}
.bar a{color:#fff;}

/*广告*/
.ad{padding-left:5px;line-height:20px;font-size:12px;}
.ad a{color:#ff0000;}

/*广播*/
.note{background:#FFF9B7;border-bottom:1px solid #9FC6EC;}
.radio{padding-left:5px;background:#FFF9B7;border-bottom:1px solid #9FC6EC;}

/*文本*/
.text{line-height:1.3;padding:3px 5px;word-wrap:break-word;}
.text img{max-width:100%;margin-right:2px;vertical-align:middle;}

/*项目*/
.item{line-height:1.3;padding-left:2px;}
.item a{color:#0051A4;}

/*名称*/
.name{padding-left:3px;line-height:20px;border-bottom:2px solid #9FC6EC;color:#000;font-weight:bold;}
.name a{color:#0051A4;}

/*标题*/
.title{padding-left:2px;line-height:20px;border-bottom:1px solid #9FC6EC;}
.title a{color:#0051A4;}

/*常见*/
.line{margin:2px 0 1px;padding:0 5px;border:1px solid #9FC6EC;}
.width{white-space:nowrap;text-overflow:ellipsis;-o-text-overflow:ellipsis;overflow:hidden;}
.deep{background:#E3EEF8;border:1px solid #9FC6EC;border-left:none;border-right:none;}

</style>
