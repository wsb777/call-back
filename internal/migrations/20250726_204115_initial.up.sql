CREATE TABLE system_roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    system_role_id INTEGER NOT NULL REFERENCES system_roles(id) ON DELETE RESTRICT
);

CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    members INTEGER NOT NULL REFERENCES system_roles(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    owners INTEGER NOT NULL REFERENCES system_roles(id) ON DELETE RESTRICT
);

INSERT INTO system_roles (id, name) VALUES (1, 'admin'), (2, 'common_user') ON CONFLICT (id) DO NOTHING;