create table contacts {
    id serial primary key,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    middle_name varchar(255),
    phone_number varchar(255) not null,
};

create table groups (
    id serial primary key,
    name varchar(255) not null
);

create table contact_group (
    contact_id integer references contacts(id),
    group_id integer references groups(id),
    primary key (contact_id, group_id)
);
