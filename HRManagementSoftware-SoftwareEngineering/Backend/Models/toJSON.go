package models

type JSON interface {
	ToJSON() []byte
}
