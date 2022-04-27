// pages/login/login.js
import request from '../../utils/request'

Page({

  data: {
    infoMess: '',
    userName: '',
    passWord: ''
  },

  //用户名和密码输入框事件
  userNameInput:function(e){
    this.setData({
      userName: e.detail.value
    })
  },
  passWordInput:function(e){
    this.setData({
      passWord: e.detail.value
    })
  },
  //登录按钮点击事件，调用参数要用：this.data.参数；
  //设置参数值，要使用this.setData({}）方法
  loginBtnClick:function() {
    if(this.data.userName.length == 0 || this.data.passWord.length == 0) {
      this.setData({
        infoMess:'温馨提示：用户名和密码不能为空！',
      })
    } else {
      let _this = this
        _this.login();
    }
  },
  async login() {
    //发起登录请求
    let response = await request.POST('/login', { 
      Username: this.data.userName,
      Password: this.data.passWord,
    })
    if(response.data.code === 200){
      wx.setStorage({ key: "uid", data: response.data.data});
      wx.switchTab({ url: '/pages/main/home/home' })
    } else {
      this.setData({
        infoMess:'登录失败!',
      })
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function () {
    console.log('onLoad')
    var that = this
    //调用应用实例的方法获取全局数据
    app.getUserInfo(function(userInfo){
      //更新数据
      that.setData({
        userInfo: userInfo
      })
    })
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