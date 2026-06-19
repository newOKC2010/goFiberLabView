package resultLabService

import (
	"log"
	"time"

	conn "view_lab/src/database/connection"
)

// InsertLog บันทึก log การดูผล lab
func InsertLog(viewerID int64, viewerCID, patientCID string, startDate, endDate time.Time, resultCount int) {
	_, err := conn.DB.Exec(
		`INSERT INTO lab_view_log (viewer_id, viewer_cid, patient_cid, start_date, end_date, result_count)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		viewerID, viewerCID, patientCID, startDate, endDate, resultCount,
	)
	if err != nil {
		log.Printf("❌ InsertLog error: %v", err)
	}
}
