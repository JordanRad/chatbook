package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	migrate "github.com/rubenv/sql-migrate"
)

// YYYYMMDDHHMMSS
const IDTimeFormat = "20060102150405"

type Tool struct {
	DB *sql.DB
}

// This method applies the configured migration source
func (t Tool) ApplyMigrations(withTestData bool) error {
	source := t.buildMigrationSource()

	migrations, err := source.FindMigrations()
	if err != nil {
		log.Fatalf("Error finding the migrations: %v", err.Error())
	}

	if withTestData {
		testDataMigration := t.buildTestDataMigrationSource()
		migrations = append(migrations, testDataMigration.Migrations...)
		source.Migrations = migrations
	}

	fmt.Println(migrations[len(migrations)-1].Id)

	for idx, migration := range migrations {
		parts := strings.Split(migration.Id, "_")
		if len(parts) < 2 {
			log.Fatalf("migration with ID:%v should have a description", parts[0])
		}

		err = checkTimestamp(parts[0])
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}

		if idx != len(migrations)-1 {
			nextMigrationID := migrations[idx+1].Id
			nextParts := strings.Split(nextMigrationID, "_")

			ok, err := isBeforeThan(parts[0], nextParts[0])
			if err != nil {
				log.Fatalf("error comparing dates: %v", err.Error())
			}

			if !ok {
				log.Fatalf("the sequence of the migrations is broken, please check the timestamp of the migration with ID: %v", nextMigrationID)
			}
		}

	}

	n, err := migrate.Exec(t.DB, "postgres", source, migrate.Up)
	if err != nil {
		log.Printf("error inserting the migrations: %v", err)
		return err
	}

	log.Printf("%v migration(s) have been applied successfully", n)
	return nil
}

// This method constructs the migrations in order and the id format is
// YYYYMMDDHHMMSS_some_description.
func (t Tool) buildMigrationSource() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "20230606193800_create_users_table",
				Up: []string{`
						CREATE TABLE IF NOT EXISTS users (
						id BIGSERIAL PRIMARY KEY,
						firstName VARCHAR,
						lastName VARCHAR,
						username VARCHAR,
						email VARCHAR NOT NULL UNIQUE,
						password VARCHAR NOT NULL
						);
					`},
				Down: []string{`DROP TABLE IF EXISTS users;`},
			},
		},
	}
}

// This method constructs the migrations for inserting the initial data in order and the id format is
// YYYYMMDDHHMMSS_some_description.
func (t Tool) buildTestDataMigrationSource() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "20241224224300_insert_test_data",
				Up: []string{`
							INSERT INTO public.users (id,firstname,lastname,email,"password") VALUES
								(1,'Christian','Doe','cd@example.com','$2a$10$MWZ99cecvC.WQ6w/l8s6JuOFWPKYh6xAq6sBavXT5kdFQbBvnt0b.'),
								(2,'Bro','Beaver','batbobura@example.com','$2a$10$MWZ99cecvC.WQ6w/l8s6JuOFWPKYh6xAq6sBavXT5kdFQbBvnt0b.'),
								(3,'Yasen','Litovski','yl@example.com','$2a$10$MWZ99cecvC.WQ6w/l8s6JuOFWPKYh6xAq6sBavXT5kdFQbBvnt0b.');
							`,
				},
			},
		},
	}
}

// This method verifies the consistency of the migration soruce
// and assures that the migration complies with the timestamp policy.
func checkTimestamp(migrationID string) error {
	_, err := time.Parse(IDTimeFormat, migrationID)
	if err != nil {
		return fmt.Errorf("incorrect format %w", err)
	}

	return nil
}

func isBeforeThan(candidate, comparer string) (bool, error) {
	toCompare, err := time.Parse(IDTimeFormat, candidate)
	if err != nil {
		return false, fmt.Errorf("incorrect format %w", err)
	}

	timeComparer, err := time.Parse(IDTimeFormat, comparer)
	if err != nil {
		return false, fmt.Errorf("incorrect format %w", err)
	}

	return toCompare.Before(timeComparer), nil
}
