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
		otp_expires_at TIMESTAMP,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Table: tokens_view_lab
	CREATE TABLE IF NOT EXISTS tokens_view_lab (
		id BIGSERIAL PRIMARY KEY,
		user_view_lab_id BIGINT NOT NULL UNIQUE,
		token TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		login_last TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_view_lab_id) REFERENCES user_view_lab(id) ON DELETE CASCADE
	);

	-- Index
	CREATE INDEX IF NOT EXISTS idx_user_cid ON user_view_lab(cid);
	CREATE INDEX IF NOT EXISTS idx_tokens_user_id ON tokens_view_lab(user_view_lab_id);
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Printf("Error creating tables: %v", err)
		return err
	}

	return nil
}
