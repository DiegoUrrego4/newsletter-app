package ports

import "github.com/DiegoUrrego4/newsletter-app/internal/domain"

type EmailSender interface {
	Send(newsletter *domain.Newsletter, recipient *domain.Recipient) error
}
