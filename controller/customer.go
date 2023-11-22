package controller

import (
	"duck-cook-auth/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Adicionar novo cliente
// @Description	Adicionar um novo cliente
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		payload	body		entity.CustomerInput	true	"Dados do Cliente"
// @Success		200		{object}	entity.Customer
// @Router		/customer [post]
func (c *Controller) CreateCustomerHandler(ctx *gin.Context) {
	var customer entity.CustomerInput
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON"})
		return
	}

	customerResult, err := c.customerUseCase.CreateCustomer(customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, _ := c.authUseCase.GenerateJWT(customerResult)

	ctx.JSON(http.StatusCreated, gin.H{
		"customer": customerResult,
		"token":    jwt,
	})
}
