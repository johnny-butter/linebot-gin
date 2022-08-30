CREATE TABLE IF NOT EXISTS foods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE,
    name_eng VARCHAR(256),
    category VARCHAR(10),
    common_names TEXT,
    code VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS food_ingredients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    amount VARCHAR(20),
    food_id INT,
    CONSTRAINT fk_food_ingredients_foods
        FOREIGN KEY(food_id)
        REFERENCES foods(id)
        ON DELETE SET NULL
);

CREATE INDEX foods_name_index ON foods (name);
CREATE INDEX food_ingredients_food_id_index ON food_ingredients (food_id);
