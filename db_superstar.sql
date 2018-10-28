/*
SQLyog Ultimate v11.24 (32 bit)
MySQL - 5.5.46-0ubuntu0.14.04.2 : Database - superstar
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`superstar` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `superstar`;

/*Table structure for table `star_info` */

DROP TABLE IF EXISTS `star_info`;

CREATE TABLE `star_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name_zh` varchar(50) NOT NULL DEFAULT '' COMMENT '中文名',
  `name_en` varchar(50) NOT NULL DEFAULT '' COMMENT '英文名',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `birthday` varchar(50) NOT NULL DEFAULT '' COMMENT '出生日期',
  `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '身高，单位cm',
  `weight` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '体重，单位kg',
  `club` varchar(50) NOT NULL DEFAULT '' COMMENT '俱乐部',
  `jersy` varchar(50) NOT NULL DEFAULT '' COMMENT '球衣号码以及主打位置',
  `country` varchar(50) NOT NULL DEFAULT '' COMMENT '国籍',
  `birthaddress` varchar(255) NOT NULL DEFAULT '' COMMENT '出生地',
  `feature` varchar(255) NOT NULL DEFAULT '' COMMENT '个人特点',
  `moreinfo` text COMMENT '更多介绍',
  `sys_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，默认值 0 正常，1 删除',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

/*Data for the table `star_info` */

insert  into `star_info`(`id`,`name_zh`,`name_en`,`avatar`,`birthday`,`height`,`weight`,`club`,`jersy`,`country`,`birthaddress`,`feature`,`moreinfo`,`sys_status`,`sys_created`,`sys_updated`) values (1,'C罗','Cristiano Ronaldo','http://www.sinaimg.cn/lf/sports/wc_2018/player/14937.jpg','1985-02-05',187,83,'皇家马德里','前锋7号','葡萄牙','','','',0,0,1529144171),(2,'梅西','Lionel Messi','http://www.sinaimg.cn/lf/sports/wc_2018/player/19054.jpg','1987-06-24',170,72,'巴塞罗那','前锋10号','阿根廷','阿根廷圣菲省罗萨里奥','','绰号“新马拉多纳”',0,0,1529135214),(3,'穆勒','Thomas Müller','http://www.sinaimg.cn/lf/sports/wc_2018/player/55634.jpg','1989-09-13',186,75,'拜仁慕尼黑','前锋13号','德国','未知','未知','',0,0,1529149343);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
