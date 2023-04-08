package runtime

import (
	"log"
	"os"
)

func init() {
	_, err := os.Stat("scripts")
	if os.IsNotExist(err) {
		log.Println("Scripts folder not found, creating one...")
		err := os.Mkdir("scripts", os.ModePerm)
		if err != nil {
			log.Println("Error creating scripts folder: ", err)
		}
	} else if err != nil {
		log.Println("Error checking scripts folder: ", err)
	}
}

func GetScripts() (files []string, err error) {
	dir, err := os.Open("scripts")
	if err != nil {
		return nil, err
	}

	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			log.Println("Error closing scripts folder: ", err)
		}
	}(dir)

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfos {
		if file.IsDir() {
			continue
		}
		files = append(files, file.Name())
	}
	return files, nil
}
