import React, { useState } from "react";
import { api } from "../api";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleChange = e => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async e => {
    e.preventDefault();
    setError("");
    try {
      await api.post("/auth/login", form);
      alert("Zalogowano!");
      navigate("/");
    } catch (err) {
      setError("Błędny email lub hasło");
    }
  };

  return (
    <div>
      <h2>Logowanie</h2>
      <form onSubmit={handleSubmit}>
        <input name="email" type="email" placeholder="Email" value={form.email} onChange={handleChange} required />
        <input name="password" type="password" placeholder="Hasło" value={form.password} onChange={handleChange} required />
        <button type="submit">Zaloguj</button>
      </form>
      <hr />
      <button
        type="button"
        onClick={() => {
          window.location.href = "http://localhost:8080/auth/google";
        }}
        style={{ background: "#4285F4", color: "white", padding: "8px 16px", border: "none", borderRadius: "4px", marginTop: "10px" }}
      >
        Zaloguj przez Google
      </button>
      {error && <p style={{color:"red"}}>{error}</p>}
      <button
        type="button"
        onClick={() => {
          window.location.href = "http://localhost:8080/auth/github";
        }}
        style={{ background: "#4285F4", color: "white", padding: "8px 16px", border: "none", borderRadius: "4px", marginTop: "10px" }}
      >
        Zaloguj przez GitHub
      </button>
      {error && <p style={{color:"red"}}>{error}</p>}
    </div>
  );
};

export default Login;