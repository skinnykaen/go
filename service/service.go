package service

import "github.com/skinnykaen/go.git/repository"

type Authorization interface {

}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}