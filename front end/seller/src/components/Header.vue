<template>
  <el-header class="el-header">
    <el-row>
      <el-col :span="21">
        <div style="margin-top: 24px;">
        <el-breadcrumb separator-class="el-icon-arrow-right">
          <el-breadcrumb-item :to="{ path: '/main/page' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item :to="navName1">{{navTitle1}}</el-breadcrumb-item>
          <el-breadcrumb-item v-if="navTitle2 !== ''">{{navTitle2}}</el-breadcrumb-item>
        </el-breadcrumb>
        </div>
      </el-col>
      <el-col :span="1">
        <el-button type="text" class="el-icon-full-screen button-1" />
      </el-col>
      <el-col :span="1">
        <el-button type="text" class="el-icon-bell button-2" />
      </el-col>
      <el-col :span="1">
        <el-dropdown trigger="click" style="margin-top: 14px;">
          <span class="el-dropdown-link">
            <el-avatar :size="30" :src="avatar.url" />
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item icon="el-icon-switch-button" @click="logout">
                退出账户
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
    </el-row>
  </el-header>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      avatar: {
        url: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png'
      },
    }
  },
  computed: {
    navTitle1: {
      get() {
        return this.$store.state.navigation.title
      },
      set(val) {
        this.$store.state.navigation.title = val
      }
    },
    navName1: {
      get() {
        return this.$store.state.navigation.path
      },
      set(val) {
        this.$store.state.navigation.path = val
      }
    },
    navTitle2: {
      get() {
        return this.$store.state.navigation.afterName
      },
      set(val) {
        this.$store.state.navigation.afterName = val
      }
    }
  },
  methods: {
    logout() {
        localStorage.clear();
        sessionStorage.clear();
        this.$store.commit("resetState");
        this.$router.push('/');
    }
  }
}
</script>

<style scoped>
.el-header {
  background-color: white;
  margin-bottom: 1px;
  box-shadow:1px 1px 1px #dcdbdb;
}
.button-1{
  font-size: 22px;
  color: #475870;
  margin-top: 7px;
}
.button-2{
  font-size: 24px;
  color: #475870;
  margin-top: 6px;
}
.el-dropdown-link {
  cursor: pointer;
  color: #409EFF;
}
</style>