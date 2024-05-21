import axios from 'axios';
import  { useEffect, useState } from 'react';
import { showError } from '../common/log/log';

export default function GetAllShowTime({ ticket_id }) {
    const [showTimes, setShowTimes] = useState([]);

    useEffect(() => {
        const fetchShowTimes = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/user/getlist/time/admin?id=${ticket_id}`);
                if (response.data.result.code === 0) {
                    setShowTimes(response.data.showtimes);
                    console.log("11",response);
                } else {
                    showError("Error fetching show times: " + response.data.result.message);
                }
            } catch (error) {
                console.error('Error fetching show times:', error);
                showError("Error fetching show times");
            }
        };

        if (ticket_id) {
            fetchShowTimes();
        }
    }, [ticket_id]);

    return showTimes;
}
