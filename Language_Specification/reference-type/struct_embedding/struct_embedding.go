package main

import "fmt"

type Notifier interface {
	notify()
}
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Struct Not Embedding
type Channel struct {
	Owner       User
	ChannelName string
}

// Struct Embedding, Inner type promotion
type Merchant struct {
	User
	ShopName string
}

func (u *User) notify() {
	fmt.Printf("Type: %T, Sending user email To %s<%s>\n", u, u.Name, u.Email)
}

// --------------------
// Polymorphic function
// --------------------
func sendNotification(n Notifier) {
	n.notify()
}
func main() {
	// Struct method, call method direct
	user := User{Name: "n", Email: "e"}
	user.notify()
	sendNotification(&user)
	// Struct Not Embedding, call method through the field owner
	channel := &Channel{Owner: user, ChannelName: "c"}
	channel.Owner.notify()
	sendNotification(&channel.Owner)
	// Struct Embedding, Inner type promotion, call method directly
	merchant := &Merchant{User: User{Name: "n", Email: "e"}, ShopName: "s"}
	merchant.notify()
	sendNotification(merchant)
}
