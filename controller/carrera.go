package controller

import (
    "github.com/labstack/echo"
    "github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/models/institucional"
    "net/http"
)

func GetCarreras(c echo.Context) error {
    // Get connection database
    db:= config.GetConnection()
    defer db.Close()

    // Execute operations
    carreras := make([]institucional.Carrera,0)
    if err := db.Find(&carreras).Error; err != nil {
        return  err
    }

    // Return final data
    return c.JSON(http.StatusOK,carreras)
}

func CreateCarrera(c echo.Context) error {
    // Get data json
    carrera := institucional.Carrera{}
    if err := c.Bind(&carrera); err != nil {
        return err
    }

    // Get connection database
    db:= config.GetConnection()
    defer db.Close()

    // Execute operations
    if err := db.Create(&carrera).Error; err != nil {
        return err
    }

    // Return final data
    return c.JSON(http.StatusCreated,carrera)
}

func UpdateCarrera(c echo.Context) error {
    // Get data json
    carrera := institucional.Carrera{}
    if err := c.Bind(&carrera); err != nil {
        return err
    }

    // Get connection database
    db:= config.GetConnection()
    defer db.Close()

    // Execute operations
    if err := db.Model(&carrera).Updates(carrera).Error; err != nil {
        return err
    }

    // Return final data
    return c.JSON(http.StatusOK,carrera)
}

func DeleteCarrera(c echo.Context) error {
    // Get data json
    carrera := institucional.Carrera{}
    if err := c.Bind(&carrera); err != nil {
        return err
    }

    // Get connection database
    db:= config.GetConnection()
    defer db.Close()

    // Execute operations
    if err := db.Delete(&carrera).Error; err != nil {
        return err
    }

    // Return final data
    return c.JSON(http.StatusNoContent,carrera)
}