package mahasiswa

import (
	"api/model"
	"api/repository/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TambahMahasiswa(c *fiber.Ctx) error {
	var requestData model.DataMahasiswa
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	requestData.ID = uuid.New().String()

	if err := db.CountDataMahasiswa(requestData.NamaLengkap); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count data in the database",
		})
	} else if err != nil && err.Error() == "data already exists" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data already exists in the database",
		})
	}

	// If data doesn't exist, proceed with insertion
	if err := db.InsertDataMahasiswa(requestData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data into the database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
		"data":    requestData,
	})
}

func GetAllDataMahasiswa(c *fiber.Ctx) error {
	filter := bson.M{}

	requestData, err := db.GetDataMahasiswaFilter(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}

func GetDataMahasiswaByUname(c *fiber.Ctx) error {
	namalengkap := c.Params("namalengkap")
	filter := bson.M{"namalengkap": namalengkap}
	requestData, err := db.GetOneDataMahasiswaFilter(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}
func DeleteMahasiswa(c *fiber.Ctx) error {
	id := c.Params("id")
	filter := bson.M{"id": id}

	requestData, err := db.DeleteMahasiswa(filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.NewError(fiber.StatusNotFound, "No document found")
		}
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data deleted successfully",
		"data":    requestData,
	})
}
func UpdateMahasiswa(c *fiber.Ctx) error {
	id := c.Params("id")

	var requestData model.DataMahasiswa
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	filter := bson.M{"id": id}
	if err := db.UpdateMahasiswa(filter, requestData); err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.NewError(fiber.StatusNotFound, "No document found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data updated successfully",
		"data":    requestData,
	})
}
