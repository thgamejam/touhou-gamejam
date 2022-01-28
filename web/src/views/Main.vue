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
      <el-carousel style=" background-color: dimgrey;margin: -60px 0 0 0" indicator-position="outside" :height="carouselHeight">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3>{{ item }}</h3>
        </el-carousel-item>
      </el-carousel>

<!--      预览图片-->
      <div style="margin: 0 auto; text-align: center">
        <div style="margin: 20px !important;">
          <el-space size="100px" wrap>
            <el-card v-for="i in 4" :key="i" style=" width: 500px; height: 300px" :body-style="{ padding: '0px' }">
              <img
                  src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                  class="image"
                  style="padding: 20px"
              />
            </el-card>
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
.image {
  width: 100%;
  display: block;
}
</style>
