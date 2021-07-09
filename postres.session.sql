create table user_list(
    id varchar,
    pwd varchar,
    name varchar,
    delyn varchar,
    dttm VARCHAR
);

select * from user_list;

insert into user_list
(
    id,
    pwd,
    name,
    delyn,
    dttm
)
values(
    'test',
    '1234',
    'testname',
    'N',
    now()
);