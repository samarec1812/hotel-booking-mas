CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    email varchar(255) not null
);

CREATE TABLE rooms
(
   id serial not null unique,
   room_name varchar(255) not null,
   description varchar(255),
   hotel_name varchar(255) not null,
   price integer not null,
   accommodates integer not null
);

CREATE TABLE bookings
(
    id      serial                                          not null unique,
    user_id int references users(id) on delete cascade      not null,
    room_id int references rooms(id) on delete cascade not null,
    arrival_date date not null,
    departure_date date not null,
    status varchar(255) not null
);