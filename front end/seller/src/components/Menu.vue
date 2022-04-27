<template>
  <el-menu default-active="/main/page"
           :uniqueOpened="true"
           :style=" { width: '200px' } "
           background-color="#324054"
           text-color="#F3F3F4" router>
    <el-menu-item index="/main/page" @click="selectMenu({title: '首页'}, {path: '',afterName: ''})">
      <i class="el-icon-s-home menu-icon"></i>
      <span>首页</span>
    </el-menu-item>
    <el-sub-menu :index="menu.name" v-for="menu in menuList" :key="menu.title">
      <template #title>
        <i :class="menu.icon" class="menu-icon"></i>
        <span>{{menu.title}}</span>
      </template>
      <div class="menu-item-back">
        <el-menu-item :index="item.name" v-for="item in menu.children" :key="item.name" @click="selectMenu(menu, item)">
            <i :class="item.icon" class="menu-icon"></i>
            <span slot="title">{{item.title}}</span>
        </el-menu-item>
      </div>
    </el-sub-menu>
  </el-menu>
</template>

<script>
export default {
  name: "Menu",
  data() {
    return {
      menuList: [
        {
          name: '/product/list',
          icon: 'el-icon-s-goods',
          title: '商品',
          children: [
            {
              submenu: 2,
              name: '/product/list',
              icon: 'el-icon-goods',
              title: '商品列表',
            },
            {
              submenu: 2,
              name: '/product/add',
              icon: 'el-icon-circle-plus-outline',
              title: '添加商品',
            },
            {
              submenu: 2,
              name: '/product/category',
              icon: 'el-icon-price-tag',
              title: '类目管理',
            }
          ]
        },
        {
          name: '/order/list',
          icon: 'el-icon-s-order',
          title: '订单',
          children: [
            {
              name: '/order/list',
              icon: 'el-icon-tickets',
              title: '订单列表',
            },
            {
              name: '/order/set',
              icon: 'el-icon-setting',
              title: '订单设置',
            }
          ]
        }
      ]
    }
  },
  methods: {
    selectMenu(menu, item) {
      let newItem;
      if (menu.title === '首页'){
        newItem = {
          title: '首页',
          path: '',
          afterName: ''
        }
      }
      if (menu.title === '商品'){
        newItem = {
          title: menu.title,
          path: menu.name,
          afterName: item.title
        }
      }
      if (menu.title === '订单'){
        newItem = {
          title: menu.title,
          path: menu.name,
          afterName: item.title
        }
      }
      if (menu.title === '用户'){
        newItem = {
          title: menu.title,
          path: menu.name,
          afterName: item.title
        }
      }
      this.$store.commit("addNav", newItem)
    }
  }
}
</script>

<style scoped>
.menu-icon {
  margin-bottom: 3px;
  color: #fff;
}
.menu-item-back{
  background-color: #222D3C;
}
</style>
<style>
.el-menu-item {
  background: none !important;
  color: #ecebeb;
}
.el-menu-item.is-active {
  background-color: #409EFF !important;
  color: #fff;
}
.el-submenu__title:hover {
  background: none !important;
}
</style>