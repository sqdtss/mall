<template>
  <div>
    <el-form ref="form" :model="query" label-width="80px">
      <div class="main-card">
        <div class="main-card-title"><h4>查询条件</h4></div>
        <div class="main-card-content">
          <el-form-item label="订单ID" prop="orderId">
            <el-input v-model="query.orderId" size="small"></el-input>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item label="订单状态" prop="orderStatus">
            <el-select v-model="query.orderStatus"
                       placeholder="请选择" size="small">
              <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item>
            <el-button @click="resetForm('form')" size="small">重置</el-button>
            <el-button type="primary" @click="queryOrder" size="small">查询</el-button>
          </el-form-item>
        </div>
      </div>
    </el-form>
    <div class="main-card">
      <div class="main-card-title"><h4>订单列表</h4></div>
      <div class="table-card-content">
        <el-table
            :highlight-current-row="true"
            :data="tableData"
            height="360"
            style="width: 100%" border stripe>
          <el-table-column
              type="selection"
              width="55">
          </el-table-column>
          <el-table-column
              prop="id"
              label="订单编号"
              width="180">
          </el-table-column>
          <el-table-column
              prop="created"
              label="提交时间"
              width="200">
            <template #default="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.created }}</span>
            </template>
          </el-table-column>
          <el-table-column
              prop="nickName"
              label="用户昵称"
              width="100">
          </el-table-column>
          <el-table-column
              prop="totalPrice"
              label="总计"
              width="100">
            <template #default="scope">
              <span style="margin-left: 5px;font-weight: 500;">¥ {{ scope.row.totalPrice }}</span>
            </template>
          </el-table-column>
          <el-table-column
              prop="status"
              label="订单状态"
              width="100">
            <template #default="scope">
              <el-tag v-if="scope.row.status === '待付款'" size="mini" type="danger">待付款</el-tag>
              <el-tag v-if="scope.row.status === '待发货'" size="mini" type="primary">待发货</el-tag>
              <el-tag v-if="scope.row.status === '配送中'" size="mini" type="primary">配送中</el-tag>
              <el-tag v-if="scope.row.status === '待收货'" size="mini" type="primary">待收货</el-tag>
              <el-tag v-if="scope.row.status === '已完成'" size="mini" type="success">已完成</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="300">
            <template #default="scope">
              <el-button
                  size="mini"
                  @click="checkOrder(scope.row)">查看订单
              </el-button>
              <el-button
                  size="mini"
                  v-if="scope.row.status === '待发货'"
                  type="primary"
                  @click="updateOrder(scope.row)">订单发货
              </el-button>
              <el-button
                  size="mini"
                  v-if="scope.row.status === '已完成'"
                  type="danger"
                  @click="deleteOrder(scope.row)">删除订单
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <div class="pagination-card">
      <el-pagination
          background
          @current-change="handleCurrentChange"
          @prev-click="handleCurrentChangePrev"
          @next-click="handleCurrentChangeNext"
          :currentPage="currentPage"
          :page-size="size"
          layout="total, prev, pager, next"
          :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  name: "OrderList",
  data() {
    return {
      query: {
        orderId: '',
        orderStatus: '',
      },
      options: [
        {
          value: '待付款',
          label: '待付款',
        },
        {
          value: '待发货',
          label: '待发货',
        },
        {
          value: '配送中',
          label: '配送中',
        },
        {
          value: '待收货',
          label: '待收货',
        },
        {
          value: '已完成',
          label: '已完成',
        }
      ],
      dialogFormVisible: false,
      tableData: null,
      currentPage: 1,
      size: 10,
      total: '',
    }
  },
  mounted() {
    this.queryOrder();
  },
  methods: {
    handleCurrentChangePrev(val) {
      this.currentPage = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.queryOrder();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.currentPage = val;
      console.log(`下一页: ${val}`);
    },

    // 重置表单
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.queryOrder();
    },

    // 查询订单列表
    queryOrder() {
      this.$axios.get('/order/list', {
        params: {
          id: this.query.orderId,
          status: this.query.orderStatus,
          pageNum: this.currentPage,
          pageSize: this.size,
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.tableData = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },

    // 查看订单详情
    checkOrder(row) {
      this.$router.push({
        name: 'orderDetail',
        params: {id: row.id}
      });
    },

    // 更新订单（订单发货）
    updateOrder(row) {
      this.$axios.put('/order', {
        status: '配送中',
        id: row.id,
      }).then((response) => {
        if (response.data.code === 200) {
          this.queryOrder();
        }
      }).catch((error) => {
        console.log(error);
      })
    },

    // 删除订单（已完成订单）
    deleteOrder(row) {
      this.$axios.delete('/order', {
        params: {
          id: row.id,
        }
      }).then((response) => {
        if (response.data.code === 200) {
          this.queryOrder();
        }
      }).catch((error) => {
        console.log(error);
      })

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
  width: 30%;
  padding-left: 3%;
  height: auto;
}

.table-card-content {
  width: 96%;
  padding: 2%;
  height: auto;
}

.pagination-card {
  width: 100%;
  position: fixed;
  bottom: 0;
  padding: 1% 2%;
  background-color: white;
  height: auto;
}
</style>