## core.autocrlf
跨平台协作开发是常有的，不统一的换行符确实对跨平台的文件交换带来了麻烦。最大的问题是，在不同平台上，换行符发生改变时，Git 会认为整个文件被修改，这就造成我们没法 diff，不能正确反映本次的修改。还好 Git 在设计时就考虑了这一点，其提供了一个 autocrlf 的配置项，用于在提交和检出时自动转换换行符，该配置有三个可选项：

- true: 提交时转换为 LF，检出时转换为 CRLF
- false: 提交检出均不转换
- input: 提交时转换为LF，检出时不转换


## core.safecrlf
如果把 autocrlf 设置为 false 时，那另一个配置项 safecrlf 最好设置为 ture。该选项用于检查文件是否包含混合换行符，其有三个可选项：

- true: 拒绝提交包含混合换行符的文件
- false: 允许提交包含混合换行符的文件
- warn: 提交包含混合换行符的文件时给出警告

全局设置
如果涉及到在多个系统平台上工作，推荐将 git 做如下配置：

```bash
git config --global core.autocrlf input
git config --global core.safecrlf true
```

本地设置
如果本地遇到fatal提交不了的情况，可以暂时把core.safecrlf设置为warn
```bash
fatal: CRLF would be replaced by LF ...
```
先提交了再说
```bash
git config core.autocrlf input
git config core.safecrlf warn
```
