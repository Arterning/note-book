/**

个人觉得效率比较高的建模方式还是用sql
因为可以大量复制粘贴
data grip 提供的 ui 建表虽然也好用
但是还是慢了一点 不能复制粘贴啊

不过对于创建主键和外键这种语法 反正不同的数据库你也记不住 干脆用datagrip生成好的多
生成之后复制为sql 然后接着复制粘贴 这样效率是比较高的
 */

create table `sa-auth`.platform
(
    id              varchar(50)  null comment 'ID',
    platform_code   varchar(50)  not null comment '平台编码',
    platform_secret varchar(100) not null comment '平台密钥',
    platform_name   varchar(50)  not null comment '平台名称',
    remark          varchar(400) null comment '备注',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);

create table `sa-auth`.role
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    role_name   varchar(50)  not null comment '平台名称',
    remark          varchar(400) null comment '备注',
    is_admin char not null comment '是否管理员',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);


create table `sa-auth`.menu
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    menu_name   varchar(100)  not null comment '菜单名称',
    menu_code   varchar(50)  not null comment '菜单编码',
    menu_icon   varchar(50)  not null comment '菜单图标',
    menu_url   varchar(200)  not null comment '菜单url',
    menu_parent_id   varchar(50)  not null comment '父级菜单ID',
    menu_level          int(10) null comment '菜单级数',
    remark          varchar(400) null comment '备注',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);



create table `sa-auth`.role_menu
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    role_id   varchar(50)  not null comment '角色ID',
    menu_id   varchar(50)  not null comment '菜单ID',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);



create table `sa-auth`.role_function
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    role_id   varchar(50)  not null comment '角色ID',
    function_id   varchar(50)  not null comment '功能ID',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);


create table `sa-auth`.group
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    group_name   varchar(50)  not null comment '组名称',
    remark          varchar(400) null comment '备注',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);


create table `sa-auth`.user_group
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    group_id   varchar(50)  not null comment '组ID',
    user_id          varchar(50) null comment '用户ID',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);


create table `sa-auth`.user_role
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    role_id   varchar(50)  not null comment '角色ID',
    user_id          varchar(50) null comment '用户ID',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);


create table `sa-auth`.group_role
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',
    group_id   varchar(50)  not null comment '组ID',
    role_id          varchar(50) null comment '角色ID',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);


create table `sa-auth`.menu_function
(
    id              varchar(50)  null comment 'ID',
    platform_id   varchar(50)  not null comment '平台ID',

    function_name   varchar(100)  not null comment '功能名称',
    function_code   varchar(50)  not null comment '功能编码',
    function_icon   varchar(50)  not null comment '功能图标',
    function_url   varchar(200)  not null comment '功能url',
    menu_id   varchar(50)  not null comment '菜单ID',

    is_active       char         not null comment '是否有效',
    create_date     datetime(6) not null comment '创建日期',
    modify_date     datetime(6) null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);
