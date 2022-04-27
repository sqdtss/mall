<template>
  <el-form ref="ruleForm" :rules="rules" :model="product" label-width="100px">
    <div class="main-card">
      <div class="main-card-title"><h4>基本信息</h4></div>
      <div class="main-card-content">
        <el-form-item label="类目" prop="categoryId">
          <el-cascader v-model="product.categoryId"
                       :options="categoryOptions"
                       placeholder="请选择"
                       @focus="categoryOption"
                       clearable/>
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="product.title" maxlength="30" style="width: 80%;" show-word-limit/>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="product.description" maxlength="30" style="width: 80%;" show-word-limit/>
        </el-form-item>
        <el-form-item label="价格" label-width="100px" prop="price">
          <el-input v-model.number="product.price" style="width: 50%;">
            <template #append>元/份</template>
          </el-input>
        </el-form-item>
        <el-form-item label="数量" label-width="100px" prop="amount">
          <el-input v-model.number="product.amount" style="width: 50%;">
            <template #append>份</template>
          </el-input>
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input v-model="product.name" style="width: 50%;"/>
        </el-form-item>
        <el-form-item label="重量" prop="weight">
          <el-input v-model.number="product.weight" style="width: 50%;">
            <template #append>g</template>
          </el-input>
        </el-form-item>
      </div>
      <div class="main-card-content">
        <el-form-item label="品牌" prop="brand">
          <el-input v-model="product.brand" style="width: 50%;"/>
        </el-form-item>
        <el-form-item label="产地" prop="origin">
          <el-input v-model="product.origin" maxlength="30" style="width: 80%;" show-word-limit/>
        </el-form-item>
        <el-form-item label="保质期" prop="shelfLife">
          <el-input v-model.number="product.shelfLife" style="width: 50%;">
            <template #append>天</template>
          </el-input>
        </el-form-item>
        <el-form-item label="净含量" prop="netWeight">
          <el-input v-model.number="product.netWeight" style="width: 50%;">
            <template #append>g</template>
          </el-input>
        </el-form-item>
        <el-form-item label="使用方式" prop="useWay">
          <el-input v-model="product.useWay" style="width: 50%;"/>
        </el-form-item>
        <el-form-item label="包装方式" prop="packingWay">
          <el-input v-model="product.packingWay" style="width: 50%;"/>
        </el-form-item>
        <el-form-item label="存储条件" prop="storageConditions">
          <el-input v-model="product.storageConditions" style="width: 50%;"/>
        </el-form-item>
      </div>
    </div>
    <div class="main-card">
      <div class="main-card-title"><h4>图文描述</h4></div>
      <div class="main-card-content">
        <el-form-item label="商品主图" prop="detailImage">
          <el-image class="image-upload" :src="product.mainImage" />
        </el-form-item>
        <el-form-item label="上传图片">
          <el-upload
              action="http://localhost:8000/web/upload"
              :headers="{'token': token}"
              :limit="1"
              :show-file-list="false"
              :on-success="handleMainImageSuccess"
              name="image">
            <el-button icon="el-icon-upload" size="mini" type="primary">点击上传</el-button>
            <template #tip>
              <div>只能上传 jpg/png 文件，且不超过 500kb</div>
            </template>
          </el-upload>
        </el-form-item>
      </div>
      <div class="main-card-content">
        <el-form-item label="商品详情" prop="detailImage">
          <el-image class="image-upload" :src="product.detailImage" alt=""/>
        </el-form-item>
        <el-form-item label="上传图片">
          <el-upload
              action="http://localhost:8000/web/upload"
              :headers="{'token': token}"
              :limit="1"
              :show-file-list="false"
              :on-success="handleDetailImageSuccess"
              name="image">
            <el-button icon="el-icon-upload" size="mini" type="primary">点击上传</el-button>
            <template #tip>
              <div>只能上传 jpg/png 文件，且不超过 500kb</div>
            </template>
          </el-upload>
        </el-form-item>
      </div>
    </div>
    <div class="main-card">
      <div class="main-card-title"><h4>商品服务</h4></div>
      <div class="main-card-content">
        <el-form-item label="发货" prop="delivery">
          <el-input v-model="product.delivery" style="width: 60%;"></el-input>
        </el-form-item>
      </div>
      <div class="main-card-content">
        <el-form-item label="保障" prop="assurance">
          <el-input v-model="product.assurance" style="width: 60%;"></el-input>
        </el-form-item>
      </div>
    </div>
    <div class="button-card">
      <el-button type="danger"
                 icon="el-icon-delete"
                 class="button-space"
                 @click="resetForm('ruleForm')">清除编辑</el-button>
      <el-button type="success"
                 icon="el-icon-folder-add"
                 class="button-space"
                 @click="submitForm('ruleForm', 2)">保存草稿</el-button>
      <el-button type="primary"
                 class="button-space"
                 icon="el-icon-circle-check"
                 @click="submitForm('ruleForm', 1)">发布商品</el-button>
    </div>
  </el-form>
