import React, { useEffect, useState } from 'react';
import DisplayPosts from './displayPost';

const RefreshPosts = () => {
  const [posts, setPosts] = useState([]);
    // const [comments, setComments] = useState([]);

  useEffect(() => {
    fetchPosts();
  }, []);

  // console.log(posts);

  const fetchPosts = () => {
    fetch("http://localhost:9090/getLatestPosts", {
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      method: "POST",
    })
      .then((response) => response.json())
      .then((data) => {
        setPosts(data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div>
      {posts.length === 0 ? (
        <p>Loading posts...</p>
      ) : (
        <DisplayPosts posts={posts} />
      )}
    </div>
  );
}

export default RefreshPosts;
