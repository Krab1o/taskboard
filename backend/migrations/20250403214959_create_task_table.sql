-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_ (
    id_ SERIAL PRIMARY KEY,
    id_project_ INTEGER NOT NULL REFERENCES project_(id_) ON DELETE CASCADE,
    id_status_ INTEGER NOT NULL REFERENCES status_(id_) ON DELETE CASCADE,
    id_priority_ INTEGER NOT NULL REFERENCES priority_(id_) ON DELETE CASCADE,
    title_ VARCHAR(255) NOT NULL,
    description_ TEXT,
    id_creator_ INTEGER NOT NULL REFERENCES user_(id_) ON DELETE CASCADE,
    id_assigned_ INTEGER NOT NULL REFERENCES user_(id_) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_;
-- +goose StatementEnd
