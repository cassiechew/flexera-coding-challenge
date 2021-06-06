package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"errors"

	"github.com/ryanchew3/flexera-coding-challenge/internal/model"
)

// ReadFile reads in the input file and groups user machines by userids
func ReadFile(s string) (map[int64][]*model.Computer, error) {
	out := make(map[int64][]*model.Computer)

	file, err := os.Open(s)
	if err != nil {
		log.Print("failed to open file")
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		var text = scanner.Text()
		var tokens = strings.Split(text, ",")

		key, e := strconv.ParseInt(tokens[1], 10, 64)
		if e != nil {
			return nil, e
		}

		id, e := strconv.ParseInt(tokens[0], 10, 64)
		if e != nil {
			return nil, e
		}

		appId, e := strconv.ParseInt(tokens[2], 10, 64)
		if e != nil {
			return nil, e
		}

		val := &model.Computer{
			Id:    id,
			AppId: appId,
			Type:  tokens[3],
		}

		_, ok := out[key]
		if !ok {
			out[key] = []*model.Computer{val}
		} else {
			out[key] = append(out[key], val)
		}
	}

	return out, nil
}
