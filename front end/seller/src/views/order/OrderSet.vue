<template>
  <div>
    <el-form ref="form" :inline="true" :model="order">
      <div class="main-card">
        <div class="main-card-title"><h4>订单设置</h4></div>
        <div class="main-card-content">
          <el-form-item label="订单超过" prop="expirationTime">
            <el-input v-model="order.expirationTime" size="small">
              <template #append>分，未付款，订单自动关闭</template>
            </el-input>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item>
            <el-button @click="resetForm('form')" size="small">重置</el-button>
            <el-button type="primary" @click="submitForm" size="small">确定</el-button>
          </el-form-item>
        </div>
      </div>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "OrderSet",
  data() {
    return {
      order: {
        expirationTime: ''
      }
    }
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    submitForm() {
    }
  }
}
</script>

<style scoped>
.main-card {
  float: left;
  width: 98%;
  height: auto;
  margin: 2% 1% 0 1%;
  border-radius: 5px;
  background-color: #FAFAFA;
}

.main-card-title {
  float: left;
  width: 96%;
  padding: 2%;
}

.main-card-title h4 {
  padding-left: 5px;
  border-left: 5px solid dodgerblue;
}

.main-card-content {
  float: left;
  width: 40%;
  padding: 5%;
  height: auto;
}
</style>