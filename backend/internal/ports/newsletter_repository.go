package ports

import "github.com/DiegoUrrego4/newsletter-app/internal/domain"

type NewsletterRepository interface {
	Save(newsletter *domain.Newsletter) error
	FindByID(id string) (*domain.Newsletter, error)
}
