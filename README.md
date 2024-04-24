## 🍽️ POS Bill Generation System (Micro-Service in go lang)

Welcome to the Go-based microservice for generating restaurant bills! This service provides a web interface where customers can select items from a menu and receive an order summary with the total cost.

### 🚀 Getting Started

1. **Installation**

   Ensure you have Go installed on your machine.

   ```bash
   go get github.com/<username>/pos-bill-generation
   ```

2. **Running the Service**

   Navigate to the project directory and execute:

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080/`.

3. **Accessing the Application**

   Open your web browser and go to `http://localhost:8080/` to access the order page.

### 🍕 Features

- **Menu Display**: View a list of delicious food items with prices.
- **Order Submission**: Select items and submit your order.
- **Order Summary**: Receive an order summary with the total cost after submission.

### 🌐 Endpoints

- `/`: Displays the order page with the food menu.
- `/order`: Handles order submission and displays the order summary.

### 📂 Directory Structure

```
pos-bill-generation/
│
├── main.go           # Main application code
├── static/           # Static files (HTML, CSS, etc.)
│   └── style.css     # CSS styles for the web interface
│
└── README.md         # Documentation for the application
```

### 🍴 Customize Menu

To customize the menu, update the `initializeMenu` function in `main.go` with your desired food items and prices.

### 🤝 Contributing

Contributions are welcome! Fork the repository and submit pull requests for improvements or new features.

### 📞 Contact Information

- **Name**: Areeb Ahmed
- **Email**: areebmobile@gmail.com
- **Instagram**: [@emareeeb](https://www.instagram.com/emareeeb/)

If you have any questions or need assistance, feel free to reach out! Bon appétit! 🍽️
---

