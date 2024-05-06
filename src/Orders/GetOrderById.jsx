import { Table, Button } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError } from '../common/log/log';
import moment from 'moment';

export default function GetOrderById() {
  const [orders, setOrders] = useState([]);
  const [selectedRow, setSelectedRow] = useState(null);
  const [page, setPage] = useState(1);
  const pageSize = 10;

  useEffect(() => {
    fetchData();
  }, [page]);

  const fetchData = async () => {
    try {
      const response = await axios.get('http://localhost:8080/manager/user/order/getlist', {
        params: {
          offset: (page - 1) * pageSize,
          limit: pageSize,
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

  const handleNextPage = () => {
    setPage(page + 1);
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
              {record.addressDetails} | {record.movie_time}
            </p>
          ),
          rowExpandable: (record) => record.name !== 'Not Expandable',
        }}
        onRow={(record) => ({
          onClick: () => handleRowClick(record),
        })}
        rowClassName={(record) => (record.key === selectedRow ? 'selected-row' : '')}
        pagination={false}
      />
      <Button onClick={handleNextPage}>Next</Button>
    </div>
  );
}
