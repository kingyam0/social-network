import React from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "./logout";

const Navbar = ({data}) => {

   const location = window.location.pathname;
   const navigate = useNavigate();
   const toLoggedIn = (evt) => {
     evt.preventDefault();
     navigate("/loggedIn");
   };
   const toUserProfile = (evt) => {
     evt.preventDefault();
     navigate("/profile");
   };
   const toAddPosts = (evt) => {
     evt.preventDefault();
     navigate("/posts");
   };
 const toHome = (evt) => {
   evt.preventDefault();
   navigate('/');
 };

   if (data) {
    return (
      <nav className="nav">
        <div onClick={toLoggedIn}>
          <div id="logo">🅂🄾🄾 </div>
          <div id="logo2"> 🄽🄴🅃</div>
        </div>

        <div id="welcome">🅆🄴🄻🄲🄾🄼🄴</div>
        <div id="user">{data.User.FirstName}</div>

        <LogoutButton />

        <div onClick={toUserProfile}>
          <img id="profile" src="src/images/prfile.png" />
        </div>

        <div onClick={toAddPosts}>
          <img id="postsIcon" src="src/images/posts.png" />
        </div>

        <div>
          <img id="notiIcon" src="src/images/noti.png" />
        </div>

        <div>
          <img id="chatIcon" src="src/images/chatIcon.png" />
        </div>
      </nav>
    );
  }

  if(location === "/register") {
  return (
    <nav className="nav">
      <div onClick={toHome}>
        <div id="logo">🅂🄾🄾 </div>
        <div id="logo2"> 🄽🄴🅃</div>
      </div>

      <a href="/login">
        <img id="logBtnRegPage" src="src/images/login.png" />
      </a>
      
    </nav>
  );
}

if(location === "/login") {
    return (
      <nav className="nav">
        <div onClick={toHome}>
          <div id="logo">🅂🄾🄾 </div>
          <div id="logo2"> 🄽🄴🅃</div>
        </div>

        <a href="/register">
          <img id="regBtnLogPage" src="src/images/register.png" />
        </a>

      </nav>
    );
}

if (location === "/") {
  return (
    <nav className="nav">
        <div id="logo">🅂🄾🄾 </div>
        <div id="logo2"> 🄽🄴🅃</div>
    </nav>
  );
}
};

export default Navbar;
