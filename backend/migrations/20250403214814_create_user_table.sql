-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_(
    id_ SERIAL PRIMARY KEY,
    given_name_ VARCHAR(127) NOT NULL,
    surname_ VARCHAR(127) NOT NULL
);

INSERT INTO user_(given_name_, surname_) VALUES
    ('Иван', 'Иванов'),
    ('Мария', 'Петрова'),
    ('Алексей', 'Сидоров'),
    ('Елена', 'Кузнецова'),
    ('Дмитрий', 'Смирнов'),
    ('Ольга', 'Попова'),
    ('Николай', 'Семенов'),
    ('Татьяна', 'Волкова'),
    ('Сергей', 'Морозов'),
    ('Анна', 'Лебедева'),
    ('Павел', 'Новиков'),
    ('Юлия', 'Федорова'),
    ('Андрей', 'Сорокин'),
    ('Наталья', 'Егорова'),
    ('Владимир', 'Зайцев');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_;
-- +goose StatementEnd
