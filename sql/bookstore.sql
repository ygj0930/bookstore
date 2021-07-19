/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50710
 Source Host           : localhost:3306
 Source Schema         : bookstore

 Target Server Type    : MySQL
 Target Server Version : 50710
 File Encoding         : 65001

 Date: 19/07/2021 17:44:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `author` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `price` float(100,0) NOT NULL,
  `sales` int(100) NOT NULL,
  `stock` int(100) NOT NULL,
  `img_path` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of books
-- ----------------------------
BEGIN;
INSERT INTO `books` VALUES (1, '解忧杂货店', '东野圭吾', 27, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (2, '边城222', '沈从文', 232, 101, 99, 'img/default.jpg');
INSERT INTO `books` VALUES (3, '中国哲学史', '冯友兰', 44, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (4, '忽然七日', ' 劳伦', 19, 104, 96, 'static/img/default.jpg');
INSERT INTO `books` VALUES (5, '苏东坡传', '林语堂', 19, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (6, '百年孤独', '马尔克斯', 29, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (7, '测试修改', '测试修改', 10, 1, 99, 'nil');
INSERT INTO `books` VALUES (8, '给孩子的诗', '北岛', 22, 102, 98, 'static/img/default.jpg');
INSERT INTO `books` VALUES (9, '为奴十二年', '所罗门', 16, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (10, '平凡的世界', '路遥', 55, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (11, '悟空传', '今何在', 14, 103, 97, 'static/img/default.jpg');
INSERT INTO `books` VALUES (12, '硬派健身', '斌卡', 31, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (13, '从晚清到民国', '唐德刚', 40, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (14, '三体', '刘慈欣', 56, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (15, '看见', '柴静', 19, 102, 98, 'static/img/default.jpg');
INSERT INTO `books` VALUES (16, '活着', '余华', 11, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (17, '小王子', '安托万', 19, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (18, '我们仨', '杨绛', 11, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (19, '生命不息,折腾不止', '罗永浩', 25, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (20, '皮囊', '蔡崇达', 24, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (21, '恰到好处的幸福', '毕淑敏', 16, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (22, '大数据预测', '埃里克', 37, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (23, '人月神话', '布鲁克斯', 56, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (24, 'C语言入门经典', '霍尔顿', 45, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (25, '数学之美', '吴军', 30, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (26, 'Java编程思想', '埃史尔', 70, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (28, '图解机器学习', '杉山将', 34, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (30, '教父', '马里奥普佐', 29, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (40, 'Go语言学习笔记', '雨痕', 51, 100, 33, '');
INSERT INTO `books` VALUES (43, 'go语言基础', 'l1ng14', 30, 10, 111, 'nil');
INSERT INTO `books` VALUES (47, '测试修改', '测试修改', 10, 1, 99, 'nil');
COMMIT;

-- ----------------------------
-- Table structure for cart_items
-- ----------------------------
DROP TABLE IF EXISTS `cart_items`;
CREATE TABLE `cart_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11,2) NOT NULL,
  `book_id` int(11) NOT NULL,
  `cart_id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `book_id` (`book_id`) USING BTREE,
  KEY `cart_id` (`cart_id`) USING BTREE,
  CONSTRAINT `cart_items_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`),
  CONSTRAINT `cart_items_ibfk_2` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of cart_items
-- ----------------------------
BEGIN;
INSERT INTO `cart_items` VALUES (92, 1, 44.00, 3, 'd7bb1aa4-9f67-407e-7ddf-c5e0c3b9e745');
INSERT INTO `cart_items` VALUES (93, 3, 81.00, 1, '0b3cbf25-a22f-4c26-4993-e5426864775f');
INSERT INTO `cart_items` VALUES (94, 2, 464.00, 2, '0b3cbf25-a22f-4c26-4993-e5426864775f');
INSERT INTO `cart_items` VALUES (95, 2, 88.00, 3, '0b3cbf25-a22f-4c26-4993-e5426864775f');
INSERT INTO `cart_items` VALUES (96, 2, 54.00, 1, '7874105e-32f8-414b-686b-644d5520e782');
INSERT INTO `cart_items` VALUES (97, 2, 464.00, 2, '7874105e-32f8-414b-686b-644d5520e782');
INSERT INTO `cart_items` VALUES (98, 1, 44.00, 3, '7874105e-32f8-414b-686b-644d5520e782');
INSERT INTO `cart_items` VALUES (99, 2, 38.00, 4, '7874105e-32f8-414b-686b-644d5520e782');
INSERT INTO `cart_items` VALUES (100, 2, 38.00, 5, '7874105e-32f8-414b-686b-644d5520e782');
INSERT INTO `cart_items` VALUES (101, 3, 81.00, 1, '36f71a1f-533a-46ac-4755-68fc1ef56256');
INSERT INTO `cart_items` VALUES (102, 4, 928.00, 2, '36f71a1f-533a-46ac-4755-68fc1ef56256');
INSERT INTO `cart_items` VALUES (103, 1, 44.00, 3, '36f71a1f-533a-46ac-4755-68fc1ef56256');
INSERT INTO `cart_items` VALUES (104, 3, 696.00, 2, 'ce1ba89a-2fc7-4843-7f1f-2cbafb06cfc7');
INSERT INTO `cart_items` VALUES (105, 1, 27.00, 1, 'ce1ba89a-2fc7-4843-7f1f-2cbafb06cfc7');
COMMIT;

-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts` (
  `id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11,2) NOT NULL,
  `user_id` int(11) NOT NULL,
  `session_id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of carts
-- ----------------------------
BEGIN;
INSERT INTO `carts` VALUES ('0b3cbf25-a22f-4c26-4993-e5426864775f', 1, 27.00, 21, '93e3e444-cd1c-4929-734c-e87edf94b34d');
INSERT INTO `carts` VALUES ('36f71a1f-533a-46ac-4755-68fc1ef56256', 8, 1053.00, 21, 'eba7527a-140a-499e-47b2-f949ab1473e0');
INSERT INTO `carts` VALUES ('7874105e-32f8-414b-686b-644d5520e782', 9, 638.00, 21, '91d8c7ea-aad1-421e-6fff-ff035b7573dd');
INSERT INTO `carts` VALUES ('ce1ba89a-2fc7-4843-7f1f-2cbafb06cfc7', 4, 723.00, 21, 'dd771c7c-d5da-4da2-6de4-ce5abcd2e97a');
INSERT INTO `carts` VALUES ('d7bb1aa4-9f67-407e-7ddf-c5e0c3b9e745', 1, 44.00, 21, '0831c2da-f343-4020-4666-295d4c27c86a');
COMMIT;

-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11,2) NOT NULL,
  `title` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `author` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `price` double(11,2) NOT NULL,
  `img_path` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `order_id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `order_id` (`order_id`) USING BTREE,
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of order_items
-- ----------------------------
BEGIN;
INSERT INTO `order_items` VALUES (3, 1, 44.00, '中国哲学史', '冯友兰', 44.00, 'static/img/default.jpg', '7d9703c4-6656-4c66-791b-5de07ccc4909');
INSERT INTO `order_items` VALUES (4, 1, 19.00, '忽然七日', ' 劳伦', 19.00, 'static/img/default.jpg', '7d9703c4-6656-4c66-791b-5de07ccc4909');
INSERT INTO `order_items` VALUES (5, 1, 20.00, '扶桑', '严歌苓', 20.00, 'static/img/default.jpg', '7d9703c4-6656-4c66-791b-5de07ccc4909');
INSERT INTO `order_items` VALUES (6, 1, 22.00, '给孩子的诗', '北岛', 22.00, 'static/img/default.jpg', '7d9703c4-6656-4c66-791b-5de07ccc4909');
INSERT INTO `order_items` VALUES (7, 3, 42.00, '悟空传', '今何在', 14.00, 'static/img/default.jpg', 'fb975d0b-1183-4ce5-7235-fc9e926db11a');
INSERT INTO `order_items` VALUES (8, 1, 55.00, '平凡的世界', '路遥', 55.00, 'static/img/default.jpg', 'fb975d0b-1183-4ce5-7235-fc9e926db11a');
INSERT INTO `order_items` VALUES (9, 1, 31.00, '硬派健身', '斌卡', 31.00, 'static/img/default.jpg', 'fb975d0b-1183-4ce5-7235-fc9e926db11a');
INSERT INTO `order_items` VALUES (10, 1, 16.00, '为奴十二年', '所罗门', 16.00, 'static/img/default.jpg', 'fb975d0b-1183-4ce5-7235-fc9e926db11a');
INSERT INTO `order_items` VALUES (11, 2, 38.00, '看见', '柴静', 19.00, 'static/img/default.jpg', 'fb975d0b-1183-4ce5-7235-fc9e926db11a');
INSERT INTO `order_items` VALUES (12, 1, 19.00, '忽然七日', ' 劳伦', 19.00, 'static/img/default.jpg', '23b85b72-1bd3-4805-6ae3-b7c808b6239b');
INSERT INTO `order_items` VALUES (13, 1, 23.00, '边城', '沈从文', 23.00, 'static/img/default.jpg', '23b85b72-1bd3-4805-6ae3-b7c808b6239b');
INSERT INTO `order_items` VALUES (14, 1, 20.00, '扶桑', '严歌苓', 20.00, 'static/img/default.jpg', '23b85b72-1bd3-4805-6ae3-b7c808b6239b');
INSERT INTO `order_items` VALUES (15, 1, 22.00, '给孩子的诗', '北岛', 22.00, 'static/img/default.jpg', '23b85b72-1bd3-4805-6ae3-b7c808b6239b');
INSERT INTO `order_items` VALUES (16, 1, 45.00, 'C语言入门经典', '霍尔顿', 45.00, 'static/img/default.jpg', 'cbe615b2-f5d3-4abe-7286-2abee3ca59fb');
INSERT INTO `order_items` VALUES (17, 1, 37.00, '大数据预测', '埃里克', 37.00, 'static/img/default.jpg', 'cbe615b2-f5d3-4abe-7286-2abee3ca59fb');
INSERT INTO `order_items` VALUES (18, 2, 38.00, '忽然七日', ' 劳伦', 19.00, 'static/img/default.jpg', 'cbe615b2-f5d3-4abe-7286-2abee3ca59fb');
COMMIT;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `create_time` datetime NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11,2) NOT NULL,
  `state` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
INSERT INTO `orders` VALUES ('23b85b72-1bd3-4805-6ae3-b7c808b6239b', '2020-11-15 16:06:27', 4, 84.00, -1, 1);
INSERT INTO `orders` VALUES ('7d9703c4-6656-4c66-791b-5de07ccc4909', '2020-11-15 15:45:57', 4, 105.00, 2, 1);
INSERT INTO `orders` VALUES ('cbe615b2-f5d3-4abe-7286-2abee3ca59fb', '2020-11-15 16:08:51', 4, 120.00, 0, 16);
INSERT INTO `orders` VALUES ('fb975d0b-1183-4ce5-7235-fc9e926db11a', '2020-11-15 16:01:07', 8, 182.00, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`session_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sessions
-- ----------------------------
BEGIN;
INSERT INTO `sessions` VALUES ('0aab42ba-a4f8-4e6e-7107-740a266e7116', '333', 21);
INSERT INTO `sessions` VALUES ('3190b06768295c752f5ee1a2eb7f401c', 'lili', 1);
INSERT INTO `sessions` VALUES ('552cb4a2-f8d7-4889-55f0-b109d3109229', '3333', 20);
INSERT INTO `sessions` VALUES ('5f9493cf-516a-4b96-7adb-3e49fd6af1bd', '333', 21);
INSERT INTO `sessions` VALUES ('ad3dc6fe8d7e714180d8186bae2a0eb2', 'lilian', 14);
INSERT INTO `sessions` VALUES ('ba117a576b68b683be736ca3a7bac600', 'lili', 1);
INSERT INTO `sessions` VALUES ('beecb3dc-b43f-4122-5de9-11920361b071', '333', 21);
INSERT INTO `sessions` VALUES ('c50e680dc5d4185a707a293263a2bc33', 'lili', 1);
INSERT INTO `sessions` VALUES ('dd771c7c-d5da-4da2-6de4-ce5abcd2e97a', '333', 21);
INSERT INTO `sessions` VALUES ('eba9dc646c47ee3e05464c95a062c449', 'lili', 1);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `PASSWORD` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`,`username`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'lili', 'e10adc3949ba59abbe56e057f20f883e', '10010@qq.com');
INSERT INTO `users` VALUES (8, 'chan', '21218cca77804d2ba1922c33e0151105', '111@qq.com');
INSERT INTO `users` VALUES (9, 'kiki', '69f8ea31de0c00502b2ae571fbab1f95', '10010@136.com');
INSERT INTO `users` VALUES (11, 'Chan Mark', '8ddcff3a80f4189ca1c9d4d902c3c909', 'asfsdfsfsfs11@foxmail.com');
INSERT INTO `users` VALUES (12, 'li xiao long', '25d55ad283aa400af464c76d713c07ad', 'l34635464213114@foxmail.com');
INSERT INTO `users` VALUES (13, 'li1ng14', '091d2f632b9187475ed9ab0adac83168', 'l1ng14@foxmail.com');
INSERT INTO `users` VALUES (14, 'lilian', '25d55ad283aa400af464c76d713c07ad', '11111111@qq.com');
INSERT INTO `users` VALUES (16, 'jack', '25d55ad283aa400af464c76d713c07ad', '55555@136.com');
INSERT INTO `users` VALUES (18, 'admin1', '123456', 'admin@123.com');
INSERT INTO `users` VALUES (19, 'admin', '123456', 'admin@123.com');
INSERT INTO `users` VALUES (20, '3333', '123456', 'admin@123.com');
INSERT INTO `users` VALUES (21, '333', '222222', '2222@222.com');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
