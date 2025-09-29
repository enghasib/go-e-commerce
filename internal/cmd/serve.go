package cmd

import (
	"fmt"

	"github.com/enghasib/server/internal/config"
	"github.com/enghasib/server/internal/infrastructure/db"
	"github.com/enghasib/server/internal/repo"
	"github.com/enghasib/server/internal/rest"
	middleware "github.com/enghasib/server/internal/rest/middlewares"
	"github.com/enghasib/server/internal/rest/product"
	"github.com/enghasib/server/internal/rest/user"
)

func Serve() {
	cnf := config.GetConfig()

	// DB connection
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("DB error:", err)
		return
	}

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	middleware := middleware.NewMiddlewares(cnf)

	userHandler := user.NewUserHandler(middleware, cnf, userRepo)
	productHandler := product.NewProductHandler(middleware, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
