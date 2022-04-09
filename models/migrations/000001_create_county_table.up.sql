CREATE TABLE IF NOT EXISTS county (
    id SERIAL PRIMARY KEY,
    name VARCHAR(12),
    cwb_id VARCHAR(12),
    UNIQUE(name, cwb_id)
);

CREATE INDEX county_name_index ON county (name);
