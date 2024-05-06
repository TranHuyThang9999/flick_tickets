import { Table, Pagination } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError } from '../common/log/log';
import moment from 'moment';

export default function GetOrderById() {
  const [orders, setOrders] = useState([]);
  const [total, setTotal] = useState(0);
  const [limit, setLimit] = useState(10);
  const [currentPage, setCurrentPage] = useState(1);

  useEffect(() => {
    fetchData();
  }, [currentPage]); // Gọi fetchData khi trang hiện tại thay đổi

  const fetchData = async () => {
    try {
      const offset = (currentPage - 1) * limit;
      const response = await axios.get('http://localhost:8080/manager/user/order/getlist', {
        params: {
          offset: offset,
          limit: limit,
        }
      });

      if (response.data.result.code === 0) {
        setOrders(response.data.orders);
        setTotal(response.data.total);
      } else if (response.data.result.code === 20) {
        // Xử lý code 20
      } else if (response.data.result.code === 4) {
        showError("Server error");
      } else {
        showError("Unknown error");
      }
    } catch (error) {
      showError("Network error");
    }
  };

  const getStatusText = (status) => {
    switch (status) {
      case 7:
        return 'Khởi tạo đơn hàng';
      case 9:
        return 'Đã thanh toán';
      case 11:
        return 'Đã hủy';
      default:
        return 'Unknown';
    }
  };

  const formatAddressDetails = (addressDetails) => {
    try {
      const parsedAddress = JSON.parse(addressDetails);
      const { cinema_name, description, conscious, district, commune, address_details } = parsedAddress;
      return `${cinema_name}, ${description}, ${conscious}, ${district}, ${commune}, ${address_details}`;
    } catch (error) {
      console.error('Error parsing address details:', error);
      return 'Thông tin địa chỉ không hợp lệ';
    }
  };

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      ellipsis: true,
      width: '10%'
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
      ellipsis: true,
      width: '20%'
    },
    {
      title: 'Ngày Phát Hành',
      dataIndex: 'release_date',
      key: 'release_date',
      render: (timestamp) => moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss'),
      ellipsis: true,
      width: '15%'
    },
    {
      title: 'Trạng Thái',
      dataIndex: 'status',
      key: 'status',
      render: (status) => getStatusText(status),
      ellipsis: true,
      width: '10%'
    },
    {
      title: 'Giá',
      dataIndex: 'price',
      key: 'price',
      ellipsis: true,
      width: '10%'
    },
    {
      title: 'Giảm Giá %',
      dataIndex: 'sale',
      key: 'sale',
      ellipsis: true,
      width: '10%'
    },
    {
      title: 'Chi Tiết Địa Chỉ',
      dataIndex: 'addressDetails',
      key: 'addressDetails',
      render: (addressDetails) => addressDetails ? JSON.parse(addressDetails).cinema_name : null,
      ellipsis: true,
      width: '20%'
    },
    {
      title: 'Đã Tạo Lúc',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (timestamp) => moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss'),
      ellipsis: true,
      width: '15%'
    },
  ];



  return (
    <div>
      <Table
        columns={columns}
        dataSource={orders.map((order, index) => ({ ...order, key: index }))}
        expandable={{
          expandedRowRender: (record) => (
            <p style={{ margin: 0, color: 'dodgerblue', paddingLeft: '10px' }}>
              {formatAddressDetails(record.addressDetails)} | {record.movie_time}
            </p>
          ),
          rowExpandable: (record) => record.name !== 'Not Expandable',
        }}
        pagination={false} // Tắt phân trang mặc định của Table
      />
      <Pagination
        total={total}
        defaultPageSize={limit}
        current={currentPage}
        onChange={(page) => setCurrentPage(page)}
      />
    </div>
  );
}
