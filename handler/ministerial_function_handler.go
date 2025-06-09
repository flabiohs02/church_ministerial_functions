package handler

import (
	"strconv"

	"church_ministerial_functions/domain"
	"church_ministerial_functions/usecase"

	"github.com/gofiber/fiber/v2"
)

type MinisterialFunctionHandler struct {
	mfService *usecase.MinisterialFunctionService
}

func NewMinisterialFunctionHandler(service *usecase.MinisterialFunctionService) *MinisterialFunctionHandler {
	return &MinisterialFunctionHandler{mfService: service}
}

func (h *MinisterialFunctionHandler) CreateMinisterialFunction(c *fiber.Ctx) error {
	var mf domain.MinisterialFunction
	if err := c.BodyParser(&mf); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.mfService.CreateMinisterialFunction(&mf); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(mf)
}

func (h *MinisterialFunctionHandler) GetMinisterialFunctionByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	mf, err := h.mfService.GetMinisterialFunctionByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ministerial function not found"})
	}

	return c.Status(fiber.StatusOK).JSON(mf)
}

func (h *MinisterialFunctionHandler) GetAllMinisterialFunctions(c *fiber.Ctx) error {
	mfs, err := h.mfService.GetAllMinisterialFunctions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(mfs)
}

func (h *MinisterialFunctionHandler) UpdateMinisterialFunction(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var mf domain.MinisterialFunction
	if err := c.BodyParser(&mf); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	mf.ID = uint(id)
	if err := h.mfService.UpdateMinisterialFunction(&mf); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(mf)
}

func (h *MinisterialFunctionHandler) DeleteMinisterialFunction(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.mfService.DeleteMinisterialFunction(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
