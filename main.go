package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/equal-collective-assignment/workflow"
	"github.com/equal-collective-assignment/xray"
)

func main() {
	tracer := xray.NewTracer("demo-run-001")

	// Demo workflow (mocked, deterministic)
	workflow.RunDemo(tracer)

	// JSON "dashboard"
	http.HandleFunc("/trace", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tracer.Trace()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("X-Ray demo running at http://localhost:8080/trace")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
