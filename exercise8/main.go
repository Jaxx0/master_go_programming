package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f6()
}

func f1() {
	file, err := os.Create("info.txt")
	if err != nil{
		log.Fatal(err)
	}
	file.Close()
}


func f2()  {
	// checking if file exists
	_, err := os.Stat("info.txt")
	// Error handling
	if err != nil{
		if os.IsNotExist(err){
			log.Fatal("File does not exist")
		}
	}

	oldPath := "info.txt"
	newPath := "data.txt"

	err = os.Rename(oldPath, newPath)
	if err != nil{
		log.Fatal(err)
	}
}

func f3()  {
	err := os.Remove("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("File removed")
}

func f4()  {
	file, err := os.Open("my_file.txt")
	if err != nil{
		log.Fatal()
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	fmt.Printf("Dara as string %s \n:", data)

}

func f5()  {
	file, err := os.Open("my_file.txt")
	if err != nil{
		log.Fatal()
	}


	defer file.Close()

	scanner := bufio.NewScanner(file)


	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil{
		log.Fatal()
	}

}

func f6()  {
	bs := []byte("The Go gopher is an iconic mascot!")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	if err != nil{
		log.Fatal()
	}

}