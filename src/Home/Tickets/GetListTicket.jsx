import React, { useEffect, useState } from 'react';
import './index.css';
import axios from 'axios';
import { showError, showWarning } from '../../common/log/log';
import { Button, Form, Input } from 'antd';

export default function GetListTicket() {
  const [tickets, setTickets] = useState([]);
  const [form] = Form.useForm();

  const handleSubmitSearch = async (values) => {
    try {
      const response = await axios.get(
        'http://localhost:8080/manager/customers/ticket',
        {
          params: values,
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        }
      );

      if (response.data.result.code === 0) {
        setTickets(response.data.list_tickets);
      } else if (response.data.result.code === 20) {
        showWarning('Không tìm thấy data');
      } else {
        showError('Lỗi server');
      }
    } catch (error) {
      showError('Lỗi server');
    }
  };

  console.log(tickets);

  return (
    <div className="wrapper">
      <Form form={form} onFinish={handleSubmitSearch}>
        <Form.Item name="id">
          <Input placeholder="ID" />
        </Form.Item>
        <Form.Item name="name">
          <Input placeholder="Name" />
        </Form.Item>
        <Form.Item name="price">
          <Input placeholder="Price" />
        </Form.Item>
        <Form.Item name="sale">
          <Input placeholder="Sale" />
        </Form.Item>
        <Form.Item name="status">
          <Input placeholder="Status" />
        </Form.Item>
        <Form.Item name="movieDuration">
          <Input placeholder="Movie Duration" />
        </Form.Item>
        <Form.Item name="age_limit">
          <Input placeholder="Age Limit" />
        </Form.Item>
        <Form.Item name="director">
          <Input placeholder="Director" />
        </Form.Item>
        <Form.Item name="actor">
          <Input placeholder="Actor" />
        </Form.Item>
        <Form.Item name="producer">
          <Input placeholder="Producer" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Search
          </Button>
        </Form.Item>
      </Form>
        {/* Display the tickets */}
        <div style={{backgroundColor:'red'}} className="ticket-list">
        {tickets.map((ticket) => (
          <div key={ticket.id} className="ticket-item">
            {/* Render ticket details */}
            <p>ID: {ticket.id}</p>
            <p>Name: {ticket.name}</p>
            {/* Render other ticket details */}
          </div>
        ))}
      </div>
    </div>
  );
}