import { createStore } from 'vuex'

const store = createStore({
    state() {
        return {

            // 动态导航
            navigation: {
                title: '',
                path: '',
                afterName: ''
            },

            // 结果页标题状态
            pageTitle: ''
        }
    },
    mutations: {
        addNav(state, item) {
            state.navigation.title = item.title;
            state.navigation.path = item.path;
            state.navigation.afterName = item.afterName;
        },
        setPageTitle(state, title) {
            state.pageTitle = title
        }
    },
    actions: {},
    modules: {}
})

export default store;