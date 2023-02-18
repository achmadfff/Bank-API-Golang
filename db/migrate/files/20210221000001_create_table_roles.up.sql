-- Table Definition ----------------------------------------------

CREATE TABLE roles (
                       id text PRIMARY KEY,
                       name text UNIQUE NOT NULL,
                       created_at timestamp NOT NULL DEFAULT now(),
                       updated_at timestamp NOT NULL DEFAULT now(),
                       deleted_at timestamp
);

-- Indices -------------------------------------------------------