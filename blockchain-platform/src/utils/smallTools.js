export function empDictValues(dictData) {
  let arr = []
  if (Object.prototype.toString.call(dictData) === '[object Object]') {
    const el = dictData
    for (let k in el) {
      let e = {
        id: k,
        name: el[k]
      }
      arr.push(e)
    }
  }
  return arr
}
