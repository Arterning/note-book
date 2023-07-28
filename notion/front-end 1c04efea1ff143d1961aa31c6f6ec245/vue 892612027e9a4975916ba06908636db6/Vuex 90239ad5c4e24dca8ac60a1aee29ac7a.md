# Vuex

****[手把手教你使用Vuex，猴子都能看懂的教程](https://juejin.cn/post/6928468842377117709)****

```jsx
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    // 定义一个name，以供全局使用
    name: '张三',
    // 定义一个number，以供全局使用
    number: 0,
    // 定义一个list，以供全局使用
    list: [
      { id: 1, name: '111' },
      { id: 2, name: '222' },
      { id: 3, name: '333' },
    ],
  },
});

export default store;

```

```jsx
import Vue from 'vue';
import App from './App';
import router from './router';
import store from './store'; // 引入我们前面导出的store对象

Vue.config.productionTip = false;

new Vue({
  el: '#app',
  router,
  store, // 把store对象添加到vue实例上
  components: { App },
  template: '<App/>',
});

```

```jsx
<template>
  <div></div>
</template>

<script>
export default {
  mounted() {
    // 使用this.$store.state.XXX可以直接访问到仓库中的状态
    console.log(this.$store.state.name);
  },
};
</script>

```

官方建议我们以上操作this.$store.state.XXX最好放在计算属性中

```jsx
export default {
  mounted() {
    console.log(this.getName);
  },
  computed: {
    getName() {
      return this.$store.state.name;
    },
  },
};
```

官方建议2： 是不是每次都写this.$store.state.XXX让你感到厌烦，你实在不想写这个东西怎么办，当然有解决方案，就像下面这样

```jsx
<script>
import { mapState } from 'vuex'; // 从vuex中导入mapState
export default {
  mounted() {
    console.log(this.name);
  },
  computed: {
    ...mapState(['name']), // 经过解构后，自动就添加到了计算属性中，此时就可以直接像访问计算属性一样访问它
  },
};
</script>

```

```jsx
...mapState({ aliasName: 'name' }),  // 赋别名的话，这里接收对象，而不是数组
```

Getter定义

```jsx
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    name: '张三',
    number: 0,
    list: [
      { id: 1, name: '111' },
      { id: 2, name: '222' },
      { id: 3, name: '333' },
    ],
  },
  // 在store对象中增加getters属性
  getters: {
    getMessage(state) {
      // 获取修饰后的name，第一个参数state为必要参数，必须写在形参上
      return `hello${state.name}`;
    },
  },
});

export default store;

```

使用Getter

```jsx
export default {
  mounted() {
    // 注意不是$store.state了，而是$store.getters
    console.log(this.$store.state.name);
    console.log(this.$store.getters.getMessage);
  },
};

```

是不是每次都写this.$store.getters.XXX让你感到厌烦，你实在不想写这个东西怎么办，当然有解决方案，官方建议我们可以使用mapGetters去解构到计算属性中，就像使用mapState一样，就可以直接使用this调用了

```jsx
<script>
import { mapState, mapGetters } from 'vuex';
export default {
  mounted() {
    console.log(this.name);
    console.log(this.getMessage);
  },
  computed: {
    ...mapState(['name']),
    ...mapGetters(['getMessage']),
  },
};
</script>

```

也可以取别名

```jsx
...mapGetters({ aliasName: 'getMessage' }),  // 赋别名的话，这里接收对象，而不是数组
```

如何修改State 需要使用Mutation

```jsx
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    name: '张三',
    number: 0,
  },
  mutations: {
    // 增加nutations属性
    setNumber(state) {
      // 增加一个mutations的方法，方法的作用是让num从0变成5，state是必须默认参数
      state.number = 5;
    },
  },
});

export default store;

```

```jsx
<script>
export default {
  mounted() {
    console.log(`旧值：${this.$store.state.number}`);
    this.$store.commit('setNumber');
    console.log(`新值：${this.$store.state.number}`);
  },
};
</script>

```

传参数给mutation

```jsx
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    name: '张三',
    number: 0,
  },
  mutations: {
    setNumber(state) {
      state.number = 5;
    },
    setNumberIsWhat(state, number) {
      // 增加一个带参数的mutations方法
      state.number = number;
    },
  },
});

export default store;

```

```jsx
<script>
export default {
  mounted() {
    console.log(`旧值：${this.$store.state.number}`);
    this.$store.commit('setNumberIsWhat', 666);
    console.log(`新值：${this.$store.state.number}`);
  },
};
</script>

```

上面的这种传参的方式虽然可以达到目的，但是并不推荐，官方建议传递一个对象进去，这样看起来更美观，对象的名字你可以随意命名，但我们一般命名为payload

```jsx
mutations: {
    setNumber(state) {
      state.number = 5;
    },
    setNumberIsWhat(state, payload) {
      // 增加一个带参数的mutations方法，并且官方建议payload为一个对象
      state.number = payload.number;
    },
  },

```

```jsx
<script>
export default {
  mounted() {
    console.log(`旧值：${this.$store.state.number}`);
    this.$store.commit('setNumberIsWhat', { number: 666 }); // 调用的时候也需要传递一个对象
    console.log(`新值：${this.$store.state.number}`);
  },
};
</script>

```

就像最开始的mapState和mapGetters一样，我们在组件中可以使用mapMutations以代替this.$store.commit('XXX')，是不是很方便呢？

```jsx
<script>
import { mapMutations } from 'vuex';
export default {
  mounted() {
    this.setNumberIsWhat({ number: 999 });
  },
  methods: {
    // 注意，mapMutations是解构到methods里面的，而不是计算属性了
    ...mapMutations(['setNumberIsWhat']),
  },
};
</script>

```

```jsx
methods: {
  ...mapMutations({ setNumberIsAlias: 'setNumberIsWhat' }), // 赋别名的话，这里接收对象，而不是数组
},
```

**这里说一条重要原则：Mutations里面的函数必须是同步操作，不能包含异步操作！（别急，后面会讲到异步）**

vuex作者不希望你将异步操作放在Mutations中，所以就给你设置了一个区域，让你放异步操作，这就是Actions

```jsx
const store = new Vuex.Store({
  state: {
    name: '张三',
    number: 0,
  },
  mutations: {
    setNumberIsWhat(state, payload) {
      state.number = payload.number;
    },
  },
  actions: {
    // 增加actions属性
    setNum(content) {
      // 增加setNum方法，默认第一个参数是content，其值是复制的一份store
      return new Promise(resolve => {
        // 我们模拟一个异步操作，1秒后修改number为888
        setTimeout(() => {
          content.commit('setNumberIsWhat', { number: 888 });
          resolve();
        }, 1000);
      });
    },
  },
});

```

```jsx
async mounted() {
  console.log(`旧值：${this.$store.state.number}`);
  await this.$store.dispatch('setNum');
  console.log(`新值：${this.$store.state.number}`);
},

```

你如果不想一直使用this.$store.dispatch('XXX')这样的写法调用action，你可以采用mapActions的方式，把相关的actions解构到methods中，用this直接调用

```jsx
<script>
import { mapActions } from 'vuex';
export default {
  methods: {
    ...mapActions(['setNum']), // 就像这样，解构到methods中
  },
  async mounted() {
    await this.setNum({ number: 123 }); // 直接这样调用即可
  },
};
</script>

```