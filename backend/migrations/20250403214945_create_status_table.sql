-- +goose Up
-- +goose StatementBegin
CREATE TABLE status_ (
    id_ SERIAL PRIMARY KEY,
    title_ VARCHAR(63)
);

INSERT INTO status_ (id_, title_) VALUES
    (1, 'Backlog'),
    (2, 'Todo'),
    (3, 'In Progress'),
    (4, 'Done');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE status_
-- +goose StatementEnd
