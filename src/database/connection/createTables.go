package database

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) error {
	schema := `
	-- Table: user_view_lab
	CREATE TABLE IF NOT EXISTS user_view_lab (
		id BIGSERIAL PRIMARY KEY,
		cid VARCHAR(13) NOT NULL UNIQUE,
		hash_cid VARCHAR(255) NOT NULL,
		role VARCHAR(20) NOT NULL DEFAULT 'user',
		status BOOLEAN NOT NULL DEFAULT true,
		full_name VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		facility_type VARCHAR(10), -- ประเภทสถานบริการ
		facility_code VARCHAR(20), -- รหัสสถานบริการ
		facility_name VARCHAR(255), -- ชื่อสถานบริการ
		otp_code VARCHAR(10),
		otp_expires_at TIMESTAMPTZ,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Table: tokens_view_lab
	CREATE TABLE IF NOT EXISTS tokens_view_lab (
		id BIGSERIAL PRIMARY KEY,
		user_view_lab_id BIGINT NOT NULL UNIQUE,
		token TEXT NOT NULL,
		expires_at TIMESTAMPTZ NOT NULL,
		login_last TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_view_lab_id) REFERENCES user_view_lab(id) ON DELETE CASCADE
	);

	-- Table: lab_view_log
	CREATE TABLE IF NOT EXISTS lab_view_log (
		id BIGSERIAL PRIMARY KEY,
		viewer_id BIGINT NOT NULL,
		viewer_cid VARCHAR(13) NOT NULL,
		patient_cid VARCHAR(13) NOT NULL,
		start_date DATE NOT NULL,
		end_date DATE NOT NULL,
		result_count INT NOT NULL DEFAULT 0,
		viewed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (viewer_id) REFERENCES user_view_lab(id) ON DELETE CASCADE
	);

	-- Index
	CREATE INDEX IF NOT EXISTS idx_user_cid ON user_view_lab(cid);
	CREATE INDEX IF NOT EXISTS idx_tokens_user_id ON tokens_view_lab(user_view_lab_id);
	CREATE INDEX IF NOT EXISTS idx_lab_view_log_viewer ON lab_view_log(viewer_id);
	CREATE INDEX IF NOT EXISTS idx_lab_view_log_patient ON lab_view_log(patient_cid);
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		return err
	}

	return nil
}
