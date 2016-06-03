package models

import (
    "strings"
)

// *****************************************************************************
// Event
// *****************************************************************************

// Event Model containing details about each event
type Event struct {
    objectID string
    objectType string
    timestamp string
    changes string
}

// EventByTimestampObjectType gets event information by timestamp
func EventByTimestampObjectType(timestamp, objectType string) {
    client := SetupRedis()
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
    // return results
}

// EventCreate creates event
// func EventCreate(firstName, lastName, email, password string) error {
    
// }