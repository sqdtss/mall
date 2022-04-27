<template>
  <div class="main">
    <div class="form">
      <img src="../assets/logo.png" class="logo" alt="logo"/>
      <el-form :model="loginForm"
               :rules="rules"
               ref="loginForm"
               class="demo-ruleForm" status-icon>
        <el-form-item prop="username">
          <el-input type="text"
                    v-model="loginForm.username"
                    autocomplete="off"
                    prefix-icon="el-icon-user"/>
        </el-form-item>
        <el-form-item prop="password">
          <el-input type="password"
                    v-model="loginForm.password"
                    autocomplete="off"
                    prefix-icon="el-icon-lock"/>
        </el-form-item>
        <el-form-item prop="captchaValue">
          <div style="display: inline-flex;">
            <el-input v-model="loginForm.captchaValue"
                      prefix-icon="el-icon-circle-check"
                      style="width: 50%;"
                      maxlength="5"/>
            <el-image :src="captchaImg" class="captchaImg" @click="getCaptcha"/>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary"
                     class="button"
                     @click="submitForm('loginForm')">登 录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    const validateUsercode = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入账号'));
      } else {
        callback();
      }
    };
    const validatePassword = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        callback();
      }
    };
    return {
      show: true,
      loginForm: {
        username: 'admin',
        password: '123',
        captchaId: '',
        captchaValue: '',
      },
      captchaImg: null,
      rules: {
        username: [
          {validator: validateUsercode, trigger: 'blur'}
        ],
        password: [
          {validator: validatePassword, trigger: 'blur'}
        ],
        captchaValue: [
          {required: true, message: '请输入验证码', trigger: 'blur'},
          {min: 4, max: 4, message: '长度为 5 个字符', trigger: 'blur'}
        ],
      }
    };
  },
  mounted() {
    this.getCaptcha();
  },
  updated() {
    this.getCaptcha();
  },
  methods: {

    // 用户登录
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$axios.post('/login', {
            username: this.loginForm.username,
            password: this.loginForm.password,
            captchaId: this.loginForm.captchaId,
            captchaValue: this.loginForm.captchaValue,
          }).then((response) => {
            localStorage.setItem("token", response.data.data.token)
            localStorage.setItem("uid", response.data.data.uid)
            this.$router.push('/home');
          }).catch((error) => {
            console.log(error);
          })
        }
      });
      this.getCaptcha();
    },

    // 获取验证码
    getCaptcha() {
      this.$axios.get('/captcha').then(response => {
        this.loginForm.captchaId = response.data.data.captchaId
        this.captchaImg = response.data.data.captchaImg
        this.loginForm.captchaValue = ''
      })
    }
  }
}
</script>

<style scoped>
.main {
  width: 100%;
  height: 100vh;
  display: flex;
  background-image: url("../assets/back.jpeg");
  background-size: 100% 60%;
  background-repeat: no-repeat;
  opacity: 0.8;
  background-color: #e8ecf1;
}
.form {
  width: 25%;
  height: 55%;
  margin: 150px auto;
  text-align: center;
  padding: 20px;
  background-color: white;
  border-radius: 8px;
  box-shadow: 1px 2px 1px #F2F6FC;
}
.logo {
  width: 80px;
  height: 80px;
  margin: 20px;
}
.captchaImg {
  width: 45%;
  height: 40px;
  padding: 0 20px;
  float: left;
}
.button {
  width: 100%;
  margin-top: 5px;
}
</style>