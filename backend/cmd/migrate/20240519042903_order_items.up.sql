
create table if not exists order_items(
ID int not null primary key auto_increment,
Status enum ("pending","cancelled","succes") default "pending",
quantity int not null,
totalPrice float not null,
ProductID int not null,
OrderID int not null,
UserID int not null,
foreign key (UserID) references users(id),
foreign key (OrderID) references orders(id),
foreign key (productID)references product(id)
)engine=innodb;
