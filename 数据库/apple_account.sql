/*
 Navicat Premium Data Transfer

 Source Server         : hym
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3306
 Source Schema         : web_apple_db

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 01/11/2023 11:34:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for apple_account
-- ----------------------------
DROP TABLE IF EXISTS `apple_account`;
CREATE TABLE `apple_account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号',
  `pwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `country_id` int DEFAULT NULL COMMENT '国家id',
  `balance` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '余额',
  `status` tinyint DEFAULT '0' COMMENT '禁用：0否 1是',
  `info` text COLLATE utf8mb4_general_ci COMMENT '信息',
  `cookies` text COLLATE utf8mb4_general_ci COMMENT 'cookies',
  `login_status` tinyint NOT NULL DEFAULT '0' COMMENT '登录：0否 1是',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of apple_account
-- ----------------------------
BEGIN;
INSERT INTO `apple_account` (`id`, `account`, `pwd`, `country_id`, `balance`, `status`, `info`, `cookies`, `login_status`, `isStop`, `updatedAt`, `createdAt`) VALUES (1, 'puxwuvyr@hotmail.com', 'Qw112212', NULL, NULL, 0, NULL, NULL, 0, 0, '2023-11-01 11:25:19', '2023-11-01 11:25:19');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
