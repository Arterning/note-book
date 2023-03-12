/**
 * @description: 组织架构树，数据处理
 * @param {Object} raw
 * @author: Hemingway
 */
export default function handleTreeData(raw) {
  const { orgInfo, childList } = raw

  raw.label = orgInfo.name
  raw.id = orgInfo.id

  if (childList && childList.length > 0) {
    childList.forEach(child => {
      handleTreeData(child)
    })
  }
}
