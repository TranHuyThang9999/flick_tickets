import axios from 'axios';
import  { useEffect, useState } from 'react'
import { showError, showWarning } from '../log/log';

export default function MovieGetAll() {

    const [movieType, setMovieType] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/user/movie/getlist');
                const data = response.data;
                if (data.result.code === 0) {
                    setMovieType(data.movie);
                    return;
                } else if (data.result.code === 20) {
                    showWarning("data emplty");
                    return;
                } else {
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
  return movieType;
}
