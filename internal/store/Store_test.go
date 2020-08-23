package store

import (
	"strconv"
	"testing"
)

func TestStore_(t *testing.T) {
	store := NewMemory()

	for i:=0; i<4; i++ {
		p, _:=store.Get(strconv.Itoa(i))
		if p.ID != strconv.Itoa(i) {
			t.Error("invalid id", p.ID)
		}
		t.Log(p)
	}
	t.Log("TestDelete")
	store.Remove(2)
	for i:=0; i<4; i++ {
		p, e:=store.Get(i)
		if e!=nil {
			t.Log(e)
			t.Log("TestDeleteOk")
			break
		}
		if p.ID != strconv.Itoa(i) {
			t.Error("not valid id")
		}
		t.Log(p)
	}
	t.Log("TestUpdate")
	productUpdate := Product{ID:"3", Model:"Update", Milage:ProductUINT{Uint:10}, Price:ProductUINT{Uint:10}, Status:"Buy", Brand: "TestUpdate" }
    p1,_:=store.Get("3")
	store.Update(productUpdate)
    p2,_:=store.Get("3")
    if p1.ID == p2.ID {
    	t.Log("test update OK")
	}
}
