package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ProjectRepository struct {
	db *pgxpool.Pool
}

type UserRepository struct {
	db *pgxpool.Pool
}

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{db: db}
}
