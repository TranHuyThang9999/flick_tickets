import { Table } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import moment from 'moment';
import { RedditCircleFilled } from '@ant-design/icons';
import { showError } from '../../common/log/log';
import './index.css';

export default function PurchaseHistory() {
    const [orders, setOrders] = useState([]);
    const [loading, setLoading] = useState(false);
    const [currentPage, setCurrentPage] = useState(1);

    useEffect(() => {
        fetchData();
    }, [currentPage]);

    const fetchData = async () => {
        try {
            setLoading(true);
            const email = localStorage.getItem('email');
            if (localStorage.getItem('email')===null){
              return;
            }
            const response = await axios.get('http://localhost:8080/manager/user/order/history', {
                params: {
                    email: email,
                },
            });
            setLoading(false);
            if (response.data.result.code === 0) {
                setOrders(response.data.order_history_entities);
            } else {
                showError(response.data.result.message);
            }
        } catch (error) {
            setLoading(false);
            showError("Network error");
        }
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

    const formatAddressDetails = (addressDetails) => {
        try {
            const parsedAddress = JSON.parse(addressDetails);
            const { cinema_name, description, conscious, district, commune, address_details } = parsedAddress;
            return `${cinema_name}, ${description}, ${conscious}, ${district}, ${commune}, ${address_details}`;
        } catch (error) {
            console.error('Error parsing address details:', error);
            return 'Thông tin địa chỉ không hợp lệ';
        }
    };

    const columns = [
        {
            title: 'Mã đơn hàng',
            dataIndex: 'id',
            key: 'id',
        },
        {
            title: 'Tên phim',
            dataIndex: 'movie_name',
            key: 'movie_name',
        },
        {
            title: 'Tại rạp',
            dataIndex: 'cinema_name',
            key: 'cinema_name',
        },
        {
            title: 'Trạng Thái',
            dataIndex: 'status',
            key: 'status',
            render: (status) => getStatusText(status),
        },
        {
            title: 'Giá',
            dataIndex: 'price',
            key: 'price',
        },
    ];

    if (!localStorage.getItem('user_name')) {
        return (
            <div>
                Vui lòng đăng nhập
                <RedditCircleFilled style={{ color: 'dodgerblue', fontSize: '30px' }} />
            </div>
        );
    }

    return (
        <div>
            <Table
                columns={columns}
                dataSource={orders}
                loading={loading}
                expandable={{
                    expandedRowRender: (record) => (
                        <p style={{ margin: 0, color: 'dodgerblue', paddingLeft: '10px' }}>
                            {formatAddressDetails(record.address_details)} | Thời gian chiếu : {moment.unix(record.movie_time).format('YYYY-MM-DD HH:mm:ss')} | Thời gian phát hành : {moment.unix(record.release_date).format('YYYY-MM-DD HH:mm:ss')} | Thời gian mua: {moment.unix(record.created_at).format('YYYY-MM-DD HH:mm:ss')}
                        </p>
                    ),
                    rowExpandable: (record) => record.name !== 'Not Expandable',
                }}
                pagination={{
                    current: currentPage,
                    onChange: (page) => setCurrentPage(page),
                }}
            />
        </div>
    );
}
