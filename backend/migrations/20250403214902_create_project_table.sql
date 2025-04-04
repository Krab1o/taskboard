-- +goose Up
-- +goose StatementBegin
CREATE TABLE project_ (
    id_ SERIAL PRIMARY KEY,
    title_ VARCHAR(255) NOT NULL
);

INSERT INTO project_(title_) VALUES
    ('Система онлайн-бронирования'),
    ('Умный дом: управление освещением'),
    ('Приложение для изучения языков'),
    ('Анализ данных продаж'),
    ('Веб-платформа для фрилансеров'),
    ('Мобильный трекер привычек'),
    ('Бот для Telegram с напоминаниями'),
    ('Сервис доставки еды'),
    ('Облачное хранилище документов'),
    ('Игра-головоломка на Unity');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE project_;
-- +goose StatementEnd
