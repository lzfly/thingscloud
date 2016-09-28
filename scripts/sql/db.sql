/*
SQLyog v10.2 
MySQL - 5.1.73 : Database - thingscloud
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`thingscloud` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `thingscloud`;

/*Table structure for table `dev_device` */

DROP TABLE IF EXISTS `dev_device`;

CREATE TABLE `dev_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `type_code` int(11) DEFAULT NULL,
  `type_name` varchar(64) DEFAULT NULL,
  `dev_model` varchar(32) DEFAULT NULL,
  `dev_ver` varchar(32) DEFAULT NULL,
  `protocol` varchar(32) DEFAULT NULL,
  `dev_name` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `is_online` int(11) DEFAULT NULL,
  `activetime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8;

/*Data for the table `dev_device` */

/*Table structure for table `dev_deviceattr` */

DROP TABLE IF EXISTS `dev_deviceattr`;

CREATE TABLE `dev_deviceattr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `attr_name` varchar(64) DEFAULT NULL,
  `attr_code` int(11) DEFAULT NULL,
  `datatype` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

/*Data for the table `dev_deviceattr` */

insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (2,'开关',1001,'bool');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (3,'开度',1002,'float% ');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (4,'浓度',1003,'float');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (5,'总量',1004,'float');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (6,'亮度',1005,'long');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (7,'颜色',1006,'long');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (8,'色温',1007,'long');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (9,'温度',1008,'float');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (10,'湿度',1009,'float');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (11,'窗帘开关停',1010,'int');
insert  into `dev_deviceattr`(`id`,`attr_name`,`attr_code`,`datatype`) values (12,'报警',1011,'bool');

/*Table structure for table `dev_deviceattrinfo` */

DROP TABLE IF EXISTS `dev_deviceattrinfo`;

CREATE TABLE `dev_deviceattrinfo` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `attr_code` int(11) DEFAULT NULL,
  `attr_name` varchar(64) DEFAULT NULL,
  `attr_permission` varchar(32) DEFAULT NULL,
  `attr_value_ctrl` varchar(64) DEFAULT NULL,
  `is_control` int(11) DEFAULT NULL,
  `attr_value_cur` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `service_code` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8;

/*Data for the table `dev_deviceattrinfo` */

/*Table structure for table `dev_devicectrl` */

DROP TABLE IF EXISTS `dev_devicectrl`;

CREATE TABLE `dev_devicectrl` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `attr_code` int(11) DEFAULT NULL,
  `attr_value` varchar(64) DEFAULT NULL,
  `source` varchar(32) DEFAULT NULL,
  `is_control` int(11) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `ctrltime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicectrl` */

/*Table structure for table `dev_devicedriver` */

DROP TABLE IF EXISTS `dev_devicedriver`;

CREATE TABLE `dev_devicedriver` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `driver_name` varchar(64) DEFAULT NULL,
  `driver_ver` varchar(32) DEFAULT NULL,
  `driver_disc` varchar(128) DEFAULT NULL,
  `driver_url` varchar(128) DEFAULT NULL,
  `active_state` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicedriver` */

/*Table structure for table `dev_devicegroup` */

DROP TABLE IF EXISTS `dev_devicegroup`;

CREATE TABLE `dev_devicegroup` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(64) DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicegroup` */

insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (2,'sleep','guest');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (3,'out','guest');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (4,'meeting','guest');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (5,'entertainment','guest');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (6,'sleep','test');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (7,'out','test');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (8,'meeting','test');
insert  into `dev_devicegroup`(`id`,`group_name`,`username`) values (9,'entertainment','test');


/*Table structure for table `dev_devicegroup_bind_device` */

DROP TABLE IF EXISTS `dev_devicegroup_bind_device`;

CREATE TABLE `dev_devicegroup_bind_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(64) DEFAULT NULL,
  `device_sn` varchar(64) DEFAULT NULL,
  `device_name` varchar(64) DEFAULT NULL,
  `type_code` int(11) DEFAULT NULL,
  `attr_code` int(11) DEFAULT NULL,
  `attr_value` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicegroup_bind_device` */

/*Table structure for table `dev_devicereport` */

DROP TABLE IF EXISTS `dev_devicereport`;

CREATE TABLE `dev_devicereport` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `device_sn` varchar(64) DEFAULT NULL,
  `attr_code` int(20) DEFAULT NULL,
  `attr_value` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `reporttime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicereport` */

/*Table structure for table `dev_devicetype` */

DROP TABLE IF EXISTS `dev_devicetype`;

CREATE TABLE `dev_devicetype` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type_name` varchar(64) DEFAULT NULL,
  `type_code` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

/*Data for the table `dev_devicetype` */

insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (2,'窗帘',2);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (3,'PM2.5',3);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (4,'甲醛报警',4);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (5,'燃气报警',5);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (6,'烟雾报警',6);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (7,'电表',7);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (8,'水表',8);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (9,'温湿度',9);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (10,'亮度色温灯',15);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (11,'开关',16);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (12,'插座',17);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (13,'门磁',18);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (14,'人体红外',19);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (15,'燃气表',20);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (16,'热表',21);
insert  into `dev_devicetype`(`id`,`type_name`,`type_code`) values (32,'摄像头',10);

