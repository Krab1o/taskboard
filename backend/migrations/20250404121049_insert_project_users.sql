-- +goose Up
-- +goose StatementBegin
INSERT INTO user_project_ (id_user_, id_project_) VALUES
    (1, 1),
    (2, 1),
    (3, 1),
    (4, 1),
    (5, 2),
    (6, 2),
    (7, 2),
    (8, 2),
    (9, 3),
    (10, 3),
    (11, 3),
    (12, 3),
    (13, 4),
    (14, 4),
    (15, 4),
    (1, 5),
    (2, 5),
    (3, 5),
    (4, 5),
    (5, 6),
    (6, 6),
    (7, 6),
    (8, 6),
    (9, 7),
    (10, 7),
    (11, 7),
    (12, 7),
    (13, 8),
    (14, 8),
    (15, 8),
    (1, 9),
    (2, 9),
    (3, 9),
    (4, 9),
    (5, 10),
    (6, 10),
    (7, 10),
    (8, 10);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_project_;
-- +goose StatementEnd
