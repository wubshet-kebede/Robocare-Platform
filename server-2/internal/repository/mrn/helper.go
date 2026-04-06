package mrn

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
)
func EnsureSequenceExists(hospitalID uuid.UUID) error {
	safeID := strings.ReplaceAll(hospitalID.String(), "-", "_")
	seqName := "admission_seq_" + safeID

	// Use double quotes for the sequence name
	query := fmt.Sprintf(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_class WHERE relname = '%s'
			) THEN
				EXECUTE 'CREATE SEQUENCE "%s" START 1';
			END IF;
		END
		$$;
	`, seqName, seqName)

	return db.DB.Exec(query).Error
}