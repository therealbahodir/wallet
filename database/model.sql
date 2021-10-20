create database wallet;

create table users(
	user_id serial not null primary key,
	digest varchar(40) not null,
	created_at timestamp with time zone default current_timestamp,
	is_identified bool not null,
	balance decimal(15,2) default 0
);


create table replenishments (
	replenishment_id serial not null primary key,
	user_id int not null references users(user_id),
	amount decimal(15,2) not null,
	received_at timestamp with time zone default current_timestamp
);