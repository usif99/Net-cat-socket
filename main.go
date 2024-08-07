package main

import(
	RUN "usif"
	"fmt"
	"os"
	
) 

var DefaultPort ="8989"

func main(){
	if len(os.Args) <=1{
		RUN.RunServer(DefaultPort)

	}else if (len(os.Args) == 2) && !(RUN.TestNumber(os.Args[1])){
		fmt.Println("Enter the port Number")
	
		}else if (len(os.Args) == 2) && (RUN.TestNumber(os.Args[1])){
			RUN.RunServer(os.Args[1])

		}else {
			fmt.Println("NOT CORRECT USED")
			fmt.Println("[USAGE]: ./TCPChat $port")
		}

}