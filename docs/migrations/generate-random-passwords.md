You can use this to generate a user with a generated random password:

```sql
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- create a user with a random 20 character password
DO $$
DECLARE
    pass TEXT := encode(gen_random_bytes(15), 'base64');
BEGIN
    EXECUTE format('CREATE ROLE david LOGIN PASSWORD %L', pass);
    RAISE NOTICE 'Generated password: %', pass;
END $$;
```

If you execute this in `psql`, you will see the password in the output.

However, if you run this with the `migrate` CLI, you will not see the output.

Therefore, we need a different way.

I did it with a helper script:

```sql
CREATE ROLE david LOGIN PASSWORD ${PG_PASSWORD_DAVID};
GRANT developer TO david;
```

```sh
export PG_PASSWORD_DAVID="$(openssl rand -base64 24)"

echo "export PG_PASSWORD_DAVID=\"$PG_PASSWORD_DAVID\""

envsubst \
  < migrations/000001_create_roles.up.sql \
  > migrations/000001_create_roles.up.sql.temp

mv \
  migrations/000001_create_roles.up.sql.temp \
  migrations/000001_create_roles.up.sql

# after running the migration:
git restore migrations/000001_create_roles.up.sql
```
