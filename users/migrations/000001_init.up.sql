CREATE TABLE IF NOT EXISTS
    users (
        id SERIAL primary key not null unique,
        name varchar(255) not null,
        email varchar(255) not null unique,
        passwordhash varchar(255) not null
    );

CREATE INDEX IF NOT EXISTS idx_email ON users (email);