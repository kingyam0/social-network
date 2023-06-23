import { useState, useEffect } from "react";
// import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { WebSocketProvider } from "./context/WebSocketContext";
import "./App.css";

import Login from "./components/login";
import Home from "./components/home";
import Register from "./components/registration/register";
import LoggedIn from "./components/loggedIn/loggedIn";
import Profile from "./components/profile/profile";
import Posts from "./components/posts/posts";
import Logout from "./components/logout";
// import NavbarAfterLogin from "./components/navbarAfterLogin";



function App() {

  const [data, setData] = useState({
    firstName: "",
    lastName: "",
    nickname: "",
    AboutMe: "",
    avatar: "",
    email: "",
    DOB: "",
  });

  const [websocket, setWebsocket] = useState(null);

  return (

    <BrowserRouter>
      <Routes>
        <Route path='/' index element={<Home />} />
        <Route path='/login' index element={<Login setData={setData} data={data} setWebsocket={setWebsocket} />} />
        <Route path='/register' index element={<Register />} />
        <Route path='/loggedIn' index element={<LoggedIn setData={setData} data={data} websocket={websocket} />} />
        <Route path='/logout' index element={<Logout />} />
        <Route path='/profile' index element={<Profile setData={setData} data={data} />} />
        <Route path='/posts' index element={<Posts setData={setData} data={data} />} />
      </Routes>
    </BrowserRouter>

  )
}

export default App;
