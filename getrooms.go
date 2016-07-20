package main

import (
  "net/http"
  "log"
  "github.com/kr/pretty"
  "io/ioutil"
  "encoding/json"
)

func getrooms(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Access-Control-Allow-Origin", "*") // FIXME: Unknown Cross Domain

  // If it is a get request user should view main page
  if r.Method == "GET"{
    http.Redirect(w, r, "/", http.StatusFound)
  } else {
    // TODO: Get Params from JSON
    r.ParseForm()
    if r.Body != nil {
      bodyBytes, _ := ioutil.ReadAll(r.Body)
      r.Body.Close()
      if r.Header.Get("Content-Type") == "application/json" {
        var data map[string]interface{}
        json.Unmarshal(bodyBytes, &data)
        log.Printf("JSON: %# v\n", pretty.Formatter(data))
      }
    }

    w.Header().Set("Content-Type", "application/json");
    rooms, err := globalRoomStore.GetAll()
    if err != nil{
      //FIXME
      w.Write([]byte("{\"error\": {\"text\":\"Database error\"}, \"code\":501}}\n"))
    } else {
      w.Write([]byte(rooms))
      //FIXME
      globalRoomStore.SaveAll()
    }
  }
}
