<template>
  <div class="index">
        {{error}}
        
        <IndexHeader />

                
<div class="module-title">【家族管理】</div>
<a href="/bbs/forum.html?bid=12">家族天地</a>
<a href="/bbs/forum.html?bid=8001">家服团议事厅</a>
<a href="/bbs/gov/">政府</a> 

<div style="background-image: url(/wml/image/bbs/jzsy.gif);background-color: #91e09d;background-repeat: no-repeat;height: 25px;"></div>
<div class="bodule-title">
  <a href="/bbs/family_add.html">创建家族</a>.<a href="/bbs/family_pending_list.html">招兵买马</a>.<br/>
</div> 
加入家族一起在3GQQ家园社区闯荡吧，我们为你推荐:<br/>最佳匹配:<a href="/bbs/family.aspx?id=2460">优乐美</a>(<a href="/bbs/family_join.aspx?id=2460">快速加入</a>)<br/><a href="/bbs/family_add.aspx">申请我的家族</a><br/>
<div class="module-title">推荐家族</div> 
  <a href="/bbs/family/13022.html">★★★超级小可爱</a> . <a href="/bbs/family/2453.html">莲花坞</a> . <a href="/bbs/family/13021.html">有事吗？</a> . <a href="/bbs/family/2489.html">云涧</a> . <a href="/bbs/family/13025.html">缘份</a><br/><a href="/bbs/family/2452.html">清风小筑</a> . <a href="/bbs/family/2442.html">旧时光</a> . <a href="/bbs/family/2418.html">诗书烫酒</a> . <a href="/bbs/family/2425.html">粤香阁</a> . <a href="/bbs/family/2491.html">明教</a><br/>
  
<div class="module-title">家族舞台</div> 
  1.<a href="/bbs/topic/1298482.html">哈哈</a>(12阅/<a href="/bbs/reply_list/1298482.html">3回</a>)<br/>2.<a href="/bbs/topic/1298468.html">试帖的</a>(26阅/<a href="/bbs/reply_list/1298468.html">9回</a>)<br/>3.<a href="/bbs/topic/1298467.html">头晕了</a>(23阅/<a href="/bbs/reply_list/1298467.html">3回</a>)<br/>4.<a href="/bbs/topic/1298466.html">   不如</a>(24阅/<a href="/bbs/reply_list/1298466.html">6回</a>)<br/>5.<a href="/bbs/topic/1298463.html">腿上被咬了好多包</a>(51阅/<a href="/bbs/reply_list/1298463.html">28回</a>)<br/>6.<a href="/bbs/topic/1298458.html">小鲤鱼</a>(23阅/<a href="/bbs/reply_list/1298458.html">6回</a>)<br/>7.<a href="/bbs/topic/1298450.html">  我家的牛回家了</a>(55阅/<a href="/bbs/reply_list/1298450.html">10回</a>)<br/>8.<a href="/bbs/topic/1298447.html">【会员最新动态】</a>(34阅/<a href="/bbs/reply_list/1298447.html">1回</a>)<br/>9.<a href="/bbs/topic/1298441.html">九现神龙 （35800000）</a>(72阅/<a href="/bbs/reply_list/1298441.html">10回</a>)<br/>10.<a href="/bbs/topic/1298440.html">连云寨</a>(30阅/<a href="/bbs/reply_list/1298440.html">3回</a>)<br/>
<div class="module-title">家族排行榜</div> 
  <a href="/bbs/family_top.html?bt=1">规模排行榜</a> <a href="/bbs/family_top.html?bt=2">活跃排行榜</a><br/>
  <a href="/bbs/family_top.html?bt=3">人气排行榜</a> <a href="/bbs/family_top.html?bt=4">声望排行榜</a><br/>
  




<div class="module-title">【特区家族】</div> 

<a href="/bbs/family/2425.html">粤香阁</a>


<div class="module-title">家族特色</div> 
  <a href="/bbs/family_cat.html?cid=11">同城同乡</a>.<a href="/bbs/family_cat.html?cid=12">青春校园</a>.<a href="/bbs/family_cat.html?cid=13">打工生涯</a><br/>
  <a href="/bbs/family_cat.html?cid=14">明星粉丝</a>.<a href="/bbs/family_cat.html?cid=15">情感男女</a>.<a href="/bbs/family_cat.html?cid=16">娱乐八卦</a><br/>
  <a href="/bbs/family_cat.html?cid=17">军人风采</a>.<a href="/bbs/family_cat.html?cid=18">舞文弄墨</a>.<a href="/bbs/family_cat.html?cid=19">游戏动漫</a><br/>
  <a href="/bbs/family_cat.html?cid=20">科技数码</a>.<a href="/bbs/family_cat.html?cid=21">时尚生活</a>.<a href="/bbs/family_cat.html?cid=22">其他综合</a><br/>

<div class="module-title">
家族搜索
</div>
<div class="module-content">

<form action="bbs/family_search.html" method="get">
  
  <input type="text" name="wd" emptyok="true" maxlength="30" value="" />
<input type="submit" value="搜索" />
</form>
</div>
<br/>
<a href="/">3GQQ</a>><a href="/bbs/">社区广场</a>>家族首页<br/>
<a href="/bbs/family_add.html">创建家族</a>.<a href="/bbs/family_pending_list.html">招兵买马</a>.<a href="/bbs/forum.html?bid=12">家族天地</a><br/>


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
        // let userinfo = this.$userinfo.get()
        // if(userinfo){
        //     this.userinfo = userinfo
        //     this.username = this.userinfo.username
        // }
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