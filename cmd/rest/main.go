package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	authHandler "github.com/bgildson/t10-challenge/api/rest/auth/handler"
	authSerializer "github.com/bgildson/t10-challenge/api/rest/auth/serializer"
	authSerializerJson "github.com/bgildson/t10-challenge/api/rest/auth/serializer/json"
	utilSerializer "github.com/bgildson/t10-challenge/api/rest/util/serializer"
	utilSerializerJson "github.com/bgildson/t10-challenge/api/rest/util/serializer/json"
	authRepositoryJwt "github.com/bgildson/t10-challenge/pkg/auth/repository/jwt"
	authRepositoryPostgres "github.com/bgildson/t10-challenge/pkg/auth/repository/postgres"
	authService "github.com/bgildson/t10-challenge/pkg/auth/service"
)

func main() {
	// refresh env variables
	godotenv.Load()

	// load env configurations
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtExpiresWith := os.Getenv("JWT_EXPIRES_WITH")
	jwtExpiresWithSeconds, err := strconv.Atoi(jwtExpiresWith)
	if err != nil {
		log.Fatalf(`could not handle JWT_EXPIRES_WITH as int: %v`, err)
	}

	// shared dependencies
	db, err := sqlx.Connect("postgres", databaseURL)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	// login resources
	jwtTokenRepository := authRepositoryJwt.NewTokenRepository(
		jwtSecretKey, int64(jwtExpiresWithSeconds),
	)
	userRepository := authRepositoryPostgres.NewUserRepository(db)
	authService := authService.NewAuthService(jwtTokenRepository, userRepository)

	// login serializers
	jsonContentType := "application/json"
	loginPayloadSerializers := map[string]authSerializer.LoginPayloadSerializer{
		jsonContentType: authSerializerJson.LoginPayloadSerializer,
	}
	loginResultSerializers := map[string]authSerializer.LoginResultSerializer{
		jsonContentType: authSerializerJson.LoginResultSerializer,
	}
	loginErrorSerializers := map[string]utilSerializer.ErrorSerializer{
		jsonContentType: utilSerializerJson.ErrorSerializer,
	}

	// auth handlers
	authLoginHandler := authHandler.NewLoginHandler(
		authService,
		loginPayloadSerializers,
		loginResultSerializers,
		loginErrorSerializers,
	)

	r := chi.NewRouter()

	// middlewares injection
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/auth/login", authLoginHandler.Handle)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatal(err)
	}
}
