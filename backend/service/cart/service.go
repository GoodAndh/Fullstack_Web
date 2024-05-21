package cart

import (
	"context"
	"fmt"
	"fullstack_toko/backend/model/web"
)

func (h *Handler) createOrders(ctx context.Context, ps *web.Product, cart []web.CartCheckoutItem, userID int) (int, float64, error) {

	if err := checkUsers(ps, userID); err != nil {
		return 0, 0, err
	}

	// store ps into map
	productMap := make(map[int]web.Product)
	productMap[ps.Id] = *ps

	if err := checkIfinStock(cart, productMap); err != nil {
		return 0, 0, err
	}

	totalPrice := calculateTotalPrice(cart, productMap)

	//create order

	orderID, err := h.orderService.CreateOrders(ctx, userID)
	if err != nil {
		return 0, 0, err
	}

	for _, item := range cart {
		err := h.orderService.CreateOrderitems(ctx, &web.OrderitemsCreatePayload{
			Status:     "pending",
			Quantity:   item.Quantity,
			TotalPrice: totalPrice,
			ProductID:  item.ProductID,
			OrderID:    orderID,
			UserID:     userID,
		})
		if err != nil {
			return 0, 0, err
		}
	}

	return orderID, totalPrice, nil
}

func calculateTotalPrice(cart []web.CartCheckoutItem, pMap map[int]web.Product) float64 {
	var total float64

	for _, item := range cart {
		product := pMap[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	return total
}

func checkIfinStock(cart []web.CartCheckoutItem, pMap map[int]web.Product) error {
	if len(cart) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cart {
		product, ok := pMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available ", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product stock %d left ,want %d ", product.Quantity, item.Quantity)
		}

	}
	return nil
}

func checkUsers(ps *web.Product, userID int) error {
	pMap := make(map[int]web.Product)
	pMap[ps.Userid] = *ps
	_, ok := pMap[userID]
	if ok {
		return fmt.Errorf("you cant order your own product")
	} else {
		return nil
	}
}
