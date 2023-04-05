package handler

import (
	"log"
	"my_api/internal/model"
	"my_api/internal/model/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signupReq struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *Handler) Singup(c *gin.Context) {
	var req signupReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := h.TokenService.NewPairFromUser(c, u, "")
	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
