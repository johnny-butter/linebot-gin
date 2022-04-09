CREATE TABLE IF NOT EXISTS district (
    id SERIAL PRIMARY KEY,
    name VARCHAR(12),
    county_id INT,
    CONSTRAINT fk_district_county
        FOREIGN KEY(county_id)
        REFERENCES county(id)
        ON DELETE SET NULL,
    UNIQUE(county_id, name)
);

CREATE INDEX district_name_index ON district (name);
