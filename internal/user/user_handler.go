package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/util"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user CreateUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExistMsg, isUserExist := h.Service.CheckUserExist(c, &user)
	if !isUserExist {
		c.JSON(http.StatusBadRequest, util.Response{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: userExistMsg,
		})
		return
	}

	validUserReqMsg, isUserReqDataValid := HandleSignUpError(&user)
	if !isUserReqDataValid {
		c.JSON(http.StatusBadRequest, util.Response{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: validUserReqMsg,
		})
		return
	}

	response, err := h.Service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, util.Response{
		Data:    response,
		Status:  http.StatusOK,
		Message: "account created!",
	})
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentCookie, _ := c.Cookie("jwt")
	if len(currentCookie) > 0 {
		c.JSON(http.StatusBadRequest, util.Response{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: "you are logged in already!",
		})
		return
	}

	c.SetCookie("jwt", u.accessToken, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, util.Response{
		Data:    u,
		Status:  http.StatusOK,
		Message: "login successfully",
	})
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, util.Response{
		Data:    nil,
		Status:  http.StatusOK,
		Message: "logout successfully",
	})
}
