// custom-tab-bar/index.js
Component({
  data: {
    active: 0,
    list: [
      {
        "url": "/pages/main/home/home",
        "icon": {
          "normal": "/images/home-o.png",
          "active": "/images/home.png"
        },
        "text": "首页"
      },
      {
        "url": "/pages/main/classify/classify",
        "icon": {
          "normal": "/images/sort-o.png",
          "active": "/images/sort.png"
        },
        "text": "分类"
      },
      {
        "url": "/pages/main/cart/cart",
        "icon": {
          "normal": "/images/cart-o.png",
          "active": "/images/cart.png"
        },
        "text": "购物车"
      },
      {
        "url": "/pages/main/mine/mine",
        "icon": {
          "normal": "/images/user-o.png",
          "active": "/images/user.png"
        },
        "text": "我的"
      }
    ]
  },
  methods: {
    onChange(e) {
        console.log(e,'e')
        this.setData({ active: e.detail });
        wx.switchTab({
          url: this.data.list[e.detail].url
        });
    },
    init() {
        const page = getCurrentPages().pop();
        this.setData({
        　  active: this.data.list.findIndex(item => item.url === `/${page.route}`)
        });
    }
   }
  })