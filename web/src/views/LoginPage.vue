<template>
  <div class="loginDiv">
<!--    logo-->
    <div style="width: 100px;height: 100px;margin: 0 auto;background-color: #A7A7A7">
      我是logo
    </div>
    <div style="margin: 40px">
      <el-input style="margin: 0 0 20px;" v-model="input" placeholder="电子邮箱或手机号码" />
      <el-input v-model="password" type="password" placeholder="密码" />
      <div style="margin: 10px">
        <div style="text-align: left;float: left;">
          <a class="loginText">忘记密码</a>
        </div>
        <div style="text-align: right;">
          <a class="loginText">立即注册</a>
        </div>
      </div>
      <el-button style="width: 100%;margin: 10px 0 0 0;" type="success">登录</el-button>
      <el-divider style="opacity: 0.8;margin: 30px 0 30px 0;">
        <p style="font-size:5px;">或用以下方式登录</p>
      </el-divider>
      <div style="font-size: 30px; color: #A7A7A7">
        <i class="ri-github-fill"></i>
        <i class="ri-qq-fill" style="margin: 0 20px 0 20px;"></i>
        <i class="ri-wechat-fill"></i>
      </div>
    </div>
  </div>
</template>

<script>
import {ref} from "vue";
import 'remixicon/fonts/remixicon.css'
export default {
  name: "LoginPage",
  setup(){
    const input = ref('')
    const password = ref('')
    return{
      input,
      password,
    }
  }
}
</script>

<style scoped>
.loginText{
  font-size:12px;
  text-decoration:underline;
  color: #A7A7A7;
}
.loginDiv{
  text-align: center;
  /*background-color: white;*/
  width: 500px;
  height: 500px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%,-50%);
}
/deep/input::-webkit-input-placeholder {
  color: #A7A7A7;
  font-size: 15px;
}

/deep/.el-input__inner {
  background-color: transparent !important;
  border: 1px solid #A7A7A7;
}
/deep/.el-divider__text{
  background-color: #131313;
  color: #A7A7A7;
}
</style>
