<template>
  <el-container class="el-container">
    <el-aside width="200px" class="el-aside">
      <Menu/>
    </el-aside>
    <el-container style="display: inline-block;">
      <Header/>
      <el-main class="el-main">
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Menu from "../components/Menu";
import Header from "../components/Header";

export default {
  name: "Home",
  components: {Header, Menu},
  data() {
    return {}
  }
}
</script>

<style scoped>
.el-container {
  height: 100vh;
  width: 100vw;
}
.el-aside {
  background-color: #2F4156;
}
.el-main {
  color: #333;
  height: 91.5%;
  padding: 0;
  background-color: white;
}
</style>