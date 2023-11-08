package store

import (
	tableemploye "ipr-savelichev/internal/store/tables/employe"
	tabletask "ipr-savelichev/internal/store/tables/task"
	tabletool "ipr-savelichev/internal/store/tables/tool"
	tableuser "ipr-savelichev/internal/store/tables/user"
)

type Store interface {
	Task() tabletask.Task
	Tool() tabletool.Tool
	Employe() tableemploye.Employe
	User() tableuser.User
}
