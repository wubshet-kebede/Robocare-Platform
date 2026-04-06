package mrn

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wubshet-kebede/robocare-platform/server-2/internal/db"
)
func GenerateMRN(hospitalID uuid.UUID) (string, error) {
	var nextVal int64
	safeID := strings.ReplaceAll(hospitalID.String(), "-", "_")
	seqName := fmt.Sprintf("admission_seq_%s", safeID)
	if err := EnsureSequenceExists(hospitalID); err != nil {
		return "", err
	}

	err := db.DB.Raw(
		fmt.Sprintf(`SELECT nextval('%s')`, seqName),
	).Scan(&nextVal).Error

	if err != nil {
		return "", err
	}

	year := time.Now().Year()
	mrn := fmt.Sprintf("ADM-%d-%06d", year, nextVal)

	return mrn, nil
}