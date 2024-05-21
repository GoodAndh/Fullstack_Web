
create table if not exists product(
Id int not null primary key auto_increment,
url_image varchar(255) default "not setting",
name varchar(255) not null default "not setting",
deskripsi text,
Category enum("electric","consumable","etc") default "etc", 
price float not null default 0,
quantity int not null ,
userid int not null,
foreign key (userid) references users(id)
)engine=innodb;
