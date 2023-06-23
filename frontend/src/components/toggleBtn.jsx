import React, { useState } from "react";
import "./privacyToggle.css";

const PrivacyToggle = () => {
  const [isPublic, setIsPublic] = useState(false);

  const handlePrivacy = () => {
    setIsPublic(!isPublic);
    let data = { Privacy: !isPublic ? "private" : "public" };

    fetch("http://localhost:9090/updateprivacy", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      },
      credentials: "include",
    })
      .then(() => {
        setIsPublic(!isPublic);
      })
      .catch((error) => {
        console.error(error);
      });
  };
  return (
    <div onClick={handlePrivacy}>
      <button id="Toggle" className="button" onClick={handlePrivacy}>{isPublic ? "Public" : "Private"}</button>
    </div>
  );
};

export default PrivacyToggle;