package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	tmprcli "go.temporal.io/sdk/client"
	"temporal-client/model"
)

func main() {
	tprCli, err := tmprcli.Dial(tmprcli.Options{
		HostPort: "localhost:7233",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		// execute weather workflow with the city name from request query
		cityName := r.URL.Query().Get("city")
		if cityName == "" {
			http.Error(w, "city name is required", http.StatusBadRequest)
			return
		}

		we, err := tprCli.ExecuteWorkflow(r.Context(), tmprcli.StartWorkflowOptions{
			ID:        "weather_workflow",
			TaskQueue: "weather",
		}, "weather-workflow", cityName)
		if err != nil {
			http.Error(w, "unable to start workflow", http.StatusInternalServerError)
			return
		}

		// wait for workflow to complete
		var result []model.WeatherData
		if err := we.Get(r.Context(), &result); err != nil {
			http.Error(w, "unable to get workflow result", http.StatusInternalServerError)
			return
		}

		// convert result to json in key-value pais
		response := make(map[string]interface{})
		for _, data := range result {
			response[cityName] = data
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "unable to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}) // curl -X GET http://localhost:8080/weather?city=Cairo

	server := &http.Server{Addr: ":8083", Handler: mux}
	err = server.ListenAndServe()
	if err != nil {
		return
	}

}
