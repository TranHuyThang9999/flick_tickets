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

    const columns = [
        {
            title: 'Id',
            dataIndex: 'id',
            key: 'id',
        },
        {
            title: 'Tên vé',
            dataIndex: 'name',
            key: 'name',
        },
        {
            title: 'Số vé tối đa',
            dataIndex: 'max_ticket',
            key: 'max_ticket',
        },
        {
            title: 'Số lượng',
            dataIndex: 'quantity',
            key: 'quantity',
        },
        {
            title: 'Mô tả',
            dataIndex: 'description',
            key: 'description',
        },
        // {
        //     title: 'Giảm giá',
        //     dataIndex: 'sale',
        //     key: 'sale',
        // },
        // {
        //     title: 'Ngày phát hành',
        //     dataIndex: 'release_date',
        //     key: 'release_date',
        // },
        // {
        //     title: 'Trạng thái',
        //     dataIndex: 'status',
        //     key: 'status',
        // },
        // {
        //     title: 'Ghế đã chọn',
        //     dataIndex: 'selected_seat',
        //     key: 'selected_seat',
        // },
        // {
        //     title: 'Thời lượng phim',
        //     dataIndex: 'movieDuration',
        //     key: 'movieDuration',
        // },
        {
            title: 'Chi tiết',
            key: 'details',
            render: (_, record) => (
                <Button type="link" onClick={() => showDrawer(record.id)}>
                    Xem chi tiết
                </Button>
            ),
        },
    ];

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
            <Table dataSource={tickets} columns={columns} rowKey="id" />

            <Drawer

                title="Chi tiết vé"
                width={1300}
                onClose={onClose}
                visible={open}
                bodyStyle={{
                    paddingBottom: 80
                }}
            >
                {selectedTicketId && <GetTicketById id={selectedTicketId} />}
            </Drawer>
        </>
    );
}