package data_contracts

type Action string

const (
	CREATE Action = "CREATE"
	UPDATE Action = "UPDATE"
	DELETE Action = "DELETE"
)

type CommandType string

const (
	CREATE_USER_DTO CommandType = "CreateUserDTO"
	UPDATE_USER_DTO CommandType = "UpdateUserDTO"
	DELETE_USER_DTO CommandType = "DeleteUserDTO"
)

type Command interface {
	GetAction() Action
	GetCommandType() CommandType
}
