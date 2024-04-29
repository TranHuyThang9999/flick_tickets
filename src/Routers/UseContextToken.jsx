import React, { useEffect, useState } from 'react';
import axios from 'axios';

export default function UseContextToken() {
    const [data, setData] = useState('');

    useEffect(() => {
        const fetchData = async () => {
            try {
                const token = localStorage.getItem('token');
                const response = await axios.get(`http://localhost:8080/manager/customer/${token}`);
                if (response.data.Errors === 4) {
                } else {
                    setData(response.data);
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, []);



    return data;
}
