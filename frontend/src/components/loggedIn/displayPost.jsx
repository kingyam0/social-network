import React, { useEffect, useState } from "react";
import expandComments from "../expandComments";

// const RefreshComments = () => {
//   const [Comments, setComments] = useState([]);

//   useEffect(() => {
//     fetchComments();
//   }, []);

// console.log(Comments);

// RefreshComments()
const DisplayPosts = ({posts}) => {
  console.log("POSTS IN DISPLAy POSTS", posts.PostID);
  const [comments, setComments] = useState([]);

  const convertDateTime = (date) => {
    const dateObject = new Date(date);
    const options = { timeStyle: "short", dateStyle: "short" };
    return dateObject.toLocaleString(undefined, options);
  };

  const [comContent, setcomContent] = useState({
    comContent: "",
    postId: 0,
  });

  const [error, setError] = useState();

  //   const toLoggedIn = () => {navigate('/loggedIn')}

  const handleCommentSubmit = () => {
    // navigate("/loggedIn");

    // console.log("COMMENT CONTENT: ", comContent);

    fetch("http://localhost:9090/comments", {
      method: "POST",
      body: JSON.stringify(comContent),
      header: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
      .then((response) => response.json())
      .then((comData) => {
        setcomData(comData);

        console.log("THIS IS DATA AFTER SETDATA:", comData);

        if (comData) {
          console.log("Comment Created Successfully.");
          // fetchComments(data);
        } else {
          console.loge(error("Comment creation failed. Please try again."));
        }
        return comData;
      })

      .catch((error) => {
        console.error(error);
      });
  };



//  useEffect(() => {
//   console.log("POST>POSTID", posts.postID);
//     getComments(posts.postID);
//   }, []);

  const renderPosts = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const reversedPosts = [...posts].reverse();
    const renderedPosts = [];

    // const reversedComments = [...comments].reverse();

    for (let i = 0; i < posts.length; i++) {
      const post = reversedPosts[i];
      console.log("DISPLAY POSTS", post);

        const getComments = (postID) => {
          console.group("getCommnets", postID);
          fetch("http://localhost:9090/sendComments", {
            headers: {
              Accept: "application/json",
              "Content-Type": "application/json",
            },
            method: "POST",
            body: JSON.stringify({
              postId: postID,
            }),
            credentials: "include",
          })
            .then((response) => response.json())
            .then((commentData) => {
              setComments(commentData);
              console.log("comments Received:", commentData);
            })
            .catch((error) => {
              console.log(error);
            });
        };

      const handleImageClick = (postId) => {
        if (reversedPosts[i]) {
          // console.log(postId);
          setcomContent((data) => {
            return { postId: postId, comContent: data.comContent };
          });
          setIsModalOpen(true);
        }
      };

      const closeModal = (postId) => {
        setIsModalOpen(false);
      };
      
      

      renderedPosts.push(
        <div className="posts1" id={post.PostID} key={post.PostID}>
          <div>
            <div className="author-category-wrap">
              <div className="name-timestamp-wrap">
                <div className="name"> Author: {post.Author}</div>
                <div className="title">Title: {post.PostTitle}</div>

                <div className="Privacy">Privacy: {post.PostPrivacy}</div>
              </div>
            </div>

            <div className="PostCate">Category: {post.PostCategory}</div>
          </div>
          <div className="PostContent">
            <div>{post.PostContent}</div>
          </div>
          <div>
            <img id="PhotoUpPost" src={post.ImageDataUrl} />
          </div>

          {/* ############################AddCommentIcon and Div################################################################## */}
          <div>
            <img
              id="addComIcon"
              src="src/images/chat3d.png"
              onClick={() => {
                console.log("Hello", post.PostID);
                getComments(post.PostID);
                handleImageClick(post.PostID);
              }}
            ></img>
            {isModalOpen && (
              <div className="modalComment">
                <div id="commentsTitle">COMMENTS</div>

                <div className="modalComment-Content">
                  {comments.map((comment) => {
                    {
                      console.log("DISPLAY COMMENTS", comment);
                    }
                    return (
                      <>
                        {" "}
                        <div key={comment.commentID} className="comment">
                          <div id="DisplayCommentBlock">
                            <div className="author">
                              Author: {comment.Author}
                            </div>
                            <div className="body">
                              Content{comment.comContent}
                            </div>
                            <div className="timestamp">
                              {convertDateTime(comment.CommentTime)}
                            </div>

                            {/* <div>
                              <img
                                src={post.User.Avata}
                                id="profile-picture"
                                width="35px"
                                alt="Profile"
                              />
                            </div> */}
                          </div>
                        </div>
                      </>
                    );
                  })}
                  ;
                  {
                    <>
                      {/* <div id="ComTitle">Comment Here...</div> */}
                      <div id="comBlock" className="comments">
                        <label id="comContent" htmlFor="comContent">
                          Comment Here:
                        </label>
                        <input
                          id="addCommentText"
                          type="text"
                          value={comContent.content}
                          onChange={(event) =>
                            setcomContent((prev) => {
                              return {
                                postId: prev.postId,
                                comContent: event.target.value,
                              };
                            })
                          }
                        />

                        <div id="addComment" type="submit">
                          <img
                            id="addCommentIcon"
                            src="src/images/plus.jpg"
                            onClick={() => handleCommentSubmit(post.postID)}
                          />
                        </div>
                      </div>
                    </>
                  }
                  <a onClick={closeModal}>
                    <img id="closeBtn" src="src/images/close.png" />
                  </a>
                </div>
              </div>
            )}
          </div>
          {/* ################################AddCommentIcon and Div################################################################## */}

          <div className="timestamp">
            Created: {convertDateTime(post.PostTime)}
          </div>
        </div>
      );
    }

    return renderedPosts;
  };

  return <div className="post">{renderPosts()}</div>;
};

export default DisplayPosts;
