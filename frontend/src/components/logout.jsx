import React from 'react';
import { useNavigate } from 'react-router-dom';

const Logout = () => {

    const navigate = useNavigate();
    const handleLogout = () => {

        fetch('http://localhost:9090/logout', {
            method: 'POST',
            credentials: 'include',
        })
            .then(() => {
            })
            .catch((error) =>
                console.error(error));
        navigate('/');
    };

    return (
        <div onClick={handleLogout}>
            <img id="logout" src="src/images/lock.png" alt="Logout" />
        </div>
    );
}

export default Logout;
