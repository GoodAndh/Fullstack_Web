

create table if not exists users(
Id int not null primary key auto_increment,
Name varchar(255) not null ,
Username varchar(255) not null,
Password varchar(255) not null,
Email varchar(255) not null,
unique (email),
unique (username)
)engine=innodb;