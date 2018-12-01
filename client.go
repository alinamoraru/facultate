package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	msg_to_send := []string{""}
	var search string



		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Inceputul intervalului: ")
		a, _ := reader.ReadString('\n')
		msg_to_send = append(msg_to_send, a)
		
		fmt.Print("Finalul intervalului: ")
		b, _ := reader.ReadString('\n')
		msg_to_send = append(msg_to_send, b)

		fmt.Print("Numarul de elemente necesare: ")
		nr, _ := reader.ReadString('\n')
		msg_to_send = append(msg_to_send, nr)

		re := regexp.MustCompile("[0-9]+")
		search = re.FindAllString(nr, -1)[0]

		number_of_elements, err := strconv.Atoi(search)
    	if err == nil {}

		for i := 1; i <= number_of_elements; i++ {

			fmt.Print("Introduceti un numar: ")

			number, _ := reader.ReadString('\n')

			msg_to_send = append(msg_to_send, number)

		}

		msg_to_send = append(msg_to_send, ".")

		newString := strings.Join(msg_to_send, " ")

		fmt.Print("Clientul a trimis datele: " + "\n" + newString + "\n")

		// send to socket
		fmt.Fprintf(conn, newString)

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('.')
		fmt.Print("Clientul a primit raspunsul: " + "\n" + message + "\n")

}