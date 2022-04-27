// pages/browse/browse.js
import request from '../../utils/request'

Page({

  data: {
    productIds: [],
    browseInfo: null,
    show: null
  },

  onChange(event) {
    this.setData({ productIds: event.detail });
  },

  toggle(event) {
    const { index } = event.currentTarget.dataset;
    const checkbox = this.selectComponent(`.checkboxes-${index}`);
    checkbox.toggle();
  },

  noop() {},

  delete: async function() {
    let res = await request.DELETE('/browse/delete', {
      userId: wx.getStorageSync('uid'),
      productIds: this.data.productIds
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
    let res = await request.GET('/browse/list', {
      userId: wx.getStorageSync('uid')
    })
    this.setData({
      browseInfo: res.data.data
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