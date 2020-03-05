-- MySQL dump 10.13  Distrib 5.7.27, for macos10.14 (x86_64)
--
-- Host: localhost    Database: Beq
-- ------------------------------------------------------
-- Server version	5.7.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(45) NOT NULL,
  `LastName` varchar(45) NOT NULL,
  `Email` varchar(80) NOT NULL,
  `SecretKey` varchar(450) NOT NULL,
  `BirthDay` date NOT NULL,
  `Gender` enum('MALE','FEMALE') NOT NULL DEFAULT 'MALE',
  `Role` enum('STUDENT','TEACHER','ADMIN') NOT NULL DEFAULT 'STUDENT',
  `isActive` tinyint(4) NOT NULL DEFAULT '1',
  `DateCreated` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID_UNIQUE` (`ID`),
  UNIQUE KEY `Email_UNIQUE` (`Email`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,'Kalana','Dhanajaya','Kalana2@gmail.com','$2a$10$.MFaaICc0.Ea3xl3bUFeue/xZIDQ/dMlefqRYoHg2pmSK76/hy.K6','2020-10-10','MALE','STUDENT',1,'2020-01-19 11:40:56'),(6,'Kalana','Dhanajaya','Kalana3@gmail.com','$2a$10$5AQ0Z0zifBO8IaviTcYRde0atQCab4.lcWIo.RXr6.91MKdFQ.0L6','2020-10-10','MALE','STUDENT',1,'2020-01-19 11:50:33'),(9,'Kalana','Dhanajaya','kalifef@gmail.com','$2a$10$zEpT9Tq1B7hy90LZdSPLqOaIU4Lqhx3/dW0sRyyAFY7h6i9XI.KEy','2020-10-10','MALE','STUDENT',1,'2020-01-19 12:17:13'),(29,'Kalana','Dhanajaya','kalife2@gmail.com','$2a$10$IJAQ5Jyx9B6Ri/9BAayZMuLUbZU62zEJE/7jAVwnWjOalLJ/C7jHG','2020-10-10','MALE','STUDENT',1,'2020-01-20 00:38:31'),(38,'Kalana','Dhanajaya','kalife@gmail.com','$2a$10$N1tMkXMCmn3MZRxS1euwIODeQdBi/1/9IywgtTcWO1zJEI35OIRfO','2020-10-10','MALE','STUDENT',1,'2020-01-20 14:00:36');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-03 23:13:48
