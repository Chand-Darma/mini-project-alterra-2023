package routes

import (
	"project/constants"
	"project/controllers"
	m "project/middleware"

	"github.com/labstack/echo"
	echomid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddlewares(e)

	//admin route
	e.POST("/admins", controllers.CreateAdminController)
	e.GET("/login", controllers.LoginAdminController)
	e.GET("/admins", controllers.GetAdminsController)
	e.GET("/admins/:id", controllers.GetAdminsIDController)
	e.PUT("/admins/:id", controllers.UpdateAdminController)
	e.DELETE("/admins/:id", controllers.DeleteAdminController)

	// petugas route
	e.POST("/petugas", controllers.CreatePetugasController)
	e.GET("/petugas", controllers.GetPetugasController)
	e.GET("/petugas/:id", controllers.GetPetugasIDController)
	e.PUT("/petugas/:id", controllers.UpdatePetugasIDController)
	e.DELETE("/petugas/:id", controllers.DeletePetugasIDController)

	// anggota route
	e.POST("/anggota", controllers.CreateAnggotaController)
	e.GET("/anggota", controllers.GetAnggotaController)
	e.GET("/anggota/:id", controllers.GetAnggotaIDController)
	e.PUT("/anggota/:id", controllers.UpdateAnggotaIDController)
	e.DELETE("/anggota/:id", controllers.DeleteAnggotaIDController)

	// buku route
	e.POST("/buku", controllers.CreateBukuController)
	e.GET("/buku", controllers.GetBukuController)
	e.GET("/buku/:id", controllers.GetBukuIDController)
	e.PUT("/buku/:id", controllers.UpdateBukuIDController)
	e.DELETE("/buku/:id", controllers.DeleteBukuIDController)

	// peminjam route
	e.POST("/peminjaman", controllers.CreatePeminjamController)
	e.GET("/peminjaman", controllers.GetPeminjamanController)
	e.GET("/peminjaman/:id", controllers.GetPeminjamanIDController)
	e.PUT("/peminjaman/:id", controllers.UpdatePeminjamIDController)
	e.DELETE("/peminjaman/:id", controllers.DeletePeminjamIDController)

	// pengembalian route
	e.POST("/pengembalian", controllers.CreatePengembalianController)
	e.GET("/pengembalian", controllers.GetPengembalianController)
	e.GET("/pengembalian/:id", controllers.GetPengembalianIDController)
	e.PUT("/pengembalian/:id", controllers.UpdatePengembalianIDController)
	e.DELETE("/pengembalian/:id", controllers.DeletePengembalianIDController)

	// denda route
	e.POST("/denda", controllers.CreateDendaController)
	e.GET("/denda", controllers.GetDendaController)
	e.GET("/denda/:id", controllers.GetDendaIDController)
	e.PUT("/denda/:id", controllers.UpdateDendaIDController)
	e.DELETE("/denda/:id", controllers.DeleteDendaIDController)

	//jwt auth admin
	jwtAuth := e.Group("/jwt")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/admins", controllers.GetAdminsController)
	jwtAuth.GET("/login", controllers.LoginAdminController)
	jwtAuth.GET("/admins/:id", controllers.GetAdminsIDController)
	jwtAuth.PUT("/admins/:id", controllers.UpdateAdminController)
	jwtAuth.DELETE("/admins/:id", controllers.DeleteAdminController)

	//jwt auth petugas
	// petugas route
	jwtAuth.POST("/petugas", controllers.CreatePetugasController)
	jwtAuth.GET("/petugas", controllers.GetPetugasController)
	jwtAuth.GET("/petugas/:id", controllers.GetPetugasIDController)
	jwtAuth.PUT("/petugas/:id", controllers.UpdatePetugasIDController)
	jwtAuth.DELETE("/petugas/:id", controllers.DeletePetugasIDController)

	return e
}
