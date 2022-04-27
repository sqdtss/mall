// pages/address/edit/edit.js
import { areaList } from '@vant/area-data';
import request from '../../../utils/request';
import Dialog from '@vant/weapp/dialog/dialog';

Page({
  data: {
    show: false,
    areaList,
    addressId: '',
    name: '',
    mobile: '',
    postalCode: '',
    province: '',
    city: '',
    district: '',
    detailedAddress: '',
    isDefault: '',
    area: '',
    checked: false,
  },
  onChange({detail}){
    this.setData({ checked: detail })
    console.log(detail)
    if (detail) {
      this.setData({ isDefault: 1 })
    } else {
      this.setData({ isDefault: 2 })
    }
  },
  showPopup() { this.setData({ show: true }) },
  onClickcancel() { this.setData({ show: false }) },
  onClickConfirm(event) {
    this.setData({ show: false });
    this.setData({ province: event.detail.values[0].name });
    this.setData({ city: event.detail.values[1].name });
    this.setData({ district: event.detail.values[2].name });
    this.setData({ area: event.detail.values[0].name + " " + event.detail.values[1].name + " " + event.detail.values[2].name});
  },
  deleteAddress(event) {
    let _this = this;
    Dialog.confirm({
      message: '确定删除吗？',
    }).then(() => {
      _this.delete(event);
      _this.onShow();
    });
  },
  async delete(event) {
    let res = await request.DELETE('/address/delete', { 
      addressId: this.data.addressId
    })
    if(res.data.code === 200){
      wx.navigateBack({
        url: '/pages/address/list/list',
      })
    }
  },
  async submitForm(options) {
    let res = await request.PUT('/address/update',{
      name: this.data.name,
      mobile: this.data.mobile,
      postalCode: this.data.postalCode,
      postalCode: this.data.postalCode,
      province: this.data.province,
      city: this.data.city,
      district: this.data.district,
      detailedAddress: this.data.detailedAddress,
      isDefault: this.data.isDefault,
      id: this.data.addressId,
      userId: wx.getStorageSync('uid')
    })
    if(res.data.code === 200){
      wx.navigateBack({ url: '/pages/address/list/list' })
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: async function (options) {
    let res = await request.GET('/address/info',{ addressId: options.id });
    if(res.data.code === 200){
      var response = res.data.data;
      this.setData({
        addressId: response.id,
        name: response.name,
        mobile: response.mobile,
        postalCode: response.postalCode,
        province: response.province,
        city: response.city,
        district: response.district,
        detailedAddress: response.detailedAddress,
        area: response.province + ' ' + response.city + ' ' + response.district,
        isDefault: response.isDefault
      })
      if (this.data.isDefault === 1){
        this.setData({ checked: true })
      }
    }
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