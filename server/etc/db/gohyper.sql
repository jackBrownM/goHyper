/*
Navicat MySQL Data Transfer

Source Server         : 本地mysql5
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : gohyper

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2024-09-18 23:07:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for system_auth_admin
-- ----------------------------
DROP TABLE IF EXISTS `system_auth_admin`;
CREATE TABLE `system_auth_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `no` varchar(20) NOT NULL COMMENT '用户编号',
  `dept_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '部门ID',
  `post_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '岗位ID',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户账号',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '用户密码',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '用户头像',
  `role` varchar(200) NOT NULL DEFAULT '' COMMENT '角色主键',
  `salt` varchar(20) NOT NULL DEFAULT '' COMMENT '加密盐巴',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_multipoint` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '多端登录: 0=否, 1=是',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `last_login_ip` varchar(20) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `last_login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `delete_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统管理成员表';

-- ----------------------------
-- Records of system_auth_admin
-- ----------------------------
INSERT INTO `system_auth_admin` VALUES ('1', '', '1', '0', 'admin', 'admin', '7fac2474740becfaf1ecbdd6cc8fb076', '/api/static/backend_avatar.png', '0', '5Xar0', '0', '1', '0', '0', '127.0.0.1', '1726669006', '1642321599', '1660287325', '0');
INSERT INTO `system_auth_admin` VALUES ('2', '', '0', '0', 'test', ' test12345678', '5f11ec1c6a01df995e4b7d22164b45fe', '', '0', 'dKSX8', '1', '1', '0', '1', '', '0', '0', '0', '1726667975');

-- ----------------------------
-- Table structure for system_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `system_auth_dept`;
CREATE TABLE `system_auth_dept` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '部门名称',
  `duty` varchar(30) NOT NULL DEFAULT '' COMMENT '负责人名',
  `mobile` varchar(30) NOT NULL DEFAULT '' COMMENT '联系电话',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序编号',
  `is_stop` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `is_delete` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除: 0=否, 1=是',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `delete_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统部门管理表';

-- ----------------------------
-- Records of system_auth_dept
-- ----------------------------
INSERT INTO `system_auth_dept` VALUES ('1', '0', '默认部门', '康明', '18327647788', '10', '0', '0', '1649841995', '1660190949', '0');

-- ----------------------------
-- Table structure for system_auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_auth_menu`;
CREATE TABLE `system_auth_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `menu_type` char(2) NOT NULL DEFAULT '' COMMENT '权限类型: M=目录，C=菜单，A=按钮',
  `menu_name` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menu_icon` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `menu_sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '菜单排序',
  `perms` varchar(100) NOT NULL DEFAULT '' COMMENT '权限标识',
  `paths` varchar(100) NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(200) NOT NULL DEFAULT '' COMMENT '前端组件',
  `selected` varchar(200) NOT NULL DEFAULT '' COMMENT '选中路径',
  `params` varchar(200) NOT NULL DEFAULT '' COMMENT '路由参数',
  `is_cache` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否缓存: 0=否, 1=是',
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示: 0=否, 1=是',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=702 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统菜单管理表';

-- ----------------------------
-- Records of system_auth_menu
-- ----------------------------
INSERT INTO `system_auth_menu` VALUES ('1', '0', 'C', '工作台', 'el-icon-Monitor', '50', 'index:console', 'workbench', 'workbench/index', '', '', '1', '1', '0', '1650341765', '1668672757');
INSERT INTO `system_auth_menu` VALUES ('100', '0', 'M', '权限管理', 'el-icon-Lock', '44', '', 'permission', '', '', '', '0', '1', '0', '1650341765', '1662626201');
INSERT INTO `system_auth_menu` VALUES ('101', '100', 'C', '管理员', 'local-icon-wode', '0', 'system:admin:list', 'admin', 'permission/admin/index', '', '', '1', '1', '0', '1650341765', '1663301404');
INSERT INTO `system_auth_menu` VALUES ('102', '101', 'A', '管理员详情', '', '0', 'system:admin:detail', '', '', '', '', '0', '1', '0', '1650341765', '1660201785');
INSERT INTO `system_auth_menu` VALUES ('103', '101', 'A', '管理员新增', '', '0', 'system:admin:add', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('104', '101', 'A', '管理员编辑', '', '0', 'system:admin:edit', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('105', '101', 'A', '管理员删除', '', '0', 'system:admin:del', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('106', '101', 'A', '管理员状态', '', '0', 'system:admin:disable', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('110', '100', 'C', '角色管理', 'el-icon-Female', '0', 'system:role:list', 'role', 'permission/role/index', '', '', '1', '1', '0', '1650341765', '1663301451');
INSERT INTO `system_auth_menu` VALUES ('111', '110', 'A', '角色详情', '', '0', 'system:role:detail', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('112', '110', 'A', '角色新增', '', '0', 'system:role:add', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('113', '110', 'A', '角色编辑', '', '0', 'system:role:edit', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('114', '110', 'A', '角色删除', '', '0', 'system:role:del', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('120', '100', 'C', '菜单管理', 'el-icon-Operation', '0', 'system:menu:list', 'menu', 'permission/menu/index', '', '', '1', '1', '0', '1650341765', '1663301388');
INSERT INTO `system_auth_menu` VALUES ('121', '120', 'A', '菜单详情', '', '0', 'system:menu:detail', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('122', '120', 'A', '菜单新增', '', '0', 'system:menu:add', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('123', '120', 'A', '菜单编辑', '', '0', 'system:menu:edit', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');
INSERT INTO `system_auth_menu` VALUES ('124', '120', 'A', '菜单删除', '', '0', 'system:menu:del', '', '', '', '', '0', '1', '0', '1650341765', '1650341765');

-- ----------------------------
-- Table structure for system_auth_perm
-- ----------------------------
DROP TABLE IF EXISTS `system_auth_perm`;
CREATE TABLE `system_auth_perm` (
  `id` varchar(100) NOT NULL DEFAULT '' COMMENT '主键',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色菜单表';

-- ----------------------------
-- Records of system_auth_perm
-- ----------------------------

-- ----------------------------
-- Table structure for system_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `system_auth_role`;
CREATE TABLE `system_auth_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '角色排序',
  `is_disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否禁用: 0=否, 1=是',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色管理表';

-- ----------------------------
-- Records of system_auth_role
-- ----------------------------
INSERT INTO `system_auth_role` VALUES ('1', '审核员', '审核数据', '0', '0', '1668679451', '1668679468');

-- ----------------------------
-- Table structure for system_config
-- ----------------------------
DROP TABLE IF EXISTS `system_config`;
CREATE TABLE `system_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` varchar(30) DEFAULT '' COMMENT '类型',
  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '键',
  `value` text COMMENT '值',
  `create_time` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统全局配置表';

-- ----------------------------
-- Records of system_config
-- ----------------------------
INSERT INTO `system_config` VALUES ('1', 'storage', 'default', 'local', '1660620367', '1662620927');
INSERT INTO `system_config` VALUES ('2', 'storage', 'local', '{\"name\":\"本地存储\"}', '1660620367', '1662620927');
INSERT INTO `system_config` VALUES ('3', 'storage', 'qiniu', '{\"name\":\"七牛云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('4', 'storage', 'aliyun', '{\"name\":\"阿里云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', '1660620367', '1662620071');
INSERT INTO `system_config` VALUES ('5', 'storage', 'qcloud', '{\"name\":\"腾讯云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\",\"region\":\"\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('6', 'sms', 'default', 'aliyun', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('7', 'sms', 'aliyun', '{\"name\":\"阿里云短信\",\"alias\":\"aliyun\",\"sign\":\"\",\"appKey\":\"\",\"secretKey\":\"\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('8', 'sms', 'tencent', '{\"name\":\"腾讯云短信\",\"alias\":\"tencent\",\"sign\":\"\",\"appId\":\"\",\"secretId\":\"\",\"secretKey\":\"\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('9', 'sms', 'huawei', '{\"name\":\"华为云短信\",\"alias\":\"huawei\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('10', 'website', 'name', '软兴内部后台', '1660620367', '1724596493');
INSERT INTO `system_config` VALUES ('11', 'website', 'logo', '/api/static/backend_logo.png', '1660620367', '1724596493');
INSERT INTO `system_config` VALUES ('12', 'website', 'favicon', '/api/static/backend_favicon.ico', '1660620367', '1724596493');
INSERT INTO `system_config` VALUES ('13', 'website', 'backdrop', '/api/static/backend_backdrop.png', '1660620367', '1724596493');
INSERT INTO `system_config` VALUES ('14', 'website', 'copyright', '[{\"name\":\"LikeAdmin开源系统\",\"link\":\"http://www.beian.gov.cn\"}]', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('15', 'website', 'shopName', '软兴开源系统', '1631255140', '1724596493');
INSERT INTO `system_config` VALUES ('16', 'website', 'shopLogo', '/api/static/shop_logo.png', '1631255140', '1724596493');
INSERT INTO `system_config` VALUES ('17', 'protocol', 'service', '{\"name\":\"服务协议\",\"content\":\"\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('18', 'protocol', 'privacy', '{\"name\":\"隐私协议\",\"content\":\"\"}', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('19', 'tabbar', 'style', '{\"defaultColor\":\"#4A5DFF\",\"selectedColor\":\"#EA5455\"}', '1660620367', '1662544900');
INSERT INTO `system_config` VALUES ('20', 'search', 'isHotSearch', '0', '1660620367', '1662546997');
INSERT INTO `system_config` VALUES ('30', 'h5_channel', 'status', '1', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('31', 'h5_channel', 'close', '0', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('32', 'h5_channel', 'url', '', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('40', 'mp_channel', 'name', '', '1660620367', '1662551403');
INSERT INTO `system_config` VALUES ('41', 'mp_channel', 'primaryId', '', '1660620367', '1662551403');
INSERT INTO `system_config` VALUES ('42', 'mp_channel', 'appId', '', '1660620367', '1662551403');
INSERT INTO `system_config` VALUES ('43', 'mp_channel', 'appSecret', '', '1660620367', '1662551403');
INSERT INTO `system_config` VALUES ('44', 'mp_channel', 'qrCode', '', '1660620367', '1662551403');
INSERT INTO `system_config` VALUES ('50', 'wx_channel', 'appId', '', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('51', 'wx_channel', 'appSecret', '', '1660620367', '1660620367');
INSERT INTO `system_config` VALUES ('55', 'oa_channel', 'name', '', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('56', 'oa_channel', 'primaryId', ' ', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('57', 'oa_channel', 'qrCode', '', '1662551337', '1662551337');
INSERT INTO `system_config` VALUES ('58', 'oa_channel', 'appId', '', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('59', 'oa_channel', 'appSecret', '', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('60', 'oa_channel', 'url', '', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('61', 'oa_channel', 'token', '', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('62', 'oa_channel', 'encodingAesKey', '', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('63', 'oa_channel', 'encryptionType', '1', '1660620367', '1662551337');
INSERT INTO `system_config` VALUES ('64', 'oa_channel', 'menus', '[]', '1631255140', '1663118712');
INSERT INTO `system_config` VALUES ('70', 'login', 'loginWay', '1,2', '1660620367', '1662538771');
INSERT INTO `system_config` VALUES ('71', 'login', 'forceBindMobile', '0', '1660620367', '1662538771');
INSERT INTO `system_config` VALUES ('72', 'login', 'openAgreement', '1', '1660620367', '1662538771');
INSERT INTO `system_config` VALUES ('73', 'login', 'openOtherAuth', '1', '1660620367', '1662538771');
INSERT INTO `system_config` VALUES ('74', 'login', 'autoLoginAuth', '1,2', '1660620367', '1662538771');
INSERT INTO `system_config` VALUES ('80', 'user', 'defaultAvatar', '/api/static/default_avatar.png', '1660620367', '1662535156');
