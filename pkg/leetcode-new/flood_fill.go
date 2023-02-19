package main

import "errors"

type Cell struct {
	i, j int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	queue := NewQueue()
	visited := make([][]bool, len(image))
	for i, row := range image {
		visited[i] = make([]bool, len(row))
	}

	queue.Enqueue(Cell{sr, sc})

	for queue.Len() > 0 {
		cell, err := queue.Dequeue()
		if err != nil {
			return nil
		}
		if visited[cell.i][cell.j] {
			continue
		}

		// top
		if cell.i-1 >= 0 && image[cell.i-1][cell.j] == image[cell.i][cell.j] {
			queue.Enqueue(Cell{
				i: cell.i - 1,
				j: cell.j,
			})
		}

		// down
		if cell.i+1 < len(image) && image[cell.i+1][cell.j] == image[cell.i][cell.j] {
			queue.Enqueue(Cell{
				i: cell.i + 1,
				j: cell.j,
			})
		}

		// left
		if cell.j-1 >= 0 && image[cell.i][cell.j-1] == image[cell.i][cell.j] {
			queue.Enqueue(Cell{
				i: cell.i,
				j: cell.j - 1,
			})
		}

		// Right
		if cell.j+1 < len(image[0]) && image[cell.i][cell.j+1] == image[cell.i][cell.j] {
			queue.Enqueue(Cell{
				i: cell.i,
				j: cell.j + 1,
			})
		}

		visited[cell.i][cell.j] = true
		image[cell.i][cell.j] = color
	}

	return image
}

type Queue struct {
	cells []Cell
}

func NewQueue() *Queue {
	return &Queue{
		cells: make([]Cell, 0),
	}
}

func (q *Queue) Enqueue(cell Cell) {
	q.cells = append(q.cells, cell)
}

func (q *Queue) Dequeue() (Cell, error) {
	if q.Len() == 0 {
		return Cell{}, errors.New("empty queue")
	}
	cell := q.cells[0]
	q.cells = q.cells[1:]
	return cell, nil
}

func (q *Queue) Len() int {
	return len(q.cells)
}
