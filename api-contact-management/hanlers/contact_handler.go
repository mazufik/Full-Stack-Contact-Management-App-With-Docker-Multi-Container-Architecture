package handlers

import (
	"api-contact-management/requests"
	"api-contact-management/responses"
	"api-contact-management/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ContactHandler handles HTTP requests related to contact operations.
type ContactHandler struct {
	service services.ContactService
}

// NewContactHandler creates a new instance of ContactHandler with the provided ContactService.
func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{service}
}

// CreateContact handles the creation of a new contact.
func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req requests.ContactRequest

	// Bind the JSON payload to the ContactRequest struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Use the service layer to create a new contact.
	contact, err := h.service.CreateContact(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Respond with the created contact and a success message.
	c.JSON(http.StatusCreated, responses.APIResponse{
		Code:    "CREATED",
		Message: "Contact created successfully",
		Data:    responses.ContactResponseFromModel(contact),
	})
}

// GetContacts retrieves all contacts.
func (h *ContactHandler) GetContacts(c *gin.Context) {
	// Fetch all contacts using the service layer.
	contacts, err := h.service.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Convert the contact models to response formats.
	var contactResponses []responses.ContactResponse
	for _, contact := range contacts {
		contactResponses = append(contactResponses, responses.ContactResponseFromModel(&contact))
	}

	// Respond with the list of contacts.
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contacts retrieved successfully",
		Data:    contactResponses,
	})
}

// GetContact retrieves a single contact by its ID.
func (h *ContactHandler) GetContact(c *gin.Context) {
	// Retrieve the 'id' parameter from the URL.
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	// Fetch the contact by ID using the service layer.
	contact, err := h.service.GetContactByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.APIResponse{
			Code:    "NOT_FOUND",
			Message: "Contact not found",
			Data:    nil,
		})
		return
	}

	// Respond with the contact details.
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contact retrieved successfully",
		Data:    responses.ContactResponseFromModel(contact),
	})
}

// UpdateContact updates an existing contact by its ID.
func (h *ContactHandler) UpdateContact(c *gin.Context) {
	// Retrieve the 'id' parameter from the URL.
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	var req requests.ContactRequest

	// Bind the JSON payload to the ContactRequest struct.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Use the service layer to update the contact.
	contact, err := h.service.UpdateContact(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Respond with the updated contact and a success message.
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contact updated successfully",
		Data:    responses.ContactResponseFromModel(contact),
	})
}

// DeleteContact removes a contact by its ID.
func (h *ContactHandler) DeleteContact(c *gin.Context) {
	// Retrieve the 'id' parameter from the URL.
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.APIResponse{
			Code:    "BAD_REQUEST",
			Message: "Invalid ID",
			Data:    nil,
		})
		return
	}

	// Use the service layer to delete the contact.
	err = h.service.DeleteContact(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.APIResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Respond with a success message.
	c.JSON(http.StatusOK, responses.APIResponse{
		Code:    "SUCCESS",
		Message: "Contact deleted successfully",
		Data:    nil,
	})
}
