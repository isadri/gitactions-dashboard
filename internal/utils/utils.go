package utils

import "io"

func Close(data io.ReadCloser) {
	log := GetLogger()

	if err := data.Close(); err != nil {
		log.Panic(err)
	}
}
