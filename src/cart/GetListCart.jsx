import axios from 'axios';
import moment from 'moment';
import React, { useEffect, useState } from 'react';
import { showError } from '../common/log/log';
import { Table } from 'antd';
import DeleteCartbyId from './DeleteCartbyId';
import { ToolFilled } from '@ant-design/icons';
import './index.css';
export default function GetListCart() {
    const [carts, setCarts] = useState([]);
    const user_name = localStorage.getItem('user_name');
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/cart/getlist', {
                    params: {
                        user_name: user_name,
                    }
                });
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
        <div >
            <Table dataSource={carts}>
                <Table.Column title="ID" dataIndex="id" key="id" />
                <Table.Column title="Vị trí ghế" dataIndex="seats_position" key="seats_position" />
                <Table.Column title="Giá" dataIndex="price" key="price" />
                <Table.Column
                    title="Rạp chiếu phim"
                    dataIndex="cinema_name"
                    key="cinema_name"
                    render={(cinemaName, record) => (
                        <span>
                            {cinemaName} - {record.district}, {record.commune}, {record.address_details}
                        </span>
                    )}
                />
                <Table.Column title="Tên phim" dataIndex="movie_name" key="movie_name" />
                <Table.Column title="Thời lượng phim" dataIndex="movie_duration" key="movie_duration" />
                <Table.Column title="Giới hạn tuổi" dataIndex="age_limit" key="age_limit" />
                <Table.Column 
                    title="Thời gian chiếu" 
                    dataIndex="movie_time" 
                    key="movie_time" 
                    render={(movieTime) => (
                        <span>{moment(movieTime).format("YYYY-MM-DD HH:mm:ss")}</span>
                    )}
                />                <Table.Column
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
