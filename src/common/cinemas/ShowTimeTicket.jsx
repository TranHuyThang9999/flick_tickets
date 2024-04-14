import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showError, showWarning } from '../log/log';
import { Table } from 'antd';
export default function ShowTimeTicket({ id }) {

    const [showTime, setShowTime] = useState([]);
    try {
        useEffect(() => {
            const fetchData = async () => {
                try {
                    const response = await axios.get(`http://localhost:8080/manager/user/getlist/time?id=${id}`);
                    const data = response.data;
                    if (data.result.code === 0) {
                        setShowTime(data.showtimes);
                        return;
                    } else if (data.result.code === 20) {
                        showWarning("data emplty");
                        return;
                    } else {
                        console.error('Error retrieving showtimes');
                        return;
                    }
                } catch (error) {
                    console.error('Error fetching data:', error);
                    showError("error ver 1");
                    return;
                }
            };

            fetchData();
        }, [id]);
    } catch (error) {
    }
    function formatTimestamp(timestamp) {
        const date = new Date(timestamp * 1000); // Nhân 1000 để chuyển đổi từ milliseconds sang seconds
      
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
      
        return `${year}-${month}-${day} ${hours}:${minutes}`;
      }
      
    console.log(showTime);
    const columns = [
        { title: 'ID', dataIndex: 'id', key: 'id' },
        { title: 'Tên rạp', dataIndex: 'cinema_name', key: 'cinema_name' },
        { title: 'Thời gian chiếu', dataIndex: 'movie_time', key: 'movie_time' },
        { title: 'Ngày tạo', dataIndex: 'created_at', key: 'created_at',  render: (created_at) => formatTimestamp(created_at),  },
    ];
  
return (
    <div>
        <Table dataSource={showTime} scroll={{ x: 100 }} columns={columns} />
    </div>
)
}
