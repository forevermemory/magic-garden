

<template>
    <div class="footer">
    <p>
    <a href="#">3GQQ家园社区</a>-<a href="#">社区广场</a>-<a href="#">导航</a>-<a href="#">退出</a><br>
    </p>
    <p>
    我的:<a href="#">地盘</a>.<a href="#">空间</a>.<a href="#">家园</a><br>
    </p>
    <p>
    小Q报时：{{now}}<br>
    </p>
    </div>
</template>

<script>
export default {
    name: 'Footer',
    props: {
    },
    data() {
        return {
            timeInterval:null,
            timeFmt:"yyyy-MM-dd HH:mm:ss",
            now: ''
        }
    },
    mounted(){
        this.startUpdateTime()
    },
    destroyed(){
        clearInterval(this.timeInterval)
    },
    methods:{
        startUpdateTime(){
            this.timeInterval = setInterval(()=>{
                this.now = this.format(new Date())
            },1000)
        },
        format(data) {
            let fmt = this.timeFmt
            var o = {
                "M+": data.getMonth() + 1, //月份 
                "d+": data.getDate(), //日 
                "H+": data.getHours(), //小时 
                "m+": data.getMinutes(), //分 
                "s+": data.getSeconds(), //秒 
                "q+": Math.floor((data.getMonth() + 3) / 3), //季度 
                "S": data.getMilliseconds() //毫秒 
            };
            if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (data.getFullYear() + "").substr(4 - RegExp.$1.length));
            for (var k in o)
            if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
            return fmt;
        },

    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
