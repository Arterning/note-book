docker run -d  -p 443:443 -p 8080:80 -p 222:22 --name gitlab --restart always -v /home/ninghuang/gitlab/data:/var/opt/gitlab gitlab/gitlab-ce
