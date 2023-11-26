CREATE TABLE auth.users
(
    id             bigserial
        PRIMARY KEY,
    username       text NOT NULL
        UNIQUE,
    first_name     text,
    last_name      text,
    password       text,
    email          text NOT NULL
        CONSTRAINT users_pk
            UNIQUE,
    is_active      boolean DEFAULT FALSE,
    is_staff       boolean DEFAULT FALSE,
    is_supper_user boolean DEFAULT FALSE,
    date_joined    timestamp WITH TIME ZONE,
    last_login     timestamp WITH TIME ZONE
);
