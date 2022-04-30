/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : localhost:3306
 Source Schema         : cloudrestaurant

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 30/04/2022 14:47:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for food_category
-- ----------------------------
DROP TABLE IF EXISTS `food_category`;
CREATE TABLE `food_category`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `image_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `link_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_in_serving` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of food_category
-- ----------------------------

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sell_count` int(11) NULL DEFAULT NULL,
  `price` float NULL DEFAULT NULL,
  `old_price` float NULL DEFAULT NULL,
  `shop_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO `goods` VALUES (1, '小小鲜肉包', '滑蛋牛肉粥(1份)+小小鲜肉包(4只)', '', 14, 25, 29, 1);
INSERT INTO `goods` VALUES (2, '滑蛋牛肉粥+小小鲜肉包', '滑蛋牛肉粥(1份)+小小鲜肉包(3只)', '', 6, 35, 41, 1);
INSERT INTO `goods` VALUES (3, '滑蛋牛肉粥+绿甘蓝馅饼', '滑蛋牛肉粥(1份)+绿甘蓝馅饼(1张)', '', 2, 25, 30, 1);
INSERT INTO `goods` VALUES (4, '茶香卤味蛋', '咸鸡蛋', '', 688, 2.5, 3, 1);
INSERT INTO `goods` VALUES (5, '韭菜鸡蛋馅饼(2张)', '韭菜鸡蛋馅饼', '', 381, 10, 12, 1);
INSERT INTO `goods` VALUES (6, '小小鲜肉包+豆浆套餐', '小小鲜肉包(8只)装+豆浆(1杯)', '', 335, 9.9, 11.9, 479);
INSERT INTO `goods` VALUES (7, '翠香炒素饼', '咸鲜翠香素炒饼', '', 260, 17.9, 20.9, 485);
INSERT INTO `goods` VALUES (8, '香煎鲜肉包', '咸鲜猪肉鲜肉包', '', 173, 10.9, 12.9, 486);

-- ----------------------------
-- Table structure for member
-- ----------------------------
DROP TABLE IF EXISTS `member`;
CREATE TABLE `member`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mobile` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `register_time` bigint(20) NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `balance` double NULL DEFAULT NULL,
  `is_active` tinyint(4) NULL DEFAULT NULL,
  `city` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of member
-- ----------------------------
INSERT INTO `member` VALUES (2, 'test', '', '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08', 1651299605, '', 0, 0, '');

-- ----------------------------
-- Table structure for shop
-- ----------------------------
DROP TABLE IF EXISTS `shop`;
CREATE TABLE `shop`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `promotion_info` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `address` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` tinyint(4) NULL DEFAULT NULL,
  `longitude` double NULL DEFAULT NULL,
  `latitude` double NULL DEFAULT NULL,
  `image_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_new` tinyint(1) NULL DEFAULT NULL,
  `is_premium` tinyint(1) NULL DEFAULT NULL,
  `rating` float NULL DEFAULT NULL,
  `rating_count` int(11) NULL DEFAULT NULL,
  `recent_order_num` int(11) NULL DEFAULT NULL,
  `minimum_order_amount` int(11) NULL DEFAULT NULL,
  `delivery_fee` int(11) NULL DEFAULT NULL,
  `opening_hours` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 489 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of shop
-- ----------------------------
INSERT INTO `shop` VALUES (1, '嘉禾一品（温都水城）', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市昌平区宏福苑温都水城F1', '13437850035', 1, 116.36868, 40.10039, '', 1, 1, 4.7, 961, 106, 20, 5, '8:30/20:30');
INSERT INTO `shop` VALUES (479, '杨国福麻辣烫', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市市蜀山区南二环路天鹅湖万达广场8号楼1705室', '13167583411', 1, 117.22124, 31.81948, '', 1, 1, 4.2, 167, 755, 20, 5, '8:30/20:30');
INSERT INTO `shop` VALUES (485, '好适口', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市海淀区西二旗大街58号', '12345678901', 1, 120.65355, 31.26578, '', 1, 1, 4.6, 576, 58, 20, 5, '8:30/20:30');
INSERT INTO `shop` VALUES (486, '东来顺旗舰店', '老北京正宗涮羊肉,非物质文化遗产', '北京市天河区东圃镇汇彩路38号1领汇创展商务中心401', '13544323775', 1, 113.41724, 23.1127, '', 1, 1, 4.2, 372, 542, 20, 5, '09:00/21:30');
INSERT INTO `shop` VALUES (487, '北京酒家', '北京第一家传承300年酒家', '北京市海淀区上下九商业步行街内', '13257482341', 0, 113.24826, 23.11488, '', 1, 1, 4.2, 871, 923, 20, 5, '8:30/20:30');
INSERT INTO `shop` VALUES (488, '和平鸽饺子馆', '吃饺子就来和平鸽饺子馆', '北京市越秀区德政中路171', '17098764762', 1, 113.27521, 23.12092, '', 1, 1, 4.2, 273, 483, 20, 5, '8:30/20:30');

-- ----------------------------
-- Table structure for sms_code
-- ----------------------------
DROP TABLE IF EXISTS `sms_code`;
CREATE TABLE `sms_code`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `biz_id` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `code` varchar(6) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_time` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sms_code
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
