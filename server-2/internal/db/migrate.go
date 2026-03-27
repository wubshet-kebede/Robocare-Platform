package db

import (
	"log"

	"github.com/wubshet-kebede/robocare-platform/server-2/internal/model"
)
func Migrate() {
	createEnumSQL := `
		DO $$ 
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
				CREATE TYPE role_enum AS ENUM ('admin', 'staff');
			END IF;
		END
		$$;
	`

	if err := DB.Exec(createEnumSQL).Error; err != nil {
		log.Fatalf("Failed to create enum type: %v", err)
	}
	err := DB.AutoMigrate(
		&model.Hospital{},
		&model.User{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}