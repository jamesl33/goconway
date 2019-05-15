/* This file is part of goconway.

Copyright (C) 2019, James Lee <jamesl33info@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

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
