import React, { useEffect } from "react";



const CheckCookie = ({ setData, data }) => {
  const checkCookie = () => {
    let cookie = document.cookie;
    // data.User.FirstName = "Zoey";
    if (data.User) {
      console.log("Not empty", data);
    } else {
      let cookieValue = document.cookie.split("=")[1];

      if (cookie != "") {
        let data = {
          cookieValue: cookieValue,
        };

        let options = {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        };

        fetch("http://localhost:9090/checkCookie", options)
          .then((response) => response.json())
          .then((data) => {
            setData(data);
            console.log("FROM FETCH", data);
          })
          .catch((error) => {
            console.error("Error:", error);
          });
      } else {
        navigate("/");
      }

      console.log("USER DATA", data.User);
      console.log("EMPTY DATA:", data);
    }
    console.log("COOKIE VALUE", cookie);
  };

  useEffect(() => {
    checkCookie();
  }, [setData, data]);


}

export default CheckCookie;