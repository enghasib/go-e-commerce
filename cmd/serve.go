package cmd

import (
	"fmt"

	"github.com/enghasib/server/config"
	"github.com/enghasib/server/infrastructure/db"
	"github.com/enghasib/server/repo"
	"github.com/enghasib/server/rest"
	middleware "github.com/enghasib/server/rest/middlewares"
	productHandler "github.com/enghasib/server/rest/product"
	userHandler "github.com/enghasib/server/rest/user"
	"github.com/enghasib/server/user"
)

func Serve() {
	cnf := config.GetConfig()

	// DB connection
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("DB error:", err)
		return
	}

	// db migration
	db.Migrate()

	//repo
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	//middleware
	middleware := middleware.NewMiddlewares(cnf)

	//domain
	userService := user.NewService(userRepo)

	//handler
	usrHandler := userHandler.NewUserHandler(middleware, cnf, userService)
	prodHandler := productHandler.NewProductHandler(middleware, productRepo)

	server := rest.NewServer(cnf, usrHandler, prodHandler)
	server.Start()
}
