package main

import (
	"aryavidyananta/Golang-Project/dto"
	"aryavidyananta/Golang-Project/internal/api"
	"aryavidyananta/Golang-Project/internal/config"
	"aryavidyananta/Golang-Project/internal/connection"
	"aryavidyananta/Golang-Project/internal/repository"
	"aryavidyananta/Golang-Project/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/cors"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()
	// Middleware CORS lengkap
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Ganti sesuai port frontend kamu
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("Anda Memerlukan Token"))
		},
	})

	//Repository
	CustomerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)
	BookRepositoy := repository.NewBook(dbConnection)
	BookStockRepository := repository.NewBookStock(dbConnection)
	MediaRepository := repository.NewMedia(dbConnection)
	BlogRepository := repository.NewBlog(dbConnection)
	StaffRepository := repository.NewStaff(dbConnection)
	//Service
	CustomerService := service.NewCustomer(CustomerRepository)
	authService := service.NewAuth(cnf, userRepository)
	BookService := service.NewBook(BookRepositoy, BookStockRepository)
	BookStockService := service.NewBookStock(BookRepositoy, BookStockRepository)
	MediaService := service.NewMedia(cnf, MediaRepository)
	BlogService := service.NewBlog(cnf, BlogRepository)
	StaffService := service.NewStaff(cnf, StaffRepository)

	api.NewCustomer(app, CustomerService, jwtMidd)
	api.NewAuth(app, authService)
	api.NewBook(app, BookService, jwtMidd)
	api.NewMedia(app, cnf, MediaService, jwtMidd)
	api.NewBookStock(app, BookStockService, jwtMidd)
	api.NewBlog(app, BlogService, jwtMidd)
	api.NewStaff(app, StaffService, jwtMidd)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}
