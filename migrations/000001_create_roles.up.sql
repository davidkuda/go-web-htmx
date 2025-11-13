-- prerequisite: createdb davidkuda

BEGIN;

create schema gokanban;

--------------------------------------------------------------------
-- Roles: Groups: Developer ----------------------------------------
-- A developer can CREATE ON SCHEMA, an app can only USAGE. --------

create role developer with nologin;

grant create on schema gokanban to developer;

grant select, insert, update, delete 
on all tables in schema gokanban
TO developer;


--------------------------------------------------------------------
-- Roles: Groups: App ----------------------------------------------

create role app with nologin;

grant usage on schema gokanban to app;

grant select, insert, update, delete 
on all tables in schema gokanban
to app;


--------------------------------------------------------------------
-- Roles: Users: (with login) --------------------------------------

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- create a user with a random 20 character password
DO $$
DECLARE
    pass TEXT := encode(gen_random_bytes(15), 'base64');
BEGIN
    EXECUTE format('CREATE ROLE david LOGIN PASSWORD %L', pass);
    RAISE NOTICE 'Generated password: %', pass;
END $$;

GRANT developer TO david;

DO $$
DECLARE
    pass TEXT := encode(gen_random_bytes(15), 'base64');
BEGIN
    EXECUTE format('CREATE ROLE gokanban LOGIN PASSWORD %L', pass);
    RAISE NOTICE 'Generated password: %', pass;
END $$;

GRANT app TO gokanban;


ALTER SCHEMA gokanban OWNER TO developer;
