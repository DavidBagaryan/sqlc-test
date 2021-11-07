package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sqlc-test/repository"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open(
		"postgres",
		"user=postgres password=qwerty dbname=sqlc_test sslmode=disable",
	)

	if err != nil {
		log.Fatal("an err on db conn", err)
	}

	m := Manage{repository.New(db)}
	//m.PrepareData(ctx)
	//m.TryJson(ctx)
	m.TryInt(ctx)
}

type Manage struct {
	*repository.Queries
}

func (m *Manage) PrepareData(ctx context.Context) {
	err := m.PrepareTables(ctx)
	if err != nil {
		log.Fatal("an err on preparing table", err)
	}

	statusHistoryJSON := []string{"new", "pending", "in-progress", "ready", "done"}
	bStatusHistoryJSON, _ := json.Marshal(statusHistoryJSON)

	for i := 0; i < 100_000; i++ {
		err = m.CreateIntTableRow(ctx, 0b1111)
		if err != nil {
			log.Fatal("an err on populating int table", err)
		}

		err = m.CreateJsonbTableRow(ctx, bStatusHistoryJSON)
		if err != nil {
			log.Fatal("an err on populating json table", err)
		}
	}
}

func (m *Manage) TryInt(ctx context.Context) {
	err := m.checkExecutionTime(ctx, "int", func(ctx context.Context) error {
		_, err := m.GetCountIntTableByStatus(ctx, int32(0b_1111))
		return err
	})

	if err != nil {
		log.Fatal("an err on populating int table", err)
	}
}

func (m *Manage) TryJson(ctx context.Context) {
	err := m.checkExecutionTime(ctx, "json", func(ctx context.Context) error {
		_, err := m.GetCountJsonbTableByStatus(ctx, "done")
		return err
	})

	if err != nil {
		log.Fatal("an err on populating json table", err)
	}
}

func (m *Manage) checkExecutionTime(ctx context.Context, benchName string, fn func(ctx context.Context) error) error {
	now := time.Now()
	err := fn(ctx)
	exec := time.Since(now)
	fmt.Println(fmt.Sprintf("%s exec time: %s", benchName, exec))

	return err
}
