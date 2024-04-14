import { Table } from 'antd';
import React, { useEffect, useState } from 'react';
import { showError, showWarning } from '../log/log';

export default function DetailedShowSchedule({ id }) {
  const [showTimeTicket, setShowTimeTicket] = useState([]);

  //   useEffect(() => {


  //   fetchData();
  // }, [id]);
  useEffect(() => {
    fetchData();
    }, [id]);
  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/manager/user/getlist/time?id=${id}`);
      const data = await response.json();
      setShowTimeTicket(data.showtimes);
      if (data.result.code === 0) {
      } else if (data.result.code === 20) {
        showWarning("Không tìm thấy bản ghi nào");
        return;
      }
      else if (data.result.code === 4) {
        showError("Lỗi server vui lòng thử lại");
        return;
      }

    } catch (error) {
      console.error('Error:', error);
      showError("Lỗi server vui lòng thử lại", error);
      return;
    }
  };

  function formatTimestamp(timestamp) {
    const date = new Date(timestamp * 1000); // Nhân 1000 để chuyển đổi từ milliseconds sang seconds

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  }


  const columns = [
    {
      title: 'Mã vé',
      dataIndex: 'ticket_id',
      key: 'ticket_id',
    },
    {
      title: 'Phòng để chiếu',
      dataIndex: 'cinema.cinema_name',
      key: 'cinema_name',
    },
    {
      title: 'Thời gian chiếu phim',
      dataIndex: 'movie_time',
      key: 'movie_time',
      render: (movie_time) => formatTimestamp(movie_time),
    },
    {
      title: 'Mô tả',
      dataIndex: 'description',
      key: 'description',
    },
    {
      title: 'Tỉnh',
      dataIndex: 'conscious',
      key: 'conscious',
    },
    {
      title: 'Huyện',
      dataIndex: 'district',
      key: 'district',
    },
    {
      title: 'Xã/Phường',
      dataIndex: 'commune',
      key: 'commune',
    },
    {
      title: 'Địa chỉ Chi tiết',
      dataIndex: 'address_details',
      key: 'address_details',
    },
  ];

  return (
    <div>
      <Table scroll={{ x: 90 }} dataSource={showTimeTicket} columns={columns} />
    </div>
  );
}
