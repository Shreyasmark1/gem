CREATE TABLE auth(
    auth_id     UUID                        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID,
    email_id    TEXT                        NOT NULL,
    token       TEXT                        NOT NULL,
    expires_at  TIMESTAMP WITH TIME ZONE    NOT NULL,
    is_expired  BOOLEAN                     NOT NULL    DEFAULT FALSE,

    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id)
);