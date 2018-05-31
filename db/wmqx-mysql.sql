-- -------------------------------------------
-- type: mysql
-- database: wmqx
-- author: phachon@163.com
-- -------------------------------------------

-- -------------------------------------------
-- wmqx user table
-- -------------------------------------------
DROP TABLE IF EXISTS `wmqx_user`;
CREATE TABLE `wmqx_user` (
  `user_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` char(18) NOT NULL DEFAULT '' COMMENT '手机',
  `role` tinyint(1) NOT NULL DEFAULT '0' COMMENT '角色 0 普通用户 1 管理员 2 超级管理员',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `wmqx_user` (`username`, `password`, `email`,  `mobile`, `role`, `is_delete`, `create_time`, `update_time`)
VALUES ('root', 'e10adc3949ba59abbe56e057f20f883e', 'root@123456.com', '1102222', '2', '0', '1500825600', '1500825600');

-- --------------------------------
-- wmqx user node table
-- --------------------------------
DROP TABLE IF EXISTS `wmqx_user_node`;
CREATE TABLE `wmqx_user_node` (
  `user_node_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户节点关系 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `node_id` int(10) NOT NULL DEFAULT '0' COMMENT '节点 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`user_node_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户节点关系表';

-- -------------------------------------------
-- wmqx node table
-- -------------------------------------------
DROP TABLE IF EXISTS `wmqx_node`;
CREATE TABLE `wmqx_node` (
  `node_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `manager_uri`  varchar(100) NOT NULL DEFAULT '' COMMENT 'wmqx 管理URI',
  `publish_uri`  varchar(100) NOT NULL DEFAULT '' COMMENT 'wmqx 发布URI',
  `token` VARCHAR(32) NOT NULL DEFAULT '' COMMENT 'token',
  `token_header_name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT 'token header name',
  `comment` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '备注',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`node_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------
-- wmqx log table
-- --------------------------------
DROP TABLE IF EXISTS `wmqx_log`;
CREATE TABLE `wmqx_log` (
  `log_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '日志id',
  `level` tinyint(3) NOT NULL DEFAULT '6' COMMENT '日志级别',
  `controller` char(100) NOT NULL DEFAULT '' COMMENT '控制器',
  `action` char(100) NOT NULL DEFAULT '' COMMENT '动作',
  `get` text NOT NULL COMMENT 'get参数',
  `post` text NOT NULL COMMENT 'post参数',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '信息',
  `ip` char(100) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `user_agent` char(200) NOT NULL DEFAULT '' COMMENT '用户代理',
  `referer` char(100) NOT NULL DEFAULT '' COMMENT 'referer',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '帐号id',
  `username` char(100) NOT NULL DEFAULT '' COMMENT '帐号名',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='行为日志表';