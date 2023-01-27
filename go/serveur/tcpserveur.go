package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// -----------------------------------------dijistra----------------------------------------
const NO_ROUTE int = 99999

func getLength(n int, length []int, last []int, matrice [][]int, s int) ([]int, []int) {
	for i := 0; i < n; i++ {
		if matrice[s][i] != NO_ROUTE {
			if length[s]+matrice[s][i] < length[i] {
				last[i] = s
				length[i] = length[s] + matrice[s][i]
			}
		}
	}
	return length, last
}
func allPass(m []bool) bool {
	for i := 0; i < len(m); i++ {
		if m[i] == false {
			return false
		}
	}
	return true
}

func dijistra(Nb_sommets int, tab [][]int, src int) ([]int, []int) {
	for i := 0; i < Nb_sommets; i++ {
		for j := 0; j < Nb_sommets; j++ {
			if tab[i][j] == -1 {
				tab[i][j] = NO_ROUTE
			}
		}
	}
	//le length de route est infini au d'abord sauf src est 0
	length_de_route := make([]int, Nb_sommets)
	route_pass := make([]int, Nb_sommets)
	for i := 0; i < Nb_sommets; i++ {
		length_de_route[i] = NO_ROUTE
		route_pass[i] = NO_ROUTE
	}
	length_de_route[src] = 0
	//noter si les sommets sont passe
	sommets_pass := make([]bool, Nb_sommets)
	for i := 0; i < Nb_sommets; i++ {
		sommets_pass[i] = false
	}
	var next_src int
	var min int
	for allPass(sommets_pass) == false {
		min = NO_ROUTE
		for i := 0; i < Nb_sommets; i++ {
			if min > length_de_route[i] && sommets_pass[i] == false {
				min = length_de_route[i]
				next_src = i

			}
		}
		print("source:", next_src)
		length_de_route, route_pass = getLength(Nb_sommets, length_de_route, route_pass, tab, next_src)
		sommets_pass[next_src] = true
	}
	//fmt.Print(sommets_pass)
	return length_de_route, route_pass
}

