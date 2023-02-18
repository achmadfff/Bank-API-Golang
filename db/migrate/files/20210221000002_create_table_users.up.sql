-- Table Definition ----------------------------------------------
CREATE TABLE users (
                       id text PRIMARY KEY,
                       role_id text NOT NULL REFERENCES roles,
                       name text NOT NULL,
                       position text,
                       email text,
                       phone text,
                       encrypted_password text NOT NULL,
                       token text,
                       created_at timestamp NOT NULL DEFAULT now(),
                       updated_at timestamp NOT NULL DEFAULT now(),
                       deleted_at timestamp
);

-- Indices -------------------------------------------------------
CREATE UNIQUE INDEX email_index ON users (email text_ops);
CREATE INDEX user_role_index ON users (role_id text_ops);