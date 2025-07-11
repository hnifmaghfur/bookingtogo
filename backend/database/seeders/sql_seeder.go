package seeders

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

// RunSQLSeeder membaca file SQL dari filePath dan mengeksekusinya di database.
func RunSQLSeeder(db *gorm.DB, filePath string) {
	// Membaca konten file SQL
	sqlContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Gagal membaca file SQL seeder %s: %v", filePath, err)
	}

	// Mengeksekusi konten SQL
	if err := db.Exec(string(sqlContent)).Error; err != nil {
		log.Fatalf("Gagal menjalankan SQL seeder dari %s: %v", filePath, err)
	}
	fmt.Printf("Berhasil menjalankan SQL seeder dari %s\n", filePath)
}
