// pages/tabbar/cart/cart.js
import Dialog from '@vant/weapp/dialog/dialog';
import request from '../../../utils/request'

Page({

  data: {
    cartItem: null,
    totalPrice: null,
    show: null,
    submitBarShow: null,
    description: '购物车竟然是空的'
  },

  // 查看商品详情
  onClick: function (event) {
    wx.navigateTo({
      url: '/pages/product/detail/detail?id=' + event.currentTarget.id
    })   
  },

  // 点击删除按钮，出现提示
  onClickDelete: function (event) {
    let _this = this;
    Dialog.confirm({
      message: '确定删除吗？',
    }).then(() => {
      _this.delete(event);
    });
  },

  // 删除购物车中的商品
  async delete (event) {
    let res = await request.DELETE('/cart/delete',{
      productId: event.currentTarget.id,
      userId: wx.getStorageSync('uid')
    })
    console.log("VVVV" + event.currentTarget.id)
    if(res.data.code === 200){ 
      this.setData({ cartItem: null })
      this.onShow()
    }
  },

  // 提交订单
  async submitOrder () {
    let res = await request.POST('/order/create',{
      userId: wx.getStorageSync('uid'),
      status: '待发货',
      nickName: wx.getStorageSync('nickName')
    })
    if(res.data.code === 200){ 
      this.clearCart()
    }
  },

  async clearCart () {
    let res = await request.DELETE('/cart/clear',{
      userId: wx.getStorageSync('uid'),
    })
    if(res.data.code === 200){ 
      wx.navigateTo({
        url: '/pages/result/pay/pay',
      })
      this.setData({
        cartItem: null
      })
    }
  },


  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    
  },

  onSubmit() {
    wx.showToast({
      title: '点击结算',
      icon: 'none'
    });
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
    this.getTabBar().init();
    let res = await request.GET('/cart/info',{
      userId: wx.getStorageSync('uid')
    })
    if(res.data.code === 200){
      this.setData({
        cartItem: res.data.data.cartItem,
        totalPrice: res.data.data.totalPrice*100,
        show: false,
      })
    }
    if(this.data.cartItem === null){
      this.setData({
        show: true,
        submitBarShow: false
      })
    } else {
      this.setData({
        show: false,
        submitBarShow: true
      })
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