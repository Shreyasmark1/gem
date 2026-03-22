CREATE TABLE gems (
    gem_id      UUID                        PRIMARY KEY DEFAULT gen_random_uuid(),
    mine_id     UUID                        NOT NULL,
    title       VARCHAR(255)                NOT NULL,
    file_url    TEXT                        NOT NULL,
    teaser_text TEXT                        NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    created_by  UUID                        NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE,
    updated_by  UUID,

    CONSTRAINT fk_mine FOREIGN KEY (mine_id) REFERENCES mines(mine_id)
);

CREATE TRIGGER update_gem_modtime BEFORE UPDATE ON gems FOR EACH ROW EXECUTE FUNCTION update_timestamp();