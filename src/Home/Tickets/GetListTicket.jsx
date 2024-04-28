import React, { useState } from 'react';
import './index.css';
import { showError, showWarning } from '../../common/log/log';
import { Button, Drawer, Form, Input, Table } from 'antd';
import FetchTicketsdata from './FetchTicketsdata';
import GetTicketById from './DetailTicketById';

export default function GetListTicket() {

  const [open, setOpen] = useState(false);
  const [selectedTicketId, setSelectedTicketId] = useState(null);

  const showDrawer = (record) => {
    setSelectedTicketId(record.id);
    setOpen(true);
  };

  const onClose = () => {
    setOpen(false);
  };


  const [tickets, setTickets] = useState([]);
  const [form] = Form.useForm();

  const handleSubmitSearch = async (values) => {
    try {
      const result = await FetchTicketsdata(values);
      if (result.code === 0) {
        setTickets(result.data);
      } else if (result.code === 20) {
        showWarning('Không tìm thấy dữ liệu');
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
      <div>
        <Form className='form-find-ticket' form={form} onFinish={handleSubmitSearch}>
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
      </div>
      <div>
        {/* Display the tickets */}
        <div className="ticket-list" style={{ overflow: 'auto' }}>
          {
            tickets.length > 0 && (
              <div>
                <Table dataSource={tickets}>
                  <Table.Column title="Id" dataIndex="id" key="id" />
                  <Table.Column title="Tên vé" dataIndex="name" key="name" />
        
                  <Table.Column title="Mô tả" dataIndex="description" key="description" />

                  <Table.Column title="Chi tiết" key="details"
                    render={(_, record) => (
                      <Button type="link" onClick={() => showDrawer(record)}>
                      Xem chi tiết
                      </Button>
                    )}
                  />

                </Table>
                <Drawer
                  title="Chi tiết vé"
                  width={1300}
                  onClose={onClose}
                  visible={open}
                  bodyStyle={{
                    paddingBottom: 80
                  }}
                >
                  <div>
                  <GetTicketById id={selectedTicketId} />
                  </div>
                </Drawer>
              </div>

            )
          }
        </div>
      </div>

    </div>
  );
}