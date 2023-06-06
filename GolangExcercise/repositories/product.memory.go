package repositories

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golangexcercise/minhnq4/entities"
	"log"
	"os"
)

type ProductMemory struct {
	Products map[int]entities.Product
}

func (p *ProductMemory) GetData() {
	file, err := os.Open("C:\\Users\\Admin\\OneDrive\\Máy tính\\GolangEx\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	var jsonText string
	for scanner.Scan() {
		// do something with a line
		// fmt.Printf("line: %s\n", scanner.Text())
		jsonText = fmt.Sprintf("%s", scanner.Text())
	}
	json.Unmarshal([]byte(jsonText), &p.Products)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func (p *ProductMemory) SaveData() {
	file, err := os.OpenFile("C:\\Users\\Admin\\OneDrive\\Máy tính\\GolangEx\\data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mJson, err := json.Marshal(p.Products)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr := string(mJson)
	fmt.Println("Saving data...")
	_, errr := file.WriteString(jsonStr)
	if errr != nil {
		fmt.Println(errr)
	}

}

func (p *ProductMemory) GetAll() []*entities.Product {
	var products []*entities.Product
	for _, product := range p.Products {
		product := product
		products = append(products, &product)
	}
	return products
}
func (p *ProductMemory) GetById(id int) (*entities.Product, bool) {
	product, ok := p.Products[id]
	return &product, ok
}
func (p *ProductMemory) Create(product *entities.Product) int {
	id := p.getMaxKey() + 1
	p.Products[id] = entities.Product{ID: id, Name: product.Name, Price: product.Price}
	return id
}
func (p *ProductMemory) DeleteById(id int) {
	delete(p.Products, id)
}
func (p *ProductMemory) getMaxKey() int {
	var max int = 0
	for _, product := range p.Products {
		if max < product.ID {
			max = product.ID
		}
	}
	return max
}
