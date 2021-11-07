package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sqlc-test/repository"
	"time"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open(
		"postgres",
		"user=postgres password=qwerty dbname=sqlc_test sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	queries := repository.New(db)
	//err = queries.PrepareTables(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//status := []string{"new", "pending", "in-progress", "ready", "done"}
	//bStatus, _ := json.Marshal(status)
	//
	//for i := 0; i < 100_000; i++ {
	//	err = queries.CreateIntTableRow(ctx, 0b1111)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	err = queries.CreateJsonbTableRow(ctx, bStatus)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	//tryJson(ctx, queries)
	tryInt(ctx, queries)
}

func tryInt(ctx context.Context, queries *repository.Queries) {
	now := time.Now()
	res, _ := queries.GetCountIntTableByStatus(ctx, int32(0b_1111))
	exec := time.Since(now)
	fmt.Println(res, "int exec time: ", exec)
}

func tryJson(ctx context.Context, queries *repository.Queries) {
	now := time.Now()
	res, _ := queries.GetCountJsonbTableByStatus(ctx, "done")
	exec := time.Since(now)
	fmt.Println(res, "json exec time: ", exec)
}
