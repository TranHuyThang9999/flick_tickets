import axios from 'axios';
import moment from 'moment';
import React, { useEffect, useState } from 'react';
import { showError } from '../common/log/log';
import { Table } from 'antd';
import DeleteCartbyId from './DeleteCartbyId';
import { ToolFilled } from '@ant-design/icons';

export default function GetListCart() {
    const [carts, setCarts] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/cart/getlist');
                if (response.data.result.code === 0) {
                    setCarts(response.data.carts);
                } else {
                    showError("Lỗi cơ sở dữ liệu");
                }
            } catch (error) {
                showError("Lỗi cơ sở dữ liệu");
            }
        };
        fetchData();
    }, []);

    const handleDelete = (cartId) => {
        setCarts(prevCarts => prevCarts.filter(cart => cart.id !== cartId));
    };

    return (
        <div>
            <Table dataSource={carts}>
                <Table.Column title="ID" dataIndex="id" key="id" />
                <Table.Column title="Tên người dùng" dataIndex="user_name" key="user_name" />
                <Table.Column title="ID Lịch chiếu" dataIndex="show_time_id" key="show_time_id" />
                <Table.Column title="Vị trí ghế" dataIndex="seats_position" key="seats_position" />
                <Table.Column title="Giá" dataIndex="price" key="price" />
                <Table.Column
                    title="Ngày tạo"
                    dataIndex="created_at"
                    key="created_at"
                    render={(created_at) => moment.unix(created_at).format('DD/MM/YYYY HH:mm:ss')}
                />
                <Table.Column
                    title={<ToolFilled />}
                    dataIndex="action_delete"
                    key="action_delete"
                    render={(text, record) => (
                        <DeleteCartbyId onDelete={handleDelete} cartId={record.id} />
                    )}
                />
            </Table>
        </div>
    );
}
