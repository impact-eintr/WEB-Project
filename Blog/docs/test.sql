# 创建标签表
create table `blog_tag` (
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `name` varchar(100) default '' comment '标签名称',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100)  default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除 0是没有 1是删了',
    `state` tinyint(3) unsigned default '1' comment '0 禁用 1 可用',
    PRIMARY KEY (`id`)
) engine=InnoDB default charset=utf8mb4 comment='标签管理';

# 创建文章表
create table `blog_article` (
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `title` varchar(100) default '' comment '文章标题',
    `desc` varchar(100) default '' comment '文章简述',
    `cover_image_url` varchar(255) default '' comment '封面图片地址',
    `content` longtext comment '文章内容',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100)  default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除 0是没有 1是删了',
    `state` tinyint(3) unsigned default '1' comment '0 禁用 1 可用',
    PRIMARY KEY (`id`)
) engine=InnoDB default charset=utf8mb4 comment='文章管理';

# 创建文章标签关联表
create table `blog_article_tag` (
    `id` int(10) unsigned not null AUTO_INCREMENT,
    `blog_article_id` int(11) not null comment '文章ID',
    `tag_id` int(10) unsigned not null default '0' comment '标签ID',
    `created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100)  default '' comment '修改人',
    `deleted_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除 0是没有 1是删了',
    `state` tinyint(3) unsigned default '1' comment '0 禁用 1 可用',
    PRIMARY KEY (`id`)
) engine=InnoDB default charset=utf8mb4 comment='文章标签管理';
