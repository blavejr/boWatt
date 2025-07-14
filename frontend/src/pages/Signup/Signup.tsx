import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { signup } from "../../api";
import "./Signup.css";

const Signup = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const navigate = useNavigate();

  const handleSignup = (e: React.FormEvent) => {
    e.preventDefault();
    signup(username, password)
      .then(() => {
        navigate("/");
      })
      .catch((error) => {
        setError("Failed to create account. Please try again.");
      }
    );
  };

  return (
    <div className="App">
      <header>
        <h1>Sign Up</h1>
      </header>
      <div className="main">
        <form onSubmit={handleSignup} className="auth-form">
          <input
            type="text"
            placeholder="Choose a Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Choose a Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit">Create Account</button>
          {error && <p style={{
            color: "red"
          }}>{error}</p>}
          <p>
            Already have an account? <Link to="/login">Log in</Link>
          </p>
        </form>
      </div>
    </div>
  );
};

export default Signup;
