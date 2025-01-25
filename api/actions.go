package api

type ActionType string

type Action interface {
	GetType() ActionType
}
