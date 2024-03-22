package adapter

import (
	"flick_tickets/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostGresql struct {
	config   *configs.Configs
	postgres *gorm.DB
}

func NewpostgreDb(config *configs.Configs) *PostGresql {

	postgresDb := &PostGresql{
		config: config,
	}
	err := postgresDb.connect()
	log.Println("starting connect postgre")
	if err != nil {
		log.Fatalf("fail to connect postgre")
	}
	return postgresDb
}

func (p *PostGresql) connect() error {
	dsn := p.config.DataSource
	if dsn == "" {
		return fmt.Errorf("empty postgre uri")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgre: %w", err)
	}

	p.postgres = db
	log.Println("Connected to postgre successfully")
	return nil
}
func (p *PostGresql) CreateCollection() *gorm.DB {
	return p.postgres
}
