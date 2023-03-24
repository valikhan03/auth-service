create or replace table tb_priveleges(
    id number not null,
    operation_code number not null,
    privelege_code varchar2(100),

);


create or replace table tb_dict_services(
    id number not null,
    code number not null UNIQUE,
    p_code varchar2(300),
    p_longname not null
);


create or replace tb_dict_operations(

);