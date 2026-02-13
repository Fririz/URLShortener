package service

import (
	"sync/atomic"
	"time"

	"github.com/fririz/URLShortener/domain"
	"github.com/fririz/URLShortener/internal/dto"
)

type LinkRepository interface {
	AddLink(link *domain.Link) error
	GetLinkById(id int) (*domain.Link, error)
	GetLastId() (int, error)
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

func (ls *LinkService) CreateLink(linkDto dto.LinkDto) (string, error) {

	newId := atomic.AddUint64(&ls.currentId, 1)

	slug := ConvertIdToHex(newId)

	link := &domain.Link{
		ID:        int(newId),
		URL:       linkDto.Url,
		Slug:      slug,
		CreatedAt: time.Now().String(),
	}

	err := ls.lr.AddLink(link)
	if err != nil {
		return "", err
	}

	return slug, nil
}

func (ls *LinkService) GetLinkBySlug(slug string) (string, error) {
	id, err := ConvertHexToId(slug)
	if err != nil {
		return "", err
	}

	link, err := ls.lr.GetLinkById(id)
	if err != nil {
		return "", err
	}
	return link.URL, nil
}
