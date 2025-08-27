package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
}

// Конструктор для создания нового продукта
func NewProduct(id int, name string, quantity int) Product {
	return Product{
		ID:       id,
		Name:     name,
		Quantity: quantity,
	}
}

// Метод для возврата продукта (геттер)
func (p *Product) ReturnProduct() *Product {
	return p
}

type Storage struct {
	Products map[int]Product
}

type Warehouse interface {
	AddProduct(product Product) error
	UpdateQuantity(productID int, quantity int) error
}

// метод добавления нового товара на склад
func (st *Storage) AddProduct(product Product) error {
	if st.Products == nil {
		st.Products = make(map[int]Product)
	}

	if _, ok := st.Products[product.ID]; ok {
		return fmt.Errorf("товар с ID %d уже существует на складе", product.ID)
	}

	st.Products[product.ID] = product
	return nil
}

// Метод для получения продукта по ID
func (st *Storage) GetProduct(productID int) (Product, error) {
	product, ok := st.Products[productID]
	if !ok {
		return Product{}, fmt.Errorf("товар с ID %d не найден на складе", productID)
	}
	return product, nil
}

// метод изменения количества товара на складе
func (st *Storage) UpdateQuantity(productID int, quantity int) error {
	product, ok := st.Products[productID]
	if !ok {
		return fmt.Errorf("товар с ID %d не найден на складе", productID)
	}

	if product.Quantity+quantity < 0 {
		return fmt.Errorf("количество товара не может быть отрицательным")
	}

	product.Quantity += quantity
	st.Products[productID] = product
	return nil
}

/////////////
func main() {
	// создание содержимого склада
	storage := &Storage{
		Products: make(map[int]Product),
	}

	// создание продуктов через конструктор
	laptop := NewProduct(1, "Ноутбук", 10)
	phone := NewProduct(2, "Телефон", 20)
	tablet := NewProduct(3, "Планшет", 15)

	// добавление товаров на склад через метод AddProduct
	err := storage.AddProduct(laptop)
	if err != nil {
		fmt.Println("Ошибка при добавлении товара:", err)
		return
	}

	err = storage.AddProduct(phone)
	if err != nil {
		fmt.Println("Ошибка при добавлении товара:", err)
		return
	}
	err = storage.AddProduct(tablet)
	if err != nil {
		fmt.Println("Ошибка при добавлении товара:", err)
		return
	}

	// получение продукта через метод ReturnProduct (из самого продукта)
	returnedLaptop := laptop.ReturnProduct()
	fmt.Println("Продукт через ReturnProduct:", returnedLaptop)

	// получение продукта со склада через GetProduct
	storedProduct, err := storage.GetProduct(1)
	if err != nil {
		fmt.Println("Ошибка при получении товара:", err)
		return
	}
	fmt.Println("Продукт со склада:", storedProduct)

	// изменение количества товаров на складе
	err = storage.UpdateQuantity(1, -5) // продали 5 ноутбуков
	if err != nil {
		fmt.Println("Ошибка при изменении количества товара:", err)
		return
	}

	// проверка изменений
	updatedProduct, err := storage.GetProduct(1)
	if err != nil {
		fmt.Println("Ошибка при получении товара:", err)
		return
	}
	fmt.Println("Обновленный продукт:", updatedProduct)

	// вывод всех продуктов на складе
	fmt.Println("\nВсе продукты на складе:")
	for id, product := range storage.Products {
		fmt.Printf("ID: %d, %s - %d шт.\n", id, product.Name, product.Quantity)
	}
}

/*
	// добавление товаров на склад
	err := storage.AddProduct(Product{ID: 1, Name: "Ноутбук", Quantity: 10})
	if err != nil {
		fmt.Print("Ошибка при добавлении товара: ", err)
		return
	}

	fmt.Println(storage.Products[1]) // {1 Ноутбук 10}

	// изменение количества товаров на складе
	err = storage.UpdateQuantity(1, -5) // продали 5 ноутбуков
	if err != nil {
		fmt.Print("Ошибка при изменении количества товара: ", err)
		return
	}

	fmt.Println(storage.Products[1]) // {1 Ноутбук 5}
}
*/
