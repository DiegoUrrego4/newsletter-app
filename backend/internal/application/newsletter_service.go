package application

import (
	"github.com/DiegoUrrego4/newsletter-app/internal/domain"
	"github.com/DiegoUrrego4/newsletter-app/internal/ports"
	"log"
)

type NewsletterService struct {
	NewsletterRepo ports.NewsletterRepository
	RecipientRepo  ports.RecipientRepository
	EmailSender    ports.EmailSender
}

func NewNewsletterService(nRepo ports.NewsletterRepository, rRepo ports.RecipientRepository, sender ports.EmailSender) *NewsletterService {
	return &NewsletterService{
		NewsletterRepo: nRepo,
		RecipientRepo:  rRepo,
		EmailSender:    sender,
	}
}

func (s *NewsletterService) CreateNewsletter(newsletter *domain.Newsletter) error {
	return s.NewsletterRepo.Save(newsletter)
}

func (s *NewsletterService) SendNewsletter(newsletterID string) error {
	newsletter, err := s.NewsletterRepo.FindByID(newsletterID)
	if err != nil {
		return err
	}

	recipients, err := s.RecipientRepo.ListAll()
	if err != nil {
		return err
	}

	for _, recipient := range recipients {
		if recipient.Subscribed {
			err := s.EmailSender.Send(newsletter, recipient)
			if err != nil {
				log.Println("Error sending email to", recipient.Email)
			}
		}
	}

	return nil
}
