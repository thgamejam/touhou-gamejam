<template>
  <headband />
  <div class="index-container">
    <Nav />
    <div class="index-wrapper">
      <div class="index-posts">
        <BlogCard />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Nav from "@/blog-venja-cc/components/Nav.vue";
import Headband from "@/blog-venja-cc/components/Headband.vue";
import BlogCard from "@/blog-venja-cc/components/BlogCard.vue";
export default {
  name: "Index",
  components: {BlogCard, Headband, Nav},
  // Override background color
  setup() {
    let parentBody = document.querySelector('body');
    if (parentBody) {
      parentBody.setAttribute('style', 'background-color: #f5f7f9; margin: 0px;');
    }
  }
}
</script>

<style scoped>
.index-container {
  width: 960px;
  height: 0;
  margin: 0 auto;
}
.index-wrapper {
  float: right;
  box-sizing: border-box;
  padding: 40px;
  width: 700px;
  background: #fff;
  min-height: 700px;
  box-shadow: initial;
  border-radius: initial;
}
.index-posts {
  padding: 40px 0 0 0;
}
</style>