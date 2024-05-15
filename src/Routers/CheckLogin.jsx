import axios from 'axios';
import React, { useState, useEffect } from 'react';
import HomeAdmin from '../dashboard/HomeaAdmin';
import PageForUser from '../Home/Page/PageForUser';

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
            localStorage.removeItem('email');
            localStorage.removeItem('user');
            localStorage.removeItem('user_name');
        }
    }, []);

    const checkToken = async (token) => {
        try {
            const response = await axios.get(`http://localhost:8080/manager/customer/${token}`);
            const { exp, id, user_name, role } = response.data;

            if (typeof exp !== 'number' || typeof id !== 'number' || typeof user_name !== 'string' || typeof role !== 'number') {
                clearLocalStorage();
                setLoading(false);
                return;
            }

            localStorage.setItem('user', JSON.stringify(response.data));
            localStorage.setItem('user_name',response.data.user_name);
            if (role === 1) {
                setIsLoggedIn(true);
                setIsLoginRole3(false);
            } else if (role === 13) {
                setIsLoggedIn(true);
                setIsLoginRole3(true);
            } else {
                setIsLoggedIn(false);
                setIsLoginRole3(false);
            }
            setLoading(false);
        } catch (error) {
            console.error(error);
            clearLocalStorage();
            setLoading(false);
        }
    };

    const clearLocalStorage = () => {
        localStorage.removeItem('email');
        localStorage.removeItem('user');
        localStorage.removeItem('user_name');
        localStorage.removeItem('token');
    };

    if (loading) {
        return <div>Loading...</div>;
    } else {
        if (isLoggedIn) {
            if (isLoginRole3) {
                return <PageForUser />;
            } else {
                return <HomeAdmin />;
            }
        } else {
            return <PageForUser />;
        }
    }
}
