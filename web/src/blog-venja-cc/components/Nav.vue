<template>
<div class="nav-container">
  <div class="nav-inner">
    <div class="nav-meta">
      <span class="nav-meta-title">Kazi-VE</span>
      <p class="nav-meta-description">Softness is your greatest weakness</p>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "Nav"
}

</script>

<style scoped>
.nav-container{
  width: 960px;
  height: 0;
  position: relative;
  margin: 0 auto;
}
.nav-inner {
  width: 240px;
  height: 295px;
  background-color: white;
  float: left;
}
.nav-meta {
  padding: 20px;
  background-color: black;
  color: white;
}
.nav-meta-title {

}
.nav-meta-description {
  color: #dddddd !important;
  margin: 10px 10px 0;
  font-size: 13px;
}
</style>