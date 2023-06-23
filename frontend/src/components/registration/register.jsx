import React, { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';
import "./registration.css";
import Navbar from "../navbar";

export default function Register() {
    const [values, setValues] = useState({
      FirstName: "",
      LastName: "",
      NickName: "",
      Email: "",
      DOB: "",
      Avatar: "",
      AboutMe: "",
      Password: "",
    });


    const [errors, setErrors] = useState({});
    const navigate = useNavigate();

    const [ProfilePhotoUp, setProfilePhotoUp] = useState(null);

    const ImageUpload = (event) => {
      const file = event.target.files[0];
      setProfilePhotoUp(file);

      // Read the image file as a data URL
      const reader = new FileReader();
      reader.onloadend = () => {
        setValues({Avatar:reader.result,});
      };

      reader.readAsDataURL(file);
    };


     const handleInput = (event) => {
       const newObj = { ...values, [event.target.name]: event.target.value };
       setValues(newObj);
       console.log("checking register      ");
       setErrors(Validation(event, values));
     }

    const Validation = (event, values) => {
       const errors = {};

       const email_pattern = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
       const password_pattern = /^([a-zA-Z0-9@*#]{8,15})$/;
       
       // checking FirstName input validity
       if (event.target.name === "FirstName") {
         if (values.FirstName === "") {
           errors.FirstName = "First Name cannot be blank";
         }
       }

       // checking FirstName input validity
       if (event.target.name === "LastName") {
         if (values.LastName === "") {
           errors.LastName = "Last Name cannot be blank";
         }
       }
       
       // checking email input validity
       if (event.target.name === "Email") {
         if (values.Email === "") {
           errors.Email = "Email cannot be empty";
         } else if (!email_pattern.test(values.Email)) {
           errors.Email = "Email is invalid";
         }
       }
       
      //  checking password input validity
       if (event.target.name === "Password") {
         if (values.Password === "") {
           errors.Password = "Password cannot be empty";
         } else if (!password_pattern.test(values.Password)) {
           errors.Password =
             "Password is invalid,it should be atleast 8 characters, with a number, a uppercase,a lowercase and one special character";
         }
       }
       return errors;
     }

    const handleRegisterSubmit = (event, setErrors) => {

        event.preventDefault();
        
        console.log("handleSubmit");
        
        
        console.log(values);
        
        fetch("http://localhost:9090/register", {
          method: "POST",
          body: JSON.stringify(values),
          // credentials: "include",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
          },
        })
          .then((response) => {
            return response.text();
          })
          .then(function (response) {
            if (response.slice(0, 5) == "ERROR") {
              if (
                response ==
                "ERROR: This email already exists, please log in instead"
              ) {
                setErrors({
                  email: "This Email already exists, please log in instead",
                });
              }
              return;
            }
            setData({ ...values });
            navigate("/login");
            return response;
          })

          .catch((error) => {
            setErrors(error);
            // do something with the error
          });
    };


    return (
      <div className="form">
        <div className="form-body">
          <Navbar />
          <h1 className="registerHere">🆁🅴🅶🅸🆂🆃🅴🆁 🅷🅴🆁🅴</h1>
          <label id="AvatarTitle" className="form__label" htmlFor="avatar">
            🅰🆅🅰🆃🅰🆁{" "}
          </label>
          <div className="avatar">
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="panda"
                  name="avatar"
                  value="src/images/a1.png"
                  onChange={handleInput}
                />
                <img src="src/images/a1.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="fox"
                  name="avatar"
                  value="src/images/a2.png"
                  onChange={handleInput}
                />
                <img src="src/images/a2.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="rabbit"
                  name="avatar"
                  value="src/images/a3.png"
                  onChange={handleInput}
                />
                <img src="src/images/a3.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="cat"
                  name="avatar"
                  value="src/images/a4.png"
                  onChange={handleInput}
                />
                <img src="src/images/a4.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="bear"
                  name="avatar"
                  value="src/images/a5.png"
                  onChange={handleInput}
                />
                <img src="src/images/a5.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="dude"
                  name="avatar"
                  value="src/images/a6.png"
                  onChange={handleInput}
                />
                <img src="src/images/a6.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="dudette"
                  name="avatar"
                  value="src/images/a7.png"
                  onChange={handleInput}
                />
                <img src="src/images/a7.png"></img>
              </label>
            </div>
            <div id="avaImg">
              <label>
                <input
                  className="avatar"
                  type="radio"
                  id="none"
                  name="none"
                  value={""}
                  onChange={handleInput}
                />
                <img src="src/images/none.png"></img>
              </label>
            </div>

            <label id="ProfilePhotoUp" htmlFor="ProfilePhotoUp"></label>
            <input
              id="ProfilePhotoUp"
              type="file"
              accept="image/*"
              onChange={ImageUpload}
            />
            {ProfilePhotoUp && (
              <img id="profileShow" src={values.Avatar} alt="uploaded" />
            )}
          </div>

          <div className="username">
            <label className="form__label" htmlFor="firstName">
              🅵🅸🆁🆂🆃 🅽🅰🅼🅴{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="text"
              id="firstName"
              name="FirstName"
              value={values.FirstName}
              onChange={handleInput}
              placeholder="First Name"
            />
          </div>

          <div className="lastname">
            <label className="form__label" htmlFor="lastName">
              🅻🅰🆂🆃 🅽🅰🅼🅴{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="text"
              id="lastName"
              name="LastName"
              value={values.LastName}
              onChange={handleInput}
              placeholder="Last Name"
            />
          </div>

          <div className="nickname">
            <label className="form__label" htmlFor="nickName">
              🅽🅸🅲🅺 🅽🅰🅼🅴{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="text"
              id="nickName"
              name="NickName"
              value={values.NickName}
              onChange={handleInput}
              placeholder="NickName"
            />
          </div>

          <div className="email">
            <label className="form__label" htmlFor="email">
              🅴🅼🅰🅸🅻{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="email"
              id="email"
              name="Email"
              value={values.Email}
              onChange={handleInput}
              placeholder="Email"
            />
          </div>

          <div className="dob">
            <label className="form__label" htmlFor="dob">
              🅳🅰🆃🅴 🅾🅵 🅱🅸🆁🆃🅷{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="date"
              id="dob"
              name="DOB"
              value={values.DOB}
              onChange={handleInput}
              placeholder="Date Of Birth"
            />
          </div>

          <div className="aboutMe">
            <label className="form__label" htmlFor="aboutMe">
              🅰🅱🅾🆄🆃 🅼🅴{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="text"
              id="aboutMe"
              name="AboutMe"
              value={values.AboutMe}
              onChange={handleInput}
              placeholder="aboutMe"
            />
          </div>

          <div className="password">
            <label className="form__label" htmlFor="password">
              🅿🅰🆂🆂🆆🅾🆁🅳{" "}
            </label>
            <br></br>
            <input
              className="form__input"
              type="password"
              id="password"
              name="Password"
              value={values.Password}
              onChange={handleInput}
              placeholder="Password"
            />
            {errors.Password && (
              <p style={{ color: "red" }}>{errors.Password}</p>
            )}
          </div>

          <div className="footer">
            <a
              onClick={(e) => {
                handleRegisterSubmit(e, setErrors);
                handleInput(e);
              }}
              type="submit"
            >
              <img id="submitBtnRegPage" src="src/images/submit.png" />
            </a>
          </div>
        </div>
      </div>
    );
}
