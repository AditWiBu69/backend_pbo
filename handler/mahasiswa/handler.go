package mahasiswa

import (
	"github.com/AditWiBu69/pakage_be/model"
	"github.com/AditWiBu69/pakage_be/repository/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertDataPresensi godoc
// @Summary Insert data.
// @Description Input data.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param request body DataMahasiswa true "Payload Body [RAW]"
// @Success 200 {object} DataMahasiswa
// @Failure 400
// @Failure 500
// @Router /api/mahasiswa [post]
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

// GetAllDataMahasiswa godoc
// @Summary Get All Data.
// @Description Mengambil semua data mahasiswa
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Success 200 {object} DataMahasiswa
// @Router /api/mahasiswa [get]
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

// GetDataMahasiswaByUname godoc
// @Summary Get By ID Data.
// @Description Ambil per ID data.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan Nama"
// @Success 200 {object} DataMahasiswa
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/mahasiswa/{id} [get]
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

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body DataMahasiswa true "Payload Body [RAW]"
// @Success 200 {object} DataMahasiswa
// @Failure 400
// @Failure 500
// @Router /api/mahasiswa/{id} [put]
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
