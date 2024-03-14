// src/components/Login.js

import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Form, Alert } from "react-bootstrap";
import { Button } from "react-bootstrap";

const Login = ({isloginUS, setUser }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
  
    // Check if email or password is empty
    if (!email.trim() || !password.trim()) {
      setError("Please enter both email and password.");
      return;
    }
  
    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        credentials: 'include',
        body: JSON.stringify({ email, password }),
      });
      if (response.ok) {
        const userData = await response.json();
        setUser(userData)
        // console.log(userData)
        
        // Now 'user' state contains the user data
      } else {
        const data = await response.json();
        setError(data.message);
        return;
      }
    
      // Uncomment the following line to set the login state
      // isloginUS(true);
      
    
      navigate("/");
    } catch (err) {
      console.error("Error during login:", err);
      setError("An unexpected error occurred. Please try again.");
    }
  };

  return (
    <div style={{display: 'flex', margin: "auto",  justifyContent: "center", alignItems: "center", height: "80vh"}}>
       <div class="bird-container">
            <div class="bird-body"></div>
          <div class="mouth"></div>
            <div class="beak"></div>
            <div class="feather"></div>
          <div class="tail"></div>
              <div class="leg"></div>
      </div>
    <div style={{ width: "500px", display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center", height: "80vh" }}>
    {/* <div> */}
      <div className="w-75 h-50 p-4 box">
        <h2 className="mb-3 text-center fw-bolder fs-1">Login</h2>
        {error && <Alert variant="danger">{error}</Alert>}
        <Form onSubmit={handleSubmit}>
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Control
              type="email"
              placeholder="Email address"
              onChange={(e) => setEmail(e.target.value)}
            />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasicPassword">
            <Form.Control
              type="password"
              placeholder="Password"
              onChange={(e) => setPassword(e.target.value)}
            />
          </Form.Group>

          <div className="d-grid gap-2">
            <Button variant="primary" type="Submit">
              Log In
            </Button>
          </div>
        </Form>
        <hr />
        <div className="p-4 box mt-3 text-center">
          Don't have an account? <Link to="/signup">Sign up</Link>
        </div>
      </div>
    </div>
    </div>
  );
};

export default Login;
