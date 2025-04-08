import React, { useEffect, useState } from 'react';

function App() {
  const [products, setProducts] = useState([]);
  const [cart, setCart] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/products')
      .then(res => res.json())
      .then(setProducts);
  }, []);

  const addToCart = (p) => setCart([...cart, p]);

  return (
    <div style={{ padding: 20 }}>
      <h2>🛍️ Магазин</h2>
      {products.map(p => (
        <div key={p.id}>
          {p.name} — {p.price}₽
          <button onClick={() => addToCart(p)}>Добавить</button>
        </div>
      ))}
      <h3>Корзина:</h3>
      <ul>
        {cart.map((item, i) => <li key={i}>{item.name} — {item.price}₽</li>)}
      </ul>
    </div>
  );
}

export default App;
