package main

import(
	"net/http"
	"encoding/json"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, q *http.Request) {
        message := map[string]string{
            "message": "hello world",
        }
        jsonMessage, err := json.Marshal(message)
        if err != nil {
            panic(err.Error())
        }
        w.Write(jsonMessage)
    })
    http.ListenAndServe("127.0.0.1:8080", mux)
}