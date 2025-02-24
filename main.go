package main

import "github.com/areeb-h/myfw"

var app = myfw.New()

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func listUsers() myfw.Paginated[User] {
	users := []User{
		{1, "Alice"}, {2, "Bob"}, {3, "Charlie"},
		{4, "David"}, {5, "Eve"},
	}
	return myfw.Paginate(users)
}

app.Get("/users", listUsers)

app.Get("/users", func() {
	return res.JSON({"message"})
})

app.Start()
