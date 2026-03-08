package handlers

import (
	"strconv"
	"time"

	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Div Rhino Trivia Time",
		"Subtitle": "Facts for funtimes with friends!",
		"Facts":    facts,
	})
}

func NewFactView(c *fiber.Ctx, view string) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)
	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
	})
}

// HealthAPI returns API health status (for load balancers, k8s probes).
func HealthAPI(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "ok",
		"service":   "trivia-api",
		"version":   "1.0.0",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

// CountFactsAPI returns the total number of facts (for pagination UIs).
func CountFactsAPI(c *fiber.Ctx) error {
	var count int64
	database.DB.Db.Model(&models.Fact{}).Count(&count)
	return c.JSON(fiber.Map{"count": count})
}

// ListFactsAPI returns facts as JSON (for API consumers).
// Query params: limit (optional, default 100), offset (optional, default 0), sort (optional: asc|desc, default desc by createdAt).
func ListFactsAPI(c *fiber.Ctx) error {
	limit := 100
	if v := c.Query("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 1000 {
			limit = n
		}
	}
	offset := 0
	if v := c.Query("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			offset = n
		}
	}
	order := "created_at DESC"
	if v := c.Query("sort"); v == "asc" {
		order = "created_at ASC"
	}
	facts := []models.Fact{}
	database.DB.Db.Order(order).Limit(limit).Offset(offset).Find(&facts)
	return c.JSON(facts)
}

// CreateFactAPI creates a fact from JSON body (for API consumers)
func CreateFactAPI(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)
	return c.Status(fiber.StatusCreated).JSON(fact)
}

// GetFactAPI returns a single fact by ID (for API consumers).
func GetFactAPI(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid fact id",
		})
	}
	var fact models.Fact
	if database.DB.Db.First(&fact, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "fact not found",
		})
	}
	return c.JSON(fact)
}

// UpdateFactAPI updates a fact by ID (for API consumers). Accepts partial updates.
func UpdateFactAPI(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid fact id",
		})
	}
	var fact models.Fact
	if database.DB.Db.First(&fact, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "fact not found",
		})
	}
	var input struct {
		Question *string `json:"question"`
		Answer   *string `json:"answer"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid JSON body",
		})
	}
	if input.Question != nil {
		fact.Question = *input.Question
	}
	if input.Answer != nil {
		fact.Answer = *input.Answer
	}
	database.DB.Db.Save(&fact)
	return c.JSON(fact)
}

// DeleteFactAPI deletes a fact by ID (for API consumers).
func DeleteFactAPI(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid fact id",
		})
	}
	result := database.DB.Db.Delete(&models.Fact{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "fact not found",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
