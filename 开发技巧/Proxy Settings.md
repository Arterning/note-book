
## Linux
note dont' use `cat > ~/.bash_pofile` , that is to override

```bash

cat >> ~/.bash_profile << EOF
function proxy_on() {
    export http_proxy=http://127.0.0.1:7890
    export https_proxy=\$http_proxy
    echo -e "终端代理已开启。"
}

function proxy_off(){
    unset http_proxy https_proxy
    echo -e "终端代理已关闭。"
}
EOF

source ~/.bash_profile

proxy_on
proxy_off
```


## Windows Terminal

```cmd
set http_proxy=http://127.0.0.1:7890
set https_proxy=http://127.0.0.1:7890

set http_proxy=  
set https_proxy=
```


## Git

```bash
# 设置
git config --global http.proxy 'socks5://127.0.0.1:7890' 
git config --global https.proxy 'socks5://127.0.0.1:7890'

# 恢复
git config --global --unset http.proxy
git config --global --unset https.proxy


git config --list

# test
ssh -T git@github.com
```



## NPM
```bash
# 设置

npm config set proxy http://127.0.0.1:7890
npm config set https-proxy http://127.0.0.1:7890

# 恢复
npm config delete proxy
npm config delete https-proxy


```




https://weilining.github.io/294.html

