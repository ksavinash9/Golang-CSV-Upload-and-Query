package routers

import (
    "fmt"
    "../controllers"
    "../core"
    "../middlewares"
    "net/http"

    "github.com/go-errors/errors"
)

// getRouter returns the routers
func GetRouter() (router *core.Router) {
    router = core.NewRouter()
    // All routes go here
    router.HandleFunc("/", LogPaincs(controllers.SayhelloName))

    router.HandleFunc("/upload/", LogPaincs(controllers.UploadCSV))

    router.HandleFunc("/query", LogPaincs(controllers.Query))
    //Static Controller
    router.PathPrefix("/").Handler(&controllers.Static{"/static/public", router})
    // All middlewares go here
    router.AddMiddleware("/", &middlewares.HTTPLogger{})

    return
}

type MyError struct {
    Msg interface{}
}

func LogPaincs(function func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if x := recover(); x != nil {
                myError := &MyError{}
                myError.Msg = errors.Wrap(x, 2).ErrorStack()
                errstr := fmt.Sprintf("%v", myError.Msg)
                fmt.Println(errstr)

            }
        }()
        function(w, r)
    }
}
