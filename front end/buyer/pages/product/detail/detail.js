// pages/product/detail/detail.js
import Toast from '@vant/weapp/toast/toast'
import request from '../../../utils/request'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    product: null,
    productId: null,
    productCount: 1,
    show: false,
    showStepper: false,
    salesServiceList: null,
    starStatus: 0,
  },
  async onClickStar() {
    let res = await request.POST('/collection/add',{ 
      productId: this.data.productId, 
      userId: wx.getStorageSync('uid')
    });
    if(res.data.code === 200){
      Toast.success('收藏成功');
    }
    if(res.data.code === 400){
      Toast.success('已收藏');
    }
  },
  onClickCart() {
    wx.switchTab({
      url: '/pages/main/cart/cart',
    });
  },
  onChange(event) {
    this.setData({ productCount: event.detail })
  },

  sorry() {
    wx.showToast({
      title: '暂无后续逻辑~',
      icon: 'none',
    });
  },
  onClickAddCart() {
    this.setData({ show: true });
  },
  onClose() {
    this.setData({ show: false });
  },

  /**
   * 点击完成时，完成，添加购物车
   */
  async onFinishAddCart() {
    this.setData({ show: false });
    let res = await request.POST('/cart/add',{
      productId: this.data.productId,
      productCount: this.data.productCount,
      userId: wx.getStorageSync('uid')
    });
    if(res.data.code === 200){
      Toast.success('添加成功');
    }
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: async function (options) {
    console.log(options.id)
    this.setData({
      productId: options.id
    });
    let res = await request.GET('/product/detail',{ 
      productId: options.id
    });
    this.setData({ product: res.data.data })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

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