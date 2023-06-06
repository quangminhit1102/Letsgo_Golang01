package main

import (
	"bufio"
	"fmt"
	"golangexcercise/minhnq4/entities"
	"golangexcercise/minhnq4/repositories"
	"os"
	"strconv"
)

var productRepo repositories.ProductRepo = &repositories.ProductMemory{Products: map[int]entities.Product{}}

func main() {
	productRepo.GetData()
	for {
		fmt.Println("===================================")
		fmt.Println("1. View All product")
		fmt.Println("2. Get by Id")
		fmt.Println("3. Create product")
		fmt.Println("4. Delete by Id")
		fmt.Println("5. Exit")
		fmt.Print("Enter your action:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		switch text {
		case "1":
			fmt.Println("View All Product!")
			viewAll()
		case "2":
			fmt.Println("Get product by Id!")
			getById()
		case "3":
			create()
			fmt.Println("Create product!")
		case "4":
			fmt.Println("Delete by Id!")
			deleteById()
		case "5":
			productRepo.SaveData()
			fmt.Println("Exiting ...")
			return
		}
	}
}
func viewAll() {
	products := productRepo.GetAll()
	if len(products) == 0 {
		fmt.Println("You have no products!")
		return
	}
	for _, product := range products {
		product := product
		fmt.Println(product)
	}
}
func getById() {
	fmt.Print("Enter product Id:")
	idStr := readText()
	idInt, _ := strconv.Atoi(idStr)
	product, ok := productRepo.GetById(idInt)
	if !ok {
		fmt.Println("Product not found!")
	} else {
		fmt.Println(product)
	}
}
func create() {
	fmt.Print("Enter name:")
	name := readText()
	fmt.Print("Enter Price:")
	priceStr := readText()
	priceInt, err := strconv.Atoi(priceStr)
	if err != nil {
		fmt.Print("Invalid Price!")
	}
	productRepo.Create(&entities.Product{Name: name, Price: priceInt})
}
func deleteById() {
	fmt.Print("Enter product Id:")
	idStr := readText()
	idInt, _ := strconv.Atoi(idStr)
	_, ok := productRepo.GetById(idInt)
	if !ok {
		fmt.Println("Product not found!")
	} else {
		productRepo.DeleteById(idInt)
	}
}

func readText() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
