-- Create user database
create database coords with encoding 'UTF8' template template0;;
grant all privileges on database coords to postgres;

-- Connect to user database
\c coords

-- Create postgis extension
create EXTENSION postgis;

-- Create table in user database
create table if not exists city_records(
	id serial primary key not null,
	title varchar(50),
	coords GEOGRAPHY(Point)
);

-- Insert test data in user table
insert into city_records (title, coords)
values('City_00001', 'POINT(135.345 52.243)');

insert into city_records (title, coords)
values('City_00002', 'POINT(-111.322 2.75)');

insert into city_records (title, coords)
values('City_00003', 'POINT(27.346 32.186)');


insert into city_records (title, coords)
values('City_00103', 'POINT(127.246 -32.186)');

insert into city_records (title, coords)
values('City_00203', 'POINT(64.346 -12.181)');

insert into city_records (title, coords)
values('City_00303', 'POINT(154.12 1.74)');

insert into city_records (title, coords)
values('City_00403', 'POINT(83.217 74.66)');

insert into city_records (title, coords)
values('City_00503', 'POINT(57.16 -2.73)');

insert into city_records (title, coords)
values('City_00603', 'POINT(62.12 42.11)');

insert into city_records (title, coords)
values('City_00703', 'POINT(-23.346 -62.2)');

insert into city_records (title, coords)
values('City_00803', 'POINT(72.5632 -12.146)');

insert into city_records (title, coords)
values('City_00903', 'POINT(52.341 53.116)');

insert into city_records (title, coords)
values('City_01003', 'POINT(82.11 53.83)');

insert into city_records (title, coords)
values('City_01103', 'POINT(-127.532 -32.12)');

insert into city_records (title, coords)
values('City_01203', 'POINT(42.32 -12.184)');

insert into city_records (title, coords)
values('City_01303', 'POINT(35.461 84.41)');

insert into city_records (title, coords)
values('City_01403', 'POINT(-23.76 -84.12)');

insert into city_records (title, coords)
values('City_01503', 'POINT(167.942 34.84)');

insert into city_records (title, coords)
values('City_01603', 'POINT(25.24 3.194)');

insert into city_records (title, coords)
values('City_01703', 'POINT(2.42 0.133)');

insert into city_records (title, coords)
values('City_01803', 'POINT(0.3342 25.681)');