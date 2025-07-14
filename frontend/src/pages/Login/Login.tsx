import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { login } from "../../api";
import "./Login.css";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

    const navigate = useNavigate();

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    login(username, password)
      .then(() => {
        navigate("/");
      })
      .catch((error) => {
        setError("Failed to log in. Please check your credentials and try again.");
      });
  };

  return (
    <div className="App">
      <header>
        <h1>Login</h1>
      </header>
      <div className="main">
        <form onSubmit={handleLogin} className="auth-form">
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit">Login</button>
          {error && <p style={{
            color: "red"
          }}>{error}</p>}
          <p>
            Don't have an account? <Link to="/signup">Sign up</Link>
          </p>
        </form>
      </div>
    </div>
  );
};

export default Login;
