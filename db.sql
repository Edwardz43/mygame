CREATE DATABASE IF NOT EXISTS `MyGame`;
USE `MyGame`;

DROP TABLE IF EXISTS `Users`;

CREATE TABLE `Users` (
  `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(50) NOT NULL COMMENT '會員名稱',
  `Created_At` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `Updated_At` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `Users_ID_NAME_IDX` (`ID`,`Name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `GameResult`;

CREATE TABLE `GameResult` (
  `ID` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `GameID` int(10) unsigned NOT NULL,
  `Run` bigint(20) unsigned NOT NULL,
  `Inn` int(10) unsigned NOT NULL,
  `Detail` varchar(500) CHARACTER SET utf16 NOT NULL,
  `Created_At` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ModTimes` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=560 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `game_results`;

CREATE TABLE `game_results` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `game_id` tinyint(3) unsigned DEFAULT NULL,
  `run` bigint(20) unsigned DEFAULT NULL,
  `inn` int(10) unsigned DEFAULT NULL,
  `detail` varchar(255) DEFAULT NULL,
  `mod_times` tinyint(3) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_game_results_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `GameInfo`;

CREATE TABLE `GameInfo` (
  `GameID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(50) NOT NULL,
  PRIMARY KEY (`GameID`)
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