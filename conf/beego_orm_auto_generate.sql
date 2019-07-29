create table `user`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `user` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(32) NOT NULL DEFAULT '' ,
        `password_hash` varchar(128) NOT NULL DEFAULT '' ,
        `mobile` varchar(11) NOT NULL DEFAULT ''  UNIQUE,
        `real_name` varchar(32) NOT NULL DEFAULT '' ,
        `id_card` varchar(20) NOT NULL DEFAULT '' ,
        `avatar_url` varchar(256) NOT NULL DEFAULT ''
    ) ENGINE=InnoDB;

create table `house`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.House`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `house` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `user_id` integer NOT NULL,
        `area_id` integer NOT NULL,
        `title` varchar(64) NOT NULL DEFAULT '' ,
        `price` integer NOT NULL DEFAULT 0 ,
        `address` varchar(512) NOT NULL DEFAULT '' ,
        `room_count` integer NOT NULL DEFAULT 1 ,
        `acreage` integer NOT NULL DEFAULT 0 ,
        `unit` varchar(32) NOT NULL DEFAULT '' ,
        `capacity` integer NOT NULL DEFAULT 1 ,
        `beds` varchar(64) NOT NULL DEFAULT '' ,
        `deposit` integer NOT NULL DEFAULT 0 ,
        `min_days` integer NOT NULL DEFAULT 1 ,
        `max_days` integer NOT NULL DEFAULT 0 ,
        `order_count` integer NOT NULL DEFAULT 0 ,
        `index_image_url` varchar(256) NOT NULL DEFAULT '' ,
        `ctime` datetime NOT NULL
    ) ENGINE=InnoDB;

create table `area`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.Area`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `area` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(32) NOT NULL DEFAULT ''
    ) ENGINE=InnoDB;

create table `facility`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.Facility`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `facility` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(32) NOT NULL DEFAULT ''
    ) ENGINE=InnoDB;

create table `house_image`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.HouseImage`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `house_image` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `url` varchar(256) NOT NULL DEFAULT '' ,
        `house_id` integer NOT NULL
    ) ENGINE=InnoDB;

create table `order_house`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.OrderHouse`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `order_house` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `user_id` integer NOT NULL,
        `house_id` integer NOT NULL,
        `begin_date` datetime NOT NULL,
        `end_date` datetime NOT NULL,
        `days` integer NOT NULL DEFAULT 0 ,
        `house_price` integer NOT NULL DEFAULT 0 ,
        `amount` integer NOT NULL DEFAULT 0 ,
        `status` varchar(255) NOT NULL DEFAULT 'WAIT_ACCEPT' ,
        `comment` varchar(512) NOT NULL DEFAULT '' ,
        `ctime` datetime NOT NULL,
        `credit` bool NOT NULL DEFAULT FALSE
    ) ENGINE=InnoDB;

create table `facility_houses`
    -- --------------------------------------------------
    --  Table Structure for `go-microservice/models.FacilityHouses`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `facility_houses` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `facility_id` integer NOT NULL,
        `house_id` integer NOT NULL
    ) ENGINE=InnoDB;
