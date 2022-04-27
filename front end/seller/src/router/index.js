import {createRouter, createWebHashHistory} from 'vue-router';
import Login from "../views/Login";
import OrderList from "../views/order/OrderList";
import AddProduct from "../views/product/AddProduct";
import ProductList from "../views/product/ProductList";
import CategoryManage from "../views/product/CategoryManage";
import Home from "../views/Home";
import ResultPage from "../views/result/ResultPage";
import OrderDetail from "../views/order/OrderDetail";
import EditProduct from "../views/product/EditProduct";
import MainPage from "../views/index/MainPage";
import OrderSet from "../views/order/OrderSet";

const routes = [
    {
        path: '/',
        component: Login,
    },
    {
        path: '/home',
        component: Home,
        redirect: '/main/page',
        children: [
            {
                path: '/main/page',
                name: 'mainPage',
                component: MainPage,
            },
            {
                path: '/product/list',
                name: 'productList',
                component: ProductList,
            },
            {
                path: '/product/add',
                name: 'addProduct',
                component: AddProduct,
            },
            {
                path: '/product/edit',
                name: 'editProduct',
                component: EditProduct,
            },
            {
                path: '/product/category',
                name: 'productCategory',
                component: CategoryManage,
            },
            {
                path: '/order/list',
                name: 'orderList',
                component: OrderList,
            },
            {
                path: '/order/set',
                name: 'orderSet',
                component: OrderSet,
            },
            {
                path: '/order/detail',
                name: 'orderDetail',
                component: OrderDetail,
            },
            {
                path: '/result/page',
                name: 'resultPage',
                component: ResultPage
            }
        ],
    },
]

const router = createRouter({
    history: createWebHashHistory(), routes
})

export default router;