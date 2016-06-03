package models
import (
    "fmt"
    "strings"
    "os"
    "encoding/csv"
)

// ParseCSVFile - parses the CSV File in stored in the memory
// The rows of the csv file are stored in Redis as Event Models
// These models are then made queryable through Query Controller
func ParseCSVFile(file string) (error) {
    f, error := os.Open(file)
    if error != nil {
        fmt.Println("Couln't open file")
        return error
    }
    defer f.Close() // this needs to be after the err check

    lines, error := csv.NewReader(f).ReadAll()
    if error != nil {
        fmt.Println("Couln't read file")
        return error
    }

    var events []Event
    client := SetupRedis()
    pipeline := client.Pipeline()

    for _, line := range lines {
        event := Event{
            objectID: line[0],
            objectType: line[1],
            timestamp: line[2],
            changes: strings.Join(line[3:], ", "),
        }
        events = append(events,event)
        redis_value := fmt.Sprintf("%s::%s::%s::&&",event.objectID, event.objectType, event.changes)
        pipeline.RPush(event.timestamp, redis_value)
    }
    _, err := pipeline.Exec()
    return err
}
