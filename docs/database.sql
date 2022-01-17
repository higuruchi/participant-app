DROP TABLE IF EXISTS `packet_logs`;
CREATE TABLE `packet_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transit_time` timestamp NOT NULL DEFAULT current_timestamp(),
  `mac_address` char(17) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- Sample data.
LOCK TABLES `packet_logs` WRITE;
INSERT INTO `packet_logs` VALUES (1,'2021-11-16 14:11:54','3c:06:30:43:3f:50'),(2,'2021-11-16 14:11:54','3c:06:30:43:3f:50');
UNLOCK TABLES;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` char(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `name` varchar(20) NOT NULL,
  `mac_address` char(17) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Sample data.
LOCK TABLES `users` WRITE;
INSERT INTO `users` VALUES ('19T325','2022-01-16 07:40:02','higuruchi','3c:06:30:43:3f:50');
UNLOCK TABLES;
