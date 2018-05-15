package main

import (
	"net"
	"fmt"
	"time"
	"strings"
	"io"
	"bytes"
	"math/rand"
	"sync"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMTTER 	   = '\t'
)

var wg sync.WaitGroup

// socket api 服务端 客户端通信
func main()  {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

// server

func serverGo()  {
	defer wg.Done()
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		fmt.Println(fmt.Sprintf("Listener Error: %s:", err.Error()))
		return
	}
	defer listener.Close()
	fmt.Println(fmt.Sprintf("Got listener for the server.(local address:%s)", listener.Addr()))
	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accpet Error: %s", err)
		}
		printServerLog("Establish ed a connection with a clinet application. (remote addredd: %s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn)  {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side")
			} else {
				printServerLog("Read Error: %s", err)
			}
			break
		}
		printServerLog("Received request: %s.", strReq)
		//
		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err:= write(conn, err.Error())
			printServerLog("Send error message(writeen %d bytes): %s", n, err)
			continue
		}
		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("The cube root of %d bytes %f.", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("Write Error: %s", err)
		}
		printServerLog("Sent response (writeen %d bytes): %s", n, respMsg)
	}
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMTTER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMTTER)
	return conn.Write(buffer.Bytes())
}

func printLog(role string, sn int, format string, args ...interface{})  {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args))
}

func printServerLog(format string, args ...interface{})  {
	printLog("Server", 0, format, args)
}

func printClientLog(sn int, format string, args ...interface{})  {
	printLog("Client", sn, format, args)
}

func strToInt32(strReq string) (int, error) {
	return 100, nil
}

func cbrt(intReq int) float64 {
	return 1.11
}



//client

func clientGo(id int)  {
	wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printClientLog(id, "Dial Error: %s", err)
		return
	}
	defer conn.Close()
	printClientLog(id, "Connected to server.(remote address: %s, local address: %s)", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

	requestNumber := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))

	// 发送数据
	for i:=0; i<requestNumber; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "Write Error: %s", err)
			continue
		}
		printClientLog(id, "Sent Request (writeen %d bytes): %d.", n, req)
	}

	// 接收数据
	for j:=0; j<requestNumber; j++ {
		strResp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientLog(id, "The connection is close by another side.")
			} else {
				printClientLog(id, "Read Error: %s", err)
			}
			break
		}
		printClientLog(id, "Received response: %s.", strResp)
	}
}