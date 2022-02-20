<template>
  <el-container>
<!--    导航栏-->
    <el-header style="z-index: 5; padding: 0;opacity: 0.8">
      <el-menu
          :default-active="activeIndex"
          class="el-menu-demo"
          mode="horizontal"
          background-color="#131313"
          text-color="#fff"
          active-text-color="#ffd04b"
          @select="handleSelect"
      >
        <el-menu-item index="1">首页</el-menu-item>
        <el-menu-item index="2">游戏</el-menu-item>
        <el-menu-item index="3">比赛</el-menu-item>
        <el-menu-item index="4">关于</el-menu-item>
      </el-menu>
    </el-header>



    <el-main style="padding: 0;overflow: visible">
<!--      走马灯-->
      <el-carousel style=" background-color: dimgrey;margin: -60px 0 0 0" :height="carouselHeight+'px'">
        <el-carousel-item v-for="item in carouselImg" :key="item">
          <el-image
              style="width: 100%;height: 100%"
              :src=item.imgURL
          />

        </el-carousel-item>
      </el-carousel>

<!--      预览图片-->
      <div style="margin: 0 auto; text-align: center">
        <div style="margin: 50px !important;">
          <el-space v-for="i in 4" :key="i" wrap>
            <div class="ani">
<!--              文字介绍栏-->
              <div class="mv" :style="{'height':cardHeight*0.5+'px','width': cardWidth+'px'}">
                <h2 style="margin: 20px">《这是标题》</h2>
                <h4 style="margin: 20px">&nbsp;&nbsp;&nbsp;&nbsp;这是介绍这是介绍这是介绍这是介绍这是介绍这是介绍这是介绍这是介绍这是介绍这是介绍这是介绍</h4>
                <div style="position:absolute;right:0px;bottom:10px;width:100px;color: white">
                  <a style="color: rgba(255,255,255,0.8)">
                    <h5>了解详情</h5>
                  </a>
                </div>
              </div>
<!--              展示图片-->
              <img
                  class="image"
                  :style="{'width': cardWidth+'px','height':cardHeight+'px', 'margin':'10px'}"
                  src="https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg"
              />
            </div>
          </el-space>
        </div>
      </div>

    </el-main>



    <el-footer>Footer</el-footer>
  </el-container>
</template>

<script lang="ts">
import {defineComponent, onBeforeMount, ref} from 'vue'
export default defineComponent({
  setup(){
    const carouselHeight=ref(0)
    const cardHeight=ref(0)
    const cardWidth=ref(0)
    const activeIndex = ref('1')
    let carouselImg = ref([
        {
          imgURL: "https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg",
          router: "",
          text: "",
          des: ""
        },
      {
        imgURL: "https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg",
        router: "",
        text: "",
        des: ""
      },
    ])


    const handleSelect = (key: string, keyPath: string[]) => {
      console.log(key, keyPath)
    }


    function setHeight(){
      carouselHeight.value=window.innerWidth/16*9
      cardWidth.value=window.innerWidth/2*0.8
      cardHeight.value=window.innerWidth/2*0.8/16*9

    }


    /**
     * 初始化设置高度
     */
    onBeforeMount(() => {
     setHeight()
    })

    /**
     * 动态自适应走马灯高度
     */
    window.addEventListener('resize',setHeight)

    return{
      activeIndex,
      cardHeight,
      cardWidth,
      carouselImg,
      carouselHeight,
      handleSelect
    }
  }
})
</script>

<style scoped>
.el-menu.el-menu--horizontal {
  border-bottom: solid 1px #131313;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
  text-align: center;
}

/*展示框文字描述动画*/
.ani:hover>.mv{
  animation-name: cardTextAni;
  animation-duration: 1s;
  animation-fill-mode: forwards;
}


/*展示框图片边缘发光动画*/
.ani:hover>.image{
  box-shadow: 0 0 10px 5px rgba(0,0,0,0.4);
}

/*展示框文字自适应*/
.mv{
  position: absolute;
  transform: translateY(100%);
  word-break:break-word;
  text-align:left;
  background-color:rgba(28,28,35,0.78);
  color:rgba(255,255,255,0.8);
  margin:10px;
  opacity: 0;
}

@keyframes cardTextAni
{
  0%   {opacity: 0}
  100% {opacity: 1}
}

</style>
