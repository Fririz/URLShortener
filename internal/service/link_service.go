package service

import (
	"github.com/fririz/URLShortener/domain"
	"github.com/fririz/URLShortener/internal/dto"
)

type LinkRepository interface {
	AddLink(link *domain.Link) error
	GetLinkById(id string) (*domain.Link, error)
}

type LinkService struct {
	currentId int
	lr        LinkRepository
}

func NewLinkService(lr LinkRepository) (*LinkService, error) {
	return &LinkService{lr: lr}, nil
}

func (ls *LinkService) CreateLink(dto.LinkDto) (dto.LinkDto, error) {
	ConvertIdToHex(ls.currentId)

}
