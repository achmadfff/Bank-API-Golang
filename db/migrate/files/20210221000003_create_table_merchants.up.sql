CREATE TABLE merchants (
                       id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
                       name text NOT NULL,
                       phone text,
                       address text,
                       balance bigint,
                       created_at timestamp NOT NULL DEFAULT now(),
                       updated_at timestamp NOT NULL DEFAULT now(),
                       deleted_at timestamp
);

CREATE UNIQUE INDEX id_merchant_index ON merchants (id);