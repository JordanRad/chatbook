package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/JordanRad/chatbook/services/internal/databases/postgresql"

	_ "github.com/lib/pq"
)

func main() {
	conn := postgresql.ConnectToDatabase("chat-dev", "123456", "localhost", 5434, "chat-dev-db")

	// Prepare the insert query with multiple value sets
	insertQuery := "INSERT INTO messages (conversation_id, sender_id, ts, content) VALUES "

	startTime := time.Now().Add(-10 * 24 * time.Hour)
	for i := 0; i < 500*20; i++ {
		sqlValues := generateSQLValues(startTime)

		toInsert := strings.Join(sqlValues, ", ")

		insertQuery += toInsert
		insertQuery += ";"

		// Execute the bulk insert query
		_, err := conn.Exec(insertQuery)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Chunk number %d has been inserted", i+1)

		insertQuery = "INSERT INTO messages (conversation_id, sender_id, ts, content) VALUES "
		startTime = startTime.Add(-2 * time.Hour)
	}

}
func generateRandomChatMessage() string {
	// List of possible chat message content
	possibleWords1 := []string{
		"Hello",
		"How",
		"Whats",
		"Nice",
		"Have",
		"Good",
		"Great",
		"Awesome",
		"Amazing",
		"Hi",
		"Greetings",
		"Hey",
	}

	possibleWords2 := []string{
		"are",
		"is",
		"you",
		"to",
		"day",
		"meet",
		"see",
		"doing",
		"going",
		"up",
		"it",
		"going",
	}

	possibleWords3 := []string{
		"doing?",
		"up?",
		"today?",
		"doing today?",
		"going on?",
		"having a good day?",
		"doing well?",
		"to see you!",
		"to meet you!",
		"to hear from you!",
		"to chat with you!",
		"to catch up with you!",
	}
	randomIndex1 := rand.Intn(len(possibleWords1))
	randomIndex2 := rand.Intn(len(possibleWords2))
	randomIndex3 := rand.Intn(len(possibleWords3))

	// Build the random chat message
	words := []string{
		possibleWords1[randomIndex1],
		possibleWords2[randomIndex2],
		possibleWords3[randomIndex3],
	}
	randomMessage := strings.Join(words, " ")

	// Return the random chat message
	return randomMessage
}

func generateSQLValues(startTime time.Time) []string {
	const (
		conversationID string = "a9e1393b-1f26-4e10-b44d-18e7e0be8682"
		chunkSize      int    = 3000
	)

	participants := []string{"7ffb8e45-ba8c-4766-b44c-713abd9b10e5", "7ffb8e45-ba8c-4766-b44c-713abd9b10e4"}

	sqlValues := make([]string, 0, chunkSize)

	currentTime := startTime
	for i := 0; i < chunkSize; i++ {
		senderID := participants[0]
		if i%2 == 0 {
			senderID = participants[1]
		}
		sqlValues = append(sqlValues, fmt.Sprintf("('%s','%s','%s','%s')", conversationID, senderID, currentTime.Format("2006-01-02 15:04:05.999999"), generateRandomChatMessage()))
		currentTime = currentTime.Add(1 * time.Second)
	}

	return sqlValues
}
