CREATE TABLE extractions (
    extraction_id   UUID                        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID                        NOT NULL,
    gem_id          UUID                        NOT NULL,
    drop_id         UUID                        NOT NULL,
    tract_id        VARCHAR(50),
    extracted_at    TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_prospector FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_gem_extraction FOREIGN KEY (gem_id) REFERENCES gems(gem_id),
    CONSTRAINT fk_drop_extraction FOREIGN KEY (drop_id) REFERENCES drops(drop_id),

    UNIQUE(user_id, gem_id)
);