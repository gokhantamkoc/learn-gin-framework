package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const companyNotFound string = "Company not found!"
const companyExists string = "Company does exist!"

type Company struct {
	Id   int64  `json:"id"`
	Name string `json:name`
}

var companies []Company = []Company{}

func GetCompany(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateBadRequestResponse(err))
		return
	}
	for i := 0; i < len(companies); i++ {
		if companies[i].Id == id {

			c.JSON(200, CreateSuccessResponse(companies[i]))
			return
		}
	}
	c.JSON(http.StatusBadRequest, CreateBadRequestResponse(errors.New(companyNotFound)))
}

func ListCompanies(c *gin.Context) {
	c.JSON(200, CreateSuccessResponse(companies))
}

func CreateCompany(c *gin.Context) {
	var company Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, CreateBadRequestResponse(err))
		return
	}
	for i := 0; i < len(companies); i++ {
		if companies[i].Id == company.Id {
			c.JSON(http.StatusBadRequest, CreateBadRequestResponse(errors.New(companyExists)))
			return
		}
	}
	companies = append(companies, company)
	c.JSON(200, CreateSuccessResponse(company))
}

func UpdateCompany(c *gin.Context) {
	var company Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, CreateBadRequestResponse(err))
		return
	}
	for i := 0; i < len(companies); i++ {
		if companies[i].Id == company.Id {
			companies[i].Name = company.Name
			c.JSON(http.StatusOK, CreateSuccessResponse(companies[i]))
			return
		}
	}
	c.JSON(http.StatusBadRequest, CreateBadRequestResponse(errors.New(companyNotFound)))
}
