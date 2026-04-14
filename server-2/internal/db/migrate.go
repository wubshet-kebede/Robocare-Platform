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
				CREATE TYPE role_enum AS ENUM ('admin', 'doctor', 'surgeon', 'nurse', 'receptionist','nurse', 'lab_technician', 'pharmacist', 'er_doctor', 'er_nurse','it_support', 'robot_operator', 'head_nurse');
			END IF;
		END
		$$;
	`

	if err := DB.Exec(createEnumSQL).Error; err != nil {
		log.Fatalf("Failed to create enum type: %v", err)
	}
	createDeptTypeSQL := `
	    DO $$
		BEGIN
		    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'department_type_enum') THEN
			    CREATE TYPE department_type_enum AS ENUM ('clinical', 'administrative', 'custom','laboratory', 'emergency','pharmacy' );
			END IF;
		END
		$$;
	`

	if err := DB.Exec(createDeptTypeSQL).Error; err != nil {
		log.Fatalf("Failed to create department type enum: %v", err)
	}
	err := DB.AutoMigrate(
		&model.Hospital{},
		&model.User{},
		&model.Invitation{},
		&model.Department{},
		&model.Patient{},
		&model.Admission{},
		&model.Room{},
		&model.Robot{},
		&model.VitalSign{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}