package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("storage/mainfile/data.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("File ni ochishda xatolik yuzaga keldi: ", err)
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Inputni o'qishda xatolik yuzaga keldi: ", err)
		log.Fatal(err)
	}
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Faylga yozishda xatolik yuzaga keldi: ", err)
		log.Fatal(err)
	}

	nameFile := fmt.Sprintf("storage/changes/data_%s.txt", time.Now().Format("2006-01-02_154813"))
	newfile, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Yangi file ochishda xatolik yuzaga keldi: ", err)
		log.Fatal(err)
	}
	defer newfile.Close()
	mainfile, err := os.Open("storage/mainfile/data.txt")
	if err != nil {
		fmt.Println("Failed to open main file: ", err)
		log.Fatal(err)
	}


	_, err = io.Copy(newfile, mainfile)
	if err != nil {
		log.Fatal("File ni copy qilishda xatolik yuzaga keldi:  ", err)
	}


}
