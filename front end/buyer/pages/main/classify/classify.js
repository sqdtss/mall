// pages/tabbar/classify/classify.js
import request from '../../../utils/request'

Page({

  /**
   * 页面的初始数据
   */
  data: {
    items: [],
    mainActiveIndex: 0,
    product: null
  },

  async onClickNav({ detail = {} }) {
    this.setData({
      mainActiveIndex: detail.index || 0,
    });
    let res = await request.GET('/product/list',{
      categoryId: this.data.items[this.data.mainActiveIndex].id
    });
    if (res.data.code === 200){
      this.setData({
        product: res.data.data
      })
    }
  },

  sorry:function () {
    wx.showToast({ title: '暂无后续逻辑~', icon: 'none',});
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
    let res = await request.GET('/category/option');
    this.setData({ items: res.data.data })
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