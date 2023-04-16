package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		//handle
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//读取输入
		readString, _ := inputReader.ReadString('\n')
		trim := strings.Trim(readString, "\r\n")
		_, err := conn.Write([]byte(trim))
		if err != nil {
			log.Fatalf("错误：%v", err)
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalf("错误：%v", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
