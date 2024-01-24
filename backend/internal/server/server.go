package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/CrYptOz007/Fusion/internal/database"
	"github.com/CrYptOz007/Fusion/internal/routes"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	"github.com/labstack/echo/v4"
	echoMiddleWare "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	DB  *gorm.DB
	App *echo.Echo
	types.Groups
}

// Init initializes the server
func (s *Server) Init(connection *database.Connection) {
	s.DB = connection.Db
	s.App = echo.New()

	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	go s.echoInit(wg)
}

// echoInit initializes the echo server
func (s *Server) echoInit(wg *sync.WaitGroup) {
	defer wg.Done()

	s.App.Use(echoMiddleWare.CORSWithConfig(echoMiddleWare.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "ResponseType", "Cache-Control"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodConnect},
	}))

	dbClient := s.DB

	s.App.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			utils.SetLocal[*gorm.DB](c, "dbClient", dbClient)
			return next(c)
		}
	})

	// Register routes
	routes.Register(s.App, s.Groups)

	// Start the server
	if err := s.App.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
