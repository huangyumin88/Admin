import { defineStore } from 'pinia'

export const useKeepAliveStore = defineStore('keepAlive', {
  state: () => {
    return {
      appContainerExclude: [], //不允许缓存的路由路径列表，这里主要用于实现缓存刷新（动态设置页面组件名称name时，用路径命名，故这里面填写路径）
      appContainerMax: 10   //缓存组件最大数量
    }
  },
  getters: {
    appContainerInclude: (state): string[] => {
      const include: string[] = []
      useRouter().getRoutes().forEach((item) => {
        if (item.meta.keepAlive) {
          //include.push(item.components.default.name)
          include.push(item.path)
        }
      })
      return include
    },
  },
  actions: {
    /**
     * 删除不允许缓存的组件
     * @param {*} path  路径
     */
    removeAppContainerExclude(path: string) {
      this.appContainerExclude = this.appContainerExclude.filter((item) => {
        return item !== path
      })
    },
    /**
     * 刷新菜单标签
     *      实现流程：
     *          1：app-container.vue文件内component标签加上判断是否允许缓存，允许才显示界面（v-if="userStore.cacheRoute.exclude.indexOf(route.path) === -1"）
     *          2：设置路由不允许缓存，不显示页面
     *          3：打开路由，路由后置守卫afterEach中重新设置成允许缓存，显示页面
     * @param {*} path  菜单标签的路由路径
     */
    refreshMenuTab(path: string) {
      this.appContainerExclude.push(path)
      const currentPath = getCurrentPath()
      if (path === currentPath) {
        useRouter().push(path)
      }
    },
  }
})
