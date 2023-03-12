It's simple
not very care about whether i can run it directly
so i put all file into on directory.

这样下来，我们再配合模糊匹配查找日志，效果不就刚刚的了。
cat -n info.log |grep "a415ad50dbf84e99b1b56a31aacd209c"
或者
grep -10 'a415ad50dbf84e99b1b56a31aacd209c' info.log   （10是指上下10行）

copy from https://mp.weixin.qq.com/s/4pBulXc6pm3OtXySiaDIRA
