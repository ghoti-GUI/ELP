package main

import "fmt"

const NO_ROUTE int = 99999

func getLength(n int, length []int, last []int, matrice [][]int, s int) ([]int, []int) {
	//func getLength(n int, matrice [][]int, s int) ([]int, []int) {
	//n := len(matrice)
	//last := []int{9999, 99999, 99999}
	//length := []int{0, 99999, 99999}
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
	//remplacer -1 dans le matrice par NO_ROUTE
	//Nb_sommets := int(len(tab))
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
	//判断哪些点被经过了。noter si les sommets sont passe
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
		//length_de_route, route_pass = getLength(Nb_sommets, tab, next_src)
		sommets_pass[next_src] = true
	}
	fmt.Print(sommets_pass)
	return length_de_route, route_pass
}

func main() {
	//test
	tab := [][]int{
		{-1, 1, 5},
		{1, -1, 2},
		{5, 2, -1},
	}
	source := 0
	lenth := 3
	cost, road := dijistra(lenth, tab, source)
	fmt.Print("road :", road)
	fmt.Print("cost:", cost)

}
