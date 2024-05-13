-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    brand VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
-- +goose StatementEnd
