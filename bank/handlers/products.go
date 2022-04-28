package handlers

import (
	"log"
	"net/http"
	"github.com/ahmedmahmo/learn/bank/data"
)

type Product struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost{
		p.addProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Products -X GET")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
	}
}

func (p *Product) addProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Products -X POST")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	}
	p.l.Printf("Product: %#v", prod)
	data.AddProducts(prod)
}