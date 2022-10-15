#!/bin/bash
# sudo add-apt-repository ppa:gnome3-team/gnome3 -y
# sudo add-apt-repository ppa:vala-team -y
# sudo add-apt-repository ppa:gnome3-team/gnome3-staging -y
# sudo aptitude install cmake valac libgtk-3-dev libgee-0.8-dev libclutter-gtk-1.0-dev libclutter-1.0-dev libwebkit2gtk-4.0-dev libclutter-gst-3.0-dev
# sudo apt install cmake valac libgtk-3-dev libgee-0.8-dev libclutter-gtk-1.0-dev libclutter-1.0-dev libwebkit2gtk-4.0-dev libclutter-gst-3.0-dev

# git clone https://github.com/cheesecakeufo/komorebi.git

cd komorebi && rm -rf build
mkdir build && cd build
cmake .. && sudo make install && ./komorebi