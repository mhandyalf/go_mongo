package controllers

import (
	"context"
	"go_mongo/database"
	"go_mongo/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func GetEmployees(c echo.Context) error {
	collection := database.GetDB().Collection("employees")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get employees")
	}

	employees := []models.Employee{}

	if err = cursor.All(context.TODO(), &employees); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get employees")
	}

	return c.JSON(http.StatusOK, employees)
}

// CreateEmployee membuat data karyawan baru
func CreateEmployee(c echo.Context) error {
	employee := new(models.Employee)
	if err := c.Bind(employee); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := validate.Struct(employee); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	collection := database.GetDB().Collection("employees")
	_, err := collection.InsertOne(context.TODO(), employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create employee")
	}

	return c.JSON(http.StatusCreated, employee)
}

// GetEmployee mengambil data karyawan berdasarkan ID
func GetEmployee(c echo.Context) error {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	collection := database.GetDB().Collection("employees")
	var employee models.Employee
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&employee)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Employee not found")
	}

	return c.JSON(http.StatusOK, employee)
}

// UpdateEmployee memperbarui data karyawan berdasarkan ID
func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	employee := new(models.Employee)
	if err := c.Bind(employee); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := validate.Struct(employee); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	collection := database.GetDB().Collection("employees")
	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": employee},
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update employee")
	}

	return c.JSON(http.StatusOK, employee)
}

// DeleteEmployee menghapus data karyawan berdasarkan ID
func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	collection := database.GetDB().Collection("employees")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete employee")
	}

	return c.NoContent(http.StatusNoContent)
}
