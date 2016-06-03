package controllers
import (
    "net/http"
    "fmt"
    "strings"
    "../models"
)

// Query Controller
// 
// Handles the querying of events
// Query is made through passing the query parameters
// Params - timestamp, type
//      timestamp - unix based timestamp
//      type - ObjectType

func Query(w http.ResponseWriter, r *http.Request) {
    fmt.Println("QUERY AREA", r.Method, r)

    timestamp := r.URL.Query().Get("timestamp")
    objectType := r.URL.Query().Get("type")

    client := models.SetupRedis()
    res := client.LRange(timestamp,0,-1)

    res_array := res.Val()
    var results []string

    for _, result := range res_array {
        result_values := strings.Split(result, "::")
        cur_object_type := result_values[1]
        if objectType == cur_object_type {
            change_val := result_values[2]
            results = append(results, change_val)
        } else if objectType == "*" {
            change_val := result_values[2]
            results = append(results, "{" + "\"" + cur_object_type + "\"" + ": " + change_val + "}")
        }
    }
    results_string := strings.Join(results, ", ")
    results_string = "[" + results_string + "]"
    fmt.Fprintln(w, results_string)
}
