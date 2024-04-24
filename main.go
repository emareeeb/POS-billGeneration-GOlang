package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Food struct {
	Name   string
	Charge int
	Next   *Food
}

var (
	foodMenu  *Food
	totalCost int
)

func main() {
	foodMenu = initializeMenu() // Initialize the food menu

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/static/", staticHandler)

	fmt.Println("Starting server on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Restaurant Order</title>
				<style>
					body {
						font-family: 'Arial', sans-serif;
						background-color: #f5f5f5;
						text-align: center;
					}
					.container {
						margin-top: 50px;
						background-color: #fff;
						border-radius: 10px;
						padding: 20px;
						box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
						width: 60%;
						margin: 0 auto;
					}
					h1, h2 {
						color: #333;
					}
					form {
						margin-top: 20px;
					}
					label {
						font-size: 18px;
						margin-right: 10px;
					}
				
					input[type="text"] {
						padding: 10px;
						font-size: 12px;
						border: 2px solid grey;
						border-radius: 17px;
						width: 200px;
					}
					.menu-item {
						display: flex;
						align-items: center;
						padding: 10px 0;
						border-bottom: 1px solid #ccc;
						margin-right : 50px;
					}
					.menu-item label {
						flex-grow: 1;
						margin-right: 10px;
					}
					input[type="submit"] {
						padding: 10px 20px;
						font-size: 15px;
						background-color: #ff6347;
						color: #fff;
						border: none;
						border-radius: 30px;
						cursor: pointer;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Welcome to Namma Bengaluru Bites!</h1>
					<form action="/order" method="post">
						<label for="name">Please enter your name</label><br>
						<input type="text" id="name" name="name" required>
						<h2>Menu:</h2>
		`)
		displayMenuHTML(w, foodMenu) // Display menu options
		fmt.Fprintf(w, `
						<input type="submit" value="Submit Order">
					</form>
				</div>
			</body>
			</html>
		`)
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.Form.Get("name")
		if name == "" {
			http.Error(w, "Please enter a name.", http.StatusBadRequest)
			return
		}

		// Prepare the order summary HTML
		var orderSummary strings.Builder
		orderSummary.WriteString(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Order Summary</title>
				<style>
					body {
						font-family: 'Arial', sans-serif;
						background-color: #f5f5f5;
						text-align: center;
					}
					.container {
						margin-top: 50px;
						background-color: #fff;
						border-radius: 10px;
						padding: 20px;
						box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
						width: 60%;
						margin: 0 auto;
					}
					h1, h2 {
						color: #333;
					}
					ul {
						list-style-type: none;
						padding: 0;
						text-align: left;
					}
					li {
						font-size: 16px;
						margin-bottom: 10px;
						display: flex;
						justify-content: space-between;
						align-items: center;
						border-bottom: 1px solid #ddd; /* Add border-bottom for separation */
						padding-bottom: 5px; /* Optional: add padding below each item */
					}
					li:last-child {
						border-bottom: none; /* Remove bottom border for the last item */
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Hello, `)
		orderSummary.WriteString(name) // Append the customer's name
		orderSummary.WriteString(`!</h1> <!-- Displaying the customer's name -->
					<h2>Order Summary:</h2>
					<ul>
		`)

		// Process selected items
		totalCost := 0
		items := r.Form["items"]
		for _, itemStr := range items {
			itemNum, err := strconv.Atoi(itemStr)
			if err != nil || itemNum <= 0 || itemNum > countMenuItems(foodMenu) {
				orderSummary.WriteString(fmt.Sprintf("<li>Invalid item number: %s</li>", itemStr))
				continue
			}
			item := getItemByNumber(foodMenu, itemNum)
			if item != nil {
				// Format each item in the bill with name on the left and price on the right
				orderSummary.WriteString(fmt.Sprintf(`
					<li>
						<span>%s</span>
						<span>Rs %d/-</span>
					</li>
				`, item.Name, item.Charge))
				totalCost += item.Charge
			} else {
				orderSummary.WriteString(fmt.Sprintf("<li>Item not found for number: %s</li>", itemStr))
			}
		}

		// Complete the HTML response with total cost
		orderSummary.WriteString(fmt.Sprintf(`
					</ul>
					<h2 style="margin-top: 20px;">Total Cost: Rs %d</h2>
				</div>
			</body>
			</html>
		`, totalCost))

		// Send the response
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(orderSummary.String())) // Send the complete HTML response
	}
}

func displayMenuHTML(w http.ResponseWriter, menu *Food) {
	fmt.Fprintf(w, `<ul style="text-align: left;">`)
	for i, current := 1, menu; current != nil; current = current.Next {
		fmt.Fprintf(w, `
			<li class="menu-item" style="background-color: #f9f9f9; padding: 10px; border-radius: 5px;">
				<input type="checkbox" id="item%d" name="items" value="%d">
				<label for="item%d">%s</label>
				<span>Rs %d/-</span>
			</li>
		`, i, i, i, current.Name, current.Charge)
		i++
	}
	fmt.Fprintf(w, `</ul>`)
}

func initializeMenu() *Food {
	idli := createFood("Idli", 50)
	dosa := createFood("Dosa", 80)
	vada := createFood("Vada", 60)
	upma := createFood("Upma", 70)
	poha := createFood("Poha", 65)
	choleBhature := createFood("Chole-Bhature", 120)
	choleKulche := createFood("Chole-Kulche", 100)
	paratha := createFood("Paratha", 90)

	idli.Next = dosa
	dosa.Next = vada
	vada.Next = upma
	upma.Next = poha
	poha.Next = choleBhature
	choleBhature.Next = choleKulche
	choleKulche.Next = paratha

	return idli
}

func createFood(name string, charge int) *Food {
	return &Food{
		Name:   name,
		Charge: charge,
		Next:   nil,
	}
}

func countMenuItems(menu *Food) int {
	count := 0
	for current := menu; current != nil; current = current.Next {
		count++
	}
	return count
}

func getItemByNumber(menu *Food, num int) *Food {
	i := 1
	for current := menu; current != nil; current = current.Next {
		if i == num {
			return current
		}
		i++
	}
	return nil
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
