tar -xzf protobuf-2.1.0.tar.gz
cd protobuf-2.1.0
./configure --prefix=$INSTALL_DIR
make
make check
make install