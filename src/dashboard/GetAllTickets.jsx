import { Button, Drawer, Table } from 'antd';
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import GetTicketById from './GetTicketById';
import './index.css';

export default function GetAllTickets() {
    const [tickets, setTickets] = useState([]);
    const [open, setOpen] = useState(false);
    const [selectedTicketId, setSelectedTicketId] = useState(null);
    const [size, setSize] = useState();

    useEffect(() => {
        axios
            .get('http://localhost:8080/manager/customers/ticket')
            .then(response => {
                const listTickets = response.data.list_tickets;
                setTickets(listTickets);
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);

    console.log(tickets);


    const showDrawer = (id) => {
        setSelectedTicketId(id);
        setSize('large');
        setOpen(true);
    };

    const onClose = () => {
        setOpen(false);
    };

    return (
        <>
            <Table dataSource={tickets} rowKey="id">
                <Table.Column title="Id" dataIndex="id" key="id" />
                <Table.Column title="Tên vé" dataIndex="name" key="name" />
                <Table.Column
                    title="Số vé tối đa"
                    dataIndex="max_ticket"
                    key="max_ticket"
                />
                <Table.Column
                    title="Số lượng"
                    dataIndex="quantity"
                    key="quantity"
                />
                <Table.Column
                    title="Mô tả"
                    dataIndex="description"
                    key="description"
                />
                <Table.Column
                    title="Chi tiết"
                    key="details"
                    render={(_, record) => (
                        <Button type="link" onClick={() => showDrawer(record.id)}>
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
                {selectedTicketId && (
                    <div>
                        <GetTicketById id={selectedTicketId} />
                    </div>
                )}
            </Drawer>
        </>
    );
}