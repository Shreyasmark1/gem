CREATE TABLE drops (
    -- Tracking & Identity
    drop_id         UUID                        PRIMARY KEY     DEFAULT gen_random_uuid(),
    gem_id          UUID                        NOT NULL,
    drop_code       VARCHAR(50)                 NOT NULL, 
    label           VARCHAR(50)                 NOT NULL,

    -- Constraints
    max_claims      INTEGER                     NOT NULL,              
    current_claims  INTEGER                     NOT NULL        DEFAULT 0,
    expires_at      TIMESTAMP WITH TIME ZONE,
    access_level    VARCHAR(50)                 NOT NULL        DEFAULT 'public',

    created_at      TIMESTAMP WITH TIME ZONE    NOT NULL        DEFAULT CURRENT_TIMESTAMP,
    created_by      UUID                        NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE,
    updated_by      UUID,

    CONSTRAINT fk_gem FOREIGN KEY (gem_id) REFERENCES gems(gem_id),
    CONSTRAINT chk_access_level CHECK (access_level IN ('public', 'subscriber'))
);

CREATE TRIGGER update_drop_modtime BEFORE UPDATE ON drops FOR EACH ROW EXECUTE FUNCTION update_timestamp();