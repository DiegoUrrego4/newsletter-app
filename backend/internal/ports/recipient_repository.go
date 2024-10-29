package ports

import "github.com/DiegoUrrego4/newsletter-app/internal/domain"

type RecipientRepository interface {
	Add(recipient *domain.Recipient) error
	Remove(email string) error
	FindByEmail(email string) (*domain.Recipient, error)
	ListAll() ([]*domain.Recipient, error)
}
