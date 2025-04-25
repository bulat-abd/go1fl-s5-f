package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, datum := range dataset {
		err := dp.Parse(datum)
		if err != nil {
			log.Print(err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(result)
	}
}
