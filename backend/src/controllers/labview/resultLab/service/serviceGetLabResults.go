package resultLabService

import (
	"log"
	"time"

	resultLabUtils "view_lab/src/controllers/labview/resultLab/utils"
	conn "view_lab/src/database/connection"
)

// GetLabResults ดึงผล lab ตาม CID และช่วงวันที่ แล้ว group by pt_name + order_date
func GetLabResults(cid string, startDate, endDate time.Time) ([]resultLabUtils.LabGroupedResult, error) {
	log.Printf("📋 SQL params → cid=%q start=%s end=%s",
		cid, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	rows, err := conn.DB.Query(
		`SELECT
			CONCAT(pt.pname, ' ', pt.fname, '  ', pt.lname) AS pt_name,
			TO_CHAR(lh.order_date, 'YYYY-MM-DD') AS order_date,
			li.lab_items_name,
			lr.lab_order_result
		FROM lab_head lh
		LEFT JOIN patient pt ON pt.hn = lh.hn
		LEFT JOIN lab_order lr ON lr.lab_order_number = lh.lab_order_number
		LEFT JOIN lab_items li ON li.lab_items_code = lr.lab_items_code
		WHERE pt.cid = $1
		AND lh.order_date BETWEEN $2 AND $3
		AND li.lab_items_name IS NOT NULL
		ORDER BY lh.order_date ASC, li.lab_items_name`,
		cid, startDate, endDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// group by pt_name + order_date
	type groupKey struct{ ptName, orderDate string }
	indexMap := make(map[groupKey]int)
	var grouped []resultLabUtils.LabGroupedResult

	for rows.Next() {
		var raw resultLabUtils.LabResultItem
		if err := rows.Scan(&raw.PtName, &raw.OrderDate, &raw.LabItemName, &raw.LabResult); err != nil {
			log.Printf("❌ Scan error: %v", err)
			return nil, err
		}

		key := groupKey{raw.PtName, raw.OrderDate}
		if idx, exists := indexMap[key]; exists {
			grouped[idx].Items = append(grouped[idx].Items, resultLabUtils.LabItem{
				LabItemName: raw.LabItemName,
				LabResult:   raw.LabResult,
			})
		} else {
			indexMap[key] = len(grouped)
			grouped = append(grouped, resultLabUtils.LabGroupedResult{
				PtName:    raw.PtName,
				OrderDate: raw.OrderDate,
				Items: []resultLabUtils.LabItem{{
					LabItemName: raw.LabItemName,
					LabResult:   raw.LabResult,
				}},
			})
		}
	}

	if err := rows.Err(); err != nil {
		log.Printf("❌ rows.Err: %v", err)
		return nil, err
	}

	log.Printf("📊 query ได้ %d กลุ่ม", len(grouped))
	return grouped, nil
}
