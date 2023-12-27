## set registry

```bash
npm config set registry http://registry.npm.taobao.org
npm get registry 

##注意如果你换成淘宝镜像的话，会影响你发布模块，这时候需要换回npm官网的镜像
npm config set registry https://registry.npmjs.org
npm config delete proxy


yarn config set registry https://registry.npm.taobao.org -g
yarn config get registry
```


## set proxy 
建议最好不要这么干，设置国内镜像就足够了
```bash
yarn config set proxy  http://127.0.0.1:33210
yarn config set https-proxy http://127.0.0.1:33210
```

