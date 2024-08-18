package users

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gofiber-templ-bench/pkg/models"
	"gofiber-templ-bench/utils"
	"gofiber-templ-bench/views"
	"gofiber-templ-bench/views/forms"
	"gorm.io/gorm"
	"log/slog"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) Index(c *fiber.Ctx) error {
	vm := views.IndexViewModel{
		EditLink: func(id uint) string {
			return fmt.Sprintf("/user/%d/edit", id)
		},
	}
	err := h.DB.Model(models.User{}).Preload("Address").Find(&vm.Users).Error
	if err != nil {
		return err
	}
	return utils.Render(c, views.IndexView(vm))
}

// RenderForm renders the form for the given user after preloading related data
func (h *Handler) RenderForm(c *fiber.Ctx, user *models.User) error {
	// enhance user with preload of address
	// TRY 1: this does not work and overwrites data
	// if err := db.Preload("Address").First(&user, c.Params("id")).Error; err != nil {
	//	 return err
	// }

	// TRY 2: this works but is not very elegant / scalable
	if user.Address.ID == 0 {
		h.DB.First(&user.Address, user.AddressID)
	}
	c.Set("Content-Type", "text/html")
	return utils.Render(c, forms.Form(user))
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("need valid id")
	}
	var user models.User
	h.DB.First(&user, id)
	return h.RenderForm(c, &user)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("need valid id")
	}

	// Get the original user
	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		return err
	}

	// Load form into model
	if err := c.BodyParser(&user); err != nil {
		slog.Error("Error parsing form", "err", err)
		return err
	}

	// Validation
	if len(user.Validate()) > 0 {
		c.Status(fiber.StatusUnprocessableEntity)
		return h.RenderForm(c, &user)
	}

	// Save Operation
	if err := h.DB.Save(&user).Error; err != nil {
		slog.Error("Error saving user", "err", err)
		return h.RenderForm(c, &user)
	}

	return c.Redirect("/")
}
