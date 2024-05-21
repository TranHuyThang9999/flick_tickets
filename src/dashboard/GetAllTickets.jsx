import { Button, Drawer, Space, Table } from 'antd';
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import GetTicketById from './GetTicketById';
import './index.css';
import DeleteTicketById from './DeleteTicketById';
import UpdateTicketById from './UpdateTicketById';

export default function GetAllTicketsForAdmin() {
    const [tickets, setTickets] = useState([]);
    const [open, setOpen] = useState(false);
    const [selectedTicketId, setSelectedTicketId] = useState(null);
    const [openUpdateTicket, setOpenUpdateTicket] = useState([]);

    useEffect(() => {
        axios
            .get('http://localhost:8080/manager/customers/ticket')
            .then(response => {
                const listTickets = response.data.list_tickets;
                setTickets(listTickets);
                // Khởi tạo mảng openUpdateTicket với giá trị mặc định là false cho mỗi vé
                setOpenUpdateTicket(new Array(listTickets.length).fill(false));
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);

    const handleDeleteTicketById = (ticketId) => {
        setTickets(prevCarts => prevCarts.filter(ticket => ticket.id !== ticketId));
    };

    const showDrawer = (id) => {
        setSelectedTicketId(id);
        setOpen(true);
    };

    const onClose = () => {
        setOpen(false);
    };

    const showDrawerUpdate = (index) => {
        setOpenUpdateTicket(prevState => {
            const newState = [...prevState];
            newState[index] = true;
            return newState;
        });
    };

    const onCloseUpdate = (index) => {
        setOpenUpdateTicket(prevState => {
            const newState = [...prevState];
            newState[index] = false;
            return newState;
        });
    };

    return (
        <>
            <Table dataSource={tickets} rowKey="id">
                <Table.Column title="Id" dataIndex="id" key="id" />
                <Table.Column title="Tên phim" dataIndex="name" key="name" />
                <Table.Column title="Mô tả" dataIndex="description" key="description" />
                <Table.Column title='Giá vé' dataIndex='price' key='price' />
                <Table.Column
                    title="Chi tiết"
                    key="details"
                    render={(_, record, index) => (
                        <Button type="link" onClick={() => showDrawer(record.id)}>
                            Xem chi tiết vé
                        </Button>
                    )}
                />
                <Table.Column
                    title='Action'
                    key='operation'
                    render={(_, record, index) => (
                        <>
                            <Space direction="horizontal" size="middle">
                                <Button
                                    style={{ width: '160px' }}
                                    onClick={() => showDrawerUpdate(index)} // Truyền index vào showDrawerUpdate
                                >
                                    Chỉnh sửa thông tin vé
                                </Button>
                                <DeleteTicketById ticketId={record.id} onDelete={() => handleDeleteTicketById(record.id)} />

                            </Space>

                            <Drawer
                                width={800}
                                onClose={() => onCloseUpdate(index)} // Truyền index vào onCloseUpdate
                                open={openUpdateTicket[index]} // Sử dụng openUpdateTicket[index]
                            >
                                <UpdateTicketById ticketId={record.id} />
                            </Drawer>
                        </>
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
                {selectedTicketId && (
                    <div>
                        <GetTicketById id={selectedTicketId} />
                    </div>
                )}
            </Drawer>
        </>
    );
}
