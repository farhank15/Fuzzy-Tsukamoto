package database_test

import (
	"go-tsukamoto/internal/database"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestDatabase(*testing.T) func() {
	// Simpan variabel lingkungan asli
	origHost := os.Getenv("BLUEPRINT_DB_HOST")
	origPort := os.Getenv("BLUEPRINT_DB_PORT")
	origUsername := os.Getenv("BLUEPRINT_DB_USERNAME")
	origPassword := os.Getenv("BLUEPRINT_DB_PASSWORD")
	origDatabase := os.Getenv("BLUEPRINT_DB_DATABASE")
	origSchema := os.Getenv("BLUEPRINT_DB_SCHEMA")

	// Set variabel lingkungan untuk database tes
	// Perbarui nilai-nilai ini untuk menunjuk ke instance PostgreSQL tes Anda
	os.Setenv("BLUEPRINT_DB_HOST", "localhost")
	os.Setenv("BLUEPRINT_DB_PORT", "5432")
	os.Setenv("BLUEPRINT_DB_USERNAME", "postgres")
	os.Setenv("BLUEPRINT_DB_PASSWORD", "postgres")
	os.Setenv("BLUEPRINT_DB_DATABASE", "gotsukamoto")
	os.Setenv("BLUEPRINT_DB_SCHEMA", "public")

	// Kembalikan fungsi cleanup
	return func() {
		// Kembalikan variabel lingkungan asli
		os.Setenv("BLUEPRINT_DB_HOST", origHost)
		os.Setenv("BLUEPRINT_DB_PORT", origPort)
		os.Setenv("BLUEPRINT_DB_USERNAME", origUsername)
		os.Setenv("BLUEPRINT_DB_PASSWORD", origPassword)
		os.Setenv("BLUEPRINT_DB_DATABASE", origDatabase)
		os.Setenv("BLUEPRINT_DB_SCHEMA", origSchema)
	}
}

func TestDatabaseConnection(t *testing.T) {
	// Lewati jika kita tidak berada di lingkungan tes dengan database yang tersedia
	if os.Getenv("SKIP_DB_TESTS") == "true" {
		t.Skip("Skipping database tests")
	}

	// Setup lingkungan database tes
	cleanup := setupTestDatabase(t)
	defer cleanup()

	// Buat layanan database baru
	dbService := database.New()
	defer dbService.Close()

	// Periksa koneksi
	status := dbService.Health()["status"]
	assert.Equal(t, "up", status, "Database connection should be 'up'")
}
