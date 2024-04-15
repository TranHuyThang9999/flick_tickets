import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showError, showWarning } from '../log/log';

export default function CinemaGetAllByRoomName({ roomName }) {

    const [cinema, setCinema] = useState([]);


    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/user/getByName?name=${roomName}`);
                const data = await response.json();
                if (data.result.code === 0) {
                    setCinema(data.cinema);
                    return;
                }
                else if (data.result.code === 20) {
                    showWarning("Không tìm thấy bản ghi nào");
                    return;
                } else {
                    showError("Lỗi server");
                    return;
                }
            } catch (error) {
                console.error('Lỗi server:', error);
                return;
            }
        };

        fetchData();
    }, [roomName]);


    if (!cinema.id) {
        return <div>Đang tải...</div>;
    }

    return cinema;
}
