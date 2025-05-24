import React from "react";
import { useCart } from "../context/CartContext";

const Payment = () => {
  const { clearCart } = useCart();

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
