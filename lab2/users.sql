create database app;
create user app;
alter user app with password 'pass';
grant all on database app to app;
\connect app;
create table if not exists users(
  id serial primary key,
  name varchar(50),
  email varchar(100)
);
grant all on all tables in schema public to app;
grant all on all sequences in schema public to app;
