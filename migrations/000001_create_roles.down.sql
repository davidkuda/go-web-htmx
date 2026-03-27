begin;

drop role gokanban;
drop role david;

revoke all on schema gokanban from application;

alter default privileges
for role developer
in schema gokanban
revoke all on tables from application;

alter default privileges
for role developer
in schema gokanban
revoke all on sequences from application;

drop role application;

alter database gokanban OWNER TO transcribo;

DROP OWNED BY developer;

drop role developer;
drop schema gokanban;

commit;
