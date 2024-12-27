package main

import (
	"fmt"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, req *http.Request) {
	resp := "지금당장떠나면아무도다치지않는다그러지않으면너희는모두죽어탐정놀이도이젠끝이다오지말아야할곳에발을들였군현실로돌아가면잊지말고전해라스텔라론헌터가너희의마지막을배웅했다는것을소탕시작액션원집행목표고정즉시처단프로토콜통과초토화작전집행깨어났군한참이나기다렸다우린전에만난적이있지난스텔라론헌터샘이다일찍이네앞에나타나사실을알려주고싶었어하지만예상보다방해물이많더군열한차례시도했지만모두실패로끝났지그러는사이나도모르게이세계와긴밀히연결되어각본의구속에서벗어날수없게됐다엘리오말대로우리는이꿈의땅에서잊을수없는수확을얻게될테지나에겐그와카프카처럼사람의마음을꿰뚫어보는통찰력도은랑과블레이드처럼뛰어난특기도없다내가잘하는것들대부분은불쌍히여길필요없는악당에게만적용되지그러니내가사용할수있는수단도단하나뿐이다네게보여주기위한거야내전부를"
	fmt.Fprint(w, resp)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	// Create a channel to signal completion
	done := make(chan bool)

	// Handle the request in a goroutine
	go func() {
		response(w, req)
		done <- true
	}()

	// Wait for the response to complete
	<-done
}

func main() {
	http.HandleFunc("/", handleRequest)

	// Start server with error handling
	fmt.Println("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

	// Log fatal error if server fails
	log.Fatal("Server shutdown unexpectedly")
}