</template>

<script>
export default {
  name: "AddProduct",
  data() {
    return {
      product: {
        categoryId: '',
        title: '',
        description: '',
        price: '',
        amount: '',
        mainImage: '',
        delivery: '',
        assurance: '',
        name: '',
        weight: '',
        brand: '',
        origin: '',
        shelfLife: '',
        netWeight: '',
        useWay: '',
        packingWay: '',
        storageConditions: '',
        detailImage: '',
        status: '',
      },
      rules: {
        categoryId: [
          {
            required: true,
            message: '请选择一个类目',
          },
        ],
        title: [
          {
            required: true,
            message: '请输入一个标题',
          }
        ],
        price: [
          {
            required: true,
            message: '价格不能为空'
          },
          {
            type: 'number',
            message: '价格必须为数字'
          }
        ],
        amount: [
          {
            required: true,
            message: '数量不能为空'
          },
          {
            type: 'number',
            message: '数量必须为数字'
          }
        ],
        mainImage: {
          required: true,
          message: '请上传一张图片',
        },
        delivery: {
          required: true,
          message: '请输入发货信息',
        },
        assurance: {
          required: true,
          message: '请输入保障信息',
        },
        name: {
          required: true,
          message: '请输入商品名称',
        },
        weight: {
          required: true,
          message: '请输入商品重量',
        },
        shelfLife: {
          required: true,
          message: '请输入保质期',
        },
        netWeight: {
          required: true,
          message: '请输入净含量',
        },
        detailImage: {
          required: true,
          message: '请上传一张图片',
        },
      },
      categoryOptions: null,
      dialogVisible: false,
      token: '',
    }
  },
  mounted() {
    this.token = localStorage.getItem('token')
  },
  methods: {

    // 获取类目选项
    categoryOption() {
      this.$axios.get('/category/option').then((response) => {
        this.categoryOptions = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },

    // 商品主图上传成功响应结果处理
    handleMainImageSuccess(response) {
      if (response.code === 200) {
        this.product.mainImage = response.data;
      }
    },

    // 商品详情图片上传成功响应结果处理
    handleDetailImageSuccess(response) {
      if (response.code === 200) {
        this.product.detailImage = response.data;
      }
    },

    // 重置表单
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },

    // 表单提交（保存或发布商品）
    submitForm(formName, status) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$axios.post('/product', {
            categoryId: this.product.categoryId[1],
            title: this.product.title,
            description: this.product.description,
            price: this.product.price,
            amount: this.product.amount,
            mainImage: this.product.mainImage,
            delivery: this.product.delivery,
            assurance: this.product.assurance,
            name: this.product.name,
            weight: this.product.weight,
            brand: this.product.brand,
            origin: this.product.origin,
            shelfLife: this.product.shelfLife,
            netWeight: this.product.netWeight,
            useWay: this.product.useWay,
            packingWay: this.product.packingWay,
            storageConditions: this.product.storageConditions,
            detailImage: this.product.detailImage,
            status: status
          }).then((response) => {
            if (response.data.code === 200) {
              if (status === 0) {
                this.$store.commit('setPageTitle', '保存成功')
              } else {
                this.$store.commit('setPageTitle', '发布成功')
              }
              this.$router.push({name: 'resultPage'})
            }
          }).catch((error) => {
            console.log(error)
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
.main-card {
  float: left;
  width: 90%;
  height: auto;
  margin: 2% 5% 0 5%;
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
  width: 50%;
  padding: 2% 0;
  height: auto;
}
.image-upload {
  width: 100px;
  height: 100px;
  border-radius: 6px;
}

.button-card {
  float: left;
  width: 90%;
  height: auto;
  margin: 2% 5% 0 5%;
  padding-bottom: 50px;
  text-align: center;
  border-radius: 5px;
}

.button-space {
  margin: 2%;
  width: 16% !important;
}
</style>