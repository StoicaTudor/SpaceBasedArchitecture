package main

import (
	"DataWriter/data_contracts"
	"DataWriter/data_supplier_receiver"
	"DataWriter/environment"
	"DataWriter/util"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func connectMariaDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv(string(environment.MariaDBUser)),
		os.Getenv(string(environment.MariaDBPassword)),
		os.Getenv(string(environment.MariaDBHost)),
		os.Getenv(string(environment.MariaDBPort)),
		os.Getenv(string(environment.MariaDBDBName)),
	)
	return sql.Open("mysql", dsn)
}

func processMessage(message kafka.Message, mariaDB *sql.DB) {
	// Insert into MariaDB
	_, err := mariaDB.Exec("INSERT INTO users (data) VALUES (?)", string(message.Value))
	if err != nil {
		log.Printf("Error inserting into MariaDB: %v", err)
	}
}

func consumeKafka() {
	// Connect to MariaDB
	mariaDB, err := connectMariaDB()
	if err != nil {
		log.Fatalf("Error connecting to MariaDB: %v", err)
	}
	defer mariaDB.Close()

	// Kafka reader configuration
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv(string(environment.KafkaBroker))},
		Topic:   os.Getenv(string(environment.KafkaTopic)),
		GroupID: os.Getenv(string(environment.KafkaGroupID)),
	})

	// Consume messages from Kafka
	for {
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process each message asynchronously
		go processMessage(msg, mariaDB)
	}
}

// Produce random strings to Kafka every 2 seconds
func produceRandomMessages() {
	// Kafka writer configuration
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv(string(environment.KafkaBroker))},
		Topic:   os.Getenv(string(environment.KafkaTopic)),
	})

	// Send a random message to Kafka every 2 seconds
	for {
		// Generate a random string
		randomString := util.GenerateRandomString(10)

		// Create Kafka message
		message := kafka.Message{
			Value: []byte(randomString),
		}

		// Write message to Kafka
		err := kafkaWriter.WriteMessages(context.Background(), message)
		if err != nil {
			log.Printf("Error writing message to Kafka: %v", err)
		} else {
			log.Printf("Produced message: %s", randomString)
		}

		// Wait for 2 seconds before producing the next message
		time.Sleep(2 * time.Second)
	}
}

type PrintConsumer struct{}

func (printConsumer *PrintConsumer) Consume(command data_contracts.Command) {
	switch castedCommand := command.(type) {
	case *data_contracts.UserCreateDTO:
		fmt.Println("UserCreateDTO - ID:", castedCommand.ID)
		fmt.Println("UserCreateDTO - Name:", castedCommand.Name)
		fmt.Println("UserCreateDTO - Balance:", castedCommand.Balance)
	case *data_contracts.UserUpdateDTO:
		fmt.Println("UserUpdateDTO - ID:", castedCommand.ID)
		fmt.Println("UserUpdateDTO - Name:", castedCommand.Name)
		fmt.Println("UserUpdateDTO - Balance:", castedCommand.Balance)
	case *data_contracts.UserDeleteDTO:
		fmt.Println("UserDeleteDTO - ID:", castedCommand.ID)
	default:
		fmt.Println("Unknown Command Type")
	}
}

func main() {
	var wg sync.WaitGroup

	// LEARN: export function from file -> start with uppercase
	environment.Load()

	// ---------------------
	//log.Println("Starting Kafka consumer...")
	//
	//// Start the Kafka producer in the background
	//go produceRandomMessages()
	//
	//// Start consuming messages from Kafka
	//consumeKafka()
	// ---------------------

	dataSupplier, _ := data_supplier_receiver.GetDataSupplier()
	wg.Add(1)
	dataSupplier.Supply(&PrintConsumer{}, &wg)
	wg.Wait()
}
