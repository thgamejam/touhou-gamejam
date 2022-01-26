<template>
  <el-container>
    <el-header style="z-index: 5; padding: 0;opacity: 0.8">
      <el-menu
          :default-active="activeIndex2"
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
      <el-carousel style=" background-color: dimgrey;margin: -60px 0 0 0" indicator-position="outside" height="500px">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3>{{ item }}</h3>
        </el-carousel-item>
      </el-carousel>
    </el-main>
    <el-footer>Footer</el-footer>
  </el-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const activeIndex2 = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
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
