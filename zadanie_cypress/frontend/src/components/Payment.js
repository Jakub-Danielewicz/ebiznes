import React from "react";
import { useCart } from "../context/CartContext";

const Payment = () => {
  const { items, clearCart } = useCart();

  const handlePayment = () => {
    clearCart();
    alert("Zamówienie złożone!");
  };

  return(
    <div>
    <h2>Płatność</h2>
    <button onClick={handlePayment}>Zamów</button>
    </div>
  );
};

export default Payment;
