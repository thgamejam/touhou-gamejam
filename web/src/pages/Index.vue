<template>
  <headband/>
  <div class="common-container">
    <Nav/>
    <Toggle/>
    <div class="common-wrapper">
      <div class="common-content">
        <BlogCard/>
        <BlogCard/>
        <BlogCard/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Toggle from "@/components/Toggle.vue";
import Nav from "@/components/Nav.vue";
import Headband from "@/components/Headband.vue";
import BlogCard from "@/components/BlogCard.vue";
</script>

<style scoped>
</style>