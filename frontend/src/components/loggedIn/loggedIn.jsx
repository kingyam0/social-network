import React from "react";
import './loggedIn.css'
import RefreshPosts from "./refreshPost";
import CheckCookie from "../checkCookie";
import Navbar from "../navbar";

export default function LoggedIn({ setData, data, websocket }) {

    if (websocket) {
        websocket.onopen = () => {
            console.log('WEBSOCKET connected!');
            websocket.send('hello');
        }
    };

    return (
        //wrapped in a fragment (<></>)

        <>
            <CheckCookie setData={setData} data={data} />
            {data.User ? (
                <>
                    <Navbar data={data} />

                    <div className="PostBlock">
                        <div id="postTitleonPage">POSTS</div>
                        <RefreshPosts />
                    </div>
                </>
            ) : (
                <></>
            )}
        </>
    );
}