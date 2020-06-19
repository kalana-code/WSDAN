package main

import (
	"Beq/api/genaral/model"
	"Beq/api/genaral/utils"
	dispurserQueue "Beq/dispurser/db"
	packethandlerUtil "Beq/packethandler/utils"
	routes "Beq/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/joho/godotenv"
)

func server() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	// Create user name
	userDb := utils.GetUserStore()
	user := model.UserInfo{
		FirstName: "Kalana",
		LastName:  "Dhanajaya",
		Email:     "Admin",
		Password:  "$2a$10$.MFaaICc0.Ea3xl3bUFeue/xZIDQ/dMlefqRYoHg2pmSK76/hy.K6",
		Role:      "ADMIN",
		Gender:    "MALE",
		BirthDay:  "1993/10/12",
	}

	userDb.AddUser(user)

	port := os.Getenv("PORT")

	// // Handle routes
	http.Handle("/", routes.Handlers())

	// // serve
	log.Println("INFO: [SV]: Server is Online @Port:" + port)

	// log.Fatal(http.ListenAndServeTLS(":"+port, "server.crt", "server.key", nil))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func packetHandler() {
	log.Println("INFO: [PH]: Packet Handler is Activeted")
	// packethandler.PacketController()
}

func requestDispurser(task *dispurserQueue.JobQueue) {

	log.Println("INFO: [RD]: Request Dispurser Is Activeted")
	for {
		task.Dispurse()
	}

}

func main() {
	log.Println("INFO: [CO]: Controller -- ")
	queue := dispurserQueue.GetRequestQueue()
	err := packethandlerUtil.IptableInitializer()
	if err != nil {
		log.Println("ERROR: [PH]: Error when initializing iptables")
	}
	go server()
	go packetHandler()
	go requestDispurser(queue)
	exit()
}

func exit() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		fmt.Println("Exit ")
		end_waiter.Done()
	}()
	end_waiter.Wait()
}
