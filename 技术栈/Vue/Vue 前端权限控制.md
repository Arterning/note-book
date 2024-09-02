# Vue 前端权限控制

我们可以注册Vue插件

```javascript
export default {
  install(Vue) {
    Vue.prototype.checkPer = (value) => {
      if (value && value instanceof Array && value.length > 0) {
        const roles = store.getters && store.getters.roles
        const permissionRoles = value
        return roles.some(role => {
          return permissionRoles.includes(role)
        })
      } else {
        console.error(`need roles! Like v-permission="['admin','editor']"`)
        return false
      }
    }
  }
}
```


使其生效
```javascript
Vue.use(checkPer)
```

使用
```javascript
<el-table-column v-if="checkPer(['admin','dept:edit','dept:del'])">
```


此外还可以自定义Vue指令 `v-permissions` 首先定义Permission.js
```javascript
import store from '@/store'

export default {
  inserted(el, binding) {
    const { value } = binding
    const roles = store.getters && store.getters.roles
    if (value && value instanceof Array) {
      if (value.length > 0) {
        const permissionRoles = value
        const hasPermission = roles.some(role => {
          return permissionRoles.includes(role)
        })
        if (!hasPermission) {
          el.parentNode && el.parentNode.removeChild(el)
        }
      }
    } else {
      throw new Error(`使用方式： v-permission="['admin','editor']"`)
    }
  }
}

```

接着使用`Vue.directive` 使其生效

``` javascript
import permission from './permission'

const install = function(Vue) {
  Vue.directive('permission', permission)
}

if (window.Vue) {
  window['permission'] = permission
  Vue.use(install); // eslint-disable-line
}

permission.install = install
export default permission

```

使用
```javascript
permission: {
        add: ['admin', 'dict:add'],
        edit: ['admin', 'dict:edit'],
        del: ['admin', 'dict:del']
}

...

<el-button v-permission="permission.add"/>
```