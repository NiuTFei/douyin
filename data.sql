DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `username`        varchar(128)        NOT NULL DEFAULT '' COMMENT '用户名',
    `password`      varchar(128)        NOT NULL DEFAULT '' COMMENT '密码',
    `name`      varchar(128)        NOT NULL DEFAULT '' COMMENT '昵称',
    `follow_count`       int(10)             NOT NULL DEFAULT 0 COMMENT '关注数',
    `follower_count`       int(10)             NOT NULL DEFAULT 0 COMMENT '粉丝数',
    `is_follow`       int(10)             NOT NULL DEFAULT 0 COMMENT '是否关注，0-未关注，1-已关注',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `author_id`        bigint(20)        NOT NULL DEFAULT 0 COMMENT '视频作者ID',
    `play_url`      varchar(128)        NOT NULL DEFAULT '' COMMENT '播放地址',
    `cover_url`      varchar(128)        NOT NULL DEFAULT '' COMMENT '封面地址',
    `favorite_count`       int(10)             NOT NULL DEFAULT 0 COMMENT '点赞数',
    `comment_count`       int(10)             NOT NULL DEFAULT 0 COMMENT '评论数',
    `is_favorite`       tinyint(1)             NOT NULL DEFAULT 0 COMMENT '是否点赞，0-未点赞，1-已点赞',
    `title`      varchar(128)        NOT NULL DEFAULT '' COMMENT '视频标题',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
    `favorite_video_id` bigint(20) UNSIGNED NOT NULL COMMENT '喜欢的视频id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '评论用户id',
    `video_id` bigint(20) UNSIGNED NOT NULL COMMENT '被评论视频id',
    `content` varchar(255) NOT NULL COMMENT '评论内容',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';

DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `from_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
    `to_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '被关注id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='关注表';


DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `from_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '发送消息用户id',
    `to_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '接收消息用户id',
    `content` varchar(255) NOT NULL COMMENT '消息内容',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '消息发送时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='消息表';