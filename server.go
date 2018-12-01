package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	"regexp"
)


func sumOfDigits (value int, c chan int) {
	var sum int = value % 10
	value = value/10
	for ; value >0 ; {
		sum = sum + value%10
		value = value / 10
	}
	c <- sum
}

func main() {

	var search string
	var element string

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	var a,b, mArithmetic int	
	mArithmetic = 0


		message, _ := bufio.NewReader(conn).ReadString('.')

		fmt.Print("Serverul a primit datele:" + "\n" + message + "\n")

		messageSlice := strings.Split(message, " ")
		
		l := len(messageSlice)
		ch := make(chan int, l)
		
		for i := 1; i < len(messageSlice); i++ {

			if messageSlice[i] != "." {

				re := regexp.MustCompile("[0-9]+")
				element = messageSlice[i]
				search = re.FindAllString(element, -1)[0]

				el, err := strconv.Atoi(search)
    			if err == nil {}

    			ch <- el
			}
		}

		a = <- ch
		b = <- ch
		<- ch

		go func(int) {
			nr := 0
			c := make(chan int)
			l := len(ch)
			
			for j := 1; j <= l; j++ {

				key := <- ch
				go sumOfDigits(key, c)
				sum := <- c

				if sum >= a && sum <= b {
					mArithmetic = mArithmetic + key
					nr ++
				}
			}

			 mArithmetic = mArithmetic / nr
		
		}(mArithmetic)

		time.Sleep(5 * time.Second)
		fmt.Println("Serverul a intors raspunsul ca medie aritmetica: ", mArithmetic)
		newString := strconv.Itoa(mArithmetic)
		conn.Write([]byte(newString + "."))
		
}
	