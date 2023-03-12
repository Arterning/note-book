export default {
  // set 设置 、替换、增加
  set: function (key, val) {
    sessionStorage.setItem(key, JSON.stringify(val))
  },

  remove: function (key) {
    sessionStorage.removeItem(key)
  },

  get: function (key) {
    return JSON.parse(sessionStorage.getItem(key))
  },

  // add 增加同键值的 value 值 注意 String、[]
  add: function (key, addVal) {
    const oldVal = Storage.get(key)
    const newVal = oldVal.concat(addVal)
    Storage.set(key, newVal)
  },

  getAll: function () {
    for (var i = 0; i < sessionStorage.length; i++) {
      console.log(
        sessionStorage.key(i) +
          ' === ' +
          sessionStorage.getItem(sessionStorage.key(i))
      )
    }
  },

  // 清除本地存储
  removeAll: function () {
    sessionStorage.clear()
  }
}
