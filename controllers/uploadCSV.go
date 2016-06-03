package controllers
import (
    "net/http"
    "fmt"
    "time"
    "crypto/md5"
    "os"
    "io"
    "strings"
    "../models"
    "strconv"
)

// UploadCSV Controller 
// 
// UploadCSV - uploads CSV File in system memory and calls 
// ParseCSVFile for parsing the csv and storing the 
// values in Redis
func UploadCSV(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        fmt.Println("TOKEN ", token)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        filename_array := strings.Split(handler.Filename, "/")
        path, _ := os.Getwd()
        filename := path + filename_array[len(filename_array)-1]
        dst, err := os.Create(filename)
        defer dst.Close()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //copy the uploaded file to the destination file
        if _, err := io.Copy(dst, file); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        models.ParseCSVFile(filename)
        fmt.Fprintf(w, "File Successfully Uploaded!")
    }
}

