package utils

import "log"

func CheckSimpleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}