-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 01, 2021 at 04:02 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `restrodb`
--

-- --------------------------------------------------------

--
-- Table structure for table `dishes`
--

CREATE TABLE `dishes` (
  `id` int(11) UNSIGNED NOT NULL,
  `restaurant_id` int(11) UNSIGNED NOT NULL DEFAULT 0,
  `admin_id` int(11) UNSIGNED NOT NULL DEFAULT 1,
  `dish_name` varchar(255) DEFAULT NULL,
  `price` decimal(11,2) UNSIGNED DEFAULT NULL,
  `description` text DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `dishes`
--

INSERT INTO `dishes` (`id`, `restaurant_id`, `admin_id`, `dish_name`, `price`, `description`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 1, 'Rice', '200.00', 'Nice Food!', 1, '2021-11-28 17:42:31.237', '2021-11-28 17:42:31.237', NULL),
(2, 0, 1, '', '0.00', '', 1, '2021-11-28 18:05:31.362', '2021-11-28 18:05:31.362', NULL),
(3, 0, 1, 'Rice22', '200.00', 'Nice Food!', 1, '2021-11-29 11:47:13.215', '2021-11-29 11:47:13.215', NULL),
(4, 1, 1, 'Rice2', '200.00', 'Nice Food!', 1, '2021-11-29 11:51:03.420', '2021-11-29 11:51:03.420', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `resources_to_roles`
--

CREATE TABLE `resources_to_roles` (
  `role_id` int(11) UNSIGNED NOT NULL DEFAULT 0,
  `resource_id` int(11) UNSIGNED NOT NULL DEFAULT 0
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `resources_to_roles`
--

INSERT INTO `resources_to_roles` (`role_id`, `resource_id`) VALUES
(1, 1),
(2, 2),
(3, 2),
(2, 3),
(3, 3),
(2, 4),
(2, 5),
(2, 6),
(2, 7),
(3, 7),
(2, 8),
(3, 8),
(2, 9),
(2, 10),
(2, 11),
(2, 12),
(3, 12),
(2, 13),
(3, 13),
(2, 14),
(3, 14),
(2, 15),
(2, 16),
(2, 17),
(2, 18),
(2, 19),
(2, 23),
(2, 24);

-- --------------------------------------------------------

--
-- Table structure for table `restaurants`
--

CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL,
  `admin_id` int(11) NOT NULL DEFAULT 1,
  `restaurant_name` varchar(64) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `latitude` decimal(12,6) NOT NULL DEFAULT 0.000000,
  `longitude` decimal(12,6) NOT NULL DEFAULT 0.000000,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `restaurants`
--

INSERT INTO `restaurants` (`id`, `admin_id`, `restaurant_name`, `address`, `latitude`, `longitude`, `status`, `created_at`, `deleted_at`, `updated_at`) VALUES
(1, 1, 'Food Resturent', '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674', 1, '2021-11-28 18:06:12.883', NULL, '2021-11-28 18:06:12.883'),
(2, 1, 'Food Resturent1', '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674', 1, '2021-11-29 11:08:46.252', NULL, '2021-11-29 11:08:46.252'),
(3, 1, 'Food Resturent1', '2681 Overlook Point, Escondido, CA 92029', '33.103265', '33.103265', 1, '2021-11-29 11:09:27.662', NULL, '2021-11-29 11:09:27.662'),
(4, 1, 'Food Resturent2', '2681 Overlook Point, Escondido, CA 92029', '33.103265', '33.103265', 1, '2021-11-29 11:51:36.953', NULL, '2021-11-29 11:51:36.953');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `admin_id` int(11) NOT NULL DEFAULT 1,
  `name` varchar(64) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `admin_id`, `name`, `email`, `password`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 'SKP', 'skp@gmail.com', '11111', 1, '2021-11-29 11:15:49.827', '2021-11-29 11:15:49.827', NULL),
(2, 1, 'SKP2', 'skp2@gmail.com', '11111', 1, '2021-11-29 11:52:31.732', '2021-11-29 11:52:31.732', NULL),
(3, 0, 'SKP3', 'skp3@gmail.com', '11111', 1, '2021-11-30 21:08:17.409', '2021-11-30 21:08:17.409', NULL),
(4, 0, 'SKP4', 'skp4@gmail.com', '11111', 1, '2021-11-30 21:22:05.901', '2021-11-30 21:22:05.901', NULL),
(5, 0, 'SKP5', 'skp5@gmail.com', '11111', 1, '2021-11-30 21:26:01.107', '2021-11-30 21:26:01.107', NULL),
(6, 0, 'SKP5', 'skp6@gmail.com', '11111', 1, '2021-11-30 21:28:53.552', '2021-11-30 21:28:53.552', NULL),
(7, 1, 'SKP5', 'skp7@gmail.com', '11111', 1, '2021-11-30 21:30:52.025', '2021-11-30 21:30:52.025', NULL),
(8, 1, 'SKP8', 'skp8@gmail.com', '11111', 1, '2021-11-30 21:35:44.162', '2021-11-30 21:35:44.162', NULL),
(9, 1, 'SKP9', 'skp9@gmail.com', '11111', 1, '2021-11-30 21:40:39.901', '2021-11-30 21:40:39.901', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `user_addresses`
--

CREATE TABLE `user_addresses` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT 0,
  `address` varchar(255) DEFAULT NULL,
  `latitude` decimal(12,6) NOT NULL DEFAULT 0.000000,
  `longitude` decimal(12,6) NOT NULL DEFAULT 0.000000
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user_addresses`
--

INSERT INTO `user_addresses` (`id`, `user_id`, `address`, `latitude`, `longitude`) VALUES
(1, 1, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(2, 1, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(3, 2, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(4, 2, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(5, 3, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(6, 3, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(7, 4, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(8, 4, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(9, 5, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(10, 5, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(11, 6, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(12, 6, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(13, 7, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(14, 7, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(15, 8, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675'),
(16, 8, '2681 Overlook Point, Escondido, CA 92029', '33.103264', '-117.127674'),
(17, 9, '2681 Overlook Point, Escondido, CA 92029', '33.103265', '-117.127675');

-- --------------------------------------------------------

--
-- Table structure for table `user_privileges`
--

CREATE TABLE `user_privileges` (
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0,
  `role_id` int(11) UNSIGNED NOT NULL DEFAULT 0
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user_privileges`
--

INSERT INTO `user_privileges` (`user_id`, `role_id`) VALUES
(1, 1),
(1, 2),
(1, 3),
(2, 1),
(2, 2),
(2, 3),
(3, 1),
(3, 2),
(3, 3),
(4, 1),
(4, 2),
(4, 3),
(5, 1),
(5, 2),
(5, 3),
(6, 1),
(6, 2),
(6, 3),
(7, 1),
(7, 2),
(7, 3),
(8, 1),
(8, 2),
(8, 3),
(9, 3);

-- --------------------------------------------------------

--
-- Table structure for table `user_resources`
--

CREATE TABLE `user_resources` (
  `id` int(11) UNSIGNED NOT NULL,
  `resource_name` varchar(64) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user_resources`
--

INSERT INTO `user_resources` (`id`, `resource_name`) VALUES
(1, 'all'),
(2, 'GetUsers'),
(3, 'GetUserByID'),
(4, 'CreateUser'),
(5, 'UpdateUser'),
(6, 'DeleteUser'),
(7, 'GetRestaurants'),
(8, 'GetRestaurantByID'),
(9, 'CreateRestaurant'),
(10, 'UpdateRestaurant'),
(11, 'DeleteRestaurant'),
(12, 'GetDishes'),
(13, 'GetDishesByRestaurantID'),
(14, 'GetDishByID'),
(15, 'CreateDish'),
(16, 'UpdateDish'),
(17, 'DeleteDish'),
(18, 'GetResources'),
(19, 'GetResourceByID'),
(20, 'CreateResource'),
(21, 'UpdateResource'),
(22, 'DeleteResource'),
(23, 'GetRoles'),
(24, 'GetRoleByID'),
(25, 'CreateRole'),
(26, 'UpdateRole'),
(27, 'DeleteRole');

-- --------------------------------------------------------

--
-- Table structure for table `user_roles`
--

CREATE TABLE `user_roles` (
  `id` int(11) UNSIGNED NOT NULL,
  `role` varchar(64) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user_roles`
--

INSERT INTO `user_roles` (`id`, `role`) VALUES
(1, 'Admin'),
(2, 'Sub-Admin'),
(3, 'User');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dishes`
--
ALTER TABLE `dishes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `admin_id` (`admin_id`),
  ADD KEY `restaurant_id` (`restaurant_id`) USING BTREE,
  ADD KEY `idx_dishes_deleted_at` (`deleted_at`);

--
-- Indexes for table `resources_to_roles`
--
ALTER TABLE `resources_to_roles`
  ADD PRIMARY KEY (`resource_id`,`role_id`) USING BTREE;

--
-- Indexes for table `restaurants`
--
ALTER TABLE `restaurants`
  ADD PRIMARY KEY (`id`),
  ADD KEY `admin_id` (`admin_id`),
  ADD KEY `idx_restaurants_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `admin_id` (`admin_id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- Indexes for table `user_addresses`
--
ALTER TABLE `user_addresses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `user_privileges`
--
ALTER TABLE `user_privileges`
  ADD PRIMARY KEY (`user_id`,`role_id`) USING BTREE;

--
-- Indexes for table `user_resources`
--
ALTER TABLE `user_resources`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_roles`
--
ALTER TABLE `user_roles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `dishes`
--
ALTER TABLE `dishes`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `restaurants`
--
ALTER TABLE `restaurants`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `user_addresses`
--
ALTER TABLE `user_addresses`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `user_resources`
--
ALTER TABLE `user_resources`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT for table `user_roles`
--
ALTER TABLE `user_roles`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
