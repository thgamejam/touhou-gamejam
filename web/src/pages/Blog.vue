<template>
  <headband/>
  <div class="common-container">
    <Nav/>
    <Toggle/>
    <div class="common-wrapper">
      <div class="common-content">
        <div class="blog-container">

          <div class="blog-title not-selectable">
            <h1>这是一个标题1abc</h1>
          </div>

          <div class="blog-meta meta not-selectable">
            <i class="el-icon-date"> 发表于 xxxx-xx-xx </i>
            <span class="blog-meta-divider">|</span>
            <i class="el-icon-folder"> 归档于 ﹢﹢﹢ </i>
            <span class="blog-meta-divider">|</span>
            <i class="el-icon-view"> 666 </i>
          </div>

          <div class="blog-divider divider"/>

          <div class="blog-body">

            <h3>随记</h3>
            <p>莫因渴望你没有的，</p>
            <p>而错过你已拥有的；</p>
            <p>要知道，</p>
            <p>你现在拥有的，</p>
            <p>也曾是你渴望的。</p>
            <p>◆孟德斯鸠</p>

          </div>

          <div class="blog-divider divider"/>

          <div class="blog-footer">

          </div>

        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Nav from "@/components/Nav.vue";
import Headband from "@/components/Headband.vue";
import Toggle from "@/components/Toggle.vue";
</script>

<style scoped>
.meta {
  margin-bottom: 10px;
}
.divider {
  margin: 40px auto;
}
</style>