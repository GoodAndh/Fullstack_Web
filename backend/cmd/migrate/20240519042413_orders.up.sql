

create table if not exists orders(
ID int not null primary key auto_increment,
lastUpdated timestamp default current_timestamp on update current_timestamp,
UserID int not null,
foreign key (UserID) references users(id)
)engine=innodb;