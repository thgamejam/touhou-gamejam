<template>
  <div class="signin-container center bg-white">
    <div class="signin-form center">
      <div class="signin-form-account">
        <p class="bold">账号</p><input>
      </div>
      <div class="signin-form-password">
        <p class="bold">密码</p><input>
      </div>
      <div class="signin-helper">
        <a class="signin-helper-forget-password left" href="">忘记密码</a> <a class="signin-helper-register right" href="">立即注册</a>
      </div>

      <div style="margin-top: 50px">
        <a href="" class="basic-button signin-login-btn not-selectable">登入</a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginRegister",
  components: {},
  setup() {
    let parentBody = document.querySelector('body');
    if (parentBody) {
      parentBody.setAttribute('style', 'background-color: #f5f7f9; margin: 0px;');
    }
  }
}
</script>

<style scoped>
p {
  display: inline-block;
  font-size: 18px;
  margin: auto;
  padding: 0 10px;
}
input {
  padding: 0 10px;
  background-color: #fff;
  border-radius: 4px;
  border: 1px solid #dcdfe6;
  box-sizing: border-box;
  color: #606266;
  display: inline-block;
  font-size: inherit;
  height: 100%;
  line-height: 40px;
  outline: 0;
  width: 80%;
}
.signin-container {
  width: 500px;
  height: 350px;
  border-radius: 20px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.2);
}
.signin-form {
  transform: translate(0, -10%);
  width: 75%;
  height: 50%;
}
.signin-form-account,
.signin-form-password {
  height: 40px;
  margin: 15px auto;
}
.signin-helper {
  margin: 15px auto;
}
.signin-helper-forget-password {
  margin-left: 18px;
}
.signin-helper-register {
  margin-right: 15px;
}
.signin-login-btn {
  font-size: 18px;
}
</style>