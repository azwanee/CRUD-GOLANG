package controllers

import (
	"crud-api/configs"
	"crud-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Menampilkan Semua Products
func ReadAllCategorys(c echo.Context) (err error) {
	var responses []models.CategoryResponse

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

	const readAllCategoryQuery = `
	SELECT
		id, category_name
	FROM
		categories 
	`

	rows, err := db.QueryContext(c.Request().Context(), readAllCategoryQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading all category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	for rows.Next() {
		var response models.CategoryResponse

		err = rows.Scan(
			&response.ID,
			&response.CategoryName,
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

// Menampilkan Produk By Id
func ReadDetailCategorys(c echo.Context) (err error) {
	var response models.CategoryResponse

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

	const readAllCategoryQuery = `
	SELECT
		categories.id, categories.category_name
	FROM
		categories
	WHERE
		categories.id = ?
		
		`
	// LEFT JOIN categories ON products.category_id = categories.id
	// LEFT JOIN categories ON product.category_id = category_id = categories.id

	row := db.QueryRowContext(c.Request().Context(), readAllCategoryQuery, id)

	err = row.Scan(
		&response.ID,
		&response.CategoryName,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading detail category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successss all category!",
		"page":    nil,
		"data":    response,
		"error":   nil,
	})
}

// Menambah Data Product
func CreateCategory(c echo.Context) (err error) {
	var request models.CategoryRequest

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

	const createCategoryQuery = `
	INSERT INTO categories
	(category_name)
	VALUES
	(?)
	`

	fmt.Println(request)

	_, err = db.ExecContext(c.Request().Context(), createCategoryQuery,
		request.CategoryName,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error creating data category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Succes creating data category!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})

}

// Mengubah Data Product
func UpdateCategory(c echo.Context) (err error) {
	var request models.CategoryRequest

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

	const updateCategoryQuery = `
	UPDATE categories set category_name = ?
		where id = ?
	`

	fmt.Println(request)

	_, err = db.ExecContext(c.Request().Context(), updateCategoryQuery,
		request.CategoryName,
		request.ID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error update data category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Succes update data category!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})

}

// Menghapus Data Product
func DeleteCategory(c echo.Context) (err error) {
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed Connecting to database",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failded  converting id",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	const deleteCategoryQuery = `
	DELETE
	FROM 
	categories 
	where
	id = ?`

	_, err = db.ExecContext(c.Request().Context(), deleteCategoryQuery, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failded  delete category",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully Delete category",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})
}
