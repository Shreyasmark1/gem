CREATE TABLE users(
  user_id       UUID                      PRIMARY KEY    DEFAULT gen_random_uuid(),
  
  full_name     TEXT                      NOT NULL,
  email_id      TEXT                      NOT NULL,
  password_hash TEXT                      NOT NULL,
  
  created_at    TIMESTAMP WITH TIME ZONE  NOT NULL       DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP WITH TIME ZONE
);

CREATE TRIGGER update_user_modtime BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_timestamp();