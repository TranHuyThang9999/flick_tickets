import  { useState, useEffect } from 'react';
import axios from 'axios';
import { showError, showWarning } from '../log/log';

export default function CinemasGetAll() {
    const [cinemas, setCinemas] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/user/get/cinema');
                const data = response.data;
                if (data.result.code === 0) {
                    setCinemas(data.cinemas);
                    return;
                }else if(data.Result.code===20){
                    showWarning("data emplty");
                    return;
                }else {
                    console.error('Error retrieving cinemas');
                    return;
                }
            } catch (error) {
                console.error('Error fetching data:', error);
                showError("error ver");
                return;
            }
        };

        fetchData();
    }, []);

    return cinemas;
}