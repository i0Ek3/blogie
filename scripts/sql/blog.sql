-- create database first
CREATE
    DATABASE
    IF
    NOT EXISTS blogie DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

-- use database
use blogie;

-- create tag table
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT 'tag name',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'created on',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modified on',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deleted on',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'is deleted? 0:undelete 1:deleted',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'state 0:disable 1:enable',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='tag management';

-- create article table
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT 'article title',
  `desc` varchar(255) DEFAULT '' COMMENT 'article desc',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT 'cover image url',
  `content` longtext COMMENT 'article content',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'created on',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modified on',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deleted on',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'is deleted? 0:undelete 1:deleted',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'state 0:disable 1:enable',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article management';

-- create article_tag table
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT 'article ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'tag ID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'created on',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modified on',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deleted on',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'is deleted? 0:undelete 1:deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='associate article and tag';


-- create auth table
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'created on',
  `created_by` varchar(100) DEFAULT '' COMMENT 'creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'modified on',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'deleted on',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'is deleted? 0:undelete 1:deleted',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='auth management';

INSERT INTO `blogie`.`blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES (1, 'i0Ek3', 'blogie', 0, 'i0Ek3', 0, '', 0, 0);