package services

import (
	"api-contact-management/models"
	"api-contact-management/repositories"
	"api-contact-management/requests"

	"github.com/go-playground/validator/v10"
)

// ContactService defines the business logic interface for contact operations.
type ContactService interface {
	// CreateContact creates a new contact based on the provided request.
	CreateContact(req *requests.ContactRequest) (*models.Contact, error)
	// GetAllContacts retrieves all non-deleted contacts.
	GetAllContacts() ([]models.Contact, error)
	// GetContactByID retrieves a single contact by its ID.
	GetContactByID(id uint) (*models.Contact, error)
	// UpdateContact updates an existing contact identified by its ID.
	UpdateContact(id uint, req *requests.ContactRequest) (*models.Contact, error)
	// DeleteContact marks a contact as deleted based on its ID.
	DeleteContact(id uint) error
}

// contactService is the concrete implementation of ContactService.
type contactService struct {
	repository repositories.ContactRepository
	validate   *validator.Validate
}

// NewContactService creates a new instance of ContactService with the provided ContactRepository.
func NewContactService(repository repositories.ContactRepository) ContactService {
	return &contactService{
		repository: repository,
		validate:   validator.New(),
	}
}

// CreateContact creates a new contact based on the provided ContactRequest.
func (s *contactService) CreateContact(req *requests.ContactRequest) (*models.Contact, error) {
	// Validate input
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Map request to Contact model
	contact := models.Contact{
		FullName: req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Message:  req.Message,
	}

	// Persist the contact using the repository
	err := s.repository.Create(&contact)
	return &contact, err
}

// GetAllContacts retrieves all non-deleted contacts from the repository.
func (s *contactService) GetAllContacts() ([]models.Contact, error) {
	return s.repository.FindAll()
}

// GetContactByID retrieves a single contact by its ID.
func (s *contactService) GetContactByID(id uint) (*models.Contact, error) {
	return s.repository.FindByID(id)
}

// UpdateContact updates an existing contact identified by its ID based on the provided ContactRequest.
func (s *contactService) UpdateContact(id uint, req *requests.ContactRequest) (*models.Contact, error) {
	// Validate input
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Retrieve the existing contact
	contact, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update contact fields
	contact.FullName = req.Name
	contact.Email = req.Email
	contact.Phone = req.Phone
	contact.Message = req.Message

	// Persist the updated contact using the repository
	err = s.repository.Update(contact)
	return contact, err
}

// DeleteContact marks a contact as deleted based on its ID.
func (s *contactService) DeleteContact(id uint) error {
	// Retrieve the contact to be deleted
	contact, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	// Mark the contact as deleted
	return s.repository.Delete(contact)
}
