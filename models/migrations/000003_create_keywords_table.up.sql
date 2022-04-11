CREATE TABLE IF NOT EXISTS keywords (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE,
    method VARCHAR(30),
    description TEXT,
    usage TEXT
);

CREATE INDEX keywords_name_index ON keywords (name);
