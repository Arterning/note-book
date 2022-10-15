## 1.大概思路

1. git pull 下来代码
2. 构建然后上传war包。

```bash

cd /home/taotao/code

#pull down
git pull origin dev


cd /home/taotao/code/java/java_taotao商城/taotao-manager
mvn clean install  -DskipTests
cp /home/taotao/code/java/java_taotao商城/taotao-manager/taotao-manager-web/target/taotao-manager-web.war /home/taotao/taotao/taotao_manage_tomat/webapps


cd /home/taotao/code/java/java_taotao商城/taotao-portal
mvn clean install  -DskipTests
cp /home/taotao/code/java/java_taotao商城/taotao-portal/target/ROOT.war /home/taotao/taotao/taotao_portal_tomcat/webapps


cd /home/taotao/code/java/java_taotao商城/taotao-rest
mvn clean install  -DskipTests
cp /home/taotao/code/java/java_taotao商城/taotao-rest/target/ROOT.war /home/taotao/taotao/taotao_rest_tomcat/webapps


cd /home/taotao/code/java/java_taotao商城/taotao-search
mvn clean install  -DskipTests
cp /home/taotao/code/java/java_taotao商城/taotao-search/target/ROOT.war /home/taotao/taotao/taotao_search_tomcat/webapps



cd /home/taotao/code/java/java_taotao商城/taotao-sso
mvn clean install  -DskipTests
cp /home/taotao/code/java/java_taotao商城/taotao-sso/target/ROOT.war /home/taotao/taotao/taotao_sso_tomcat/webapps



```

