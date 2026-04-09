create table if not exists users (
    id serial primary key,
    name text not null,
    email text not null
);

insert into users (name, email)
values
    ('john', 'john@example.com'),
    ('mary', 'mary@example.com')
on conflict do nothing;
