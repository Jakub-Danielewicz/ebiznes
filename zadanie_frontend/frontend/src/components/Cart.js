import React from "react"
import { useCart } from "../context/CartContext";
import { Link } from "react-router-dom";

const Cart = () => {
  const { items, removeItem } = useCart();

  return(
    <div>
      <h2>Koszyk</h2>
      {items.length === 0 ? (
	<p>Koszyk jest pusty</p>
      ): (
      items.map(p => (
	<div key={p.id}>
	  <span>{p.name} - {p.price} zł</span>
	  <button onClick={() => removeItem(p.id)}>
	    Usuń
	  </button>
	</div>
      ))
	)}
      <Link to="/payment"> Przejdź do płatności</Link>
    </div>
  );
};

export default Cart;
