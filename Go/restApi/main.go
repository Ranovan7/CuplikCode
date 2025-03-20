package main

import (
    "log"
    "net/http"
    "ranovan7/restApi/db"
    "ranovan7/restApi/handlers"
)

func main() {
    // Initialize the database
    if err := db.InitDB(); err != nil {
        log.Fatalf("Could not initialize database: %v", err)
    }
    defer db.DB.Close()

    // Define routes
    http.HandleFunc("/register", handlers.RegisterUser(db.DB))
    http.HandleFunc("/login", handlers.LoginUser(db.DB))

    // Start the server
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
