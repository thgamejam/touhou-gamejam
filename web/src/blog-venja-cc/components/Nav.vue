<template>
<div class="nav-container">
  <div class="nav-inner">
    <div class="nav-meta">
      <span class="nav-meta-title">Kazi-VE</span>
      <p class="nav-meta-description">Softness is your greatest weakness</p>
    </div>
    <div class="nav-menu">
      <ul class="nav-menu-list not-selectable">
        <li class="nav-menu-list-item">
          <a href="" ref="section">
            <i class="el-icon-s-home"> 首页 </i>
          </a>
        </li>
        <li class="nav-menu-list-item">
          <a href="" ref="section">
            <i class="el-icon-s-flag"> 标签 </i>
          </a>
        </li>
        <li class="nav-menu-list-item">
          <a href="" ref="section">
            <i class="el-icon-s-grid"> 分类 </i>
          </a>
        </li>
        <li class="nav-menu-list-item">
          <a href="" ref="section">
            <i class="el-icon-s-management"> 归档 </i>
          </a>
        </li>
      </ul>
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
.nav-menu {
  margin-top: 20px;
  padding-left: 0;
  text-align: center;
}
.nav-menu-list {
  display: block;
  padding-inline-start: 0;
}
.nav-menu-list-item {
  list-style: none;
  text-align: left;
  margin: 0;
  line-height: 2.1;
}
a {
  outline: none;
  text-decoration: none;
  color: #555;
  display: block;
  font-size: 13px;
  position: relative;
  box-sizing: border-box;
  padding: 5px 20px;
  text-align: left;
  line-height: inherit;
  transition-property: background-color;
  transition-duration: 0.2s;
  transition-timing-function: ease-in-out;
  transition-delay: 0s;
}
</style>