package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	ctx := context.Background()
	db, err := sql.Open("postgres", "postgres://app:pass@db:5432/app?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	if err := db.PingContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	if err := http.ListenAndServe("0.0.0.0:80", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			rows, err := db.QueryContext(ctx, "select id, name, email from users order by id")
			if err != nil {
				fmt.Printf("error db query: %s", err)
				return
			}
			defer rows.Close()

			var result []users
			for rows.Next() {
				var user users
				if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
					fmt.Printf("error scan row: %s", err)
					return
				}
				result = append(result, user)
			}
			if err := rows.Err(); err != nil {
				fmt.Printf("error rows iteration: %s", err)
				return
			}

			jsonUsers, err := json.Marshal(result)
			if err != nil {
				fmt.Printf("error marshal json: %s", err)
				return
			}
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Write(jsonUsers)
		}),
	); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start HTTP server: %v\n", err)
		os.Exit(1)
	}
}
