package controllers

import (
	"net/http"
	"order-management/database"
	"order-management/models"
	"order-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateOrderInput struct {
	StoreID            uint    `json:"store_id" binding:"required"`
	MerchantOrderID    string  `json:"merchant_order_id"`
	RecipientName      string  `json:"recipient_name" binding:"required"`
	RecipientPhone     string  `json:"recipient_phone" binding:"required"`
	RecipientAddress   string  `json:"recipient_address" binding:"required"`
	RecipientCity      uint    `json:"recipient_city" binding:"required"`
	RecipientZone      uint    `json:"recipient_zone" binding:"required"`
	RecipientArea      uint    `json:"recipient_area" binding:"required"`
	DeliveryType       uint    `json:"delivery_type" binding:"required"`
	ItemType           uint    `json:"item_type" binding:"required"`
	ItemQuantity       uint    `json:"item_quantity" binding:"required"`
	ItemWeight         float64 `json:"item_weight" binding:"required"`
	AmountToCollect    float64 `json:"amount_to_collect" binding:"required"`
	SpecialInstruction string  `json:"instruction"`
	ItemDescription    string  `json:"item_description"`
}

func CreateOrder(c *gin.Context) {
	var input CreateOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		handleValidationError(c, err)
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	deliveryFee := calculateDeliveryFee(input.RecipientCity, input.ItemWeight)
	codFee := calculateCODFee(input.AmountToCollect)

	order := models.Order{
		UserID:             userID.(uint),
		StoreID:            input.StoreID,
		MerchantOrderID:    input.MerchantOrderID,
		RecipientName:      input.RecipientName,
		RecipientPhone:     input.RecipientPhone,
		RecipientAddress:   input.RecipientAddress,
		RecipientCity:      input.RecipientCity,
		RecipientZone:      input.RecipientZone,
		RecipientArea:      input.RecipientArea,
		DeliveryType:       input.DeliveryType,
		ItemType:           input.ItemType,
		ItemQuantity:       input.ItemQuantity,
		ItemWeight:         input.ItemWeight,
		SpecialInstruction: input.SpecialInstruction,
		ItemDescription:    input.ItemDescription,
		AmountToCollect:    input.AmountToCollect,
		CODFee:             codFee,
		DeliveryFee:        deliveryFee,
		TotalFee:           deliveryFee + codFee,
		OrderStatus:        "Pending",
		OrderType:          "Delivery",
		OrderConsignmentID: utils.GenerateConsignmentID(),
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
}

func handleValidationError(c *gin.Context, err error) {
	validationErrors := make(map[string][]string)

	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range fieldErrors {
			fieldName := fieldError.Field()
			message := "The '" + fieldName + "' is required."
			validationErrors[fieldName] = append(validationErrors[fieldName], message)
		}
	} else {
		validationErrors["error"] = []string{err.Error()}
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Please fix the given errors",
		"type":    "error",
		"code":    422,
		"errors":  validationErrors,
	})
}

func calculateDeliveryFee(city uint, weight float64) float64 {
	baseDeliveryFee := 60.0
	if city != 1 {
		baseDeliveryFee = 100.0
	}

	if city == 1 {
		if weight <= 0.5 {
			return baseDeliveryFee
		} else if weight <= 1 {
			return baseDeliveryFee + 10
		} else {
			return baseDeliveryFee + 10 + ((weight - 1) * 15)
		}
	}
	return baseDeliveryFee + ((weight - 1) * 15)
}

func calculateCODFee(amount float64) float64 {
	return 0.01 * amount
}
