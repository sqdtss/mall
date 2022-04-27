<template>
  <div>
    <div class="card-container">
      <el-row :span="24" :gutter="30">
        <el-col :span="8">
          <el-card shadow="never">
            <div style="display: flex;">
              <img src="../../assets/goods.png" class="card-img" alt="logo"/><br>
              <div style="display: block;">
                <span class="card-span-1">商品数</span>
                <span class="card-span-2">{{ statistics.goodsCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="never">
            <div style="display: flex;">
              <img src="../../assets/order.png" class="card-img" alt="logo"/><br>
              <div style="display: block;">
                <span class="card-span-1">订单量</span>
                <span class="card-span-2">{{ statistics.orderCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="never">
            <div style="display: flex;">
              <img src="../../assets/amount.png" class="card-img" alt="logo"/><br>
              <div style="display: block;">
                <span class="card-span-1">交易金额（元）</span>
                <span class="card-span-2">{{ statistics.amount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div style="padding: 30px">
      <el-row :span="24" :gutter="30">
        <el-col :span="8">
          <div class="main-card">
            <div class="main-card-title"><h4>今日订单</h4></div>
            <div class="main-card-content">
              <el-descriptions direction="vertical" :column="2" border>
                <el-descriptions-item v-for="item in todayOrder" :key="item.title" :label="item.title">
                  {{item.data}}
                </el-descriptions-item>
              </el-descriptions><br>
              <el-alert title="订单数据实时统计，更新于 8：32" type="info" show-icon />
            </div>
          </div>
        </el-col>
        <el-col :span="16">
          <div class="main-card">
            <div class="main-card-title"><h4>每周数据总览</h4></div>
            <div class="main-card-content">
              <div id="main" style="width: 100%;height:300px;"></div>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "MainPage",
  data() {
    return {
      statistics: {
        goodsCount: '',
        orderCount: '',
        amount: '',
        visitorCount: ''
      },
      todayOrder: [
        {
          title: '待付款',
          data: ''
        },
        {
          title: '待发货',
          data: ''
        },
        {
          title: '配送中',
          data: ''
        },
        {
          title: '待收货',
          data: ''
        },
        {
          title: '已完成',
          data: ''
        },
      ],
    }
  },
  mounted() {
    this.getDataOverviewInfo()
    this.getTodayOrderInfo();
    const echarts = require('echarts/lib/echarts');
    require('echarts/lib/component/title');
    require('echarts/lib/component/toolbox');
    require('echarts/lib/component/tooltip');
    require('echarts/lib/component/grid');
    require('echarts/lib/component/legend');
    require('echarts/lib/chart/line');
    let chartDom = document.getElementById('main');
    let myChart = echarts.init(chartDom);
    let option;

    // 获取本周数据总览
    this.$axios.get("/week/data/info").then(response => {
      if (response.data.code === 200) {
        let res = response.data.data
        option = {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'cross',
              label: {
                backgroundColor: '#6a7985'
              }
            }
          },
          legend: {
            data: ['订单量', '销售额']
          },
          toolbox: {
            feature: {
              saveAsImage: {}
            }
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          xAxis: [
            {
              type: 'category',
              boundaryGap: false,
              data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
            }
          ],
          yAxis: [
            {
              type: 'value',
              name: '订单量',
              min: 0,
              max: function (value) {
                return value.max * 2;
              },
              position: 'left',
            },
            {
              type: 'value',
              name: '销售额',
              min: 0,
              max: function (value) {
                return value.max * 2;
              },
              position: 'right',
            }
          ],
          series: [
            {
              name: '订单量',
              type: 'line',
              stack: 'Total',
              yAxisIndex: 0,
              areaStyle: {},
              emphasis: {
                focus: 'series'
              },
              data: res.orders
            },
            {
              name: '销售额',
              type: 'line',
              stack: 'Total',
              yAxisIndex: 1,
              areaStyle: {},
              emphasis: {
                focus: 'series'
              },
              data: res.amount
            }
          ]
        };
        option && myChart.setOption(option);
      }
    })
  },
  methods: {

    // 获取数据统计信息（商品数、订单量、交易金额）
    getDataOverviewInfo() {
      this.$axios.get("/data/overview/info").then(response => {
        if (response.data.code === 200) {
          let res = response.data.data;
          this.statistics.goodsCount = res.goodsCount;
          this.statistics.orderCount = res.orderCount;
          this.statistics.amount = res.amount;
          if (res.goodsCount === 0 || res.orderCount === 0 || res.amount === 0 ) {
            this.statistics.goodsCount = '--';
            this.statistics.orderCount = '--';
            this.statistics.amount = '--';
          }
        }
      })
    },

    // 获取今日订单数据统计信息
    getTodayOrderInfo() {
      this.$axios.get("/today/order/data/info").then(response => {
        if (response.data.code === 200) {
          let res = response.data.data.data;
          for (let i = 0; i < 5; i++) {
            this.todayOrder[i].data = res[i]
          }
        }
      })
    },
  }
}
</script>

<style scoped>
.card-container {
  border: none;
  background-color: #f0f3f6;
  padding: 30px;
}
.card-img {
  width: 60px;
  height: 60px;
}
.card-span-1 {
  padding-top: 3px;
  padding-left: 10px;
}
.card-span-2 {
  display: block;
  font-size: 24px;
  font-weight: 500;
  padding-left: 10px;
  padding-top: 8px;
}

.main-card {
  float: left;
  width: 100%;
  height: auto;
  border-radius: 5px;
  background-color: #FAFAFA;
}

.main-card-title {
  float: left;
  padding: 12px;
}

.main-card-title h4 {
  padding-left: 5px;
  border-left: 5px solid dodgerblue;
}

.main-card-content {
  float: left;
  width: 90%;
  padding: 5%;
  height: auto;
}
</style>