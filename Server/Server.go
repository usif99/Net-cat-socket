package server_test

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"usif"
)

func RunServer(PortNumber string) {
	ln, err := net.Listen("tcp", ":"+PortNumber)
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	PrintWelcomArtTer()
	fmt.Println("Listening on the port: " + PortNumber)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConnection(conn)
	}
}

type User struct {
	username string
	conn     net.Conn
}

var (
	AllUsers     []*User
	SavedMessage = []string{}
	Mutexes      sync.Mutex
)

func handleConnection(conn net.Conn) {
	// Check the users in the server
	if len(AllUsers) < 10 {

		// Ensure the connection is closed when the function returns.
		// defer conn.Close()

		// Print the welcom art
		PrintWelcomArt(conn)

		// Ask for the username
		conn.Write([]byte("[ENTER YOUR NAME]:"))

		user := &User{username: TakeUserName(conn), conn: conn}

		// Add the user to the list
		Mutexes.Lock()
		AllUsers = append(AllUsers, user)
		Mutexes.Unlock()

		// seve to history
		for _, m := range SavedMessage {
			conn.Write([]byte(m))
		}

		// Send to all the users
		for _, u := range AllUsers {
			u.conn.Write([]byte(user.username + " has joined our chat...\r\n"))
		}
		fmt.Println(TimeNow() + "[" + user.username + "]: " + "has joined our chat...")

		Mutexes.Lock()
		SavedMessage = append(SavedMessage, user.username+" has joined our chat...\r\n")
		Mutexes.Unlock()

		// Create a buffer to read incoming data.
		reader := bufio.NewReader(conn)

		// This loop will handle the massges betwen the user
		for {
			// Read data until a newline is encountered.
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(TimeNow() + "[" + user.username + "]: " + "has left our chat...")

				// Remove user from AllUsers
				var newUsers []*User
				for _, u := range AllUsers {
					if u.conn != conn {
						Mutexes.Lock()
						newUsers = append(newUsers, u)
						Mutexes.Unlock()
					}
				}
				AllUsers = newUsers

				// Broadcast leaving message
				for _, u := range AllUsers {
					u.conn.Write([]byte(user.username + " has left the chat\r\n"))
				}

				Mutexes.Lock()
				SavedMessage = append(SavedMessage, user.username+" has left the chat...\r\n")
				Mutexes.Unlock()
				return // Exit the loop since user has left

			}

			response := TimeNow() + "[" + user.username + "]:" + message

			// Process the message (here we simply convert it to uppercase).
			Mutexes.Lock()
			SavedMessage = append(SavedMessage, response)
			Mutexes.Unlock()

			for _, u := range AllUsers {
				u.conn.Write([]byte(response))
			}
		}
	} else {
		fmt.Println("The server is full")
	}
}
