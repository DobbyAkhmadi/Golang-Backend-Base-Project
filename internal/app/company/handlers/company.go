package handlers

import (
	"backend/internal/app/company/models"
	"backend/internal/app/company/repository"
	"backend/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

var companyRepository repository.CompanyRepository // Declare a variable `companyRepository` of type `repository.CompanyRepository`

func init() {
	// Initialize the `companyRepository` variable by assigning a new instance of `repository.CompanyRepository`
	companyRepository = repository.NewCompanyRepository()
}

// GetAllCompanies godoc
// @Summary      GetAllCompanies gets all repository information
// @Description  GetAllCompanies gets all repository information
// @Tags         Company
// @Accept       json
// @Produce      json
// @Request
// @Success      200  {object} 	utils.Response
// @Security 	 oauth2[identity_api]
// @Router       /api/v1/company [get]
func GetAllCompanies(c *fiber.Ctx) error {
	// Retrieve all companies from the company repository
	companies := companyRepository.FindAll()

	// Prepare the response object
	response := utils.Response{
		Code:    http.StatusOK,
		Body:    companies,
		Title:   "GetAllCompanies",
		Message: "All Companies",
	}

	// Return the response as JSON
	return c.Status(response.Code).JSON(response)
}

// GetSingleCompany godoc
// @Summary      GetSingleCompany Gets single company information
// @Description  GetSingleCompany Gets single company information
// @Tags         Company
// @Accept       json
// @Produce      json
// @Request
// @Param        id 		path 	string  	true  "id UUID"
// @Success      200  {object} 	utils.Response
// @Failure 	 406  {object}  utils.ErrorResponse
// @Failure 	 404  {object}  utils.ErrorResponse
// @Security 	 oauth2[identity_api]
// @Router       /api/v1/company/{id} [get]
func GetSingleCompany(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.FormValue("id"))

	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting company information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	company, err := companyRepository.FindByID(id)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting company information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if company == nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := utils.Response{
		Code:    http.StatusOK,
		Body:    company,
		Title:   "OK",
		Message: "Company information",
	}
	return c.Status(resp.Code).JSON(resp)

}

// AddNewCompany godoc
// @Summary      AddNewCompany adds new company
// @Description  AddNewCompany adds new company
// @Tags         Company
// @Accept       json
// @Produce      json
// @Request
// @Success      200  {object} 	utils.Response
// @Failure 	 406  {object}  utils.ErrorResponse
// @Failure 	 404  {object}  utils.ErrorResponse
// @Failure 	 500  {object}  utils.ErrorResponse
// @Security 	 oauth2[identity_api]
// @Router       /api/v1/company [post]
func AddNewCompany(c *fiber.Ctx) error {
	company := &models.Company{}

	err := c.BodyParser(company)

	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing company body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := companyRepository.Save(*company)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	company, err = companyRepository.FindByID(id)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if company == nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := utils.Response{
		Code:    http.StatusOK,
		Body:    company,
		Title:   "OK",
		Message: "new company added successfully",
	}
	return c.Status(resp.Code).JSON(resp)

}

// UpdateCompany godoc
// @Summary      UpdateCompany updates a company by company id
// @Description  UpdateCompany updates a company by company id
// @Tags         Company
// @Accept       json
// @Produce      json
// @Request
// @Success      200  {object} 	utils.Response
// @Failure 	 406  {object}  utils.ErrorResponse
// @Failure 	 404  {object}  utils.ErrorResponse
// @Failure 	 500  {object}  utils.ErrorResponse
// @Security 	 oauth2[identity_api]
// @Router       /api/v1/company [put]
func UpdateCompany(c *fiber.Ctx) error {
	company := &models.Company{}

	err := c.BodyParser(company)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing company body information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing company ID. (it should be an integer)",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	updatingCompany, err := companyRepository.FindByID(id)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting company information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if updatingCompany == nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	company.ID = id

	err = companyRepository.Update(*company)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in updating company information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	company, err = companyRepository.FindByID(id)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly updated company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if company == nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := utils.Response{
		Code:    http.StatusOK,
		Body:    company,
		Title:   "UpdateCompany",
		Message: "company updated successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

// DeleteCompany godoc
// @Summary      DeleteCompany deletes the company from db
// @Description  DeleteCompany deletes the company from db
// @Tags         Company
// @Accept       json
// @Produce      json
// @Request
// @Success      200  {object} 	utils.Response
// @Failure 	 406  {object}  utils.ErrorResponse
// @Failure 	 404  {object}  utils.ErrorResponse
// @Failure 	 500  {object}  utils.ErrorResponse
// @Security 	 oauth2[identity_api]
// @Router       /api/v1/company/{id} [delete]
func DeleteCompany(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in getting company information",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	company, err := companyRepository.FindByID(id)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	if company == nil {
		errorResp := utils.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	err = companyRepository.Delete(*company)
	if err != nil {
		errorResp := utils.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in deleting company object",
		}

		return c.Status(errorResp.Code).JSON(errorResp)
	}

	resp := utils.Response{
		Code:    http.StatusOK,
		Body:    "company deleted successfully",
		Title:   "OK",
		Message: "company deleted successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}
