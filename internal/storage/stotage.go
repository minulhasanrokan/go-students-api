package storage

import "github.com/minulhasanrokan/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int64) (int64, error)
	GetStudentById(id int64) (types.Student, error)
}
