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
      <el-carousel style=" background-color: dimgrey;margin: -60px 0 0 0" :height="carouselHeight">
        <el-carousel-item v-for="item in 4" :key="item">
          <img
              style="width: 100%;height: 100%"
              src="https://fuss10.elemecdn.com/a/3f/3302e58f9a181d2509f3dc0fa68b0jpeg.jpeg"
          />
          <h3>aaaaaaaaaaaa</h3>
        </el-carousel-item>
      </el-carousel>

<!--      预览图片-->
      <div style="margin: 0 auto; text-align: center">
        <div style="margin: 50px !important;">
          <el-space v-for="i in 4" :key="i" wrap>
            <div>
              <img
                  :style="{'width': cardWidth,'height':cardHeight, 'margin':'10px'}"
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
    const carouselHeight=ref('')
    const cardHeight=ref('')
    const cardWidth=ref('')
    const activeIndex = ref('1')


    const handleSelect = (key: string, keyPath: string[]) => {
      console.log(key, keyPath)
    }


    function setHeight(){
      carouselHeight.value=window.innerWidth/16*9+'px'
      cardWidth.value=window.innerWidth/2*0.8+'px'
      cardHeight.value=window.innerWidth/2*0.8/16*9+'px'

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

</style>
