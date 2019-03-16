package main

import "fmt"

type Cell struct {
	cost int
	quantity int
	valid bool
	visited bool
	factory int
	warehouse int
}

type Path struct {
	cells []Cell
}

func (p *Path) append(c Cell) {
	p.cells = append(p.cells, c)
}



func steppingStone(problem *[][] Cell, m int, n int) {


	for i:=0; i < m; i++ {
		for j:=0; j < n; j++ {
			if (*problem)[i][j].quantity == 0 { //check if empty
				fmt.Println()
				marginalCost((*problem)[i][j], problem, m, n)
				fmt.Println()
			} else {
				continue // if non-empty check next cell
			}
		}

	}

}

func marginalCost(cell Cell ,problem *[][]Cell, m int, n int) Path { //ADD MARGINAL COST TO THE PATH ??
	//copy problem array
	temp := make([][]Cell, len(*problem))
	for i := range *problem {
		temp[i] = make([] Cell, len((*problem)[i]))
		copy(temp[i], (*problem)[i])
	}

	path := Path{make([]Cell, 0)}
	path.append(cell)


	for { //break when closed
		closed := true
		for i:=0; i < m; i++ {
			for j := 0; j < n; j++ {
				if !hasNeighbours(temp, cell, temp[i][j], m, n) && temp[i][j].valid{
					temp[i][j].visited = true
					temp[i][j].valid = false
					closed = false
				}
			}
		}

		if closed {
			break
		}

	}

	temp[cell.factory][cell.warehouse].valid = true
	for i:=0; i < m; i++ {
		for j := 0; j < n; j++ {
			if temp[i][j].valid{
				fmt.Printf("%d-%d ",temp[i][j].quantity, temp[i][j].cost)
			}
		}
	}

	return *new(Path)

}

func hasNeighbours(temp [][]Cell, start Cell, c Cell, m int, n int) bool{
	hasHorizontal := false
	hasVertical := false

	if c.quantity == 0 {
		return false
	}
	//horizontal
	if start.factory == c.factory && c.warehouse != start.warehouse {
		hasHorizontal = true
	} else {
		for i := 0; i < n; i++ {
			if temp[c.factory][i].quantity > 0 && i != c.warehouse && temp[c.factory][i].valid {
				hasHorizontal = true
				break
			}
		}
	}

	//vertical
	if start.warehouse == c.warehouse && c.factory != start.factory {
		hasVertical = true
	} else {
		for j := 0; j < n; j++ {
			if temp[j][c.warehouse].quantity > 0 && j != c.factory && temp[j][c.warehouse].valid{ //CHECK IF SAME COLUMN ??
				hasVertical = true
				break
			}
		}
	}

	return hasVertical && hasHorizontal


}

func main() {
	problem := [][]Cell{{Cell{6,0,true,false,0,0},
	Cell{8,25,true,false,0,1},
	Cell{10,125,true,false,0,2}},

	{Cell{7,0,true,false,1,0},
	Cell{11,0,true,false,1,1},
	Cell{11,175,true,false,1,2}},

	{Cell{4,200,true,false,2,0},
	Cell{5,75,true,false,2,1},
	Cell{12,0,true,false,2,2}} }

	bool := hasNeighbours( problem , problem[0][0],problem[2][0], 3,3)
	fmt.Println(bool)

	steppingStone(&problem,3,3)



}
