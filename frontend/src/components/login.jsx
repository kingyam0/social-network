import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "./navbar";

export default function Login({ socket, setSocket, data, setData, setWebsocket }) {
    const [Username, setUsername] = useState("");
    const [Password, setPassword] = useState("");
    const [error, setError] = useState();
    const navigate = useNavigate();

    const handleSubmit = (event) => {
        console.log("handleSubmit");
        event.preventDefault();
        console.log(Username, Password);

        fetch("http://localhost:9090/login", {
            method: "POST",
            body: JSON.stringify({ Username, Password }),
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json",
            },
            credentials: "include",
        })
            .then((response) => response.json())
            .then((data) => {
                setData(data);
                console.log("LOGGEDIN OBJECT DATA", data);

                if (data) {
                    console.log("Successful Login!");
                    let sc = new WebSocket('ws://localhost:9090/ws');
                    if (sc) { // if socket is not null
                        setWebsocket(sc);
                        console.log("WebSocket has been set..");
                        navigate('/loggedIn')
                        localStorage.setItem('userData', JSON.stringify(data));
                    }
                } else {
                    // add alert  not ok
                    // notyf.error("The login details you entered are incorrect.");
                }
                return data;
            })
            .catch((err) => {
                setError(err);
            });
    };

    return (
      <div className="login-wrapper">
        <Navbar />
        <h1 className="headings">🅻🅾🅶🅸🅽 🅷🅴🆁🅴</h1>
        <form onSubmit={handleSubmit}>
          <label>
            <p className="form__label">🆄🆂🅴🆁🅽🅰🅼🅴</p>
            <input
              id="textBox"
              type="text"
              value={Username}
              onChange={(event) => setUsername(event.target.value)}
            />
          </label>
          <label>
            <p className="form__label">🅿🅰🆂🆂🆆🅾🆁🅳</p>
            <input
              id="textBox"
              type="password"
              value={Password}
              onChange={(event) => setPassword(event.target.value)}
            />
          </label>
          <div>
            <a onClick={handleSubmit} type="submit">
              <img id="submitBtn" src="src/images/submit.png" />
            </a>
          </div>
        </form>
      </div>
    );
}
