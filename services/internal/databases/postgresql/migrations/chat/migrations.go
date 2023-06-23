package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
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
				Id: "20230617161100_init",
				Up: []string{`
				-- Enable the uuid-ossp extension for generating UUIDs
				CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

				-- Enable the TimescaleDB extension
				CREATE EXTENSION IF NOT EXISTS timescaledb;

				-- Create the conversations table
				CREATE TABLE IF NOT EXISTS conversations (
				  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY
				);
				
				-- Create the conversations_participants table
				CREATE TABLE IF NOT EXISTS conversations_participants (
				  id SERIAL PRIMARY KEY,
				  conversation_id UUID NOT NULL,
				  participant_id UUID NOT NULL,
				  FOREIGN KEY (conversation_id) REFERENCES conversations (id),
				  CONSTRAINT uq_conversation_participant UNIQUE (conversation_id, participant_id)
				);
				
				-- Create the messages table
				CREATE TABLE IF NOT EXISTS messages (
				  id UUID DEFAULT uuid_generate_v4(),
				  conversation_id UUID NOT NULL,
				  sender_id UUID NOT NULL,
				  content TEXT NOT NULL,
				  ts TIMESTAMP NOT NULL DEFAULT NOW(),
				  PRIMARY KEY (id, ts),
				  FOREIGN KEY (conversation_id) REFERENCES conversations (id)
				);
				
				-- Create a hypertable on the messages table for time-series data
				SELECT create_hypertable('messages', 'ts');
					`},
				Down: []string{
					`
					-- Drop the hypertable on the messages table
					SELECT drop_hypertable('messages');

					-- Drop the messages table
					DROP TABLE IF EXISTS messages;

					-- Drop the conversations_participants table
					DROP TABLE IF EXISTS conversations_participants;

					-- Drop the conversations table
					DROP TABLE IF EXISTS conversations;
					`,
				},
			},
			{
				Id: "20230617183700_add_index_for_full_text_search",
				Up: []string{
					`
					CREATE EXTENSION IF NOT EXISTS pg_trgm;
					CREATE INDEX IF NOT EXISTS messages_content_fts_idx ON messages USING gin (to_tsvector('english', content));
					`,
				},
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
				-- Insert statements for conversations table
				INSERT INTO conversations (id) VALUES
				    ('a9e1393b-1f26-4e10-b44d-18e7e0be8682'),
				    ('2d3e26be-47fd-4f2b-84fd-2675ffae2c6c'),
				    ('7c273a7f-dbf5-4c1a-a3c2-8b0e25a6fb87'),
				    ('5f2c6d0e-0e7d-4a72-9d01-6c0e8bdc8354'),
				    ('3aee9b2e-9346-4ea4-8b6b-21b7b3c2a41b');

				-- Insert statements for conversations_participants table
				INSERT INTO conversations_participants (conversation_id, participant_id) VALUES
				    ('a9e1393b-1f26-4e10-b44d-18e7e0be8682', '7ffb8e45-ba8c-4766-b44c-713abd9b10e4'),
				    ('a9e1393b-1f26-4e10-b44d-18e7e0be8682', '7ffb8e45-ba8c-4766-b44c-713abd9b10e5'),
				    ('2d3e26be-47fd-4f2b-84fd-2675ffae2c6c', '7ffb8e45-ba8c-4766-b44c-713abd9b10e5'),
				    ('2d3e26be-47fd-4f2b-84fd-2675ffae2c6c', '7ffb8e45-ba8c-4766-b44c-713abd9b10e9'),
				    ('7c273a7f-dbf5-4c1a-a3c2-8b0e25a6fb87', '7ffb8e45-ba8c-4766-b44c-713abd9b10e3'),
				    ('7c273a7f-dbf5-4c1a-a3c2-8b0e25a6fb87', '7ffb8e45-ba8c-4766-b44c-713abd9b10e9'),
				    ('5f2c6d0e-0e7d-4a72-9d01-6c0e8bdc8354', '7ffb8e45-ba8c-4766-b44c-713abd9b10e3'),
				    ('5f2c6d0e-0e7d-4a72-9d01-6c0e8bdc8354', '7ffb8e45-ba8c-4766-b44c-713abd9b10e4'),
				    ('5f2c6d0e-0e7d-4a72-9d01-6c0e8bdc8354', '7ffb8e45-ba8c-4766-b44c-713abd9b10e9'),
				    ('3aee9b2e-9346-4ea4-8b6b-21b7b3c2a41b', 'a4049425-75e6-4b88-9ca3-2a069a57e5ae');
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
