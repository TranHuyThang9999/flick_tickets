import axios from 'axios';
import React, { useState, useEffect } from 'react';
import GetListTicket from '../Home/Tickets/GetListTicket';
import HomeAdmin from '../dashboard/HomeaAdmin';
import FormLogin from '../dashboard/FormLogin';

export default function CheckLogin() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isLoginRole3, setIsLoginRole3] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            checkToken(token);
        } else {
            setLoading(false);
        }
    }, []);

    const checkToken = async (token) => {
        try {
            const response = await axios.get(`http://localhost:8080/manager/customer/${token}`);
            if (response.data.Errors === 4) {
                setLoading(false);
            } else {
                localStorage.setItem('user', JSON.stringify(response.data));
                if (response.data.role === 3) {
                    setIsLoggedIn(true);
                    setIsLoginRole3(false);
                } else if (response.data.role === 13) {
                    setIsLoggedIn(true);
                    setIsLoginRole3(true);
                } else {
                    setIsLoggedIn(false);
                    setIsLoginRole3(false);
                }
                setLoading(false);
            }
        } catch (error) {
            console.error(error);
            setLoading(false);
        }
    }

    if (loading) {
        return <div>Loading...</div>;
    } else {
        if (isLoggedIn) {
            if (isLoginRole3) {
                return <GetListTicket />;
            } else {
                return <HomeAdmin />;
            }
        } else {
            return <FormLogin />;
        }
    }
}
