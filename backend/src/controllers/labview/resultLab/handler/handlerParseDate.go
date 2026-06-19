package resultLabHandler

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDateString แปลง string "YYYY-MM-DD" → time.Time รองรับทั้งปีพ.ศ. (>=2500) และค.ศ.
func ParseDateString(dateStr string) (time.Time, error) {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("รูปแบบวันที่ไม่ถูกต้อง ต้องเป็น YYYY-MM-DD")
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil || year <= 0 {
		return time.Time{}, fmt.Errorf("ปีไม่ถูกต้อง")
	}

	// ถ้าปี >= 2500 ถือว่าเป็น พ.ศ. → แปลงเป็น ค.ศ.
	if year >= 2500 {
		year -= 543
	}

	t, err := time.Parse("2006-01-02", fmt.Sprintf("%04d-%s-%s", year, parts[1], parts[2]))
	if err != nil {
		return time.Time{}, fmt.Errorf("วันที่ไม่ถูกต้อง: %w", err)
	}

	return t, nil
}
