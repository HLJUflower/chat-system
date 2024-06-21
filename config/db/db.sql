USE chat;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE IF NOT EXISTS `user` (
    `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` varchar(20) DEFAULT 'user' COMMENT '用户名称',
    `user_real_id` bigint(20) DEFAULT NULL COMMENT '用户身份证号',
    `user_real_name` varchar(20) DEFAULT 'user_name' COMMENT '用户姓名',
    `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
    `status` char(1) DEFAULT '0' COMMENT '状态（1正常 0离线）',
    `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（1代表存在 0代表删除）',
    `create_time` datetime DEFAULT NULL COMMENT '注册时间',
    `login_time` datetime DEFAULT NULL COMMENT '登录时间',
    ·login_ip· varchar(15) DEFAULT NULL COMMENT '登录IP',
    `remark` varchar(500) DEFAULT NULL COMMENT '备注',
    PRIMARY KEY (`user_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8 COMMENT='用户信息表';
