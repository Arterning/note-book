const data = [
    { 'name': 'a', 'title': 'a', 'date': 'dfd' },
    { 'name': 'b', 'title': 'b', 'date': 'dfd' },
    { 'name': 'c', 'title': 'c', 'date': 'dfd' }
]

// 工作中遇到的非常多的一个场景就是处理后端数据，已让其适配组件的数据格式
// 一般都是数组 对象 的组装
const items = []
data.forEach(e =>{
    console.log(e);
    items.push({"value": e.name, "label": e.title})
})
console.log(items);