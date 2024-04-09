import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showWarning } from '../log/log';
import { Image } from 'antd';

export default function DisplayTickets() {

    const [tickets, setTickets] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/customers/ticket');
                const data = response.data;
                if (response.data.result.code === 0) {
                    setTickets(data.tickets);
                    return;
                } else if (response.data.result.code === 20) {
                    showWarning("Không thấy bản ghi nào");
                    return;
                } else {
                    console.error('Error retrieving cinemas:', response.data.result.message);
                    return;
                }
            } catch (error) {
                console.error('Error fetching data:', error);
                return;
            }
        };

        fetchData();
    }, []);
    function formatTimestamp(timestamp) {
        const date = new Date(timestamp * 1000); // Nhân 1000 để chuyển đổi từ milliseconds sang seconds

        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}`;
    }
    return (
        <div>
            {tickets.map((item) => (
                <div key={item.ticket.id}>
                    <p>Ticket ID: {item.ticket.id}</p>
                    <p>Ticket Name: {item.ticket.name}</p>
                    <p>Price: {item.ticket.price}</p>
                    {/* ... và các thông tin khác về vé */}

                    <h4>Show Times:</h4>
                    <ul>
                        {item.show_times.map((showTime) => (
                            <li key={showTime.id}>
                                <p>Show Time ID: {showTime.id}</p>
                                <p>Cinema Name: {showTime.cinema_name}</p>
                                <p>Movie Time: {formatTimestamp(showTime.movie_time)}</p>                                {/* ... và các thông tin khác về showtime */}
                            </li>
                        ))}
                    </ul>

                    <h4>List URLs:</h4>
                    <ul>
                        {item.list_url.map((url) => (
                            <li key={url.id}>
                                <p>URL ID: {url.id}</p>
                                <p>
                                    <Image width={100} src={url.url} alt="Ticket Image" />

                                </p>
                                {/* ... và các thông tin khác về URL */}
                            </li>
                        ))}
                    </ul>
                </div>
            ))}
        </div>
    )
}
