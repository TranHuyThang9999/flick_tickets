import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Button, Input, Modal, Select, Space, Table, DatePicker } from 'antd';
import { DollarOutlined } from '@ant-design/icons';
import RevenueOrder from './RevenueOrder';
import StatisticalBar from './StatisticalBar';

export default function OrderStatistics() {
  const [orders, setOrders] = useState([]);
  const [status, setStatus] = useState(0); // Default status
  const [modalOpenStatistics, setModalOpenStatistics] = useState(false);
  const [modalOpenStatisticsbar, setModalOpenStatisticsbar] = useState(false);
  const [searchParams, setSearchParams] = useState({
    id: '',
    movie_name: '',
    price: '',
    cinema_name: '',
    movie_time: '',
    email: '',
    created_at: null
  });

  useEffect(() => {
    fetchOrders();
  }, [status]);

  const fetchOrders = (params = {}) => {
    axios.get('http://localhost:8080/manager/user/order/getlist', {
      params: {
        status: status,
        id: params.id || '',
        movie_name: params.movie_name || '',
        price: params.price || '',
        cinema_name: params.cinema_name || '',
        email: params.email || '',
        created_at: params.created_at ? params.created_at.unix() : 0
      }
    })
    .then(response => {
      if (response.data.result.code === 0) {
        setOrders(response.data.orders);
      }
    })
    .catch(error => {
      console.error('There was an error fetching the order list!', error);
    });
  };

  const statusLabels = {
    9: 'Đã thanh toán',
    11: 'Đã hủy'
  };

  const handleGetAllOrders = () => {
    setStatus(0);
  };

  const handleSearch = () => {
    fetchOrders(searchParams);
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setSearchParams(prevParams => ({
      ...prevParams,
      [name]: value
    }));
  };


  const columns = [
    {
      title: 'ID Đơn hàng',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: 'Tên phim',
      dataIndex: 'movie_name',
      key: 'movie_name',
    },
    {
      title: 'Tên rạp',
      dataIndex: 'cinema_name',
      key: 'cinema_name',
    },
    {
      title: 'Ngày phát hành',
      dataIndex: 'release_date',
      key: 'release_date',
      render: (text) => new Date(text * 1000).toLocaleString(),
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
    },
    {
      title: 'Trạng thái',
      dataIndex: 'status',
      key: 'status',
      render: (status) => statusLabels[status] || status,
    },
    {
      title: 'Giá',
      dataIndex: 'price',
      key: 'price',
    },
    {
      title: 'Ghế',
      dataIndex: 'seats',
      key: 'seats',
    },
    {
      title: 'Giảm giá',
      dataIndex: 'sale',
      key: 'sale',
    },
    {
      title: 'Thời gian chiếu phim',
      dataIndex: 'movie_time',
      key: 'movie_time',
      render: (text) => new Date(text * 1000).toLocaleString(),
    },
    {
      title: 'Ngày mua',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (text) => new Date(text * 1000).toLocaleString(),
    },
  ];

  return (
    <div>
      <h1>Thống kê đơn hàng</h1>
      <Space direction="vertical" style={{ width: '100%' }}>
        <Space>
          <Select
            style={{ width: 200 }}
            defaultValue='9'
            onChange={(value) => setStatus(value)}
            options={[
              { value: '9', label: 'Đơn hàng đã thanh toán' },
              { value: '11', label: 'Đơn hàng đã hủy' },
            ]}
          />
          <Button onClick={handleGetAllOrders}>Toàn bộ đơn hàng</Button>
          <Button onClick={() => setModalOpenStatistics(true)}>
            Doanh thu phòng vé<DollarOutlined />
          </Button>
          <Button onClick={()=>setModalOpenStatisticsbar(true)}>Thống kê doanh thu vé đã bán theo thời gian</Button>
          <Modal
            width={1000}
            title='Tính doanh thu'
            footer={null}
            open={modalOpenStatistics}
            onOk={() => setModalOpenStatistics(false)}
            onCancel={() => setModalOpenStatistics(false)}
          >
            <RevenueOrder />
          </Modal>
          <Modal
            width={1200}
            footer
            open={modalOpenStatisticsbar}
            onOk={() => setModalOpenStatisticsbar(false)}
            onCancel={() => setModalOpenStatisticsbar(false)}
          >
            <StatisticalBar/>
          </Modal>
        </Space>
        <Space>
          <Input
            name="id"
            placeholder='Nhập mã đơn hàng'
            onChange={handleInputChange}
          />
          <Input
            name="movie_name"
            placeholder='Nhập tên phim'
            onChange={handleInputChange}
          />
          <Input
            name="price"
            placeholder='Nhập giá'
            onChange={handleInputChange}
          />
          <Input
            name="cinema_name"
            placeholder='Nhập tên phòng'
            onChange={handleInputChange}
          />
          <Input
            name="email"
            placeholder='Email khách hàng'
            onChange={handleInputChange}
          />
          
          <Button onClick={handleSearch}>Tìm kiếm</Button>
        </Space>
        <Table dataSource={orders} columns={columns} rowKey="id" />
      </Space>
    </div>
  );
}
