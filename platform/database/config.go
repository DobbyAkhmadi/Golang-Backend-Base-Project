package database

import (
	"backend/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

var err error

func Connect() (*gorm.DB, error) {
	// Connect to db using GORM
	dbType := config.Config.GetString("DB.TYPE")
	switch dbType {
	case "postgres":
		return ConnectPostgres()
	case "mysql":
		return ConnectMySQL()
	case "sqlserver":
		return ConnectSQLServer()
	default:
		log.Fatal("Unknown database type specified in DB.TYPE configuration")
		return nil, fmt.Errorf("unknown database type")
	}
}

func ConnectPostgres() (*gorm.DB, error) {
	// Connect to PostgreSQL using GORM
	host := config.Config.GetString("DB.ADDRESS")
	port := config.Config.GetString("DB.PORT")
	user := config.Config.GetString("DB.USERNAME")
	password := config.Config.GetString("DB.PASSWORD")
	dbName := config.Config.GetString("DB.DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return DB, nil
}

func ConnectMySQL() (*gorm.DB, error) {
	// Connect to MySQL using GORM
	host := config.Config.GetString("DB.ADDRESS")
	port := config.Config.GetString("DB.PORT")
	user := config.Config.GetString("DB.USERNAME")
	password := config.Config.GetString("DB.PASSWORD")
	dbName := config.Config.GetString("DB.DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return DB, nil
}

func ConnectSQLServer() (*gorm.DB, error) {
	// Connect to SQL Server using GORM
	host := config.Config.GetString("DB.ADDRESS")
	port := config.Config.GetString("DB.PORT")
	user := config.Config.GetString("DB.USERNAME")
	password := config.Config.GetString("DB.PASSWORD")
	dbName := config.Config.GetString("DB.DATABASE")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return DB, nil
}
