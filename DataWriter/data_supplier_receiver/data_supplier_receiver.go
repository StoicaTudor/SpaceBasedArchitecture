package data_supplier_receiver

import (
	"DataWriter/data_contracts"
	"DataWriter/environment"
	"DataWriter/util"
	"fmt"
	"os"
	"sync"
	"time"
)

type DataSupplier interface {
	Supply(util.Consumer[data_contracts.Command], *sync.WaitGroup)
}

type MockSupplier struct{}
type KafkaSupplier struct{}

func (mockSupplier *MockSupplier) Supply(consumer util.Consumer[data_contracts.Command], wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			data := data_contracts.GetRandomCommand()
			consumer.Consume(data)
			time.Sleep(2 * time.Second)
		}
	}()
}

func (kafkaSupplier *KafkaSupplier) Supply(consumer util.Consumer[data_contracts.Command], wg *sync.WaitGroup) {
	fmt.Println("Kafka Supplier Received Data:")
}

func GetDataSupplier() (DataSupplier, error) {
	switch os.Getenv(string(environment.DataSupplySource)) {
	case "kafka":
		return &KafkaSupplier{}, nil
	case "mock":
		return &MockSupplier{}, nil
	default:
		return nil, fmt.Errorf("unknown environment: %s", environment.DataSupplySource)
	}
}
