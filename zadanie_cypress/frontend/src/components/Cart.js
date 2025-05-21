import React, { useEffect } from "react";
import { useCart } from "../context/CartContext";
import { Link } from "react-router-dom";

const Cart = () => {
  const { items, removeItem, fetchCartItems } = useCart();
useEffect(() => {
    fetchCartItems();
  }, []); 
  if (items.length === 0) {
    return (
      <div>
        <h2>Koszyk</h2>
        <p>Koszyk jest pusty</p>
      </div>
    );
  }

  return (
    <div>
      <h2>Koszyk</h2>
      {items.map(cartItem => (
        <div key={cartItem.ID}>
          <span>
            {cartItem.product.name} - {cartItem.product.price} zł
          </span>
          <button onClick={() => removeItem(cartItem.ID)}>
            Usuń
          </button>
        </div>
      ))}
      <Link to="/payment"> Przejdź do płatności</Link>
    </div>
  );
};

export default Cart;
