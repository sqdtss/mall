// pages/collection/collection.js
import request from '../../utils/request'

Page({
  data: {
    product: null,
    show: null
  },
  onClick: function (event) {
    wx.navigateTo({
      url: '/pages/product/detail/detail?id=' + event.currentTarget.id
    })   
  },
  async deleteAll() {
    let res = await request.DELETE('/collection/delete', {
      userId: wx.getStorageSync('uid')
    })
    if(res.data.code === 200){
      this.onShow()
    }
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
    let res = await request.GET('/collection/list', {
      userId: wx.getStorageSync('uid')
    })
    this.setData({
      product: res.data.data
    })
    if(res.data.data === null){
      this.setData({ show: true })
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