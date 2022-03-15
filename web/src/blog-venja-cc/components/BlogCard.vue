<template>
<div class="blog-card-container">

  <div class="blog-card-title">
    <h1><a class="blog-card-link" href="" itemprop="url">这是一个标题1abc</a></h1>
  </div>

  <div class="blog-card-meta">
    <i class="el-icon-date"> 发表于 xxxx-xx-xx  </i>
    <span class="blog-card-meta-divider">|</span>
    <i class="el-icon-folder"> 归档于 ﹢﹢﹢ </i>
    <span class="blog-card-meta-divider">|</span>
    <i class="el-icon-view"> 666 </i>
  </div>

  <div class="blog-card-body">

    <h3>随记</h3>
    <p>莫因渴望你没有的，</p>
    <p>而错过你已拥有的；</p>
    <p>要知道，</p>
    <p>你现在拥有的，</p>
    <p>也曾是你渴望的。</p>
    <p>◆孟德斯鸠</p>

    <div class="blog-card-body-button">
      <a href="" class="blog-card-link-btn">阅读全文 »</a>
    </div>
  </div>

  <div class="blog-card-footer" />

</div>
</template>

<script>
export default {
  name: "BlogCard"
}
</script>

<style scoped>
h1 {
  margin: 0;
  padding: 0;
  font-size: 22px;
  font-weight: 400;
  line-height: 1.5;
  display: block;
}
h3 {
  padding-top: 10px;
  font-size: 18px;
  margin: 20px 0 15px;
  line-height: 1.5;
}
a {
  text-decoration: none;
  position: relative;
  outline: none;
  color: #555;
}
p {
  margin: 0;
  font-size: 14px;
}

.blog-card-container {
  line-height: 2;
}
.blog-card-meta {
  margin: 3px 0 60px 0;
  color: #999;
  font-size: 12px;
  text-align: center;
}
.blog-card-meta-divider {
  margin: 0 0.5em;
}

.blog-card-title {

}
.blog-card-body {
  text-align: justify;
  min-height: 180px;
}
.blog-card-body-button {
  margin-top: 40px;
  text-align: center;
}
.blog-card-footer {
  display: block;
  margin: 80px auto 60px;
  width: 8%;
  height: 1px;
  background: #ccc;
  text-align: center;
}

.blog-card-link::before {
  content: "";
  position: absolute;
  width: 100%;
  height: 2px;
  left: 0;
  bottom: 0;
  background-color: #000;
  visibility: hidden;
  -webkit-transform: scaleX(0);
  -moz-transform: scaleX(0);
  -ms-transform: scaleX(0);
  -o-transform: scaleX(0);
  transform: scaleX(0);
  transition-duration: 0.2s;
  transition-timing-function: ease-in-out;
  transition-delay: 0s;
}
.blog-card-link:hover::before {
  visibility: visible;
  -webkit-transform: scaleX(1);
  -moz-transform: scaleX(1);
  -ms-transform: scaleX(1);
  -o-transform: scaleX(1);
  transform: scaleX(1);
}

.blog-card-link-btn {
  display: inline-block;
  padding: 0 20px;
  font-size: 14px;
  color: #555;
  background: #fff;
  border: 2px solid #555;
  text-decoration: none;
  border-radius: 2px;
  line-height: 2;
}
.blog-card-link-btn:hover {
  border-color: #222;
  color: #fff;
  background: #222;
  transition-property: background-color, color, border-color;
  transition-duration: 0.2s;
  transition-timing-function: ease-in-out;
  transition-delay: 0s;
}
</style>