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
