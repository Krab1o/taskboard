-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_project_(
    id_user_ INTEGER NOT NULL REFERENCES user_(id_) ON DELETE CASCADE,
    id_project_ INTEGER NOT NULL REFERENCES user_(id_) ON DELETE CASCADE,
    PRIMARY KEY(id_user_, id_project_)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_project_;
-- +goose StatementEnd
