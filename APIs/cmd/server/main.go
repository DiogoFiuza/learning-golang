package main

import (
	"github.com/DiogoFiuza/learning-golang/APIs/configs"
	_ "github.com/DiogoFiuza/learning-golang/APIs/docs"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/entity"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/infra/database"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title           Golang Course
// @version         1.0
// @description     This is a project to learn Golang
// @termsOfService  http://swagger.io/terms/

// @contact.name   Diogo Fiuza
// @contact.email  difipalm@gmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.basic  ApiKeyAuth
// @in header
// @name Autorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)
	}
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter() // https://github.com/go-chi/chi/
	r.Use(middleware.Logger)
	r.Use(LogRequest)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("JWTExpireIn", config.JWTExpiration))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/doc.json"), //The url pointing to API definition
	))

	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
