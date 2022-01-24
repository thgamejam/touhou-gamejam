<template>
  <a-layout>
    <a-layout-header :style="{ backgroundColor:'#121212', position: 'fixed', zIndex: 1, width: '100%' }">
      <div class="logo" />
      <a-menu
          v-model:selectedKeys="selectedKeys"
          mode="horizontal"
          :style="{ color:'#FFFFFF', backgroundColor:'#121212', lineHeight: '64px' }"
      >
        <a-menu-item key="1">首页</a-menu-item>
        <a-menu-item key="2">游戏</a-menu-item>
        <a-menu-item key="3">比赛</a-menu-item>
        <a-menu-item key="4">关于</a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content :style="{ backgroundColor:'#FFFFFF', padding: '0 50px', marginTop: '64px' }">


    </a-layout-content>
    <a-layout-footer :style="{ backgroundColor:'#121212', textAlign: 'center' }">

    </a-layout-footer>
  </a-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
export default defineComponent({
  components:{

  },
  setup() {
    return {
      selectedKeys: ref<string[]>(['1']),
    };
  },
});
</script>

<style scoped>

.ant-menu-horizontal{
  border-bottom: 1px solid #131313;
}

</style>
