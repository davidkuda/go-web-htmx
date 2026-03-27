-- prerequisite: createdb gokanban

-- developers inherit from group developer.
-- go web app runs with bellevue which inherits from application.
-- migrations should be run with developer.

BEGIN;

create schema gokanban;

ALTER DATABASE gokanban
SET search_path = gokanban, public;

--------------------------------------------------------------------
-- Roles: Groups: Developer ----------------------------------------

create role developer with nologin;

alter schema gokanban owner to developer;

grant create on schema gokanban to developer;

grant select, insert, update, delete 
on all tables in schema gokanban
TO developer;


--------------------------------------------------------------------
-- Roles: Groups: App ----------------------------------------------

create role application with nologin;

grant usage on schema gokanban to application;

grant select, insert, update, delete 
on all tables in schema gokanban
to application;


--------------------------------------------------------------------
-- Default privileges: developer => application: -------------------

-- every time developer creates a new table, application will
-- receive a grant as specified in:
ALTER DEFAULT PRIVILEGES
FOR ROLE developer
IN SCHEMA gokanban
GRANT SELECT, INSERT, UPDATE, DELETE
ON TABLES
TO application;

-- also consider sequences:
-- USAGE:  allows nextval(), currval(), lastval()
-- SELECT: allows currval() and reading the sequence
--         via SELECT directly.
-- UPDATE: allows nextval() and setval() – modifying
--         the sequence’s current value.
ALTER DEFAULT PRIVILEGES
FOR ROLE developer
IN SCHEMA gokanban
GRANT USAGE, SELECT, UPDATE
ON SEQUENCES
TO application;


--------------------------------------------------------------------
-- Roles: Users: (with login) --------------------------------------

CREATE ROLE david LOGIN PASSWORD '${PG_PASSWORD_DAVID}';
GRANT developer TO david;

CREATE ROLE gokanban LOGIN PASSWORD '${PG_PASSWORD_GOKANBAN}';
GRANT application TO gokanban;

COMMIT;
