package main

import (
	"errors"
	"strconv"
	"strings"
)

type Cmd struct {
	x, y    int
	r, g, b byte
}

func ParseCmd(size int, msg []byte) (Cmd, error) {
	str := string(msg)
	parts := strings.Split(str, " ")
	cmd := Cmd{}

	if len(parts) != 5 {
		return cmd, errors.New("should be of the format 'x y r g b'")
	}

	args := make([]int, 5)

	for i, p := range parts {
		a, err := strconv.Atoi(p)

		if err != nil {
			return cmd, err
		}

		if a < 0 {
			return cmd, errors.New("no arguments should be negative")
		}

		if i < 2 && a >= size {
			return cmd, errors.New("x, y should not be larger than size of canvas")
		}

		if i > 1 && a > 255 {
			return cmd, errors.New("r, g, b should not be larger than 255")
		}

		args[i] = a
	}

	cmd.x = args[0]
	cmd.y = args[1]
	cmd.r = byte(args[2])
	cmd.g = byte(args[3])
	cmd.b = byte(args[4])

	return cmd, nil
}
