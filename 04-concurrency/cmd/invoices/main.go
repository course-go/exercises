package main

import (
	"github.com/course-go/04-concurrency/internal/consumer"
	"github.com/course-go/04-concurrency/internal/invoice"
	"github.com/course-go/04-concurrency/internal/logger"
	"github.com/course-go/04-concurrency/internal/producer"
)

func main() {
	log := logger.NewDebug()
	g := invoice.NewGenerator(log)
	_ = producer.New(log, g)
	s := invoice.NewStorage(log)
	_ = consumer.New(log, s)

	// TODO: Wire the rest of dependecies and run the components.
}
