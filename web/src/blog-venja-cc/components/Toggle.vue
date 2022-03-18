<template>
<div class="sidebar-container bg-white" :class="affixTop ? 'sidebar-affix-top' : 'sidebar-affix'">
  <div class="sidebar-inner">
    <div class="sidebar-profile" v-show="logon">
      <!--<img class="sidebar-profile-avatar" alt="" />-->
      <div class="sidebar-profile-avatar bg-black"/>
      <p class="sidebar-profile-username black bold">Kazi-VE</p>
      <p class="sidebar-profile-description">Unity3D & Unreal Engine 4</p>
    </div>
    <div class="sidebar-profile-login" v-show="!logon">
      <div class="sidebar-profile-avatar bg-black"/>
      <p class="sidebar-profile-not-login bold black">未登入</p>
      <p class="sidebar-profile-login-register">
        <button class="sidebar-profile-login-btn" @click="onButtonDown">
          登入
        </button>
        <button class="sidebar-profile-register-btn" @click="onButtonDown">
          注册
        </button>
      </p>
    </div>
    <button @click="onButtonDown">Login/Logout</button>
  </div>
</div>
</template>

<script lang="ts">
import {ref} from "vue";

export default {
  name: "Toggle",
  setup() {
    let affixTop = ref(false);
    let logon = ref(false);

    window.addEventListener('scroll', onScroll)

    function onScroll() {
      affixTop.value = window.scrollY >= 300;
    }

    function onButtonDown() {
      logon.value = !logon.value;
    }

    return {
      affixTop,
      logon,
      onButtonDown
    }
  },
}
</script>

<style scoped>
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
  min-height: 240px;
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
  margin: 0 auto;
}
.sidebar-profile-description {
  height: 20px;
}
.sidebar-profile-login-register {
}
.sidebar-profile-register-btn,
.sidebar-profile-login-btn {
  cursor: pointer;
  color: #999;
  border: solid transparent 1px;
  background-color: transparent;
  margin: -1px 2px;
}
.sidebar-profile-register-btn:hover,
.sidebar-profile-login-btn:hover
{
  color: black;
  border: black solid 1px;
  transition-duration: 0.2s;
  transition-timing-function: ease-in-out;
  transition-delay: 0s;
}
</style>