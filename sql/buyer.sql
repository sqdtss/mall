/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : mall

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 15/11/2021 18:37:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `buyer`;
CREATE TABLE `buyer` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id（主键）',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名称',
  `password` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
  `status` tinyint(20) NOT NULL DEFAULT '1' COMMENT '用户状态，1表示正常，0表示暂停',
  `created` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '创建时间',
  `updated` char(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100037 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `buyer` VALUES (100030, 'admin', '123', 1, '2021-07-16 10:35:02', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
