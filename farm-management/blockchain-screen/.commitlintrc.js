const types = [
  'build', // 构建执行
  'chore', // 构建工具相关
  'ci', // CI 相关
  'docs', // 文档更新
  'feat', // 新功能
  'fix', // bug 修复
  'perf', // 性能优化
  'refactor', // 功能重构
  'revert', // 回滚操作
  'style', // 样式变动
  'test' // 单元测试
]

module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [2, 'always', types],
    'type-case': [0],
    'type-empty': [0],
    'scope-empty': [0],
    'scope-case': [0],
    'subject-full-stop': [0, 'never'],
    'subject-case': [0, 'never'],
    'header-max-length': [0, 'always', 72]
  }
}
