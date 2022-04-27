// pages/home/home.js
import request from '../../../utils/request'

Page({
  data: {
    background: [
      '/images/banner-1.png', 
      '/images/banner-2.png', 
      '/images/banner-3.png'
    ],
    product: null
  },
  async onClick(event) {
    wx.navigateTo({
      url: '/pages/product/detail/detail?id=' + event.currentTarget.id
    })
    await request.POST('/browse/save', {
      productId: event.currentTarget.id,
      userId: wx.getStorageSync('uid')
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    
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
    let res = await request.GET('/product/list');
    this.setData({
      product: res.data.data
    })
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