begin;

alter database gokanban reset search_path;

alter default privileges
for role developer
in schema gokanban
revoke all on tables from application;

alter default privileges
for role developer
in schema gokanban
revoke all on sequences from application;

revoke all on schema gokanban from application;
revoke application from gokanban;
revoke developer from david;

drop owned by gokanban;
drop owned by david;
drop owned by application;
drop owned by developer;

drop schema if exists gokanban;

drop role if exists gokanban;
drop role if exists david;
drop role if exists application;
drop role if exists developer;

commit;
