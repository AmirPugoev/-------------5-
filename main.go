package main

import (
"fmt"
"os"
)

type Furniture struct {
Name string
Height int
Width int
Depth int
Material string
}
type Appliances struct { //бытовая техника 
	Name string
	Height int
	Width int
	Depth int
	Material string
	}
func writeToFile(f Furniture) error {
fileName := f.Name + ".txt"
file, err := os.Create(fileName)
if err != nil {
return err
}
defer file.Close()

_, err = fmt.Fprintf(file, "Высота: %d см\nШирина: %d см\nГлубина: %d см\nМатериал: %s", f.Height, f.Width, f.Depth, f.Material)
if err != nil {
return err
}
return nil
}
func writeToFile_1(v Appliances) error {
	fileName := v.Name + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
	return err
	}
	defer file.Close()
	
	_, err = fmt.Fprintf(file, "Высота: %d см\nШирина: %d см\nГлубина: %d см\nМатериал: %s", v.Height, v.Width, v.Depth, .Material)
	if err != nil {
	return err
	}
	return nil
	}
func main() {
furniture := []Furniture{
Furniture{Name: "Кровать", Height: 200, Width: 160, Depth: 220, Material: "Дерево"},
Furniture{Name: "Шкаф", Height: 220, Width: 120, Depth: 50, Material: "ДСП"},
Furniture{Name: "Стол", Height: 75, Width: 120, Depth: 70, Material: "Стекло"},
}
appliances:= []Appliances{
	Appliances{Name: "тостер", Height: 200, Width: 160, Depth: 220, Material: "алюминий"},
	Appliances{Name: "телевизор", Height: 220, Width: 120, Depth: 50, Material: "пластик"},
	Appliances{Name: "холодильник", Height: 75, Width: 120, Depth: 70, Material: "метал"},
}

for _, f := range furniture {
err := writeToFile(f)
if err != nil {fmt.Println(err)}
for _, v := range appliances {
	err := writeToFile(Furniture(v)) 
	if err != nil { fmt.Println(err)}
}
	
}
}


