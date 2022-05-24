use touhou_gamejam;

create table account (
    id          int unsigned        auto_increment  primary key,
    tel_code    smallint unsigned   not null    default 0   comment '国际电话区号',
    phone       char(11)            not null    default ''  comment '电话号',
    email       char(32)            not null    default ''  comment '邮箱',
    uuid        binary(16)          not null                comment 'uuid',
    password    binary(20)          not null                comment '密码哈希值',
    status      tinyint unsigned    not null    default 0   comment '账号状态',
    ctime       datetime            not null                comment '创建时间',
    mtime       datetime            not null                comment '修改时间'
);

create index idx_account_tel_code_and_phone on account (tel_code, phone);

create unique index idx_account_email ON account (email);
