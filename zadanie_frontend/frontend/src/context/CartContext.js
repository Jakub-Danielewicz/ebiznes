import React, { createContext, useContext, useEffect, useState } from "react";
import { api } from "../api";

const CartContext = createContext();

export const CartProvider = ({ children }) => {
  const [cartId, setCartId] = useState(null);
  const [items, setItems] = useState([]);

  useEffect(() => {
    api.post("/cart")
      .then(res => setCartId(res.data.id))
      .catch(err => console.error("Nie udało się utworzyć koszyka:", err));
  }, []);

  const addItem = (product, quantity = 1) => {
    if (!cartId){
      console.log("cartId:", cartId);
      return;
    }

    console.log("Wywołuję API");
    api.post(`/cart/${cartId}/items`, {productId: product.id, quantity: quantity })
      .then(() => setItems([...items, product]))
      .catch(err => console.error("Nie dodano produktu:", err));
  };

  const removeItem = (productId) => {
    if (!cartId) return;

    api.delete(`/cart/${cartId}/items/${productId}`)
      .then(() => setItems(items.filter(item => item.id !== productId)))
      .catch(err => console.error("Nie usunięto produktu:", err));
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

  return (
    <CartContext.Provider value={{ cartId, items, addItem, removeItem, clearCart }}>
      {children}
    </CartContext.Provider>

  );
};

export const useCart = () => useContext(CartContext);
