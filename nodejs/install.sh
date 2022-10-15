https://npm.taobao.org/mirrors/node/

wget https://npm.taobao.org/mirrors/node/v12.14.1/node-v12.14.1-linux-x64.tar.xz
sudo mv node /usr/local/node
sudo ln -s /usr/local/node/bin/node /usr/bin/node
sudo ln -s /usr/local/node/bin/npm /usr/bin/npm
sudo ln -s /usr/local/node/bin/cnpm /usr/bin/cnpm

# usr/local/node/lib/node_modules/jquery/node_modules

npm config set registry https://registry.npm.taobao.org

npm install datatables.net    # Core library
npm install datatables.net-dt # Styling

