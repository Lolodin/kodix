package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
)

var id = 0
type Memory struct {
sync.Mutex
productList map[string]*Product
}
type ProductUINT struct {
	Uint uint
}

type Product struct {
	ID string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Status string `json:"status"` // *
	Price ProductUINT `json:"price"`
	Milage ProductUINT `json:"milage"`
}
type List struct {
	ProductList []*Product `json:"product_list"`
}
func NewMemory() Memory {
	m:= Memory{productList: make(map[string]*Product,5)}
	product1 := Product{ID:"1", Model:"test1", Milage:ProductUINT{Uint:10}, Price:ProductUINT{Uint:10}, Status:"Sell", Brand: "TestBrand" }
	product2 := Product{ID:"2", Model:"test2", Milage:ProductUINT{Uint:10}, Price:ProductUINT{Uint:10}, Status:"Buy", Brand: "TestBrand" }
	product3 := Product{ID:"34353", Model:"test3", Milage:ProductUINT{Uint:10}, Price:ProductUINT{Uint:10}, Status:"Sell", Brand: "TestBrand" }
	product4 := Product{ID:"4", Model:"test4", Milage:ProductUINT{Uint:10}, Price:ProductUINT{Uint:10}, Status:"Buy", Brand: "TestBrand" }
	m.Set(product1)
	m.Set(product2)
	m.Set(product3)
	m.Set(product4)
	return m
}
func(m *Memory) Set(p Product) {
	m.Lock()
	p.ID = strconv.Itoa(id)
	id++
	m.productList[p.ID] = &p
	m.Unlock()
	log.Println("Product Save", p.ID)
}
func(m  Memory) Get(id string) (*Product, error) {
	m.Lock()
	p, ok := m.productList[id]
	m.Unlock()
	if ok {
		return p, nil
	}
	log.Println("Product not found", id)
	return nil, fmt.Errorf("Product not found, id: %v", id)
}
func(m *Memory) Update(product Product) error {
	m.Lock()
	_, ok := m.productList[product.ID]
	m.Unlock()
	if ok {
		m.Lock()
		m.productList[product.ID] = &product
		m.Unlock()
		return nil
	}
	return fmt.Errorf("Product not exist")
}
func(m *Memory) Remove(id string)  {
	m.Lock()
	delete(m.productList, id)
	m.Unlock()
}

func (p Product) ToJson() ([]byte) {
	b, e :=json.Marshal(p)
	if e!= nil {
		return nil
	}
	return b
}
func JsonToProduct(b []byte) Product  {
	p:= Product{}
	err:=json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("Error convert")
	}
	return p
}

func NewList(m Memory) List {
	l:= List{}
	for _, p := range m.productList {
		l.ProductList = append(l.ProductList, p)
	}
	return l
}

func (p *ProductUINT) UnmarshalJSON(data []byte) error {
	if data[0] == 34{
		err := json.Unmarshal(data[1:len(data)-1], &p.Uint)
		if err != nil {
			return errors.New("ProductUINT: UnmarshalJSON: " + err.Error())
		}
	} else {
err := json.Unmarshal(data, &p.Uint)
if err != nil {
return errors.New("ProductUINT: UnmarshalJSON: " + err.Error())
}
}
return nil
}
func (p ProductUINT) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(p.Uint)
	return json, err
}