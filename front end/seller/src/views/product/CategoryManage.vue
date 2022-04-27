<template>
  <div>
    <el-form ref="form" :model="queryCondition" label-width="100px">
      <div class="main-card">
        <div class="main-card-title"><h4>查询条件</h4></div>
        <div class="main-card-content">
          <el-form-item label="类目名称">
            <el-input size="small" v-model="queryCondition.name"/>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item label="类目等级">
            <el-select v-model="queryCondition.value" size="small">
              <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"/>
            </el-select>
          </el-form-item>
        </div>
        <div class="main-card-content">
          <el-form-item>
            <el-button size="small" type="primary" @click="queryList" style="margin-left: 28px;">查询</el-button>
            <el-button size="small" type="primary" @click="createClick" style="margin-left: 18px;">新建</el-button>
          </el-form-item>
        </div>
      </div>
    </el-form>
    <div class="main-card">
      <div class="main-card-title"><h4>类目列表</h4></div>
      <div class="table-card-content">
        <el-table
            :data="tableData"
            @selection-change="handleSelectionChange"
            height="360"
            style="width: 100%" border stripe>
          <el-table-column
              type="selection"
              width="55">
          </el-table-column>
          <el-table-column
              prop="id"
              label="编号"
              width="100">
          </el-table-column>
          <el-table-column
              prop="name"
              label="类目名称"
              width="150">
          </el-table-column>
          <el-table-column
              prop="level"
              label="类目等级"
              width="150">
          </el-table-column>
          <el-table-column
              prop="sort"
              label="排序"
              width="150">
          </el-table-column>
          <el-table-column
              label="操作">
            <template #default="scope">
              <el-button
                  size="mini"
                  @click="updateClick(scope.$index, scope.row)">修改
              </el-button>
              <el-button
                  size="mini"
                  type="primary"
                  v-if="scope.row.level === 1"
                  @click="nextLevelCategory(scope.$index, scope.row)">下一级类目
              </el-button>
              <el-button
                  size="mini"
                  type="danger"
                  @click="deleteCategory(scope.$index, scope.row)">删除
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
    <el-dialog :title="dialogTitle" width="30%" v-model="dialogFormVisible" center>
      <el-divider>{{ category.levelTitle }}</el-divider>
      <el-form :model="category">
        <el-form-item label="类目名称" :label-width="formLabelWidth">
          <el-input v-model="category.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="类目排序" :label-width="formLabelWidth">
          <el-input v-model.number="category.sort" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="primary" v-show="updateButton" @click="updateCategory">确 定</el-button>
      <el-button type="primary" v-if="this.category.level === 1 && createButton === true"
                 @click="createCategory">下一级</el-button>
      <el-button type="primary" v-if="this.category.level === 2 && createButton === true"
                 @click="createCategory">完成</el-button>
    </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "CategoryManage",
  data() {
    return {
      options: [{
        value: '1',
        label: '一级类目'
      }, {
        value: '2',
        label: '二级类目'
      }],
      queryCondition: {
        name: '',
        value: '1',
      },
      tableData: '',
      category: {
        id: '',
        name: '',
        level: 1,
        sort: '',
        parentId: 1,
        levelTitle: ''
      },
      multipleSelection: [],
      total: null,
      currentPage: 1,
      size: 5,
      dialogFormVisible: false,
      formLabelWidth: '80px',
      dialogTitle: '',
      updateButton: false,
      createButton: false
    }
  },
  mounted() {
    this.queryCategoryList();
  },
  methods: {

    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    handleCurrentChangePrev(val) {
      this.currentPage = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.queryCategoryList();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.currentPage = val;
      console.log(`下一页: ${val}`);
    },

    // 下一级类目
    nextLevelCategory(index, row) {
      console.log(index)
      this.$axios.get('/category/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size,
          parentId: row.id
        }
      }).then((response) => {
        if (response.data.code === 200) {
          this.total = response.data.data.total;
          this.tableData = response.data.data.list;
        }
      }).catch((error) => {
        console.log(error);
      })
    },

    // 更新类目
    updateClick(index, row) {
      console.log(index, row);
      this.dialogFormVisible = true;
      this.dialogTitle = '修改类目';
      this.category.name = row.name;
      this.category.sort = row.sort;
      this.updateButton = true;
      this.createButton = false;
      this.category.id = row.id;
    },

    // 创建类目
    createClick() {
      this.dialogFormVisible = true;
      this.dialogTitle = '新建品牌';
      this.createButton = true;
      this.updateButton = false;
      this.category.parentId = 1;
      this.category.levelTitle = '一级类目';
    },

    // 更新类目
    updateCategory() {
      this.$axios.put('/category', {
        id: this.category.id,
        name: this.category.name,
        sort: this.category.sort
      }).then((response) => {
        if (response.data.code === 200) {
          console.log(response)
          this.category.id = '';
          this.category.name = '';
          this.category.sort = '';
        }
        this.dialogFormVisible = false;
        this.queryCategoryList();
      }).catch((error) => {
        console.log(error);
      })
    },

    // 创建类目
    createCategory() {
      this.$axios.post('/category', {
        name: this.category.name,
        level: this.category.level,
        sort: this.category.sort,
        parentId: this.category.parentId
      }).then((response) => {
        if (response.data.code === 200) {
          this.category.level++;
          this.category.name = '';
          this.category.sort = '';
          if (this.category.level === 2) {
            this.category.levelTitle = '二级类目';
          }
          this.category.parentId = response.data.data;
          if (this.category.level === 3) {
            this.dialogFormVisible = false;
            this.category.level = 1;
            this.queryCategoryList();
          }
        }
      }).catch((error) => {
        console.log(error);
      })
    },

    // 查询类目列表
    queryCategoryList() {
      this.$axios.get('/category/list', {
        params: {
          pageNum: this.currentPage,
          pageSize: this.size,
          name: this.queryCondition.name,
          level: this.queryCondition.value
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.tableData = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },

    // 删除类目
    deleteCategory(index, row) {
      this.$axios.delete('/category', {
        params: {
          id: row.id
        }
      }).then((response) => {
        if (response.data.code === 200) {
          this.queryCategoryList();
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