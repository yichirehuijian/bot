package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func tcpPipe(conn *net.TCPConn){
	ipStr:=conn.RemoteAddr().String()
	defer func(){
		fmt.Printf("%v失去链接",ipStr)
		conn.Close()
	}()

	reader:=bufio.NewReader(conn)
	for{
		message,err:=reader.ReadString('\n')

		if err!=nil||err==io.EOF{
			break
		}
		fmt.Printf("服务端接受到的信息%v",message)
		time.Sleep(time.Second*3)

		msg:=conn.RemoteAddr().String()+"--服务端发送数据\n"
		b:=[]byte(msg)
		conn.Write(b)
	}
}

func main(){
	var tcpAddr *net.TCPAddr

	tcpAddr,_=net.ResolveTCPAddr("tcp","0.0.0.0:9999")

	tcpListener,_:=net.ListenTCP("tcp",tcpAddr)

	defer tcpListener.Close()
	for{
		tcpConn,err:=tcpListener.AcceptTCP()
		if err!=nil{
			fmt.Println(err)
			continue
		}
		go tcpPipe(tcpConn)
	}
}