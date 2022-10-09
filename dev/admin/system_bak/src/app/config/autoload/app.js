export default readonly({
    webTitle: process.env.VUE_APP_WEB_TITLE,   //网站标题
    router: {
        basePath: process.env.VUE_APP_BASE_PATH,    //设置路由基础路径
        whiteList: ['/login'],    //设置路由白名单，不用验证登录
        indexRedirect: '/index/index',    //首页'/'路由的跳转路径
        layoutName: '/',    //首页'/'路由的名称，用于动态从后台注册新路由
        cacheRoute: {   //缓存路由组件设置，即页面缓存
            constExclude: [],    //固定不做缓存的路由路径（对应的页面将永远不做缓存）
            //constExclude: ['/profile'],
            max: 10,    //缓存组件最大数量
        },
    },
    apiScene: {
        name: process.env.VUE_APP_API_SCENE_NAME,  //场景名称。作用：设置http的请求头
        code: process.env.VUE_APP_API_SCENE_CODE,  //场景标识。与后台代码配合使用，多后台情况下，复制一份前端代码，更改这个参数，可方便开发。
    },
    accessToken: {
        storage: process.env.VUE_APP_ACCESS_TOKEN_STORAGE === 'localStorage' ? localStorage : sessionStorage, //存储方式。直接填写localStorage对象和sessionStorage对象
        name: process.env.VUE_APP_ACCESS_TOKEN_NAME,   //accessToken名称。作用：设置http的请求头；在storage中存储的键名
        activeTimeName: process.env.VUE_APP_ACCESS_TOKEN_ACTIVE_TIME_NAME, //活跃时间名称。作用：在storage中存储的键名
        activeTimeout: parseInt(process.env.VUE_APP_ACCESS_TOKEN_ACTIVE_TIMEOUT),  //失活时间，大于0生效，即验证是否失活（单位：毫秒时间戳。当前时间与活跃时间相差超过该值，判定失活，删除accessToken）
    },
    http: {
        host: process.env.VUE_APP_HTTP_HOST,    //前后端域名一致时可设置为空，这样上线后就不用改
        timeout: parseInt(process.env.VUE_APP_HTTP_TIMEOUT)   //超时时间。0不限制
    }
})