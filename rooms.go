package main

import (
    "io/ioutil"
    "encoding/json"
    "os"
    "log"
)

/*
[
	//room object
	{
		name: "room name",
		location: "location name",
		equipment: [
			//list of avaiable equipment
			"equipment name",
			...
		],
		size: "NNm²", //size of the room in m²
		capacity: /suggested number of people a room can fit/,
		avail: [
			//list of times at which a room is free and can be booked,
			//between 7 am and 7 pm in 15min steps
			"HH:mm - HH:mm",
			"HH:mm - HH:mm",
			...
		],
		images: [
			//can contain 0-3 images
			"url of image",
			...
		]
	},
	...
]
*/

type Room struct {
  Name          string    `json:"name"`       // Room name
  Location      string    `json:"location"`   // Locagolantion Name
  Equipment     []string  `json:"equipment"`  // List of available equipment
  Size          string    `json:"size"`       // Size of room
  Capacity      int       `json:"capacity"`   // Number of people
  Avail         []string  `json:"avail"`      // Availability
  Images        []string  `json:"images"`     // Photos of room
}

type RoomStore interface {
  SaveAll() error
  GetAll() (string,error)
}

type FileRoomStore struct {
  filename string
  Rooms []Room
}

func (store FileRoomStore) GetAll() (string, error){
  contents, err := json.MarshalIndent(store.Rooms, "", "  ")
  if err != nil {
    return "", err
  }
  return string(contents), nil
}

func (store FileRoomStore) SaveAll() error {
  contents, err := json.MarshalIndent(store.Rooms, "", "  ")
  if  err != nil{
    return err
  }

  err = ioutil.WriteFile(store.filename, contents, 0660)
  if err != nil {
    return err
  }
  return nil
}

func NewFileRoomStore(filename string) (*FileRoomStore, error){
  store := &FileRoomStore{
    filename: filename,
  }

  contents, err := ioutil.ReadFile(filename)
  if err != nil {
    if os.IsNotExist(err){
      return store, nil
    }
    return nil, err
  }

  err = json.Unmarshal(contents, &store.Rooms)
  if err != nil {
    return nil, err
  }
  return store, nil
}

var globalRoomStore RoomStore

func init(){
  store, err := NewFileRoomStore("./data/rooms.json")
  if err != nil {
    log.Fatalf("Error creating rooms store: %s", err)
  }
  globalRoomStore = store
  globalRoomStore.SaveAll()
}
