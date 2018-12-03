drop table group_aduser cascade;
drop table group_admin cascade;
drop table gmgroup cascade;
drop table gmuser cascade;

create table gmuser(
    id          serial primary key,
    username    varchar(100) not null unique,
    rol         varchar(100) not null
);

create table gmgroup(
    id          serial PRIMARY key,
    adgroup   varchar(100) not null unique
);

create table group_admin(
    id          serial primary key,
    id_gmgroup  integer not null references gmgroup(id) ON DELETE CASCADE,
    id_gmuser   integer not null references gmuser(id) ON DELETE CASCADE,
    
    CONSTRAINT adgroup_gmuser UNIQUE (id_gmgroup, id_gmuser)
);

create table group_aduser(
    id          serial primary key,
    id_gmgroup  integer not null references gmgroup(id) ON DELETE CASCADE,
    aduser      varchar(100) not null,
    aduser_full_name varchar(100) not null,

    CONSTRAINT gmgroup_aduser UNIQUE (id_gmgroup, aduser)
);
