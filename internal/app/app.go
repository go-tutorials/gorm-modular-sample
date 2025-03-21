package app

import (
	"context"
	"github.com/core-go/health"
	hs "github.com/core-go/health/sql"
	"github.com/core-go/log"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"
	g "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"

	. "go-service/internal/user"
)

type ApplicationContext struct {
	Health *health.Handler
	User   UserHandler
}

func NewApp(ctx context.Context, cfg Config) (*ApplicationContext, error) {
	ormDB, err := gorm.Open(g.Open(cfg.Sql.DataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ormDB.AutoMigrate(&User{})
	db, err := ormDB.DB()
	if err != nil {
		return nil, err
	}
	logError := log.LogError

	userType := reflect.TypeOf(User{})
	userQuery := query.UseQuery(db, "users", userType)
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, userQuery)
	if err != nil {
		return nil, err
	}
	userRepository := NewUserRepository(ormDB)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userSearchBuilder.Search, userService, logError)

	sqlChecker := hs.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
