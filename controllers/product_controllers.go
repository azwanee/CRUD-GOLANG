package controllers

import (
	"crud-api/configs"
	"crud-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Read All Products
func ReadAllProducts(c echo.Context) (err error) {
	var responses []models.ProductResponse

	// Buat koneksi ke database
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const readAllProductsQuery = `
	SELECT
		id, product_name, price, category, description
	FROM
		products
	`

	rows, err := db.QueryContext(c.Request().Context(), readAllProductsQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading all products!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	for rows.Next() {
		var response models.ProductResponse

		err = rows.Scan(
			&response.ID,
			&response.ProductName,
			&response.Price,
			&response.Category,
			&response.Description,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Error reading all products!",
				"page":    nil,
				"data":    nil,
				"error":   err.Error(),
			})
		}

		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success reading all products!",
		"page":    nil,
		"data":    responses,
		"error":   nil,
	})
}

// Read Detail Product
func ReadDetailProducts(c echo.Context) (err error) {
	var response models.ProductResponse

	// Buat koneksi ke database
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing parameter to integer!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	const readDetailProductQuery = `
	SELECT
		products.id, products.product_name, products.price, categories.category_name, products.description
	FROM
		products
		INNER JOIN categories ON products.category_id = categories.id
	WHERE
		products.id = ?
	`

	row := db.QueryRowContext(c.Request().Context(), readDetailProductQuery, id)

	err = row.Scan(
		&response.ID,
		&response.ProductName,
		&response.Price,
		&response.Category,
		&response.Description,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading detail product!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success reading all products!",
		"page":    nil,
		"data":    response,
		"error":   nil,
	})
}

// Create Product
func CreateProduct(c echo.Context) (err error) {
	var request models.ProductRequest

	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error binding request!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const createProductQuery = `
	INSERT INTO products
	(product_name, price, category, description)
	VALUES
	(?, ?, ?, ?)
	`

	_, err = db.ExecContext(c.Request().Context(), createProductQuery,
		request.ProductName,
		request.Price,
		request.Category,
		request.Description,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error creating data product!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success creating data product!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})
}

// Update Product
func UpdateProduct(c echo.Context) (err error) {
	var request models.ProductRequest

	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{

			"message": "Error binding request",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	db, err := configs.ConnectDatabase()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{

			"message": "Error binding request",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const updateProductQuery = `
	UPDATE products set product_name = ?,
	 price = ?,
	  category = ?,
	   description = ?
	    where id = ?
	`

	fmt.Println(request)

	_, err = db.ExecContext(c.Request().Context(), updateProductQuery,
		request.ProductName,
		request.Price,
		request.Category,
		request.Description,
		request.ID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error update data product!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Succes update data product!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})

}

// Delete Product
func DeleteProduct(c echo.Context) (err error) {
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed something id!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	const DeleteProductQuery = `
	DELETE FROM
	products
	WHERE
	id = ?
	`
	_, err = db.ExecContext(c.Request().Context(), DeleteProductQuery, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed something id!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"massage": "Success Delete Form Product",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})
}
