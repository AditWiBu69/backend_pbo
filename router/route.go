package router

import (
	"api/handler/mahasiswa"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/mahasiswa", mahasiswa.TambahMahasiswa)
	api.Get("/mahasiswa", mahasiswa.GetAllDataMahasiswa)
	api.Get("/mahasiswa/:namalengkap", mahasiswa.GetDataMahasiswaByUname)
	api.Delete("/mahasiswa/:id", mahasiswa.DeleteMahasiswa)
	api.Put("/mahasiswa/:id", mahasiswa.UpdateMahasiswa)

}
