package controller

import (
	"duck-cook-auth/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Login cliente
// @Description	Cria um JWT para o cliente
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		payload	body		entity.Credentials	true	"Dados do Cliente"
// @Success		200		{object}	entity.Customer
// @Router		/auth/login [post]
func (c Controller) LoginHandler(ctx *gin.Context) {
	var auth entity.Credentials
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar o JSON"})
		return
	}
	customer, err := c.customerUseCase.GetCustomerByUser(auth.User)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Crendecias invalidas"})
		return
	}
	isValid, err := c.authUseCase.Login(auth.Pass, customer.Pass)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Não foi possivel validar o usuário"})
		return
	}
	if isValid {
		jwt, err := c.authUseCase.GenerateJWT(customer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possivel gerar o JWT"})
		}
		ctx.JSON(http.StatusOK, gin.H{"token": jwt})
		return
	}
	ctx.JSON(http.StatusForbidden, gin.H{"error": "Crendecias erradas"})
}
