package csvfile

import (
	"encoding/csv"
	"log"
	"orderdetails/data"
	"os"
	"strconv"

	"github.com/robfig/cron/v3"
)

func Loadcsvdata(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV file:", err)
	}

	for _, record := range records[1:] {
		customerID := record[2]
		productID := record[1]
		orderID := record[0]
		orderDate := record[6]
		quantitySold, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)
		paymentMethod := record[11]

		_, err := data.Db.Exec("insert into Customers (Customer_ID, Customer_Name, Customer_Email, Customer_Address) values (?, ?, ?, ?) ON DUPLICATE KEY UPDATE Customer_Name=VALUES(Customer_Name), Customer_Email=VALUES(Customer_Email), Customer_Address=VALUES(Customer_Address)",
			record[2], record[12], record[13], record[14])
		if err != nil {
			log.Println("Error inserting customer:", err)
		}

		_, err = data.Db.Exec("insert into Products (Product_ID, Product_Name, Category, Region) values (?, ?, ?, ?) ON DUPLICATE KEY UPDATE Product_Name=VALUES(Product_Name), Category=VALUES(Category), Region=VALUES(Region)",
			productID, record[3], record[4], record[5])
		if err != nil {
			log.Println("Error inserting product:", err)
		}

		_, err = data.Db.Exec("insert into Orders (Order_ID, Customer_ID, Order_Date, Total_Quantity, Total_Shipping_Cost, Payment_Method) values (?, ?, ?, ?, ?, ?)",
			orderID, customerID, orderDate, quantitySold, shippingCost, paymentMethod)
		if err != nil {
			log.Println("Error inserting order:", err)
		}

		_, err = data.Db.Exec("insert into Order_Items (Order_ID, Product_ID, Quantity_Sold, Unit_Price, Discount) values (?, ?, ?, ?, ?)",
			orderID, productID, quantitySold, unitPrice, discount)
		if err != nil {
			log.Println("Error inserting order item:", err)
		}
	}
}

func StartCronJob() {
	c := cron.New()

	c.AddFunc("0 0 * * *", func() {
		log.Println("Triggering daily data refresh...")
		Loadcsvdata("sales_data.csv")
	})
	c.Start()
}
