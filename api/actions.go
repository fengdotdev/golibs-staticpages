package api

type ActionKind string

type Action interface {
	GetKind() ActionKind
}
