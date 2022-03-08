package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"
	g "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"

	. "go-service/internal/usecase/user"
)

type ApplicationContext struct {
	Health *health.Handler
	User   UserHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	gormDB, err := gorm.Open(g.Open(conf.Sql.DataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	userType := reflect.TypeOf(User{})
	userQuery := query.UseQuery(db, "users", userType)
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, userQuery)
	if err != nil {
		return nil, err
	}
	userRepository := NewUserRepository(gormDB)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userSearchBuilder.Search, userService, logError)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
