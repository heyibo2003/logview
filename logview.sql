/*
SQLyog Ultimate v11.24 (32 bit)
MySQL - 5.7.20 : Database - logview
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`logview` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `logview`;

/*Table structure for table `env` */

DROP TABLE IF EXISTS `env`;

CREATE TABLE `env` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `envname` varchar(100) NOT NULL,
  `Create_Time` datetime DEFAULT NULL,
  `Update_Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

/*Data for the table `env` */

LOCK TABLES `env` WRITE;



UNLOCK TABLES;

/*Table structure for table `environmentals` */

DROP TABLE IF EXISTS `environmentals`;

CREATE TABLE `environmentals` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `machine` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `user` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `passwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `hostport` int(11) DEFAULT NULL,
  `domain` varchar(255) DEFAULT NULL,
  `Create_Time` datetime DEFAULT NULL,
  `Update_Time` datetime DEFAULT NULL,
  `Env_Id` int(11) NOT NULL DEFAULT '0',
  `Perject_Id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=189 DEFAULT CHARSET=utf8;

/*Data for the table `environmentals` */

LOCK TABLES `environmentals` WRITE;


/*Table structure for table `health` */

DROP TABLE IF EXISTS `health`;

CREATE TABLE `health` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `processname` varchar(100) DEFAULT NULL,
  `inspection` varchar(50) DEFAULT NULL,
  `url` varchar(255) DEFAULT 'null',
  `port` int(11) DEFAULT NULL,
  `user` varchar(100) DEFAULT NULL,
  `passwd` varchar(100) DEFAULT NULL,
  `machine` varchar(100) DEFAULT NULL,
  `hostport` int(11) DEFAULT NULL,
  `Env_Id` int(11) NOT NULL DEFAULT '0',
  `Perject_Id` int(11) NOT NULL DEFAULT '0',
  `Create_Time` datetime DEFAULT NULL,
  `Update_Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `health` */

LOCK TABLES `health` WRITE;

UNLOCK TABLES;

/*Table structure for table `perject` */

DROP TABLE IF EXISTS `perject`;

CREATE TABLE `perject` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `perject` varchar(100) NOT NULL,
  `Create_Time` datetime DEFAULT NULL,
  `Update_Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `perject` */


/*Table structure for table `software` */

DROP TABLE IF EXISTS `software`;

CREATE TABLE `software` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `domain` varchar(100) NOT NULL,
  `script` varchar(100) NOT NULL,
  `checkscript` varchar(100) NOT NULL,
  `path` varchar(255) NOT NULL,
  `port` int(10) NOT NULL,
  `softname` varchar(50) NOT NULL,
  `softpath` varchar(100) NOT NULL,
  `Env_Id` int(5) NOT NULL,
  `Perject_Id` int(5) NOT NULL,
  `Create_Time` datetime DEFAULT NULL,
  `Update_Time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8;

/*Data for the table `software` */

LOCK TABLES `software` WRITE;


/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(50) NOT NULL,
  `Birth_Date` varchar(50) NOT NULL,
  `perjectid` int(11) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `Gender` int(8) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `Phone` varchar(255) NOT NULL,
  `Create_Time` datetime DEFAULT NULL,
  `Update_Time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
