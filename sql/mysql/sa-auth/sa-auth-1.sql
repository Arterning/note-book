create table `sa-auth`.group_
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50)  not null comment '平台ID',
    group_name  varchar(50)  not null comment '组名称',
    remark      varchar(400) null comment '备注',
    is_active   char         not null comment '是否有效',
    create_date datetime(6)  not null comment '创建日期',
    modify_date datetime(6)  null comment '更新日期',
    creator     varchar(50)  not null comment '创建人  ',
    modifier    varchar(50)  null comment '更新人'
);

create index group_platform_null_fk
    on `sa-auth`.group_ (platform_id);

create table `sa-auth`.group_role
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50) not null comment '平台ID',
    group_id    varchar(50) not null comment '组ID',
    role_id     varchar(50) null comment '角色ID',
    is_active   char        not null comment '是否有效',
    create_date datetime(6) not null comment '创建日期',
    modify_date datetime(6) null comment '更新日期',
    creator     varchar(50) not null comment '创建人  ',
    modifier    varchar(50) null comment '更新人'
);

create index group_role_group_null_fk
    on `sa-auth`.group_role (group_id);

create index group_role_platform_null_fk
    on `sa-auth`.group_role (platform_id);

create index group_role_role_null_fk
    on `sa-auth`.group_role (role_id);


create table `sa-auth`.menu
(
    id             bigint auto_increment comment 'ID'
        primary key,
    platform_id    varchar(50)  not null comment '平台ID',
    menu_name      varchar(100) not null comment '菜单名称',
    menu_code      varchar(50)  not null comment '菜单编码',
    menu_icon      varchar(50)  not null comment '菜单图标',
    menu_url       varchar(200) not null comment '菜单url',
    menu_parent_id varchar(50)  not null comment '父级菜单ID',
    menu_level     int(10)      null comment '菜单级数',
    remark         varchar(400) null comment '备注',
    is_active      char         not null comment '是否有效',
    create_date    datetime(6)  not null comment '创建日期',
    modify_date    datetime(6)  null comment '更新日期',
    creator        varchar(50)  not null comment '创建人  ',
    modifier       varchar(50)  null comment '更新人'
);

create index menu_platform_null_fk
    on `sa-auth`.menu (platform_id);

create table `sa-auth`.menu_function
(
    id            bigint auto_increment comment 'ID'
        primary key,
    platform_id   varchar(50)  not null comment '平台ID',
    function_name varchar(100) not null comment '功能名称',
    function_code varchar(50)  not null comment '功能编码',
    function_icon varchar(50)  not null comment '功能图标',
    function_url  varchar(200) not null comment '功能url',
    menu_id       varchar(50)  not null comment '菜单ID',
    is_active     char         not null comment '是否有效',
    create_date   datetime(6)  not null comment '创建日期',
    modify_date   datetime(6)  null comment '更新日期',
    creator       varchar(50)  not null comment '创建人  ',
    modifier      varchar(50)  null comment '更新人'
);

create index menu_function_menu_null_fk
    on `sa-auth`.menu_function (menu_id);

create index menu_function_platform_null_fk
    on `sa-auth`.menu_function (platform_id);

create table `sa-auth`.platform
(
    id              bigint auto_increment comment 'ID'
        primary key,
    platform_code   varchar(50)  not null comment '平台编码',
    platform_secret varchar(100) not null comment '平台密钥',
    platform_name   varchar(50)  not null comment '平台名称',
    remark          varchar(400) null comment '备注',
    is_active       char         not null comment '是否有效',
    create_date     datetime(6)  not null comment '创建日期',
    modify_date     datetime(6)  null comment '更新日期',
    creator         varchar(50)  not null comment '创建人  ',
    modifier        varchar(50)  null comment '更新人'
);

create index platform_name_index_name
    on `sa-auth`.platform (id, platform_name);

create table `sa-auth`.role
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50)  not null comment '平台ID',
    role_name   varchar(50)  not null comment '平台名称',
    remark      varchar(400) null comment '备注',
    is_admin    char         not null comment '是否管理员',
    is_active   char         not null comment '是否有效',
    create_date datetime(6)  not null comment '创建日期',
    modify_date datetime(6)  null comment '更新日期',
    creator     varchar(50)  not null comment '创建人  ',
    modifier    varchar(50)  null comment '更新人'
);

create index role_foreign_id
    on `sa-auth`.role (platform_id);

create table `sa-auth`.role_function
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50) not null comment '平台ID',
    role_id     varchar(50) not null comment '角色ID',
    function_id varchar(50) not null comment '功能ID',
    is_active   char        not null comment '是否有效',
    create_date datetime(6) not null comment '创建日期',
    modify_date datetime(6) null comment '更新日期',
    creator     varchar(50) not null comment '创建人  ',
    modifier    varchar(50) null comment '更新人'
);

create index role_function_menu_function_null_fk
    on `sa-auth`.role_function (function_id);

create index role_function_platform_null_fk
    on `sa-auth`.role_function (platform_id);

create index role_function_role_null_fk
    on `sa-auth`.role_function (role_id);

create table `sa-auth`.role_menu
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50) not null comment '平台ID',
    role_id     varchar(50) not null comment '角色ID',
    menu_id     varchar(50) not null comment '菜单ID',
    is_active   char        not null comment '是否有效',
    create_date datetime(6) not null comment '创建日期',
    modify_date datetime(6) null comment '更新日期',
    creator     varchar(50) not null comment '创建人  ',
    modifier    varchar(50) null comment '更新人'
);

create index role_menu_menu_null_fk
    on `sa-auth`.role_menu (menu_id);

create index role_menu_platform_null_fk
    on `sa-auth`.role_menu (platform_id);

create index role_menu_role_null_fk
    on `sa-auth`.role_menu (role_id);

create table `sa-auth`.user_group
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50) not null comment '平台ID',
    group_id    varchar(50) not null comment '组ID',
    user_id     varchar(50) null comment '用户ID',
    is_active   char        not null comment '是否有效',
    create_date datetime(6) not null comment '创建日期',
    modify_date datetime(6) null comment '更新日期',
    creator     varchar(50) not null comment '创建人  ',
    modifier    varchar(50) null comment '更新人'
);

create index user_group_group_null_fk
    on `sa-auth`.user_group (group_id);

create index user_group_platform_null_fk
    on `sa-auth`.user_group (platform_id);

create table `sa-auth`.user_role
(
    id          bigint auto_increment comment 'ID'
        primary key,
    platform_id varchar(50) not null comment '平台ID',
    role_id     varchar(50) not null comment '角色ID',
    user_id     varchar(50) null comment '用户ID',
    is_active   char        not null comment '是否有效',
    create_date datetime(6) not null comment '创建日期',
    modify_date datetime(6) null comment '更新日期',
    creator     varchar(50) not null comment '创建人  ',
    modifier    varchar(50) null comment '更新人'
);

create index user_role_platform_null_fk
    on `sa-auth`.user_role (platform_id);

create index user_role_role_null_fk
    on `sa-auth`.user_role (role_id);

