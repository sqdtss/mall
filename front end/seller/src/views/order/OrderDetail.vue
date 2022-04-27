<template>
  <el-row>
    <el-col :span="20" :offset="2"><br><br>
      <el-descriptions direction="vertical" :column="6" border>
        <el-descriptions-item label="订单编号">{{order.id}}</el-descriptions-item>
        <el-descriptions-item label="提交时间">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ order.created }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="用户昵称">{{order.nickName}}</el-descriptions-item>
        <el-descriptions-item label="订单状态" :span="2">
          <el-tag v-if="order.status === '待付款'" size="mini" type="danger">待付款</el-tag>
          <el-tag v-if="order.status === '待发货'" size="mini" type="primary">待发货</el-tag>
          <el-tag v-if="order.status === '配送中'" size="mini" type="primary">配送中</el-tag>
          <el-tag v-if="order.status === '待收货'" size="mini" type="primary">待收货</el-tag>
          <el-tag v-if="order.status === '已完成'" size="mini" type="success">已完成</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="合计" :span="2">
          <span> ¥ </span>
          <span style="font-weight: 600;font-size: 16px;">{{order.totalPrice}}</span>
        </el-descriptions-item>
      </el-descriptions><br>
    </el-col>
    <el-col :span="20" :offset="2">
      <el-descriptions direction="vertical" :column="5" border>
        <el-descriptions-item label="收货人">{{order.name}}</el-descriptions-item>
        <el-descriptions-item label="手机号码">{{order.mobile}}</el-descriptions-item>
        <el-descriptions-item label="邮政编码">{{order.postalCode}}</el-descriptions-item>
        <el-descriptions-item label="收货地址" :span="2">
          {{order.province + order.city + order.district + order.detailedAddress}}
        </el-descriptions-item>
      </el-descriptions><br>
    </el-col>
    <el-col :span="20" :offset="2">
      <el-descriptions direction="vertical" border>
        <el-descriptions-item label="商品信息">
          <el-table
              :data="productItem"
              height="280"
              stripe>
            <el-table-column
                prop="mainImage"
                label="主图"
                width="100">
              <template #default="scope">
                <el-image
                    style="width: 50px; height: 50px"
                    :src="scope.row.mainImage"></el-image>
              </template>
            </el-table-column>
            <el-table-column
                prop="title"
                label="标题"
                width="370">
            </el-table-column>
            <el-table-column
                prop="name"
                label="名称"
                width="200">
            </el-table-column>
            <el-table-column
                prop="price"
                label="价格"
                width="200">
              <template #default="scope">
                <span>¥ {{scope.row.price}}</span>
              </template>
            </el-table-column>
          </el-table><br>
        </el-descriptions-item>
      </el-descriptions>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "OrderDetail",
  data() {
    return {
      order: {
        id: 0,
        created: '',
        nickName: '',
        status: '',
        totalPrice: '',
        name: '',
        mobile: '',
        postalCode: '',
        province: '',
        city: '',
        district: '',
        detailedAddress: '',
      },
      productItem: null
    }
  },
  mounted() {
    this.getOrderDetail();
    this.order.totalPrice = this.$route.params.totalPrice
  },
  methods: {

    // 获取订单详情
    getOrderDetail() {
      this.$axios.get('/order/detail',{
        params: {
          id: this.$route.params.id
        }
      }).then(response => {
        let res = response.data.data;
        this.order.id =  res.id;
        this.order.created = res.created;
        this.order.nickName = res.nickName;
        this.order.status = res.status;
        this.order.totalPrice = res.totalPrice;
        this.order.name = res.name;
        this.order.mobile = res.mobile;
        this.order.postalCode = res.postalCode;
        this.order.province = res.province;
        this.order.city = res.city;
        this.order.district = res.district;
        this.order.detailedAddress = res.detailedAddress;
        this.productItem = res.productItem;
      })
    }
  }
}
</script>

<style scoped>

</style>