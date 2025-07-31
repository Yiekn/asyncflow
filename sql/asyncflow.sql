CREATE DATABASE IF NOT EXISTS asyncflow;
USE asyncflow;

CREATE TABLE `schedule_cfg`
(
    `task_type`           varchar(128)      NOT NULL COMMENT '任务类型',
    `task_stage`          tinyint unsigned  NOT NULL COMMENT '任务阶段',
    `schedule_limit`      int               NOT NULL COMMENT '一次拉取多少个任务',
    `schedule_interval`   int               NOT NULL COMMENT '拉取任务的间隔，单位为秒',
    `max_processing_time` int               NOT NULL COMMENT 'Worker处于执行中的最大时间，单位为秒',
    `create_time`         timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modify_time`         timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`task_type`, `task_stage`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `task_cfg`
(
    `task_type`           varchar(128)      NOT NULL COMMENT '任务类型',
    `task_stage`          tinyint unsigned  NOT NULL COMMENT '任务阶段',
    `priority`            int               NOT NULL COMMENT '优先级，单位为秒',
    `max_retry_num`       int               NOT NULL COMMENT '最大重试次数',
    `retry_interval`      int               NOT NULL COMMENT '重试间隔，单位为秒',
    `max_retry_interval`  int               NOT NULL COMMENT '最大重试间隔，单位为秒',
    `create_time`         timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modify_time`         timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`task_type`, `task_stage`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `schedule_pos`
(
    `task_type`          varchar(128)      NOT NULL COMMENT '任务类型',
    `schedule_begin_pos` int               NOT NULL COMMENT '调度开始于几号表',
    `schedule_end_pos`   int               NOT NULL COMMENT '调度结束于几号表',
    `create_time`        timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modify_time`        timestamp         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`task_type`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;