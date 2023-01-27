package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// send file to serveur
func SendFile(filePath string, conn net.Conn) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for {
		buf := make([]byte, 2048)
		//read the file
		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("transform end")
			break
		}
		//send the file to serveur
		conn.Write(buf[:n])
		fmt.Println("uploading......")
	}
	conn.Write([]byte("ok"))
}
func Handler(conn net.Conn) {
	buf := make([]byte, 2048)
	//read the information from serveur
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	resFileName := string(buf[:n])

	//get the address and port
	//addr := conn.RemoteAddr().String()
	fmt.Println(": file is stocked as--" + resFileName)
	//tell the server the connection is ok
	conn.Write([]byte("ok"))
	//creat file
	f, err := os.Create(resFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//receive the file
	buf = make([]byte, 4096)
	for {
		n, err4 := conn.Read(buf)
		//fmt.Print(buf[:n])
		fmt.Print(err4)

		if err4 != nil {
			if err4 == io.EOF {
				fmt.Println("download end")
				//i = NO_ROUTE
				fmt.Println(": process end")
				//runtime.Goexit()
				break

			} else {
				fmt.Println("conn.Read err", err4)
			}
		}
		/*fmt.Print(buf[:n])
		fmt.Print("1")*/
		//结束协程
		/*if string(buf[:n]) == "ok" {
			fmt.Println(addr + ": 协程结束")
			runtime.Goexit()
		}*/
		revData := string(buf[:n])
		print(revData + "!!!")
		if revData == "ok" {
			//send the file
			print("ok")
			break
		}
		f.Write(buf[:n])
	}
	//defer conn.Close()
	fmt.Println("fclose")
	defer f.Close()
	fmt.Println("close")

}

func main() {
	fmt.Print("please type in the path of file：")
	//store the path in str
	var str string
	fmt.Scan(&str)
	//get the information necessary
	fileInfo, err := os.Stat(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	//creat a connection with server
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//name of the file
	fileName := fileInfo.Name()
	//send the name of file to serveur
	conn.Write([]byte(fileName))

	buf := make([]byte, 2048)
	//read the answer of serveur
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf[:n])
	if revData == "ok" {
		//send the file
		SendFile(str, conn)
	}
	//----------new connect to recieve---------
	/*	listen, err := net.Listen("tcp", ":8001")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer listen.Close()
		for {
			//wait for conection
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			//handle the connection
			defer conn.Close()
			go Handler(conn)

		}
	*/

	Handler(conn)

}
