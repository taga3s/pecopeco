-- コンテナ初期起動時に流すスクリプトを記述する
drop database if exists peco_db;
create database peco_db;
use peco_db;

-- レストランテーブル
drop table if exists restaurants;
create table restaurants (
  id varchar(256) not null primary key,
  name varchar(256) not null,
  genre varchar(256) not null,
  nearest_station varchar(256) not null,
  address varchar(256) not null,
  url varchar(256) not null,
  posted_by varchar(256) not null,
  created_at datetime not null default now()
) charset=utf8;
