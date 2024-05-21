
create table if not exists product_stat(
Id int not null primary key auto_increment,
createdAt timestamp default current_timestamp,
lastUpdated timestamp default current_timestamp on update current_timestamp,
productId int not null,
foreign key (productId)references product(id)
)engine=innodb;