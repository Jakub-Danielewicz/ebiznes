import React, { useEffect, useState } from "react";
import { api } from "../api";
import { useCart } from "../context/CartContext";

const Products = () => {
  const [products, setProducts] = useState([]);
  const { addItem } = useCart();

  useEffect(() => {
    api.get("/products")
      .then(res => setProducts(res.data))
      .catch(err => console.error(err));
      }, []);
  return (
    <div>
      <h2> Produkty </h2>
      {products.map(p => (
	<div key={p.id}>
	  <span>{p.name} - {p.price} zł - kategoria: {p.category.name} </span>
	  <button onClick={() => {addItem(p, 1); console.log("dodaję: ", p.name)}}>
	    Dodaj do koszyka
	  </button>
	</div>
      ))}
    </div>
  );
};

export default Products;
