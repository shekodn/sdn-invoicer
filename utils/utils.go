package utils

import (
  "encoding/json"
  "net/http"
  "fmt"
)


func Message(status bool, message string)(map[string]interface{}) {
    return map[string]interface{} {"status" : status, "message" : message}
}

// func Respond(w http.ResponseWriter, data map[string] interface{}) {
//     w.Header().Add("Content-Type", "application/json")
//     status := data["status"]
//     json.NewEncoder(w).Encode(status)
// }


func Respond2(w http.ResponseWriter, data map[string] interface{}) {
    // w.Header().Add("Content-Type", "application/json")
    w.Write([]byte(fmt.Sprintf("<script type='text/javascript'>alert('xss');</script>")))
    fmt.Println("data", data)
    json.NewEncoder(w).Encode(data)
}

func PrettyPrint(i interface{}) string {
  s, _ := json.MarshalIndent(i, "", "\t")
  return string(s)
}
