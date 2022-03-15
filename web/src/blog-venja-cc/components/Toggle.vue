<template>
<div class="sidebar-container" :class="affixTop ? 'sidebar-affix-top' : 'sidebar-affix'">
  <div class="sidebar-inner">
    <div class="sidebar-profile">
      <!--<img class="sidebar-profile-avatar" alt="" />-->
      <div class="sidebar-profile-avatar"/>
      <p class="sidebar-profile-username black bold">Kazi-VE</p>
      <p class="sidebar-profile-description">Unity3D & Unreal Engine 4</p>
    </div>
  </div>
</div>
</template>

<script lang="ts">
import {ref} from "vue";

export default {
  name: "Toggle",
  setup() {
    let affixTop = ref(false);

    window.addEventListener('scroll', onScroll)

    function onScroll(){
      if (window.scrollY >= 300) {
        affixTop = ref(true);
      } else {
        affixTop = ref(false);
      }
      console.log(affixTop.value)
    }

    return { affixTop }
  },
}
</script>

<style scoped>
.black {
  color: black;
}
.bold {
  font-weight: bold;
}
p {
  font-size: 13px;
}
.sidebar-affix {
  position: absolute;
  margin-top: 307px;
}
.sidebar-affix-top {
  position: fixed;
  margin-top: 12px;
}
.sidebar-container {
  width: 240px;
  height: 240px;
  background-color: white;
  float: left;
}
.sidebar-inner {
  position: relative;
  padding: 20px 10px;
  color: #999;
  text-align: center;
}
.sidebar-profile-avatar {
  height: 126px;
  width: 126px;
  background-color: black;
  margin: 0 auto;
}
</style>