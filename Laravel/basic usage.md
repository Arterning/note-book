
Create new app

```php
composer create-project laravel/laravel example-app


composer global require laravel/installer
laravel new example-app


cd example-app

php artisan serve


```


create file `database.sqlite` and then run

```php
php artisan migrate
```


```php
php artisan make:migration create_tasks_table --create=tasks  # 创建数据表迁移
php artisan make:migration alter_tasks_add_nickname --table=tasks  # 更新数据表迁移
```


我们对数据库的迁移操作都是基于 `Schema` 门面来完成（底层对应的类是 `Illuminate\Database\Schema\Builder`），比如创建数据表，需要调用该门面的 `create` 方法，该方法的第一个参数是要创建的数据表的名称，第二个参数是一个闭包，其中定义的是新增数据表的所有字段信息。

```php
Schema::create('users', function (Blueprint $table) {
    $table->increments('id');
    $table->string('name');
    $table->string('email')->unique();
    $table->timestamp('email_verified_at')->nullable();
    $table->string('password');
    $table->rememberToken();
    $table->timestamps();        
});
```

https://laravelacademy.org/post/9693


执行变更很简单，通过：

```php
php artisan migrate
```