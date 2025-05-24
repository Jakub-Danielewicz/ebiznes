import React, { createContext, useContext, useEffect, useState, useMemo } from "react";
import { api } from "../api";
import PropTypes from "prop-types";

const CartContext = createContext();

export const CartProvider = ({ children }) => {
  const [cartId, setCartId] = useState(null);
  const [items, setItems] = useState([]);

  useEffect(() => {
    api.get("/cart")
      .then(res => {
		console.log("Odświeżono koszyk:", res.data);
		setCartId(res.data.id);
		setItems(Array.isArray(res.data.cartItems) ? res.data.cartItems : []);		
	})
      .catch(err => console.error("Nie udało się utworzyć koszyka:", err));
  }, []);

  const addItem = (product, quantity = 1) => {
    if (!cartId) {
      console.log("cartId:", cartId);
      return;
    }

    console.log("Wywołuję API");
    api.post(`/cart/${cartId}/items`, { productId: product.ID, quantity: quantity })
      .then(res => {
        const cartItem = res.data;
        setItems([...items, cartItem]);
        console.log("Dodano produkt do koszyka:", cartItem);
      })
      .catch(err => console.error("Nie dodano produktu:", err));
  };

  const removeItem = (productId) => {
    if (!cartId) return;

    api.delete(`/cart/${cartId}/items/${productId}`)

      .then(() => console.log("Usunięto produkt z koszyka:", productId))
      .then(() => fetchCartItems())
      .catch(err => console.error("Nie usunięto produktu:", err));
  };

  const fetchCartItems = () => {
    if (!cartId) return;

    api.get(`/cart`)
      .then(res => {
        console.log("Pobrano produkty z koszyka:", res.data);
        setItems(res.data.cartItems);
      })
      .catch(err => console.error("Nie udało się pobrać produktów:", err));
  };

  const clearCart = () => {
    if (!cartId) return;
    
    api.delete(`/cart/${cartId}`)
      .then(() => {
	setItems([]);
	setCartId(null);
      })
      .catch(err => console.error("Nie wyczyszczono koszyka:", err));
  };

  const contextValue = useMemo(
    () => ({
      cartId,
      items,
      addItem,
      removeItem,
      clearCart,
      fetchCartItems,
    }),
    [cartId, items, addItem, removeItem, clearCart, fetchCartItems]
  );

  return (
    <CartContext.Provider value={contextValue}>
      {children}
    </CartContext.Provider>
  );
};
CartProvider.propTypes = {
  children: PropTypes.node.isRequired,
};

export const useCart = () => useContext(CartContext);
