CREATE TABLE transactions (
                           id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
                           amount text NOT NULL,
                           user_id text NOT NULL REFERENCES users,
                           merchant_id uuid NOT NULL REFERENCES merchants,
                           created_at timestamp NOT NULL DEFAULT now(),
                           updated_at timestamp NOT NULL DEFAULT now(),
                           deleted_at timestamp
);