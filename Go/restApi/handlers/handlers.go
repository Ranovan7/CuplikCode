package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "ranovan7/restApi/models"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Hash the password (you can use bcrypt or another library for this)
        hashedPassword := user.Password // Replace with actual hashing

        _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, hashedPassword)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
    }
}

func LoginUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        var storedUser models.User
        err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password)
        if err != nil {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Compare hashed password (replace with actual comparison)
        if user.Password != storedUser.Password {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Generate a token (you can use JWT or another method)
        token := "example-token" // Replace with actual token generation

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"token": token})
    }
}
