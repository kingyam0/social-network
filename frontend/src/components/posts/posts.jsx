import React, { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';
import './posts.css'
import CheckCookie from "../checkCookie";
import Navbar from "../navbar";



const Posts = ({setData, data}) => {
  console.log("PostData", data);

  const navigate = useNavigate();

  const [PostTitle, setPostTitle] = useState("");
  const [PostContent, setPostContent] = useState("");
  const [PostCategory, setPostCategory] = useState("");
  const [PostPhotoUp, setPostPhotoUp] = useState(null);
  const [ImageDataUrl, setImageDataUrl] = useState("");
  const [PostPrivacy, setPostPrivacy] = useState("");

  const [postData, setPostData] = useState({
    PostTitle: "",
    PostContent: "",
    PostCategory: "",
    ImageDataUrl: "",
    PostPrivacy: "",
  });
  const [error, setError] = useState();

  //handling image upload for posts
  const handleImageUpload = (event) => {
    const file = event.target.files[0];
    setPostPhotoUp(file);

    // Read the image file as a data URL
    const reader = new FileReader();
    reader.onloadend = () => {
      setImageDataUrl(reader.result);
    };

    reader.readAsDataURL(file);
  };

  const tologgedIn = () => {
    navigate("/loggedIn");
  };

  // const toLoggedIn = () => {navigate('/loggedIn')}

  const handlePostSubmit = (event) => {
    event.preventDefault();
    navigate("/loggedIn");

    // console.log("THE POST DETAILS:", formData);
    console.log(
      "THIS IS EPEREATELY:",
      PostTitle,
      PostContent,
      PostCategory,
      ImageDataUrl,
      PostPrivacy
    );

    // // navigate('/posts');
    // toLoggedIn;

    fetch("http://localhost:9090/posts", {
      method: "POST",
      body: JSON.stringify({
        PostTitle,
        PostCategory,
        PostContent,
        ImageDataUrl,
        PostPrivacy,
      }),
      header: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
      .then((response) => response.json())
      .then((postData) => {
        setPostData(postData);

        console.log("THIS IS DATA AFTER SETDATA:", postData);
        if (postData) {
          console.log("Post Created Successfully.");
        } else {
          console.loge(error("Post creation failed. Please try again."));
        }
        return postData;
      })

      .catch((error) => {
        console.error(error);
      });
  };

  useEffect(() => {
    if (postData) {
      console.log(postData);
    }
    if (error) {
      console.log(error);
    }
  }, [postData]);

  return (
    //wrapped in a fragment (<></>)

    <>
      <CheckCookie setData={setData} data={data} />
      {data.User ? (
        <>
          <Navbar data={data} />
      
          <div id="CreatePostTitle">🄲🅁🄴🄰🅃🄴-🄿🄾🅂🅃</div>
          <div id="postBlock" className="posts">
            <label id="PostTitle" htmlFor="PostTitle">
              Post Title:
              <input
                id="textBoxPost"
                type="text"
                value={PostTitle}
                onChange={(event) => setPostTitle(event.target.value)}
              />
            </label>
            <br></br>
            <br></br>
            <br></br>

            <label id="Post_Cat" htmlFor="PostCategory">
              Catergory:
            </label>

            <select
              id="PostCat"
              value={PostCategory}
              onChange={(event) => setPostCategory(event.target.value)}
            >
              <option value="">-- Select --</option>
              <option value="Sports">Sports</option>
              <option value="Weather">Weather</option>
              <option value="Social">Social</option>
              <option value="Gaming">Gaming</option>
              <option value="Other">Other</option>
            </select>

            <label id="PostContent" for="PostContent">
              {" "}
              Content:{" "}
            </label>
            <input
              id="textBoxPostCont"
              type="text"
              value={PostContent}
              onChange={(event) => setPostContent(event.target.value)}
            />

            <label id="PostPhotoUp" htmlFor="PostPhotoUp">
              {" "}
              Image:{" "}
            </label>
            <input
              id="PhotoUp"
              type="file"
              accept="image/*"
              onChange={handleImageUpload}
            />
            {PostPhotoUp && (
              <img id="photoShow" src={ImageDataUrl} alt="uploaded" />
            )}

            <label id="Privacy" htmlFor="Privacy">
              {" "}
              Privacy:{" "}
            </label>
            <select
              id="postprivacy"
              value={PostPrivacy}
              onChange={(event) => setPostPrivacy(event.target.value)}
            >
              <option value="">-- Select --</option>
              <option value="Public">Public</option>
              <option value="Private">Private</option>
              <option value="Other">Optional</option>
            </select>

            <div onClick={handlePostSubmit} id="addPost" type="submit">
              <img id="addPostIcon" src="src/images/send.png" />
            </div>
          </div>
        </>
      ) : (
        <></>
      )}
    </>
  );
}

export default Posts;