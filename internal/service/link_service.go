package service

import (
	"sync/atomic"
	"time"

	"github.com/fririz/URLShortener/domain"
	"github.com/fririz/URLShortener/internal/dto"
)

type LinkRepository interface {
	AddLink(link domain.Link) error
	GetLinkById(id int) (*domain.Link, error)
	GetLastId() (uint64, error)
}

type LinkService struct {
	currentId uint64
	lr        LinkRepository
}

func NewLinkService(lr LinkRepository) (*LinkService, error) {
	lastId, err := lr.GetLastId()
	if err != nil {
		return nil, err
	}
	return &LinkService{
		currentId: uint64(lastId),
		lr:        lr,
	}, nil
}

func (ls *LinkService) CreateLink(linkDto dto.LinkDto) (dto.LinkDto, error) {
	newId := atomic.AddUint64(&ls.currentId, 1)
	slug := ConvertIdToHex(uint64(newId))
	link := domain.Link{ID: int(newId), URL: linkDto.Url, Slug: slug, CreatedAt: time.Now().String()}
	err := ls.lr.AddLink(link)
	linkResponse := dto.LinkDto{Url: "localhost:8080/" + slug}
	if err != nil {
		return linkResponse, err
	}
	return linkResponse, nil
}
