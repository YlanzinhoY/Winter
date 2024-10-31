package services

import (
	"fmt"
	"log"
	"os"
)

type Scaffold struct {
	DirName string
	Path    string
	Lang    string
}

func (s *Scaffold) Scaffold() {
	scaffold := []string{"config", "controller", "dto", "enum", "exceptions", "model", "repository", "service", "util"}
	fmt.Println("Building Scaffold...")
	for _, dir := range scaffold {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", s.Path, dir), os.ModePerm)
		if err != nil {
			log.Fatal("erro ao criar os diretorios")
			return
		}
		fmt.Printf("Creating Folder %s\n", dir)
		fmt.Println("Scaffold Finish")
	}

}
