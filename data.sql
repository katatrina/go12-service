DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`
(
    `user_id`    varchar(36) NOT NULL,
    `food_id`    varchar(36) NOT NULL,
    `quantity`   int         NOT NULL,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`, `food_id`),
    KEY          `food_id` (`food_id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`
(
    `id`          varchar(36)  NOT NULL,
    `name`        varchar(100) NOT NULL,
    `description` text,
    `icon`        json DEFAULT NULL,
    `status`      enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `cities`;
CREATE TABLE `cities`
(
    `id`         varchar(36)  NOT NULL,
    `title`      varchar(100) NOT NULL,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `food_likes`;
CREATE TABLE `food_likes`
(
    `user_id`    varchar(36) NOT NULL,
    `food_id`    varchar(36) NOT NULL,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`, `food_id`),
    KEY          `food_id` (`food_id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `food_ratings`;
CREATE TABLE `food_ratings`
(
    `id`         varchar(36) NOT NULL,
    `user_id`    varchar(36) NOT NULL,
    `food_id`    varchar(36) NOT NULL,
    `point`      float DEFAULT '0',
    `comment`    text,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `food_id` (`food_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `foods`;
CREATE TABLE `foods`
(
    `id`            varchar(36)  NOT NULL,
    `restaurant_id` varchar(36)  NOT NULL,
    `category_id`   varchar(36) DEFAULT NULL,
    `name`          varchar(255) NOT NULL,
    `description`   text,
    `price`         float        NOT NULL,
    `images`        json         NOT NULL,
    `status`        enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY             `restaurant_id` (`restaurant_id`) USING BTREE,
    KEY             `category_id` (`category_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `medias`;
CREATE TABLE `medias`
(
    `id`         varchar(36)                                                   NOT NULL,
    `filename`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `cloud_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `size`       int                                                           NOT NULL,
    `ext`        varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `order_details`;
CREATE TABLE `order_details`
(
    `id`          varchar(36) NOT NULL,
    `order_id`    varchar(36) NOT NULL,
    `food_origin` json  DEFAULT NULL,
    `price`       float       NOT NULL,
    `quantity`    int         NOT NULL,
    `discount`    float DEFAULT '0',
    `status`      enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY           `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `order_trackings`;
CREATE TABLE `order_trackings`
(
    `id`         varchar(36) NOT NULL,
    `order_id`   varchar(36) NOT NULL,
    `state`      enum('waiting_for_shipper','preparing','on_the_way','delivered','cancel') NOT NULL,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`
(
    `id`          varchar(36) NOT NULL,
    `user_id`     varchar(36) NOT NULL,
    `total_price` float       NOT NULL,
    `shipper_id`  varchar(36) DEFAULT NULL,
    `status`      enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY           `user_id` (`user_id`) USING BTREE,
    KEY           `shipper_id` (`shipper_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `restaurant_foods`;
CREATE TABLE `restaurant_foods`
(
    `restaurant_id` varchar(36) NOT NULL,
    `food_id`       varchar(36) NOT NULL,
    `status`        enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`restaurant_id`, `food_id`),
    KEY             `food_id` (`food_id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `restaurant_likes`;
CREATE TABLE `restaurant_likes`
(
    `restaurant_id` varchar(36) NOT NULL,
    `user_id`       varchar(36) NOT NULL,
    `created_at`    datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`restaurant_id`, `user_id`),
    KEY             `user_id` (`user_id`)
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `restaurant_ratings`;
CREATE TABLE `restaurant_ratings`
(
    `id`            varchar(36) NOT NULL,
    `user_id`       varchar(36) NOT NULL,
    `restaurant_id` varchar(36) NOT NULL,
    `point`         float       NOT NULL DEFAULT '0',
    `comment`       text,
    `status`        enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY             `user_id` (`user_id`) USING BTREE,
    KEY             `restaurant_id` (`restaurant_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `restaurants`;
CREATE TABLE `restaurants`
(
    `id`          varchar(36)  NOT NULL,
    `owner_id`    varchar(36)  NOT NULL,
    `name`        varchar(50)  NOT NULL,
    `addr`        varchar(255) NOT NULL,
    `city_id`     varchar(36) DEFAULT NULL,
    `category_id` varchar(36) DEFAULT NULL,
    `lat` double DEFAULT NULL,
    `lng` double DEFAULT NULL,
    `cover`       json        DEFAULT NULL,
    `logo`        json        DEFAULT NULL,
    `shipping_fee_per_km` double DEFAULT '0',
    `liked_count` int         DEFAULT '0',
    `status`      enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY           `owner_id` (`owner_id`) USING BTREE,
    KEY           `city_id` (`city_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `user_addresses`;
CREATE TABLE `user_addresses`
(
    `id`         varchar(36)  NOT NULL,
    `user_id`    varchar(36)  NOT NULL,
    `city_id`    varchar(36)  NOT NULL,
    `title`      varchar(100) DEFAULT NULL,
    `icon`       json         DEFAULT NULL,
    `addr`       varchar(255) NOT NULL,
    `lat` double DEFAULT NULL,
    `lng` double DEFAULT NULL,
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `user_id` (`user_id`) USING BTREE,
    KEY          `city_id` (`city_id`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `user_device_tokens`;
CREATE TABLE `user_device_tokens`
(
    `id`            varchar(36) NOT NULL,
    `user_id`       varchar(36)  DEFAULT NULL,
    `is_production` tinyint(1) DEFAULT '0',
    `os`            enum('ios','android','web') DEFAULT 'ios' COMMENT '1: iOS, 2: Android',
    `token`         varchar(255) DEFAULT NULL,
    `status`        enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY             `user_id` (`user_id`) USING BTREE,
    KEY             `os` (`os`) USING BTREE
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`         varchar(36)  NOT NULL,
    `email`      varchar(50)  NOT NULL,
    `avatar`     json        DEFAULT NULL,
    `fb_id`      varchar(50) DEFAULT NULL,
    `gg_id`      varchar(50) DEFAULT NULL,
    `password`   varchar(100) NOT NULL,
    `salt`       varchar(50) DEFAULT NULL,
    `last_name`  varchar(50)  NOT NULL,
    `first_name` varchar(50)  NOT NULL,
    `phone`      varchar(20) DEFAULT NULL,
    `type`       enum('email_password', 'facebook', 'google') NOT NULL DEFAULT 'email_password',
    `role`       enum('user','admin','shipper') NOT NULL DEFAULT 'user',
    `status`     enum('active','inactive','deleted','pending') NOT NULL DEFAULT 'active',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB;
