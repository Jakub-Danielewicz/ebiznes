import React, { useState } from "react";
import { api } from "../api";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [form, setForm] = useState({ username: "", email: "", password: "" });
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleChange = e => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async e => {
    e.preventDefault();
    setError("");
    try {
      await api.post("/auth/register", form);
      alert("Rejestracja udana! Możesz się zalogować.");
      navigate("/auth/login");
    } catch (err) {
      setError("Błąd rejestracji");
    }
  };

  return (
    <div>
      <h2>Rejestracja</h2>
      <form onSubmit={handleSubmit}>
        <input name="username" placeholder="Nazwa użytkownika" value={form.username} onChange={handleChange} required />
        <input name="email" type="email" placeholder="Email" value={form.email} onChange={handleChange} required />
        <input name="password" type="password" placeholder="Hasło" value={form.password} onChange={handleChange} required />
        <button type="submit">Zarejestruj</button>
      </form>
      {error && <p style={{color:"red"}}>{error}</p>}
    </div>
  );
};

export default Register;