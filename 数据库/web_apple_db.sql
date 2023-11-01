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

 Date: 01/11/2023 11:53:35
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

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action` (
  `actionId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '操作ID',
  `actionName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `actionCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`) USING BTREE,
  UNIQUE KEY `actionCode` (`actionCode`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限操作表';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
BEGIN;
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (1, '权限场景-查看', 'authSceneLook', '', 0, '2023-07-01 15:38:45', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (2, '权限场景-新增', 'authSceneCreate', '', 0, '2023-07-01 15:38:46', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (3, '权限场景-编辑', 'authSceneUpdate', '', 0, '2023-07-01 15:38:49', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (4, '权限场景-删除', 'authSceneDelete', '', 0, '2023-07-01 15:38:50', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (5, '权限操作-查看', 'authActionLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (6, '权限操作-新增', 'authActionCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (7, '权限操作-编辑', 'authActionUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (8, '权限操作-删除', 'authActionDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (9, '权限菜单-查看', 'authMenuLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (10, '权限菜单-新增', 'authMenuCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (11, '权限菜单-编辑', 'authMenuUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (12, '权限菜单-删除', 'authMenuDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (13, '权限角色-查看', 'authRoleLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (14, '权限角色-新增', 'authRoleCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (15, '权限角色-编辑', 'authRoleUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (16, '权限角色-删除', 'authRoleDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (17, '平台管理员-查看', 'platformAdminLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (18, '平台管理员-新增', 'platformAdminCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (19, '平台管理员-编辑', 'platformAdminUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (20, '平台管理员-删除', 'platformAdminDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (21, '平台配置-查看', 'platformConfigLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (22, '平台配置-保存', 'platformConfigSave', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (23, '用户-查看', 'userLook', '', 0, '2023-10-14 16:03:55', '2023-10-14 15:32:37');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (24, '用户-编辑', 'userUpdate', '', 0, '2023-10-14 16:03:59', '2023-10-14 15:32:37');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (25, '苹果-查看', 'appleAccountLook', '', 0, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (26, '苹果-新增', 'appleAccountCreate', '', 0, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (27, '苹果-编辑', 'appleAccountUpdate', '', 0, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
INSERT INTO `auth_action` (`actionId`, `actionName`, `actionCode`, `remark`, `isStop`, `updatedAt`, `createdAt`) VALUES (28, '苹果-删除', 'appleAccountDelete', '', 0, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
COMMIT;

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene` (
  `actionId` int unsigned NOT NULL DEFAULT '0' COMMENT '操作ID',
  `sceneId` int unsigned NOT NULL DEFAULT '0' COMMENT '场景ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`,`sceneId`) USING BTREE,
  KEY `actionId` (`actionId`) USING BTREE,
  KEY `sceneId` (`sceneId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限操作，权限场景关联表（操作可用在哪些场景）';

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
BEGIN;
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (1, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (2, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (3, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (4, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (5, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (6, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (7, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (8, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (9, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (10, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (11, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (12, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (13, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (14, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (15, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (16, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (17, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (18, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (19, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (20, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (21, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (22, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (23, 1, '2023-10-14 15:32:37', '2023-10-14 15:32:37');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (24, 1, '2023-10-14 15:32:37', '2023-10-14 15:32:37');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (25, 1, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (26, 1, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (27, 1, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
INSERT INTO `auth_action_rel_to_scene` (`actionId`, `sceneId`, `updatedAt`, `createdAt`) VALUES (28, 1, '2023-11-01 11:21:15', '2023-11-01 10:48:40');
COMMIT;

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu` (
  `menuId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menuName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneId` int unsigned NOT NULL DEFAULT '0' COMMENT '场景ID',
  `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `level` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '层级',
  `idPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `menuIcon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `menuUrl` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `extraData` json DEFAULT NULL COMMENT '额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}',
  `sort` tinyint unsigned NOT NULL DEFAULT '50' COMMENT '排序值。从小到大排序，默认50，范围0-100',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`menuId`) USING BTREE,
  KEY `sceneId` (`sceneId`) USING BTREE,
  KEY `pid` (`pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
BEGIN;
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (1, '主页', 1, 0, 1, '0-1', 'AutoiconEpHomeFilled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 0, 0, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (2, '权限管理', 1, 0, 1, '0-2', 'AutoiconEpLock', '', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage\", \"zh-cn\": \"权限管理\"}}}', 90, 0, '2023-09-20 23:02:07', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (3, '场景', 1, 2, 2, '0-2-3', 'AutoiconEpFlag', '/auth/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"场景\"}}}', 100, 0, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (4, '操作', 1, 2, 2, '0-2-4', 'AutoiconEpCoordinate', '/auth/action', '{\"i18n\": {\"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}}', 90, 0, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (5, '菜单', 1, 2, 2, '0-2-5', 'AutoiconEpMenu', '/auth/menu', '{\"i18n\": {\"title\": {\"en\": \"Menu\", \"zh-cn\": \"菜单\"}}}', 80, 0, '2023-09-20 23:02:10', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (6, '角色', 1, 2, 2, '0-2-6', 'AutoiconEpView', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 70, 0, '2023-09-20 23:02:11', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (7, '平台管理员', 1, 2, 2, '0-2-7', 'Vant-manager-o', '/platform/admin', '{\"i18n\": {\"title\": {\"en\": \"Platform Admin\", \"zh-cn\": \"平台管理员\"}}}', 60, 0, '2023-10-14 16:18:16', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (8, '系统管理', 1, 0, 1, '0-8', 'AutoiconEpPlatform', '', '{\"i18n\": {\"title\": {\"en\": \"System Manage\", \"zh-cn\": \"系统管理\"}}}', 85, 0, '2023-09-20 23:02:14', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (9, '配置中心', 1, 8, 2, '0-8-9', 'AutoiconEpSetting', '', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 100, 0, '2023-06-24 17:30:46', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (10, '平台配置', 1, 9, 3, '0-8-9-10', '', '/platform/config', '{\"i18n\": {\"title\": {\"en\": \"Platform Config\", \"zh-cn\": \"平台配置\"}}}', 50, 0, '2023-09-20 23:02:18', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (11, '用户管理', 1, 0, 1, '0-11', 'Vant-friends', '', '{\"i18n\": {\"title\": {\"en\": \"User Manage\", \"zh-cn\": \"用户管理\"}}}', 50, 0, '2023-10-14 16:14:56', '2023-10-14 15:32:37');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (12, '用户', 1, 11, 2, '0-11-12', 'Vant-user-o', '/user/user', '{\"i18n\": {\"title\": {\"en\": \"User\", \"zh-cn\": \"用户\"}}}', 50, 0, '2023-10-14 16:17:28', '2023-10-14 15:32:37');
INSERT INTO `auth_menu` (`menuId`, `menuName`, `sceneId`, `pid`, `level`, `idPath`, `menuIcon`, `menuUrl`, `extraData`, `sort`, `isStop`, `updatedAt`, `createdAt`) VALUES (13, '苹果', 1, 0, 1, '0-13', 'AutoiconEpLink', '/apple/account', '{\"i18n\": {\"title\": {\"en\": \"Account\", \"zh-cn\": \"苹果\"}}}', 50, 0, '2023-11-01 11:21:15', '2023-11-01 11:09:15');
COMMIT;

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role` (
  `roleId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `roleName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneId` int unsigned NOT NULL DEFAULT '0' COMMENT '场景ID',
  `tableId` int unsigned NOT NULL DEFAULT '0' COMMENT '关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`) USING BTREE,
  KEY `sceneId` (`sceneId`) USING BTREE,
  KEY `tableId` (`tableId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限角色表';

-- ----------------------------
-- Records of auth_role
-- ----------------------------
BEGIN;
INSERT INTO `auth_role` (`roleId`, `roleName`, `sceneId`, `tableId`, `isStop`, `updatedAt`, `createdAt`) VALUES (1, '超级管理员', 1, 0, 0, '2023-10-14 16:22:51', '2023-06-09 12:03:30');
COMMIT;

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_platform_admin`;
CREATE TABLE `auth_role_rel_of_platform_admin` (
  `roleId` int unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `adminId` int unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`,`adminId`) USING BTREE,
  KEY `roleId` (`roleId`) USING BTREE,
  KEY `adminId` (`adminId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限角色，系统管理员关联表（系统管理员包含哪些角色）';

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_action`;
CREATE TABLE `auth_role_rel_to_action` (
  `roleId` int unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `actionId` int unsigned NOT NULL DEFAULT '0' COMMENT '操作ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`,`actionId`) USING BTREE,
  KEY `roleId` (`roleId`) USING BTREE,
  KEY `actionId` (`actionId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限角色，权限操作关联表（角色包含哪些操作）';

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------
BEGIN;
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 1, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 2, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 3, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 4, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 5, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 6, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 7, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 8, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 9, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 10, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 11, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 12, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 13, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 14, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 15, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 16, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 17, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 18, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 19, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 20, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 21, '2023-06-11 14:52:39', '2023-06-11 14:52:39');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 22, '2023-06-11 14:52:39', '2023-06-11 14:52:39');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 23, '2023-10-14 16:22:51', '2023-10-14 16:22:51');
INSERT INTO `auth_role_rel_to_action` (`roleId`, `actionId`, `updatedAt`, `createdAt`) VALUES (1, 24, '2023-10-14 16:22:51', '2023-10-14 16:22:51');
COMMIT;

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_menu`;
CREATE TABLE `auth_role_rel_to_menu` (
  `roleId` int unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `menuId` int unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`,`menuId`) USING BTREE,
  KEY `roleId` (`roleId`) USING BTREE,
  KEY `menuId` (`menuId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限角色，权限菜单关联表（角色包含哪些菜单）';

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------
BEGIN;
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 2, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 3, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 4, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 5, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 6, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 7, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 8, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 9, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 10, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 11, '2023-10-14 16:22:51', '2023-10-14 16:22:51');
INSERT INTO `auth_role_rel_to_menu` (`roleId`, `menuId`, `updatedAt`, `createdAt`) VALUES (1, 12, '2023-10-14 16:22:51', '2023-10-14 16:22:51');
COMMIT;

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_scene`;
CREATE TABLE `auth_scene` (
  `sceneId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '场景ID',
  `sceneName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `sceneConfig` json DEFAULT NULL COMMENT '配置。JSON格式：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sceneId`) USING BTREE,
  UNIQUE KEY `sceneCode` (`sceneCode`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='权限场景表';

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
BEGIN;
INSERT INTO `auth_scene` (`sceneId`, `sceneName`, `sceneCode`, `sceneConfig`, `isStop`, `updatedAt`, `createdAt`) VALUES (1, '平台后台', 'platform', '{\"signKey\": \"www.admin.com_platform\", \"signType\": \"HS256\", \"expireTime\": 14400}', 0, '2023-06-14 21:52:16', '2023-06-09 12:03:30');
INSERT INTO `auth_scene` (`sceneId`, `sceneName`, `sceneCode`, `sceneConfig`, `isStop`, `updatedAt`, `createdAt`) VALUES (2, 'APP', 'app', '{\"signKey\": \"www.admin.com_app\", \"signType\": \"HS256\", \"expireTime\": 604800}', 0, '2023-10-21 17:21:15', '2023-10-21 17:21:15');
COMMIT;

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `platform_admin`;
CREATE TABLE `platform_admin` (
  `adminId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adminId`) USING BTREE,
  UNIQUE KEY `account` (`account`) USING BTREE,
  UNIQUE KEY `phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台管理员表';

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
BEGIN;
INSERT INTO `platform_admin` (`adminId`, `phone`, `account`, `password`, `salt`, `nickname`, `avatar`, `isStop`, `updatedAt`, `createdAt`) VALUES (1, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB', '超级管理员', 'http://www.admin.com/common/20230920/1695222477127_79698554.png', 0, '2023-09-20 23:09:17', '2023-06-09 12:03:30');
COMMIT;

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS `platform_config`;
CREATE TABLE `platform_config` (
  `configId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `configKey` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置Key',
  `configValue` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置值',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE KEY `configKey` (`configKey`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台配置表';

-- ----------------------------
-- Records of platform_config
-- ----------------------------
BEGIN;
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (1, 'uploadType', 'local', '2023-10-21 17:10:45', '2023-09-20 22:59:42');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (2, 'localUploadUrl', 'http://192.168.0.200:20080/upload/upload', '2023-10-21 17:08:18', '2023-09-20 22:59:47');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (3, 'localUploadSignKey', '123456', '2023-10-21 17:08:18', '2023-09-20 22:59:56');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (4, 'localUploadFileSaveDir', '../public/', '2023-10-21 17:08:18', '2023-09-20 23:00:05');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (5, 'localUploadFileUrlPrefix', 'http://www.admin.com', '2023-10-21 17:08:18', '2023-09-20 23:00:14');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (6, 'aliyunOssHost', 'https://oss-cn-hangzhou.aliyuncs.com', '2023-10-21 17:08:18', '2023-10-21 16:55:43');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (7, 'aliyunOssBucket', 'bucket', '2023-10-21 17:08:18', '2023-10-21 16:55:43');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (8, 'aliyunOssAccessKeyId', 'accessKeyId', '2023-10-21 17:08:18', '2023-10-21 16:55:43');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (9, 'aliyunOssAccessKeySecret', 'accessKeySecret', '2023-10-21 17:08:18', '2023-10-21 16:55:43');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (10, 'aliyunOssRoleArn', 'acs:ram::xxxxxxxxxxxxxxxx:role/aliyunosstokengeneratorrole', '2023-10-21 17:08:18', '2023-10-21 16:55:43');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (11, 'aliyunOssCallbackUrl', 'https://www.xxxx.com/upload/notify', '2023-10-21 17:08:18', '2023-10-21 16:55:43');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (12, 'smsType', 'aliyunSms', '2023-10-21 17:08:48', '2023-10-21 16:55:44');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (13, 'aliyunSmsAccessKeyId', 'accessKeyId', '2023-10-21 17:08:48', '2023-10-21 16:55:44');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (14, 'aliyunSmsAccessKeySecret', 'accessKeySecret', '2023-10-21 17:08:48', '2023-10-21 16:55:44');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (15, 'aliyunSmsEndpoint', 'dysmsapi.aliyuncs.com', '2023-10-21 17:08:48', '2023-10-21 16:55:44');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (16, 'aliyunSmsSignName', 'JB Admin', '2023-10-21 17:08:48', '2023-10-21 16:55:44');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (17, 'aliyunSmsTemplateCode', 'SMS_xxxxxxxx', '2023-10-21 17:08:48', '2023-10-21 16:55:44');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (18, 'idCardType', 'aliyunIdCard', '2023-10-21 17:09:24', '2023-10-21 16:55:46');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (19, 'aliyunIdCardHost', 'http://idcard.market.alicloudapi.com', '2023-10-21 17:09:24', '2023-10-21 16:55:46');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (20, 'aliyunIdCardPath', '/lianzhuo/idcard', '2023-10-21 17:09:24', '2023-10-21 16:55:46');
INSERT INTO `platform_config` (`configId`, `configKey`, `configValue`, `updatedAt`, `createdAt`) VALUES (21, 'aliyunIdCardAppcode', 'appcode', '2023-10-21 17:09:24', '2023-10-21 16:55:46');
COMMIT;

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS `platform_server`;
CREATE TABLE `platform_server` (
  `serverId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '服务器ID',
  `networkIp` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公网IP',
  `localIp` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内网IP',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`serverId`) USING BTREE,
  UNIQUE KEY `networkIp` (`networkIp`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='平台服务器表';

-- ----------------------------
-- Records of platform_server
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userId` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '性别：0未设置 1男 2女',
  `birthday` date DEFAULT NULL COMMENT '生日',
  `address` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `idCardName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证姓名',
  `idCardNo` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证号码',
  `isStop` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userId`) USING BTREE,
  UNIQUE KEY `phone` (`phone`) USING BTREE,
  UNIQUE KEY `account` (`account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
