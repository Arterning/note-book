
一句话概括：for in是遍历（object）键名，for of是遍历（array）键值。



`for...in` 循环只遍历可枚举属性（包括它的原型链上的可枚举属性）

```js
var obj = {a:1, b:2, c:3};
    
for (let key in obj) {
  console.log(key);
}

// a
// b
// c
```



**`for...of`语句**在[可迭代对象](https://link.segmentfault.com/?enc=f27ZJinHud555rUkJJ9EiA%3D%3D.5AEAyXpN6dF7y3s3VReD7GfZGbasI5BaYDOAOe%2BdD0gnAT13vY8RKdKLmSKQwgjJS%2Ff5QdjYGxQ6aGWca3AZ7tVSGf0%2BJZPoqjXAKjXzot8%3D)（包括[`Array`](https://link.segmentfault.com/?enc=MhQvdEPLW5%2FS179XyHR9dg%3D%3D.Weq2LVjXdvoBmVedHQmLejJ73u6NZQ%2BNUNndNdTsvtzKDLOF0kv3wzWzo3T76KyhE570NDcPsHILoA6Njwh%2BiFkCnMy8sv5FKKdDJFP5Isg%3D)，[`Map`](https://link.segmentfault.com/?enc=Rz5bl8mmvIgoAr30i0c2OA%3D%3D.in37z0dBeviy1s78buHRL7BIltkQpfV1qICc2Jh5gm5qTjq9zT6V69hipwI1Vikf6TBrlsjVamimsa2zidSjw8GbbUe%2FpUnjhSA2E44OkhI%3D)，[`Set`](https://link.segmentfault.com/?enc=tsC0HOgrKybOyrip6yMB7w%3D%3D.93vOPhAKHMaJ8wBnqTGR%2FnTpy%2BzSNdwFgD97GdjKvkjoK7FreaMAo2hLY8Gh7gnOvnFOVZ%2Bp4Yk6prArDBW5OOrsgIIaE21KReQCPslLgxKVyH2MG3ZqUsTfw7YtSDKj)，[`String`](https://link.segmentfault.com/?enc=1SO7bbIzIagxJvvSPsfJAw%3D%3D.fAntAk7KtuooIBe6dfVQd%2Fezr7PVOFNa%2F3C0YsTbxfBWLHruNTwZBOM5un8NgRK%2B2VJrCo1fHnVVrgO7qekXYxZWBaQutHDLIgarM%2BT04rA%3D)，[`TypedArray`](https://link.segmentfault.com/?enc=5UIwomVondzrUsMWWaWsOw%3D%3D.Y9MmHC9pUYs0dqfzyWZ5izPrXSWokkuzGhU8aBj9igvdUVrZ6TExn%2B8ieqHr%2FC3muk0zspCHT7Jimt8%2BmXUeSWftVZEqrOm4Szo%2BJKfFfAajRRMFmwHtsQEt6kckSrtr)，[arguments](https://link.segmentfault.com/?enc=J2voBJ8aD0KkcGM%2BLP4acg%3D%3D.3goG7Hjwr7ODAofsjSGxtuw3ptEX8E3t72HJ5t5jUrQ%2FoD6g7O6ks3zMOPMkJvAQyQWu7UJN2P6alrS0hRXSr2jTdm%2FPVYtSZMVs6bgiJLZkllyRsNPtEThhLulB%2B4u%2F4eUfRsHiwJYjMTfhr0PI1w%3D%3D) 对象等等）上创建一个迭代循环，调用自定义迭代钩子，并为每个不同属性的值执行语句
```js
const array1 = ['a', 'b', 'c'];

for (const val of array1) {
  console.log(val);
}

// a
// b
// c
```

