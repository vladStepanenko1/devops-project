CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null,
    phone_number varchar(10) not null
);

INSERT INTO users (name, email, phone_number) VALUES
    ('keso101', 'keso101@mail.com', '1234567890'),
    ('keso102', 'keso102@mail.com', '1234567890'),
    ('keso103', 'keso103@mail.com', '1234567890'),
    ('keso104', 'keso104@mail.com', '1234567890'),
    ('keso105', 'keso105@mail.com', '1234567890');