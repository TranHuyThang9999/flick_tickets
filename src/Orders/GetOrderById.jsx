import { Table, Button } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError } from '../common/log/log';
import moment from 'moment';

export default function GetOrderById() {
  const [orders, setOrders] = useState([]);
  const [selectedRow, setSelectedRow] = useState(null);
  const [pagination, setPagination] = useState({ current: 1, pageSize: 10 });

  useEffect(() => {
    fetchData();
  }, [pagination]);

  const fetchData = async () => {
    try {
      const response = await axios.get('http://localhost:8080/manager/user/order/getlist', {
        params: {
          offset: (pagination.current - 1) * pagination.pageSize,
          limit: pagination.pageSize,
        }
      });

      if (response.data.result.code === 0) {
        setOrders(response.data.orders);
      } else if (response.data.result.code === 20) {
        // Handle code 20
      } else if (response.data.result.code === 4) {
        showError("Server error");
      } else {
        showError("Unknown error");
      }
    } catch (error) {
      showError("Network error");
    }
  };

  const handleRowClick = (record) => {
    setSelectedRow(record.key);
    // Do something with the clicked row
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
  }
  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'Release Date',
      dataIndex: 'release_date',
      key: 'release_date',
      render: (timestamp) => moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss'),
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: (status) => getStatusText(status),
    },
    {
      title: 'Price',
      dataIndex: 'price',
      key: 'price',
    },
    {
      title: 'Sale',
      dataIndex: 'sale',
      key: 'sale',
    },
    {
      title: 'addressDetails',
      dataIndex: 'addressDetails',
      key: 'addressDetails',
      render: (addressDetails) => addressDetails ? JSON.parse(addressDetails).cinema_name : null,
    },
    {
      title: 'Created At',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (timestamp) => moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss'),
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
        onRow={(record) => ({
          onClick: () => handleRowClick(record),
        })}
        rowClassName={(record) => (record.key === selectedRow ? 'selected-row' : '')}
        pagination={pagination}
        onChange={(pagination) => setPagination(pagination)}
      />
    </div>
  );
}
