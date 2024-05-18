import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Button, Modal, Select, Space, Table } from 'antd';
import { DollarOutlined } from '@ant-design/icons';
import RevenueOrder from './RevenueOrder';
export default function OrderStatistics() {
  const [orders, setOrders] = useState([]);
  const [status, setStatus] = useState(0); // Default status
  const [modalOpenStatistics, setModalOpenStatistics] = useState(false);

  useEffect(() => {
    axios.get('http://localhost:8080/manager/user/order/getlist', {
      params: {
        status: status
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
  }, [status]);

  const statusLabels = {
    9: 'Đã thanh toán',
    11: 'Đã hủy'
  };

  const handlerGetAllOrder = ()=>{
    setStatus(0);
  }


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
      title: 'Ngày tạo',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (text) => new Date(text * 1000).toLocaleString(),
    },
  ];

  return (
    <div>
      <h1>Thống kê đơn hàng</h1>
      <Space>
        <Space.Compact>
          <Select
             style={{
              width: 200,
            }}
            defaultValue = '9'
            onChange={(value) => setStatus(value)}
            options={[
              {
                value: '9',
                label: 'Đơn hàng đã thanh toán',
              },
              {
                value: '11',
                label: 'Đơn hàng đã hủy',
              },
            ]}
          />
          <Button onClick={handlerGetAllOrder}>Toàn bộ đơn hàng</Button>
          <Button  onClick={() => setModalOpenStatistics(true)}>Doanh thu<DollarOutlined /></Button>
          <Modal
              width={1000}
              title='Tính doanh thu'
              footer
              open={modalOpenStatistics}
              onOk={() => setModalOpenStatistics(false)}
              onCancel={() => setModalOpenStatistics(false)}
          >
            <RevenueOrder/>
          </Modal>
        </Space.Compact>
      </Space>
      <Table dataSource={orders} columns={columns} rowKey="id" />
    </div>
  );
}
