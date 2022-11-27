CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4 (),
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    created_at DATE NOT NULL,
    created_by uuid NOT NULL,
    updated_at DATE,
    updated_by uuid,
    deleted_at DATE,
    deleted_by uuid,
    PRIMARY KEY (id)
);