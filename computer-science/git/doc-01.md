# generate key
ssh-keygen -t rsa -C "comment"
-t 表示加密算法
-C 表示comment

Changing permission on your SSH private key
chmod 600 yourkeyname

# set the name and email
git config --global user.name "Arterning"
git config --global user.email "gmail.com"
