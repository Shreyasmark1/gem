CREATE TABLE mines (
    mine_id         UUID                        PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id        UUID                        NOT NULL,
    
    mine_name       VARCHAR(255)                NOT NULL,
    
    created_at      TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    created_by      UUID                        NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE,
    updated_by      UUID,

    CONSTRAINT fk_user FOREIGN KEY (owner_id) REFERENCES users(user_id)
);

CREATE TRIGGER update_mine_modtime BEFORE UPDATE ON mines FOR EACH ROW EXECUTE FUNCTION update_timestamp();