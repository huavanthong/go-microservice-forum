package models

type iBaseEntity interface {
	BeforeCreate() error
	BeforeUpdate() error
	AfterDelete() error
}
