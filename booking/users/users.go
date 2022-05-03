

package serve

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmedmahmo/learn/booking/db"
	"github.com/gorilla/mux"
)

type usersKey struct{}
type Users struct{
	l *log.Logger
}

func NewUsers(l*log.Logger) *Users {
	return &Users{l}
}

func (users *Users) Get(w http.ResponseWriter, r *http.Request)  {
	users.l.Printf("Users GET")
	lu := db.GetUsers()
	err := lu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Encode Json", http.StatusBadRequest)
		return
	}
}

// Post is a Method of Type Users that Adds users to the data package to
// Users. Post uses Middleware function and get the context from the request
// from usersKey{}
//
// See #65
func (users *Users) Post(w http.ResponseWriter, r *http.Request)  {
	users.l.Printf("Users POST")

	u := r.Context().Value(usersKey{}).(db.User)
	db.AddUser(&u)
}

func (users *Users) Put(w http.ResponseWriter, r *http.Request)  {
	users.l.Printf("Users PUT")
	w.Write([]byte("Users PUT"))
}

func (users *Users) Delete(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	
	users.l.Printf("Users DELETE %v", id)
	err = db.DeleteUser(id)
	if err != nil {
		http.Error(w, "Unable to delete Object", http.StatusBadRequest)
	}
}

func (users Users) ValidateUsersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := db.User{}
		err := u.FromJSON(r.Body)
		if err != nil {
			users.l.Println("Error deserializing User", err)
			http.Error(w, 
				"Error reading user",
				http.StatusBadRequest,
			)
			return
		}

		// Validate the User scheme 
		err = u.Validate()
		if err != nil {
			users.l.Println("Error validating User", err)
			http.Error(w,
				fmt.Sprintf("Error validating user %s", err),
				http.StatusBadRequest,
			)
			return
		}
		// Add user to the context
		ctx := context.WithValue(r.Context(), usersKey{}, u)
		r = r.WithContext(ctx)

		users.l.Printf("Middleware (%v) in Users", r.Method)
		next.ServeHTTP(w, r)
	})
}