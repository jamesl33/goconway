package main

import (
    "bytes"
    "fmt"
    "math/rand"
    "time"
)

type Field struct {
    width int
    height int
    board [][] bool
}

func NewField(width int, height int) *Field {
    board := make([][] bool, height)

    for index := range board {
        board[index] = make([] bool, width)
    }

    return &Field{width: width, height: height, board: board}
}

func (field *Field) Set(x int, y int, value bool) {
    field.board[y][x] = value
}

func (field *Field) Alive(x int, y int) bool {
    x += field.width
    x %= field.width
    y += field.height
    y %= field.height

    return field.board[y][x]
}

func (field *Field) Next(x int, y int) bool {
    alive := 0

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if (j != 0 || i != 0) && field.Alive(x + i, y + j) {
                alive++
            }
        }
    }

    return alive == 3 || alive == 2 && field.Alive(x, y)
}

type Life struct {
    a *Field
    b *Field
    width int
    height int
}

func NewLife(width int, height int) *Life {
    a := NewField(width, height)

    for index := 0; index < (width * height / 4); index++ {
        a.Set(rand.Intn(width), rand.Intn(height), true)
    }

    return &Life {
        a: a,
        b: NewField(width, height),
        width: width,
        height: height,
    }
}

func (life *Life) Step() {
    for y := 0; y < life.height; y++ {
        for x := 0; x < life.width; x++ {
            life.b.Set(x, y, life.a.Next(x, y))
        }
    }

    life.a, life.b = life.b, life.a
}

func (life *Life) String() string {
    var buf bytes.Buffer

    for y := 0; y < life.height; y++ {
        for x := 0; x < life.width; x++ {
            b := byte(' ')

            if (life.a.Alive(x, y)) {
                b = '*'
            }

            buf.WriteByte(b)
        }

        buf.WriteByte('\n')
    }

    return buf.String()
}

func main() {
    life := NewLife(80, 24)

    for i := 0; i < 300; i++ {
        life.Step()
        fmt.Print("\x0c", life)
        time.Sleep(time.Second / 30)
    }
}
