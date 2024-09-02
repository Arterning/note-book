Scroll down the httpd.conf file to the end and write the below three lines. In Apache, PHP is loaded as a module so load the module using  LoadModule, read PHP files with the file handler and finally add the PHP file located address within the double quotes

```
LoadModule php_module “C:\PHP8.1\php8apache2_4.dll”  

AddHandler application/x-httpd-php  .php

PHPIniDir “C:\PHP8.1”
```


Apache-Php
WSGI/ASGI-Python
Tomcat-Java
Nodejs-Javascript
