package mrn

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
)
func EnsureSequenceExists(hospitalID uuid.UUID) error {
	seqName := fmt.Sprintf("admission_seq_%s", hospitalID.String())

	query := fmt.Sprintf(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_class WHERE relname = '%s'
			) THEN
				CREATE SEQUENCE %s START 1;
			END IF;
		END
		$$;
	`, seqName, seqName)

	return db.DB.Exec(query).Error
}