CREATE DATABASE IF NOT EXISTS `MyGame`;
USE `MyGame`;

DROP TABLE IF EXISTS `Users`;

CREATE TABLE `Users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '會員名稱',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Users_ID_NAME_IDX` (`id`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `Lobby`;

CREATE TABLE `Lobby` (
  `game_id` smallint(5) unsigned NOT NULL COMMENT '遊戲類型',
  `run` bigint(20) unsigned NOT NULL,
  `inn` int(10) unsigned NOT NULL,
  `status` smallint(5) unsigned NOT NULL COMMENT '遊戲狀態 1.新局 2.開牌 3.結算 4.中場休息 5.維護'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO MyGame.Lobby (game_id,run,inn, `status`) VALUES 
(1, 20190914, 1, 1);

DROP TABLE IF EXISTS `GameResult`;

CREATE TABLE `GameResult` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `game_id` tinyint(3) unsigned DEFAULT NULL,
  `run` bigint(20) unsigned NOT NULL,
  `inn` int(10) unsigned NOT NULL,
  `detail` varchar(500) CHARACTER SET utf16 NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `mod_times` tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=560 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `GameInfo`;

CREATE TABLE `GameInfo` (
  `game_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`game_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

LOCK TABLES `Users` WRITE;

INSERT INTO MyGame.Users (Name,Created_At,Updated_At) VALUES 
('test0001','2019-08-20 01:36:43.000','2019-08-20 01:36:43.000')
,('test0002','2019-08-20 01:36:43.000','2019-08-20 01:39:40.000')
,('test0003','2019-08-20 04:10:22.000','2019-08-20 04:10:22.000')
,('test0004','2019-08-20 04:18:05.000','2019-08-20 04:18:05.000')
;

UNLOCK TABLES;

GRANT ALL ON *.* TO 'root'@'%' IDENTIFIED BY 'root' WITH GRANT OPTION;

FLUSH PRIVILEGES;