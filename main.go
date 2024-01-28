package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func calculateExpression(expression string) (float64, error) {
	result, err := strconv.ParseFloat(strings.ReplaceAll(expression, "÷", "/"), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	expression := r.FormValue("expression")
	result, err := calculateExpression(expression)
	if err != nil {
		http.Error(w, "Invalid expression", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%.2f", result)
}

func main() {
	// Serve calculator.html when accessing http://localhost:8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "calculator.html")
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
