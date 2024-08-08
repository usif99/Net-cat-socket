package usif

import(
"bufio"
"fmt"
"net"
"strings"
"time"
"usif"
)
func TestNumber(numbers )bool{
	for _,number := range numbers{
		if number < '0'||number > '9'{
			return false 
		}

	}
	return true
}

//Print Welcome logo

func PrintWelcome(conn net.conn){
	conn.Write([]byte("Welcome to TCP-Chat!\r\n"))
	conn.Write([]byte("         _nnnn_\r\n"))
	conn.Write([]byte("        dGGGGMMb\r\n"))
	conn.Write([]byte("       @p~qp~~qMb\r\n"))
	conn.Write([]byte("       M|@||@) M|\r\n"))
	conn.Write([]byte("       @,----.JM|\r\n"))
	conn.Write([]byte("      JS^\\__/  qKL\r\n"))
	conn.Write([]byte("     dZP        qKRb\r\n"))
	conn.Write([]byte("    dZP          qKKb\r\n"))
	conn+
	.Write([]byte("   fZP            SMMb\r\n"))
	conn.Write([]byte("   HZM            MMMM\r\n"))
	conn.Write([]byte("   FqM            MMMM\r\n"))
	conn.Write([]byte(" __| \".        |\\dS\"qML\r\n"))
	conn.Write([]byte(" |    `.       | `' \\Zq\r\n"))
	conn.Write([]byte("_)      \\.___.,|     .'\r\n"))
	conn.Write([]byte("\\____   )MMMMMP|   .'\r\n"))
	conn.Write([]byte("     `-'       `--'\r\n"))
}

func PrintWelcomeArt(){
	fmt.Println("         _nnnn_")
	fmt.Println("        dGGGGMMb")
	fmt.Println("       @p~qp~~qMb")
	fmt.Println("       M|@||@) M|")
	fmt.Println("       @,----.JM|")
	fmt.Println("      JS^\\__/  qKL")
	fmt.Println("     dZP        qKRb")
	fmt.Println("    dZP          qKKb")
	fmt.Println("   fZP            SMMb")
	fmt.Println("   HZM            MMMM")
	fmt.Println("   FqM            MMMM")
	fmt.Println(" __| \".        |\\dS\"qML")
	fmt.Println(" |    `.       | `' \\Zq")
	fmt.Println("_)      \\.___.,|     .'")
	fmt.Println("\\____   )MMMMMP|   .'")
	fmt.Println("     `-'       `--'")
}

func TakeUserName(conn net.Conn)string{
	for {
		username,_ :=bufio.NewReader(conn).ReadSting('\n')
//removes any leading or trailing whitespace from the username string
	username :=string.TrimSpace(username)

	if  username != ""{
		return username
	}else {
		conn.Write([]byte("Enter a username to join the chat\r\nTry again: "))
	}
}
}

func TimeNow()sting{
	now :=time.Now()
	formattedTime := now.Format("[2006-01-02 15:04:05]")
	return (formattedTime)
}

