package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ryanchew3/flexera-coding-challenge/internal/model"
)

func ReadFile(s string) map[int64][]*model.Computer {
	out := make(map[int64][]*model.Computer)

	file, err := os.Open(s)
	if err != nil {
		log.Fatalf("failed to open file")
		return nil
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var text = scanner.Text()
		var tokens = strings.Split(text, ",")

		key := toInt(tokens[1])
		val := &model.Computer{
			Id:    toInt(tokens[0]),
			AppId: toInt(tokens[2]),
			Type:  tokens[3],
		}

		_, ok := out[key]
		if !ok {
			in := []*model.Computer{val}
			out[key] = in
		} else {
			out[key] = append(out[key], val)
		}
	}

	return out
}

func toInt(s string) int64 {
	v, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		panic("this is not an int")
	}
	return v
}
