package repositories

import (
	"api-contact-management/models"
	"time"

	"gorm.io/gorm"
)

// ContactRepository defines the interface for contact data operations.
type ContactRepository interface {
	// Create adds a new contact to the database.
	Create(contact *models.Contact) error
	// FindAll retrieves all non-deleted contacts from the database.
	FindAll() ([]models.Contact, error)
	// FindByID retrieves a contact by its ID, ensuring it is not deleted.
	FindByID(id uint) (*models.Contact, error)
	// Update modifies an existing contact in the database.
	Update(contact *models.Contact) error
	// Delete marks a contact as deleted in the database.
	Delete(contact *models.Contact) error
}

// contactRepository is the GORM-based implementation of ContactRepository.
type contactRepository struct {
	db *gorm.DB
}

// NewContactRepository creates a new instance of ContactRepository with the provided GORM DB.
func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db}
}

// Create adds a new contact to the database.
func (r *contactRepository) Create(contact *models.Contact) error {
	return r.db.Create(contact).Error
}

// FindAll retrieves all non-deleted contacts from the database.
func (r *contactRepository) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.db.Where("deleted_at = ?", "0000-00-00 00:00:00").Find(&contacts).Error
	return contacts, err
}

// FindByID retrieves a contact by its ID, ensuring it is not deleted.
func (r *contactRepository) FindByID(id uint) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.Where("id = ? AND deleted_at = ?", id, "0000-00-00 00:00:00").First(&contact).Error
	return &contact, err
}

// Update modifies an existing contact in the database.
func (r *contactRepository) Update(contact *models.Contact) error {
	return r.db.Save(contact).Error
}

// Delete marks a contact as deleted in the database by setting the DeletedAt field.
func (r *contactRepository) Delete(contact *models.Contact) error {
	contact.DeletedAt = time.Now()
	return r.db.Save(contact).Error
}
