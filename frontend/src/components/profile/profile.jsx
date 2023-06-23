import React , {useEffect} from "react";
import { useNavigate } from 'react-router-dom';
import './profile.css';
import '../login'
import CheckCookie from "../checkCookie";
import Navbar from "../navbar";


export default function Profile({setData, data }) {
  console.log("PROFILE DATA:", data);

  const navigate = useNavigate();

  const tologgedIn = () => {
    navigate("/loggedIn");
  };

  return (
    //wrapped in a fragment (<></>)

    <>
      <CheckCookie setData={setData} data={data} />
      {data.User ? (
        <>
          <Navbar data={data} />
      

          <img id="ProfImage" src={data.User.Avatar} />

          <div className="profDiv">
            <div id="firstNAME"> 🄵🄸🅁🅂🅃🄽🄰🄼🄴 : </div>

            <div id="lastNAME"> 🄻🄰🅂🅃🄽🄰🄼🄴 : </div>

            <div id="nickNAME"> 🄽🄸🄲🄺🄽🄰🄼🄴 : </div>

            <div id="aboutME"> 🄰🄱🄾🅄🅃 🄼🄴 : </div>

            <div id="EMail"> 🄴🄼🄰🄸🄻 : </div>
          </div>

          <div className="info">
            <div id="infofirstName">{data.User.FirstName}</div>

            <div id="infolastName">{data.User.LastName}</div>

            <div id="infonickName">{data.User.NickName}</div>

            <div id="infoaboutMe">{data.User.AboutMe}</div>

            <div id="infoEmail">{data.User.Email}</div>
          </div>
        </>
      ) : (
        <></>
      )}
    </>
  );
}