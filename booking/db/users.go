package db

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

var  ErrProductNotFound = fmt.Errorf("Product not found")

type User struct{
	ID      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
}

type Users []*User

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// Validate validates the Struct User by parsing the 
// 
func (u *User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}

// GetUsers lists the users in the users Slice
func GetUsers() Users{
	return usersList
}


func AddUser(u*User) {
	u.ID = getNextID()
	usersList = append(usersList, u)
}

func UpdateUser(id int, u*User) error {
	_, pos, err := findUser(id)
	if err != nil {
		return err
	}

	u.ID = id
	usersList[pos] = u

	return nil
}

func DeleteUser(id int) error {
	_, pos, err := findUser(id)
	if err != nil {
		return err
	}
	usersList = append(usersList[:pos], usersList[pos+1:]...)
	return nil
}

func findUser(id int) (*User, int, error) {
	for i, p := range usersList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound 
}

func getNextID() int {
	lp := usersList[len(usersList)-1]
	return lp.ID + 1
}

var usersList = Users{
	&User{
		ID: 1,
		Name: "Ahmed",
		Address: "Agnesstr. 56a, 80798 München",
	},
	&User{
		ID: 2,
		Name: "Anna",
		Address: "Bergamnnstr 15, Berlin",
	},
}

