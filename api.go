package main

import (
	"database/sql"
  "math/rand"
	// "log"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "fmt"
  "strconv"
	_"github.com/go-sql-driver/mysql"
)

type Enemy struct {
	ID          int      `json:"id"`
	Attacks     string   `json:"attacks"`
  RewardClass int      `json:"rewardclass"`
  HP          int      `json:"hp"`
}

type Card struct{
  ID          int      `json:"id"`
  Name        string   `json:"name"`
  Cost        int      `json:"cost"`
  Type        string   `json:"type"`
  Effects     string   `json:"effects"`
  Color       string   `json:"color"`
  Upgraded    int      `json:"upgraded"`
  CardText    string   `json:"cardText"`
  Rarity      int      `json:"rarity"`
}

type User struct{
  Level       int      `json:"level"`
}

func GetEnemy(response http.ResponseWriter, request *http.Request){
  response.Header().Set("content-type", "application/json")
  var enemies []Enemy
  var user User
	err := json.NewDecoder(request.Body).Decode(&user)
  fmt.Println(user.Level)
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/stsclone")
  results, err := db.Query("SELECT * FROM Enemy where RewardClass =" + strconv.Itoa(user.Level))
	defer db.Close()

  for results.Next() {
  	var enemy Enemy
  	err = results.Scan(&enemy.ID, &enemy.Attacks, &enemy.RewardClass, &enemy.HP)
  	if err != nil {
  		panic(err.Error())
  	}

  	enemies = append(enemies, enemy)
  }
	json.NewEncoder(response).Encode(enemies[rand.Intn(len(enemies))])

}


func GetReward(response http.ResponseWriter, request *http.Request){
  response.Header().Set("content-type", "application/json")

  var cards []Card
  var user User
	err := json.NewDecoder(request.Body).Decode(&user)
  fmt.Println(user.Level)
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/stsclone")
  results, err := db.Query("SELECT * FROM Card where Rarity <=" + strconv.Itoa(user.Level))
	defer db.Close()

  for results.Next() {
  	var card Card
  	err = results.Scan(&card.ID, &card.Name, &card.Cost, &card.Type, &card.Effects, &card.Color, &card.Upgraded, &card.CardText, &card.Rarity)
  	if err != nil {
  		panic(err.Error())
    }
    cards = append(cards, card)

  }
	json.NewEncoder(response).Encode(cards)

}


func main() {

  fmt.Println("Starting the application...")
  router := mux.NewRouter()
	router.HandleFunc("/Enemies", GetEnemy).Methods("POST")
  router.HandleFunc("/Rewards", GetReward).Methods("POST")
  http.ListenAndServe(":12345", router)
}
