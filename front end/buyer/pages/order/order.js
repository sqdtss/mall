// pages/order/order.js
import request from '../../utils/request'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    orderList: null,
    orderStatus: ['待付款', '待发货', '配送中', '待收货'],
    optionsId: null,
    description: '',
    show: null,
    buttonName: '立即支付'
  },

  onClick: function (event) {
    wx.navigateTo({
      url: '/pages/product/detail/detail?id=' + event.currentTarget.id
    })   
  },

  async orderList(){
    let res = await request.GET('/order/list', {
      userId: wx.getStorageSync('uid')
    })
    this.setData({
      product: res.data.data
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.setData({ optionsId: options.id })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: async function () {
    let res = await request.GET('/order/list',{
      userId: wx.getStorageSync('uid'),
      status: this.data.orderStatus[this.data.optionsId]
    });
    if(res.data.code === 200){
      this.setData({ orderList: res.data.data })
    }
    if(res.data.message === '暂无订单'){
      this.setData({
        show: true,
        description: res.data.message
      })
    }
    console.log("AAAa" + this.data.optionsId)
    let _this = this
    switch (this.data.optionsId) {
      case '0': _this.setData({ buttonName: '立即支付' }); break;
      case '1': _this.setData({ buttonName: '提醒发货' }); break;
      case '2': _this.setData({ buttonName: '查看物流' }); break;
      case '3': _this.setData({ buttonName: '确认收货' }); break;
    }
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})