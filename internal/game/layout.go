package game

import (
	"fmt"
)

const (
	SpaceCoefficient  = 3
	VerticalPadding   = 50
	HorizontalPadding = 100
)

func NewDefaultEnemyLayout() EnemyLayout {
	return EnemyLayout{
		x:      150,
		y:      50,
		width:  500,
		height: 250,
		colNum: 6,
		rowNum: 3,
	}
}

type EnemyLayout struct {
	x      float32
	y      float32
	width  float32
	height float32
	colNum int
	rowNum int
}

func (e EnemyLayout) Generate() [][]Enemy {
	w, ws := e.optimalAttributes(e.width, e.colNum)
	h, hs := e.optimalAttributes(e.height, e.rowNum)
	fmt.Println(w)
	fmt.Println(h)

	var rows [][]Enemy
	y := e.y
	enemyFactory := NewGameEnemyFactory()
	for i := 0; i < e.rowNum; i++ {
		var columns []Enemy
		x := e.x
		for j := 0; j < e.colNum; j++ {
			e := enemyFactory.Create(x, y, w, h)
			x += w + ws
			columns = append(columns, e)
		}
		y += h + hs
		rows = append(rows, columns)
	}

	return rows
}

func (e EnemyLayout) optimalAttributes(dimension float32, elemNum int) (float32, float32) {
	d := dimension / float32(elemNum)
	s := d / SpaceCoefficient
	return d - s, s
}
