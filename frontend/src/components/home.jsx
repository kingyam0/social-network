import React from "react";
import Navbar from "./navbar";

        export default function Home() {
            return (
              <>
                <div>
                  <img id="welcomeBanner" src="src/images/welcome.gif"></img>
                </div>
                <div className="homeDiv">
                  <Navbar />
                  <a href="/login">
                    <img id="loginBtn" src="src/images/login.png" />
                  </a>


                  <a href="/register">
                    <img id="regBtn"  src="src/images/register.png" />
                  </a>
                </div>
              </>
            );
        }
