# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20042
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 5.7.10)
# 数据库: likeadmin
# 生成时间: 2024-09-10 05:10:06 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 system_auth_admin
# ------------------------------------------------------------

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统管理成员表';

LOCK TABLES `system_auth_admin` WRITE;
/*!40000 ALTER TABLE `system_auth_admin` DISABLE KEYS */;

INSERT INTO `system_auth_admin` (`id`, `no`, `dept_id`, `post_id`, `username`, `nickname`, `password`, `avatar`, `role`, `salt`, `sort`, `is_multipoint`, `is_disable`, `is_delete`, `last_login_ip`, `last_login_time`, `create_time`, `update_time`, `delete_time`)
VALUES
	(1,'',1,0,'admin','admin','7fac2474740becfaf1ecbdd6cc8fb076','/api/static/backend_avatar.png','0','5Xar0',0,1,0,0,'127.0.0.1',1660641347,1642321599,1660287325,0);

/*!40000 ALTER TABLE `system_auth_admin` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 system_auth_dept
# ------------------------------------------------------------

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统部门管理表';

LOCK TABLES `system_auth_dept` WRITE;
/*!40000 ALTER TABLE `system_auth_dept` DISABLE KEYS */;

INSERT INTO `system_auth_dept` (`id`, `pid`, `name`, `duty`, `mobile`, `sort`, `is_stop`, `is_delete`, `create_time`, `update_time`, `delete_time`)
VALUES
	(1,0,'默认部门','康明','18327647788',10,0,0,1649841995,1660190949,0);

/*!40000 ALTER TABLE `system_auth_dept` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 system_auth_menu
# ------------------------------------------------------------

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统菜单管理表';

LOCK TABLES `system_auth_menu` WRITE;
/*!40000 ALTER TABLE `system_auth_menu` DISABLE KEYS */;

