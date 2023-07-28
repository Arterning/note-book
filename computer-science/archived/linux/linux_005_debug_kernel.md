scripts/config --disable SYSTEM_REVOCATION_KEYS

scripts/config --disable SYSTEM_TRUSTED_KEYS

make bzImage
make
make install
they are different make target


how to install qemu in ubuntu
sudo apt install qemu-kvm qemu virt-manager virt-viewer libvirt-clients libvirt-daemon-system bridge-utils virtinst libvirt-daemon
refer:
https://www.atechtown.com/install-qemu-on-ubuntu/

