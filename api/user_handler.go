package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zacksfF/Hotel-Reservation-Backend/db"
	"github.com/zacksfF/Hotel-Reservation-Backend/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	store *db.Store
}

func NewUserHandler(store *db.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}

	if errs := params.Validate(); len(errs) > 0 {
		return c.JSON(errs)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.store.User.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.store.User.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrResourceNotFound("User")
		}

		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetBookings(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.store.User.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrResourceNotFound("user")
		}

		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.store.User.DeleteUser(c.Context(), id); err != nil {
		return err
	}

	return c.JSON(map[string]string{"deleted": id})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		id     = c.Params("id")
	)

	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}

	filter := db.Map{"_id": id}
	if err := h.store.User.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}

	return c.JSON(map[string]string{"updated": id})
}
