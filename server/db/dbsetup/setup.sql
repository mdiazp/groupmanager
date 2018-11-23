drop table group_user cascade;
drop table group_admin cascade;
drop table gmuser cascade;

create table gmuser(
    id          serial primary key,
    username    varchar(100) not null unique,
    rol         varchar(100) not null
);

create table group_admin(
    id          serial primary key,
    adgroup     varchar(100) not null,
    id_gmuser   integer not null references gmuser(id) ON DELETE CASCADE,
    
    CONSTRAINT adgroup_gmuser UNIQUE (adgroup, id_gmuser)
);

create table group_user(
    id          serial primary key,
    adgroup     varchar(100) not null,
    aduser      varchar(100) not null,

    CONSTRAINT adgroup_aduser UNIQUE (adgroup, aduser)
);
