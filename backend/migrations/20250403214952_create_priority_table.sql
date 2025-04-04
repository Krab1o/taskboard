-- +goose Up
-- +goose StatementBegin
CREATE TABLE priority_ (
    id_ SERIAL PRIMARY KEY,
    title_ VARCHAR(63)
);

INSERT INTO priority_(id_, title_) VALUES
    (1, 'Low'),
    (2, 'Medium'),
    (3, 'High'),
    (4, 'Critical');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE priority_
-- +goose StatementEnd
