package store

import "testing"

func TestProduct_ToJson(t *testing.T) {
	product1 := Product{ID:"1", Model:"test1", Milage:ProductUINT{Uint:10}, Price:ProductUINT{Uint:177}, Status:"Sell", Brand: "TestBrand" }
	b:= product1.ToJson()
	if string(b) == `{"ID":1,"brand":"TestBrand","model":"test1","price":200,"status":"Sell","milage":120}` {
		t.Log("Test Ok", string(b))
	}
}
func TestJsonToProduct(t *testing.T) {
	js:= []byte(`{"ID":1,"brand":"TestBrand","model":"test1","price":200,"status":"Sell","milage":-120}`)
	p:=JsonToProduct(js)
	t.Log(p)
	//test uint
	js = []byte(`{"ID":1,"brand":"TestBrand","model":"test1","status":"Sell","price":-200,"milage":120}`)
	p =JsonToProduct(js)
	t.Log(p)
}