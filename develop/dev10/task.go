package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type Client struct {
	host     string
	port     string
	connTime time.Duration
}

func NewClient(h, p, cTime string) *Client {
	t, err := strconv.Atoi(cTime)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{
		host:     h,
		port:     p,
		connTime: time.Duration(t) * time.Second,
	}
}

func dial(client *Client, ctx context.Context, cancel context.CancelFunc) {
	conn, err := net.DialTimeout("tcp", client.host+":"+client.port, client.connTime)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			_, err := fmt.Fprintf(conn, "time is up")
			if err != nil {
				log.Print(err)
			}
			log.Println("time is up...")
			err = conn.Close()
			if err != nil {
				log.Print(err)
			}
			return
		default:
			rd := bufio.NewReader(os.Stdin)
			fmt.Print("message: ")
			text, err := rd.ReadString('\n')
			if err != nil {
				log.Print("read error...")
			}
			_, err = fmt.Fprintf(conn, text+"\n")
			if err != nil {
				log.Print(err)
			}
			fb, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Println("from server :" + fb)
		}
	}
}

func main() {
	t := flag.String("timeout", "10", "time on work with server")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("incorrect soccet")
	}
	port := args[1]
	host := args[0]

	client := NewClient(host, port, *t)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, client.connTime)

	dial(client, ctx, cancel)

	fmt.Println("Connection closed")
}



// func myHandle(conn net.Conn) {
// 	for {
// 		message, _ := bufio.NewReader(conn).ReadString('\n')
// 		switch message {
// 		case "time is up":
// 			return
// 		default:
// 			fmt.Print("Message Received:", message)
// 			newMessage := strings.ToUpper(message)
// 			_, err := conn.Write([]byte(time.Now().String() + " " + newMessage + "\n"))
// 			if err != nil {
// 				log.Print(err)
// 			}
// 		}
// 	}
// }

// func main() {
// 	ln, _ := net.Listen("tcp", "localhost:3000")
// 	conn, _ := ln.Accept()

// 	myHandle(conn)

// 	err := ln.Close()
// 	if err != nil {
// 		log.Print(err)
// 	}
// }