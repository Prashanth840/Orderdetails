package repository

import (
	"orderdetails/data"
	"orderdetails/models"
)

func TotalRevenue(input models.Input) (float64, error) {
	var totalRevenue float64
	if err := data.Db.QueryRow(`
	select sum((unit_price - discount) * quantity_sold) from order_items ot
	join orders os ON ot.order_id = os.order_id 
	where order_date between ? AND ?
`, input.Startdate, input.Enddate).Scan(&totalRevenue); err != nil {
		return totalRevenue, err
	}
	return totalRevenue, nil
}

func TotalRevenuebyproduct(input models.Input) ([]models.TotalRevenuebyproduct, error) {
	var result []models.TotalRevenuebyproduct
	rows, err := data.Db.Query(`
		SELECT p.Product_Name, SUM((oi.unit_price - oi.discount) * oi.quantity_sold) AS revenue
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		JOIN orders o ON oi.order_id = o.order_id
		WHERE o.order_date BETWEEN ? AND ?
		GROUP BY p.Product_Name
		ORDER BY revenue DESC
	`, input.Startdate, input.Enddate)
	if err != nil {

		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var content models.TotalRevenuebyproduct
		err := rows.Scan(&content.Productname, &content.Revenue)
		if err != nil {
			return result, err
		}
		result = append(result, content)
	}
	return result, nil

}

func TotalRevenueByCategory(input models.Input) ([]models.TotalRevenueByCategory, error) {
	var result []models.TotalRevenueByCategory
	rows, err := data.Db.Query(`
		SELECT p.Category, SUM((oi.unit_price - oi.discount) * oi.quantity_sold) AS revenue
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		JOIN orders o ON oi.order_id = o.order_id
		WHERE o.order_date BETWEEN ? AND ?
		GROUP BY p.Category
		ORDER BY revenue DESC
	`, input.Startdate, input.Enddate)
	if err != nil {

		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		var content models.TotalRevenueByCategory
		err := rows.Scan(&content.Category, &content.Revenue)
		if err != nil {
			return result, err
		}
		result = append(result, content)
	}
	return result, nil
}
