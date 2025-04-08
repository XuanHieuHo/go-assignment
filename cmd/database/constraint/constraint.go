package main

import (
	"fmt"
	"log"

	"github.com/XuanHieuHo/go-assignment/config"
	"gorm.io/gorm"
)

type TableSchema struct {
	TableName   string
	Constraints map[string]string
	Indexes     map[string]string
}

func applySchemaChange(db *gorm.DB, schema TableSchema) {
	var tableExists bool
	checkTableExistsSQL := fmt.Sprintf(`SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = '%s');`, schema.TableName)
	if err := db.Raw(checkTableExistsSQL).Scan(&tableExists).Error; err != nil {
		log.Fatal("Failed to check table exist:", err)
	}

	if tableExists {
		for name, sql := range schema.Constraints {
			dropSQL := fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT IF EXISTS %s;", schema.TableName, name)
			if err := db.Exec(dropSQL).Error; err != nil {
				log.Fatalf("Failed to drop constraint '%s' on table '%s': %v", name, schema.TableName, err)
			}
			if err := db.Exec(sql).Error; err != nil {
				log.Fatalf("Failed to add/update constraint '%s' on table '%s': %v", name, schema.TableName, err)
			}
		}

		for name, sql := range schema.Indexes {
			if err := db.Exec(sql).Error; err != nil {
				log.Fatalf("Failed to add index '%s' on table '%s': %v", name, schema.TableName, err)
			}
		}
	} else {
		log.Printf("Table '%s' does not exist, skipping adding constraints and index.", schema.TableName)
	}
}

func main() {
	db := config.DatabaseConnect()
	log.Println("Running database constraint, index,...")

	schemas := []TableSchema{
		{
			TableName: "friend_ships",
			Constraints: map[string]string{
				"unique_friendship": "ALTER TABLE friend_ships ADD CONSTRAINT unique_friendship UNIQUE (user_id, friend_id);",
				"check_friendship":  "ALTER TABLE friend_ships ADD CONSTRAINT check_friendship CHECK (user_id < friend_id);",
			},
			Indexes: map[string]string{
				"idx_user_friend": "CREATE INDEX IF NOT EXISTS idx_user_friend ON friend_ships (user_id, friend_id);",
			},
		},
	}
	for _, schema := range schemas {
		applySchemaChange(db, schema)
	}
}
