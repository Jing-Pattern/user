
DROP TABLE IF EXISTS  love_user_info;
CREATE TABLE love_user_info (
                                   `id` int(20) NOT NULL AUTO_INCREMENT  COMMENT 'id',
                                   `open_id` varchar(255) NOT NULL COMMENT 'openId',
                                   `session_key` VARCHAR(255) NOT NULL COMMENT '会话密钥',
                                   `union_id` VARCHAR(50) COMMENT '微信开放平台唯一标识（可选，如果需要跨应用统一用户）',
                                   `nickname` varchar(50) NOT NULL COMMENT '昵称',
                                   `name` varchar(255) NOT NULL COMMENT '名称',
                                   `gender` INT COMMENT '用户性别（0表示未知，1表示男性，2表示女性）',
                                   `tel` int(12) COMMENT '绑定手机号',
                                   `avatar_url` VARCHAR(255)  COMMENT '用户头像url',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   INDEX idx_tel (`tel`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='用户信息表';