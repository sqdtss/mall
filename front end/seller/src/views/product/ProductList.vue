<template>
  <div>
    <el-form ref="query" :model="query">
      <div class="main-card">
        <div class="main-card-title"><h4>筛选条件</h4></div>
        <div class="main-card-content">
          <el-form-item label="编号" prop="id">
            <el-input v-model="query.id" size="small"/>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item label="标题" prop="title">
            <el-input v-model="query.title" size="small"/>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item label="类目" prop="categoryId">
            <el-cascader size="small"
                         v-model="query.categoryId"
                         :options="productOption"
                         placeholder="请选择"
                         @focus="queryCategory" clearable/>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item label="状态" prop="status">
            <el-select v-model="query.status" placeholder="请选择" size="small">
              <el-option value="1" label="已上架">已上架</el-option>
              <el-option value="2" label="未上架">未上架</el-option>
            </el-select>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item>
            <el-button @click="resetForm('query')" size="small">重置</el-button>
            <el-button type="primary" @click="queryProductList" size="small">查询</el-button>
          </el-form-item>
        </div>
      </div>
    </el-form>
    <div class="main-card">
      <div class="main-card-title"><h4>商品列表</h4></div>
      <div class="table-card-content">
        <el-table
            :highlight-current-row="true"
            :data="productList"
            height="360"
            style="width: 100%;" border stripe>
          <el-table-column
              fixed
              type="selection"
              width="40">
          </el-table-column>
          <el-table-column
              label="商品标题"
              width="260">
            <template #default="scope">
              <div style="display: inline-flex;">
                <div style="padding: 5px;">
                  <el-image style="width: 60px; height: 60px" :src="scope.row.mainImage"/>
                </div>
                <div style="padding: 5px;">
                  <div style="font-weight: 450;">{{ scope.row.title }}</div>
                  <div style="padding-top: 8px;">ID：{{ scope.row.id }}</div>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column
              prop="price"
              label="价格"
              width="100">
            <template #default="scope">
              <span>¥ {{ scope.row.price }}</span>
            </template>
          </el-table-column>
          <el-table-column
              prop="amount"
              label="库存"
              width="100">
          </el-table-column>
          <el-table-column
              prop="sales"
              label="销量"
              width="100">
          </el-table-column>
          <el-table-column
              label="是否上架"
              width="80">
            <template #default="scope">
              <el-tag size="small" type="primary" v-if="scope.row.status === 1">出售中</el-tag>
              <el-tag size="small" type="success" v-if="scope.row.status === 2">仓库中</el-tag>
            </template>
          </el-table-column>
          <el-table-column
              label="创建时间"
              width="180">
            <template #default="scope">
              <i class="el-icon-time"></i>
              <span style="margin-left: 10px">{{ scope.row.created }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150">
            <template #default="scope">
              <el-button
                  size="mini"
                  type="text"
                  @click="editProduct(scope.$index, scope.row)">编辑
              </el-button>
              <el-button
                  size="mini"
                  type="text"
                  style="color: red;"
                  @click="deleteProduct(scope.$index, scope.row)">删除
              </el-button>
              <el-button
                  size="mini"
                  type="text"
                  v-if="scope.row.status === 2"
                  @click="changeStatus(scope.row, 1)">上架
              </el-button>
              <el-button
                  size="mini"
                  type="text"
                  v-if="scope.row.status === 1"
                  @click="changeStatus(scope.row, 2)">下架
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
  name: "ProductList",
  data() {
    return {
      query: {
        id: '',
        title: '',
        categoryId: '',
        status: '',
      },
      productList: [],
      productOption: [],
      total: null,
      currentPage: 1,
      size: 6,
      tableData: null,
    }
  },
  mounted() {
    this.queryProductList();
  },
  methods: {

    handleCurrentChangePrev(val) {
      this.currentPage = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.queryProductList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.currentPage = val;
      console.log(`下一页: ${val}`);
    },
    changeStatus(row, status) {
      this.$axios.put('/product/status', {
        id: row.id,
        status: status
      }).then(() => {
        this.queryProductList();
      }).catch((error) => {
        console.log(error);
      })
    },

    // 查询类目选项
    queryCategory() {
      this.$axios.get('/category/option').then((response) => {
        this.productOption = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },
    // 按条件查询商品信息，显示商品列表
    queryProductList() {
      this.$axios.get('/product/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size,
          id: this.query.id,
          title: this.query.title,
          categoryId: this.query.categoryId[1],
          status: this.query.status
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.productList = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },

    // 重置表单
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.queryProductList();
    },

    // 编辑商品
    editProduct(index, row) {
      console.log(index)
      this.$router.push({
        name: 'editProduct',
        params: {
          id: row.id
        }
      });
    },

    // 删除商品
    deleteProduct(index, row) {
      this.$confirm('此操作将永久删除该信息, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$axios.delete('/product', {
          params: {
            id: row.id
          }
        }).then((response) => {
          this.queryProductList();
          console.log(response)
        }).catch((error) => {
          console.log(error);
        })
      }).catch();
      console.log(index, row);
    },
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
  width: 16%;
  padding-left: 4%;
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