func openfile(fN string) ([][]int, int) {
	var src int
	//src = 0
	file, err := os.Open(fN) // For read access.

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	str1, _ := reader.ReadString('\n')
	arr0 := strings.Fields(str1)
	fmt.Print("arr0:", arr0)
	var n = len(arr0)
	fmt.Print("le nombre de sommet est", len(arr0))
	var arr2 = make([][]int, n)

	for j := 0; j < n; j++ {
		arr2[j] = make([]int, n)
	}
	for {
		fmt.Println("Merci d'entre le source de", fN, ":")
		fmt.Print(src)
		if src < 0 || src >= n {
			fmt.Println("source n'existe pas, merci de entrer un nombre entre 0 et", n-1)
			os.Exit(1)
		} else {
			break
		}
	}
	for j := 0; j < n; j++ {

		arr2[0][j], _ = strconv.Atoi(arr0[j])

		//fmt.Println(arr2)
	}

	for i := 1; i < n; i++ {
		str, err := reader.ReadString('\n')
		arr := strings.Fields(str)

		for j := 0; j < n; j++ {

			arr2[i][j], _ = strconv.Atoi(arr[j])

			//fmt.Println(arr2)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}

	file.Close()
	return arr2, n

}
func godijistra(lenth int, tab [][]int, source int, fileName string) {
	cost, road := dijistra(lenth, tab, source)
	fmt.Print("road :", road)
	fmt.Print("cost:", cost)
	file, err := os.OpenFile("res"+strconv.Itoa(source)+fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	defer file.Close()
	for i := 0; i < lenth; i++ {
		str = "the costfrom" + strconv.Itoa(source) + " to " + strconv.Itoa(i) + " est " + strconv.Itoa(cost[i]) + "\n"
		file.WriteString(str)
	}

	for i := 0; i < lenth; i++ {
		if i != source {
			str = "\nthe road from " + strconv.Itoa(source) + " to " + strconv.Itoa(i) + " est \n" + strconv.Itoa(i) + "<-"
			file.WriteString(str)
			for q := i; road[q] != NO_ROUTE; {
				if road[q] == source {
					str2 := strconv.Itoa(road[q])
					q = road[q]
					file.WriteString(str2)
				} else {
					str2 := strconv.Itoa(road[q]) + "<-"
					q = road[q]
					file.WriteString(str2)
				}
			}
		}
	}
	file.WriteString("\n\n\n")
}
func Child(fileName string) {
	fmt.Print("begin child")
	tab, lenth := openfile(fileName)
	if lenth <= 10 {
		for i := 0; i < lenth; i++ {
			go godijistra(lenth, tab, i, fileName)
		}
		var wg sync.WaitGroup
		wg.Add(1) // Add one to the WaitGroup counter

		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes
			time.Sleep(time.Second)
			fmt.Println("Goroutine finished")
		}()

		fmt.Println("Waiting for goroutine to finish...")
		wg.Wait() // Block until the WaitGroup counter is zero
	} else {
		for i := 0; i <= lenth/10; i++ {
			fmt.Print(i, "go................")
			for j := 0; j < 10 && j < (lenth-10*i); j++ {
				go godijistra(lenth, tab, 10*i+j, fileName)
			}
			var wg sync.WaitGroup
			wg.Add(1) // Add one to the WaitGroup counter

			go func() {
				defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes
				time.Sleep(time.Second)
				fmt.Println("Goroutine finished")
			}()

			fmt.Println("Waiting for goroutine to finish...........")
			wg.Wait() // Block until the WaitGroup counter is zero
		}
	}

	fmt.Println("Goroutine finished")
	fmt.Println("end")
	for i := 0; i < lenth; i++ {
		fmt.Println("i=", i)
		fileToRead, err := os.Open("res" + strconv.Itoa(i) + fileName)
		if err != nil {
			panic(err)
		}
		defer fileToRead.Close()
		fmt.Println("open ok")
		// Open the file to write
		fileToWrite, err := os.OpenFile("res"+fileName, os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer fileToWrite.Close()

		// Copy the contents of the file to read to the file to write
		if _, err := io.Copy(fileToWrite, fileToRead); err != nil {
			panic(err)
		}
	}

}

func SendFile(resFileName string, conn net.Conn) {
	f, err := os.Open(resFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for {
		buf := make([]byte, 2048)
		//read the file
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Println("transform end")
			break
		}
		conn.Write(buf[:n])

		fmt.Println("uploading......")
	}
	conn.Write([]byte("ok"))
}

func Handler(conn net.Conn) {
	buf := make([]byte, 2048)
	//read the information of client
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := string(buf[:n])
	//get the ip and port
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + ": name of the file transforming is--" + fileName)
	conn.Write([]byte("ok"))
	//creat the file
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//recieve the file
	buf = make([]byte, 2048)
	//for {

	fmt.Print(buf[:n])
	fmt.Print("1")
	n, err4 := conn.Read(buf)
	if err4 != nil {
		if err4 == io.EOF {
			fmt.Println("file is recieved")
			//i = NO_ROUTE
			fmt.Println(addr + ": 协程结束")
			//runtime.Goexit()
			//break

		} else {
			fmt.Println("conn.Read err", err4)
		}
	}
	revData := string(buf[:n])
	print(revData + "!!!")
	if revData == "ok" { //test
		//send the name of file
		//break //test
	} //test

	f.Write(buf[:n])
	//}

	fmt.Println("fclose")
	defer f.Close()
	fmt.Println("close")
	Child(fileName)
	//-----------------------------send file--------------------------------------------
	resFileName := "res" + fileName

	//conn2, err := net.Dial("tcp", ":8001")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer conn2.Close()
	//conn2.Write([]byte(resFileName))
	print("1")
	conn.Write([]byte(resFileName)) //test
	print("2")
	//buf2 := make([]byte, 2048)
	buf2 := make([]byte, 2048)
	//read the ifo of client
	//n2, err := conn2.Read(buf2)
	n2, err := conn.Read(buf2) //test
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	revData2 := string(buf2[:n2]) //test
	//if revData == "ok" {
	//send the name of file
	//SendFile(resFileName, conn2)
	//}
	if revData2 == "ok" { //test
		//send the name of file
		SendFile(resFileName, conn) //test
	} //test

}

func main() {
	//creat the tcp
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		//wait for connection
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		defer conn.Close()
		go Handler(conn)

	}

}