INSERT INTO `system_auth_menu` (`id`, `pid`, `menu_type`, `menu_name`, `menu_icon`, `menu_sort`, `perms`, `paths`, `component`, `selected`, `params`, `is_cache`, `is_show`, `is_disable`, `create_time`, `update_time`)
VALUES
	(1,0,'C','工作台','el-icon-Monitor',50,'index:console','workbench','workbench/index','','',1,1,0,1650341765,1668672757),
	(100,0,'M','权限管理','el-icon-Lock',44,'','permission','','','',0,1,0,1650341765,1662626201),
	(101,100,'C','管理员','local-icon-wode',0,'system:admin:list','admin','permission/admin/index','','',1,1,0,1650341765,1663301404),
	(102,101,'A','管理员详情','',0,'system:admin:detail','','','','',0,1,0,1650341765,1660201785),
	(103,101,'A','管理员新增','',0,'system:admin:add','','','','',0,1,0,1650341765,1650341765),
	(104,101,'A','管理员编辑','',0,'system:admin:edit','','','','',0,1,0,1650341765,1650341765),
	(105,101,'A','管理员删除','',0,'system:admin:del','','','','',0,1,0,1650341765,1650341765),
	(106,101,'A','管理员状态','',0,'system:admin:disable','','','','',0,1,0,1650341765,1650341765),
	(110,100,'C','角色管理','el-icon-Female',0,'system:role:list','role','permission/role/index','','',1,1,0,1650341765,1663301451),
	(111,110,'A','角色详情','',0,'system:role:detail','','','','',0,1,0,1650341765,1650341765),
	(112,110,'A','角色新增','',0,'system:role:add','','','','',0,1,0,1650341765,1650341765),
	(113,110,'A','角色编辑','',0,'system:role:edit','','','','',0,1,0,1650341765,1650341765),
	(114,110,'A','角色删除','',0,'system:role:del','','','','',0,1,0,1650341765,1650341765),
	(120,100,'C','菜单管理','el-icon-Operation',0,'system:menu:list','menu','permission/menu/index','','',1,1,0,1650341765,1663301388),
	(121,120,'A','菜单详情','',0,'system:menu:detail','','','','',0,1,0,1650341765,1650341765),
	(122,120,'A','菜单新增','',0,'system:menu:add','','','','',0,1,0,1650341765,1650341765),
	(123,120,'A','菜单编辑','',0,'system:menu:edit','','','','',0,1,0,1650341765,1650341765),
	(124,120,'A','菜单删除','',0,'system:menu:del','','','','',0,1,0,1650341765,1650341765),
	(130,0,'M','组织管理','el-icon-OfficeBuilding',45,'','organization','','','',0,1,0,1650341765,1664416715),
	(131,130,'C','部门管理','el-icon-Coordinate',0,'system:dept:list','department','organization/department/index','','',1,1,0,1650341765,1660201994),
	(132,131,'A','部门详情','',0,'system:dept:detail','','','','',0,1,0,1650341765,1650341765),
	(133,131,'A','部门新增','',0,'system:dept:add','','','','',0,1,0,1650341765,1650341765),
	(134,131,'A','部门编辑','',0,'system:dept:edit','','','','',0,1,0,1650341765,1650341765),
	(135,131,'A','部门删除','',0,'system:dept:del','','','','',0,1,0,1650341765,1650341765),
	(140,130,'C','岗位管理','el-icon-PriceTag',0,'system:post:list','post','organization/post/index','','',1,1,0,1650341765,1660202057),
	(141,140,'A','岗位详情','',0,'system:post:detail','','','','',0,1,0,1650341765,1650341765),
	(142,140,'A','岗位新增','',0,'system:post:add','','','','',0,1,0,1650341765,1650341765),
	(143,140,'A','岗位编辑','',0,'system:post:edit','','','','',0,1,0,1650341765,1650341765),
	(144,140,'A','岗位删除','',0,'system:post:del','','','','',0,1,0,1650341765,1650341765),
	(200,0,'M','其它管理','',0,'','','','','',0,0,0,1650341765,1660636870),
	(201,200,'M','图库管理','',0,'','','','','',0,0,0,1650341765,1650341765),
	(202,201,'A','文件列表','',0,'albums:albumList','','','','',0,0,0,1650341765,1650341765),
	(203,201,'A','文件命名','',0,'albums:albumRename','','','','',0,0,0,1650341765,1650341765),
	(204,201,'A','文件移动','',0,'albums:albumMove','','','','',0,0,0,1650341765,1650341765),
	(205,201,'A','文件删除','',0,'albums:albumDel','','','','',0,0,0,1650341765,1650341765),
	(206,201,'A','分类列表','',0,'albums:cateList','','','','',0,0,0,1650341765,1650341765),
	(207,201,'A','分类新增','',0,'albums:cateAdd','','','','',0,0,0,1650341765,1650341765),
	(208,201,'A','分类命名','',0,'albums:cateRename','','','','',0,0,0,1650341765,1650341765),
	(209,201,'A','分类删除','',0,'albums:cateDel','','','','',0,0,0,1650341765,1650341765),
	(215,200,'M','上传管理','',0,'','','','','',0,0,0,1650341765,1650341765),
	(216,215,'A','上传图片','',0,'upload:image','','','','',0,0,0,1650341765,1650341765),
	(217,215,'A','上传视频','',0,'upload:video','','','','',0,0,0,1650341765,1650341765),
	(500,0,'M','系统设置','el-icon-Setting',0,'','setting','','','',0,1,0,1650341765,1662626322),
	(501,500,'M','网站设置','el-icon-Basketball',10,'','website','','','',0,1,0,1650341765,1663233572),
	(502,501,'C','网站信息','',0,'setting:website:detail','information','setting/website/information','','',0,1,0,1650341765,1660202218),
	(503,502,'A','保存配置','',0,'setting:website:save','','','','',0,0,0,1650341765,1650341765),
	(505,501,'C','网站备案','',0,'setting:copyright:detail','filing','setting/website/filing','','',0,1,0,1650341765,1660202294),
	(506,505,'A','备案保存','',0,'setting:copyright:save','','setting/website/protocol','','',0,0,0,1650341765,1650341765),
	(510,501,'C','政策协议','',0,'setting:protocol:detail','protocol','setting/website/protocol','','',0,1,0,1660027606,1660202312),
	(511,510,'A','协议保存','',0,'setting:protocol:save','','','','',0,0,0,1660027606,1663670865),
	(515,600,'C','字典管理','el-icon-Box',0,'setting:dict:type:list','dict','setting/dict/type/index','','',0,1,0,1660035436,1663226087),
	(516,515,'A','字典类型新增','',0,'setting:dict:type:add','','','','',0,1,0,1660202761,1660202761),
	(517,515,'A','字典类型编辑','',0,'setting:dict:type:edit','','','','',0,1,0,1660202842,1660202842),
	(518,515,'A','字典类型删除','',0,'setting:dict:type:del','','','','',0,1,0,1660202903,1660202903),
	(519,600,'C','字典数据管理','',0,'setting:dict:data:list','dict/data','setting/dict/data/index','/dev_tools/dict','',0,0,0,1660202948,1663309252),
	(520,515,'A','字典数据新增','',0,'setting:dict:data:add','','','','',0,1,0,1660203117,1660203117),
	(521,515,'A','字典数据编辑','',0,'setting:dict:data:edit','','','','',0,1,0,1660203142,1660203142),
	(522,515,'A','字典数据删除','',0,'setting:dict:data:del','','','','',0,1,0,1660203159,1660203159),
	(550,500,'M','系统维护','el-icon-SetUp',0,'','system','','','',0,1,0,1650341765,1660202466),
	(551,550,'C','系统环境','',0,'monitor:server','environment','setting/system/environment','','',0,1,0,1650341765,1650341765),
	(552,550,'C','系统缓存','',0,'monitor:cache','cache','setting/system/cache','','',0,1,0,1650341765,1650341765),
	(553,550,'C','系统日志','',0,'system:log:operate','journal','setting/system/journal','','',0,1,0,1650341765,1650341765),
	(555,500,'C','存储设置','el-icon-FolderOpened',6,'setting:storage:list','storage','setting/storage/index','','',0,1,0,1650341765,1663312996),
	(556,555,'A','保存配置','',0,'setting:storage:edit','','','','',0,1,0,1650341765,1650341765),
	(600,0,'M','开发工具','el-icon-EditPen',0,'','dev_tools','','','',0,1,0,1660027606,1664335701),
	(610,600,'C','代码生成器','el-icon-DocumentAdd',0,'gen:list','code','dev_tools/code/index','','',0,1,0,1660028954,1660532510),
	(611,610,'A','导入数据表','',0,'gen:importTable','','','','',0,1,0,1660532389,1660532389),
	(612,610,'A','生成代码','',0,'gen:genCode','','','','',0,1,0,1660532421,1660532421),
	(613,610,'A','下载代码','',0,'gen:downloadCode','','','','',0,1,0,1660532437,1660532437),
	(614,610,'A','预览代码','',0,'gen:previewCode','','','','',0,1,0,1660532549,1660532549),
	(616,610,'A','同步表结构','',0,'gen:syncTable','','','','',0,1,0,1660532781,1660532781),
	(617,610,'A','删除数据表','',0,'gen:delTable','','','','',0,1,0,1660532800,1660532800),
	(618,610,'A','数据表详情','',0,'gen:detail','','','','',0,1,0,1660532964,1660532977),
	(700,0,'M','素材管理','el-icon-Picture',43,'','material','','','',0,1,0,1660203293,1663300847),
	(701,700,'C','素材中心','el-icon-PictureRounded',0,'','index','material/index','','',0,1,0,1660203402,1663301493);

/*!40000 ALTER TABLE `system_auth_menu` ENABLE KEYS */;
UNLOCK TABLES;


# 转储表 system_auth_perm
# ------------------------------------------------------------

DROP TABLE IF EXISTS `system_auth_perm`;

CREATE TABLE `system_auth_perm` (
  `id` varchar(100) NOT NULL DEFAULT '' COMMENT '主键',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色菜单表';



# 转储表 system_auth_role
# ------------------------------------------------------------

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统角色管理表';

LOCK TABLES `system_auth_role` WRITE;
/*!40000 ALTER TABLE `system_auth_role` DISABLE KEYS */;

INSERT INTO `system_auth_role` (`id`, `name`, `remark`, `sort`, `is_disable`, `create_time`, `update_time`)
VALUES
	(1,'审核员','审核数据',0,0,1668679451,1668679468);

/*!40000 ALTER TABLE `system_auth_role` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
