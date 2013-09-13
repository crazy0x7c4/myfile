drop database if exists myfile;
create database myfile;
use myfile;
create table account (
	a_id int not null auto_increment,
	a_name varchar(32) not null,
	a_password varchar(32) not null,
	primary key(a_id)
);

create table folder (
	f_id int not null auto_increment,
	f_pid int not null,
	f_account_id int not null,
	f_name varchar(32) not null,
	f_create_date bigint,
	primary key(f_id)
);

create table file (
	f_id int not null auto_increment,
	f_folder_id int not null,
	f_account_id int not null,
	f_name varchar(32) not null,
	f_stuffix varchar(8),
	f_size bigint,
	f_create_date bigint,
	f_modify_date bigint,
	primary key(f_id)
);
