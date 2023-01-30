/**
 * crud
 */
function crud() {
    //往末尾增加元素:数组名.push(元素)  : 一个或多个元素,添加到末尾
    //     数组名.unshift(元素)  : 一个或多个元素,添加到开头
    let arr = [100, 90, 88, 50, 60]
    // 1.数组名 .push(元素) : 新增一个或多个元素,添加到'末尾'
    arr.push(10)//或者是arr.push(66,88)到结尾增加二个新元素
    console.log(arr)
    //2.数组名.unshift(1)
    arr.unshift(1)//或者是arr.(1,2,3)到开头增加三个新元素
    console.log(arr)


    let arr1 = [10, 20, 30, 40, 50]
    //1.删'末尾'元素:   数组名.pop()
    arr1.pop()
    console.log(arr1)//[10,20,30,40]
    //删'开头'元素:     数组名.shift()
    arr1.shift()
    console.log(arr1) //[20,30,40]
    arr1.splice(2, 3) //从第二个下标开始,删除3个元素



    // update
    //2.1 如果下标存在,则修改元素值
    arr[2] = 100
    //2.2 如果下标不存在，则是动态新增(开发中使用不多)
    arr[10] = 88

    
}