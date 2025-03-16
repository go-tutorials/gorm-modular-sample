package app

import (
	"github.com/core-go/core"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
)

type Config struct {
	Server     core.ServerConf `mapstructure:"server"`
	Sql        SqlConfig       `mapstructure:"sql"`
	Log        log.Config      `mapstructure:"log"`
	MiddleWare mid.LogConfig   `mapstructure:"middleware"`
}
type SqlConfig struct {
	DataSourceName string `yaml:"data_source_name" mapstructure:"data_source_name" json:"dataSourceName,omitempty" gorm:"column:datasourcename" bson:"dataSourceName,omitempty" dynamodbav:"dataSourceName,omitempty" firestore:"dataSourceName,omitempty"`
	Driver         string `yaml:"driver" mapstructure:"driver" json:"driver,omitempty" gorm:"column:driver" bson:"driver,omitempty" dynamodbav:"driver,omitempty" firestore:"driver,omitempty"`
}
