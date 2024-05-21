create table if not exists users_profile(
Id int not null primary key auto_increment,
url_image varchar(255) default "not setting",
Name varchar(255) default "not setting",
Deskripsi text,
userid int not null,
foreign key (userid) references users(id)
)engine=innodb;