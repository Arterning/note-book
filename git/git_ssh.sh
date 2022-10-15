#bin/sh
git remote add origin git@gitee.com:handshake_one/note.git
ssh-keygen -t rsa -C "ning"
cat /home/ning/.ssh/id_rsa.pub
ssh -T git@gitee.com
git remote -v


