import  { useState, useEffect } from 'react';
import axios from 'axios';

export default function CinemasGetAll() {
    const [cinemas, setCinemas] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/user/get/cinema');
                const data = response.data;
                if (data.result.code === 0) {
                    setCinemas(data.cinemas);
                } else {
                    console.error('Error retrieving cinemas:', data.Result.message);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    return cinemas;
}