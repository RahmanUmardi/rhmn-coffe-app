package handler

import (
	"fmt"
	"rhmn-coffe/config"
	"rhmn-coffe/middleware"
	"rhmn-coffe/repository"
	"rhmn-coffe/shared/service"
	"rhmn-coffe/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	jwtService service.JwtService
	authUc     usecase.AuthUseCase
	userUc     usecase.UserUsecase
	productUc  usecase.ProductUsecase

	app  *fiber.App
	host string
}

func (s *Server) initRoute() {
	rg := s.app.Group(config.ApiGroup)
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)

	authHandler := NewAuthHandler(s.authUc, rg)
	userHandler := NewUserHandler(s.userUc, authMiddleware, rg)
	productHandlder := NewProductHandler(s.productUc, authMiddleware, rg)

	authHandler.Route()
	userHandler.Route()
	productHandlder.Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.app.Listen(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, because of error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	jwtService := service.NewJwtService(cfg.TokenConfig)
	userUc := usecase.NewUserUsecase(userRepo)
	authUc := usecase.NewAuthUseCase(userUc, jwtService)
	productUc := usecase.NewProductUsecase(productRepo)
	app := fiber.New()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		jwtService: jwtService,
		authUc:     authUc,
		userUc:     userUc,
		productUc:  productUc,

		app:  app,
		host: host,
	}
}