/*Table structure for table `dev_driver_bind_device` */

DROP TABLE IF EXISTS `dev_driver_bind_device`;

CREATE TABLE `dev_driver_bind_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `driver_id` bigint(20) DEFAULT NULL,
  `dev_type` int(11) DEFAULT NULL,
  `dev_model` varchar(32) DEFAULT NULL,
  `dev_ver` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dev_driver_bind_device` */

/*Table structure for table `dev_gateway` */

DROP TABLE IF EXISTS `dev_gateway`;

CREATE TABLE `dev_gateway` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `gw_model` varchar(32) DEFAULT NULL,
  `gw_type` varchar(32) DEFAULT NULL,
  `gw_mac` varchar(32) DEFAULT NULL,
  `wifi_ssid` varchar(64) DEFAULT NULL,
  `wifi_pwd` varchar(32) DEFAULT NULL,
  `hw_ver` varchar(32) DEFAULT NULL,
  `sw_ver` varchar(64) DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  `operate_cap` int(11) DEFAULT NULL,
  `activetime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8;

/*Data for the table `dev_gateway` */

/*Table structure for table `dev_logicentity` */

DROP TABLE IF EXISTS `dev_logicentity`;

CREATE TABLE `dev_logicentity` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `if_device_sn` varchar(64) DEFAULT NULL,
  `if_device_name` varchar(64) DEFAULT NULL,
  `if_type_code` int(11) DEFAULT NULL,
  `if_attr_code` int(11) DEFAULT NULL,
  `if_operate_code` varchar(32) DEFAULT NULL,
  `if_attr_value` varchar(64) DEFAULT NULL,
  `if_attr_value2` varchar(64) DEFAULT NULL,
  `th_device_sn` varchar(64) DEFAULT NULL,
  `th_device_name` varchar(64) DEFAULT NULL,
  `th_type_code` int(11) DEFAULT NULL,
  `th_attr_code` int(11) DEFAULT NULL,
  `th_attr_value` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Data for the table `dev_logicentity` */

/*Table structure for table `dev_mqtt_con_state` */

DROP TABLE IF EXISTS `dev_mqtt_con_state`;

CREATE TABLE `dev_mqtt_con_state` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `dev_sn` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `conn_state` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dev_mqtt_con_state` */

/*Table structure for table `dev_mqtt_server` */

DROP TABLE IF EXISTS `dev_mqtt_server`;

CREATE TABLE `dev_mqtt_server` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `server_ip` varchar(64) DEFAULT NULL,
  `server_port` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dev_mqtt_server` */

/*Table structure for table `dev_room` */

DROP TABLE IF EXISTS `dev_room`;

CREATE TABLE `dev_room` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room_name` varchar(64) DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `dev_room` */

insert  into `dev_room`(`id`,`room_name`,`username`) values (4,'卧室','guest');
insert  into `dev_room`(`id`,`room_name`,`username`) values (5,'客厅','guest');
insert  into `dev_room`(`id`,`room_name`,`username`) values (6,'厨房','guest');

/*Table structure for table `dev_room_bind_device` */

DROP TABLE IF EXISTS `dev_room_bind_device`;

CREATE TABLE `dev_room_bind_device` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `room_name` varchar(64) DEFAULT NULL,
  `device_sn` varchar(64) DEFAULT NULL,
  `device_name` varchar(64) DEFAULT NULL,
  `type_code` int(11) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

/*Data for the table `dev_room_bind_device` */

/*Table structure for table `dev_service` */

DROP TABLE IF EXISTS `dev_service`;

CREATE TABLE `dev_service` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `service_code` int(11) DEFAULT NULL,
  `service_name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `dev_service` */

/*Table structure for table `dev_user_bind_device` */

DROP TABLE IF EXISTS `dev_user_bind_device`;

CREATE TABLE `dev_user_bind_device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `device_sn` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;

/*Data for the table `dev_user_bind_device` */


/*Table structure for table `dev_user_bind_gateway` */

DROP TABLE IF EXISTS `dev_user_bind_gateway`;

CREATE TABLE `dev_user_bind_gateway` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `gateway_sn` varchar(64) DEFAULT NULL,
  `is_master` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=158 DEFAULT CHARSET=utf8;

/*Data for the table `dev_user_bind_gateway` */

/*Table structure for table `usr_user` */

DROP TABLE IF EXISTS `usr_user`;

CREATE TABLE `usr_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `phone` varchar(32) NOT NULL,
  `mail` varchar(64) DEFAULT NULL,
  `registertime` datetime DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=225 DEFAULT CHARSET=utf8;

/*Data for the table `usr_user` */

insert  into `usr_user`(`id`,`username`,`password`,`phone`,`mail`,`registertime`,`state`) values (123,'test','123456','18610330033','luz@1234.com','2016-12-01 18:43:12',10);
insert  into `usr_user`(`id`,`username`,`password`,`phone`,`mail`,`registertime`,`state`) values (143,'luz','234','13810501286','luz@1344.com','2015-16-03 16:13:18',1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
