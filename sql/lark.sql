USE asyncflow;

INSERT INTO task_cfg (task_type, task_stage, priority, max_retry_num, retry_interval, max_retry_interval)
VALUES ('lark', 1, 0, 3, 1, 10);
INSERT INTO task_cfg (task_type, task_stage, priority, max_retry_num, retry_interval, max_retry_interval)
VALUES ('lark', 2, 0, 3, 1, 10);

INSERT INTO schedule_cfg (task_type, task_stage, schedule_limit, schedule_interval, max_processing_time)
VALUES ('lark', 1, 1, 3, 100);
INSERT INTO schedule_cfg (task_type, task_stage, schedule_limit, schedule_interval, max_processing_time)
VALUES ('lark', 2, 1, 3, 100);

INSERT INTO schedule_pos (task_type, schedule_begin_pos, schedule_end_pos)
VALUES ('lark', 1, 1);

CREATE TABLE `task_lark_1`
(
    `id`                 bigint            NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`            varchar(256)      NOT NULL COMMENT '用户id，标识用户',
    `task_id`            varchar(256)      NOT NULL COMMENT '任务id，标识任务',
    `task_type`          varchar(128)      NOT NULL COMMENT '任务类型',
    `task_stage`         tinyint unsigned  NOT NULL COMMENT '任务阶段',
    `status`             tinyint unsigned  NOT NULL COMMENT '状态',
    `priority`           int               NOT NULL COMMENT '优先级，单位为秒',
    `crt_retry_num`      int               NOT NULL COMMENT '已经重试几次了',
    `max_retry_num`      int               NOT NULL COMMENT '最大能重试几次',
    `retry_interval`     int               NOT NULL COMMENT '重试间隔，单位为秒',
    `max_retry_interval` int               NOT NULL COMMENT '最大重试间隔，单位为秒',
    `schedule_log`       varchar(4096)     NOT NULL COMMENT '调度信息记录',
    `task_context`       varchar(8192)     NOT NULL COMMENT '任务上下文',
    `order_time`         bigint            NOT NULL COMMENT '调度时间，越小调度越优先，单位为毫秒',
    `create_time`        timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modify_time`        timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_task_id` (`task_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_status` (`status`),
    KEY `idx_status_order_time` (`status`, `order_time`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;