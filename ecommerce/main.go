package ecommerce

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

type Order struct {
	ID     int
	UserID int
	Item   string
}

type UserService struct{}

func (us *UserService) GetUser(id int) (*User, error) {
	// Simulate fetching user from a database
	return &User{ID: id, Name: "John Doe"}, nil
}

type OrderService struct {
	userService *UserService
}

func (os *OrderService) PlaceOrder(userID int, item string) (*Order, error) {
	user, err := os.userService.GetUser(userID)
	if err != nil {
		return nil, err
	}

	order := &Order{
		ID:     1, // Simulate order ID generation
		UserID: user.ID,
		Item:   item,
	}

	// Simulate storing the order in a database
	return order, nil
}

func main() {
	userService := &UserService{}
	orderService := &OrderService{userService: userService}

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		// Simulate user input
		userID := 1
		item := "Book"

		order, err := orderService.PlaceOrder(userID, item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Order placed: %+v", order)
	})

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
