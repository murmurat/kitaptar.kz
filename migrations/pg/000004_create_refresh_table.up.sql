CREATE TABLE refresh_tokens (
     id UUID default uuid_generate_v4() not null
         primary key,
     user_id UUID NOT NULL,
     refresh_token VARCHAR(255)
);