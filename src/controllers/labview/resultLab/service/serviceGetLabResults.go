package resultLabService

import (
	"log"
	"time"

	resultLabUtils "view_lab/src/controllers/labview/resultLab/utils"
	conn "view_lab/src/database/connection"
)

// GetLabResults ดึงผล lab ตาม CID และช่วงวันที่ แล้ว group by order_date -> lab_items_group
func GetLabResults(cid string, startDate, endDate time.Time) ([]resultLabUtils.LabGroupedResult, error) {
	log.Printf("📋 SQL params → cid=%q start=%s end=%s",
		cid, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	rows, err := conn.DB.Query(
		`SELECT
			CONCAT(pt.pname, ' ', pt.fname, '  ', pt.lname) AS pt_name,
			TO_CHAR(lh.order_date, 'YYYY-MM-DD') AS order_date,
			COALESCE(lg.lab_items_group_name, 'อื่น ๆ') AS group_name,
			li.lab_items_name,
			lr.lab_order_result
		FROM lab_head lh
		LEFT JOIN patient pt ON pt.hn = lh.hn
		LEFT JOIN lab_order lr ON lr.lab_order_number = lh.lab_order_number
		LEFT JOIN lab_items li ON li.lab_items_code = lr.lab_items_code
		LEFT JOIN lab_items_group lg ON lg.lab_items_group_code = li.lab_items_group
		WHERE pt.cid = $1
		AND lh.order_date BETWEEN $2 AND $3
		AND li.lab_items_name IS NOT NULL
		ORDER BY lh.order_date ASC, lg.lab_items_group_code ASC, li.lab_items_name ASC`,
		cid, startDate, endDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// group by order_date -> group_name -> items
	type dateKey struct{ ptName, orderDate string }
	type categoryKey struct {
		ptName    string
		orderDate string
		groupName string
	}

	dateIndexMap := make(map[dateKey]int)
	catIndexMap := make(map[categoryKey]int)
	var grouped []resultLabUtils.LabGroupedResult

	for rows.Next() {
		var ptName, orderDate, groupName, labItemName string
		var labResult *string

		if err := rows.Scan(&ptName, &orderDate, &groupName, &labItemName, &labResult); err != nil {
			log.Printf("❌ Scan error: %v", err)
			return nil, err
		}

		dKey := dateKey{ptName, orderDate}
		dateIdx, dateExists := dateIndexMap[dKey]
		if !dateExists {
			dateIdx = len(grouped)
			dateIndexMap[dKey] = dateIdx
			grouped = append(grouped, resultLabUtils.LabGroupedResult{
				PtName:    ptName,
				OrderDate: orderDate,
				Groups:    []resultLabUtils.LabCategory{},
			})
		}

		cKey := categoryKey{ptName, orderDate, groupName}
		catIdx, catExists := catIndexMap[cKey]
		if !catExists {
			catIdx = len(grouped[dateIdx].Groups)
			catIndexMap[cKey] = catIdx
			grouped[dateIdx].Groups = append(grouped[dateIdx].Groups, resultLabUtils.LabCategory{
				GroupName: groupName,
				Items:     []resultLabUtils.LabItem{},
			})
		}

		grouped[dateIdx].Groups[catIdx].Items = append(grouped[dateIdx].Groups[catIdx].Items, resultLabUtils.LabItem{
			LabItemName: labItemName,
			LabResult:   labResult,
		})
	}

	if err := rows.Err(); err != nil {
		log.Printf("❌ rows.Err: %v", err)
		return nil, err
	}

	log.Printf("📊 query ได้ %d กลุ่มวันที่", len(grouped))
	return grouped, nil
}