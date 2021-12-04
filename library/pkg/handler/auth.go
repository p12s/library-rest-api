package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p12s/library-rest-api/library/pkg/models"
)

// @Summary Sign up
// @Tags Auth
// @Description Create account
// @ID signUp
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO неудачное решение - встроить отдельный сервис
	// непонятно как его mock-ать

	// go func() {
	// 	_, err := h.grpcLogger.Service.Log(context.Background(), &pb.LoggerRequest{
	// 		Action:    pb.LoggerRequest_REGISTER,
	// 		Entity:    pb.LoggerRequest_USER,
	// 		EntityId:  int64(id),
	// 		Timestamp: timestamppb.Now(),
	// 	})
	// 	if err != nil {
	// 		logrus.Errorf("GRPC-logging signup user: %s/n", err.Error())
	// 	}
	// }()

	// go func() {
	// 	err := h.queueLogger.Produce(pb.LoggerRequest{
	// 		Action:    pb.LoggerRequest_REGISTER,
	// 		Entity:    pb.LoggerRequest_USER,
	// 		EntityId:  int64(id),
	// 		Timestamp: timestamppb.Now(),
	// 	})
	// 	if err != nil {
	// 		logrus.Errorf("Queue-logging signup user: %s/n", err.Error())
	// 	}
	// }()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// signInInput - authorization data
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Sign in
// @Tags Auth
// @Description Sending data to get authentication jwt-token
// @ID signIn
// @Accept  json
// @Produce  json
// @Param input body handler.signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO неудачное решение - встроить отдельный сервис
	// непонятно как его mock-ать

	// go func() {
	// 	_, err := h.grpcLogger.Service.Log(context.Background(), &pb.LoggerRequest{
	// 		Action:    pb.LoggerRequest_LOGIN,
	// 		Entity:    pb.LoggerRequest_USER,
	// 		Timestamp: timestamppb.Now(),
	// 	})
	// 	if err != nil {
	// 		logrus.Errorf("GRPC-logging login user: %s/n", err.Error())
	// 	}
	// }()

	// go func() {
	// 	err := h.queueLogger.Produce(pb.LoggerRequest{
	// 		Action:    pb.LoggerRequest_LOGIN,
	// 		Entity:    pb.LoggerRequest_USER,
	// 		Timestamp: timestamppb.Now(),
	// 	})
	// 	if err != nil {
	// 		logrus.Errorf("Queue-logging login user: %s/n", err.Error())
	// 	}
	// }()

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
