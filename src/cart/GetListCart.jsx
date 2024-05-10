import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Button, Space, Table } from 'antd';
import { showError, showWarning } from '../common/log/log';
import moment from 'moment';
import { ToolFilled, RedditCircleFilled } from '@ant-design/icons'; // Import các biểu tượng từ Ant Design
import DeleteCartbyId from './DeleteCartbyId';
import Cookies from 'js-cookie';

export default function GetListCart() {
    const [listCarts, setListCarts] = useState([]);
    const [selectedRow, setSelectedRow] = useState(null);
    const [loadingPayment, setLoadingPayment] = useState(false);
    const [user, setUser] = useState(null);

    const user_name = localStorage.getItem('user_name');


    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get("http://localhost:8080/manager/customer/user/profile", {
                    params: {
                        user_name: user_name
                    }
                });
                setUser(response.data.customer);
            } catch (error) {
                console.log(error);
            }
        };

        fetchData();
    }, []);


    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/cart/getlist', {
                    params: {
                        user_name: user_name
                    }
                });
                if (response.data.result.code === 0) {
                    setListCarts(response.data.carts);
                } else if(response.data.result.code === 20){
                    return;
                }
                else {
                    showError("Lỗi từ máy chủ");
                }
            } catch (error) {
                console.error(error);
                showError("Lỗi từ máy chủ");
            }
        };
        fetchData();
    }, [user_name]);

    const handleRowClick = (record) => {
        if (selectedRow === record.key) {
            setSelectedRow(null);
        } else {
            setSelectedRow(record.key);
        }
    };

    const handleDelete = (cartId) => {
        setListCarts(prevCarts => prevCarts.filter(cart => cart.id !== cartId));
    };

    const handleBuyNow = async (cart) => {
        setLoadingPayment(true);
        try {
            // Xác định dữ liệu cần thiết từ bản ghi giỏ hàng (cart)
            const seats = cart.seats_position.split(','); // Tách chuỗi ghế ngồi thành một mảng các ghế ngồi
            const items = seats.map(seat => ({ // Duyệt qua từng ghế ngồi trong mảng và tạo đối tượng item tương ứng
                name: "Vị trí ghế : " + seat,
                quantity: seats.length, // Số lượng ghế ngồi mỗi ghế là 1
                price: cart.price // Giá của mỗi ghế là giống nhau
            }));

            const requestData = {
                amount: cart.price,
                description: 'Xin Cảm ơn',
                items: items,
                seats: cart.seats_position,
                ShowTimeId: cart.show_time_id,
                cancelUrl: "http://localhost:8080/manager/public/customer/payment/calcel",
                returnUrl: "http://localhost:8080/manager/public/customer/payment/return",
                buyerName: "John Doe",
                buyerEmail: user.email,
                buyerPhone: user.phone_number, // Sử dụng số điện thoại đã nhập vào
                // Thêm các thông tin khác cần thiết cho request thanh toán
            };

            // Gửi request thanh toán
            const response = await axios.post('http://localhost:8080/manager/public/customer/payment/pay', requestData, {
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            // Xử lý kết quả thanh toán
            const paymentResult = response.data;
            if (paymentResult.resp_order === 44) {
                setLoadingPayment(false);
                showWarning("Ghế này đã được người mua trước vui lòng chọn lại");
                return;
            }

            localStorage.setItem("order_id", response.data.orderCode);
            Cookies.set("order_id", response.data.orderCode, { expires: 30 }); // Đặt cookie với thời gian sống là 1 tháng (30 ngày)

            if (paymentResult && paymentResult.checkoutUrl) {
                window.location.href = paymentResult.checkoutUrl;
            }
        } catch (error) {
            console.error(error);
            showError('Error server 1');
        }
        setLoadingPayment(false);
    };
    if (user_name === null) {
        return (
            <div>
                Vui lòng đăng nhập
                <RedditCircleFilled style={{color:'dodgerblue',fontSize:'30px'}} />
            </div>
        )
    }
    return (
        <Table
            columns={[
                { title: 'Tên phim', dataIndex: 'movie_name', key: 'movie_name' },
                { title: 'Vị trí ghế ngồi', dataIndex: 'seats_position', key: 'seats_position' },
                { title: 'Địa chỉ', dataIndex: 'cinema_name', key: 'cinema_name' },
                {
                    title: 'Thời gian chiếu', dataIndex: 'movie_time', key: 'movie_time',
                    render: (timestamp) => moment(timestamp).format('YYYY-MM-DD HH:mm:ss')
                },
                {
                    title: <ToolFilled />,
                    key: 'action',
                    render: (text, record) => (
                        <>
                            <Space>
                                <DeleteCartbyId cartId={record.id} onDelete={() => handleDelete(record.id)} />
                                <Button onClick={() => handleBuyNow(record)} disabled={loadingPayment}>
                                    {loadingPayment ? 'Đang thanh toán...' : 'Mua Ngay'}
                                </Button>
                            </Space>

                        </>
                    ),
                },
            ]}
            expandable={{
                expandedRowRender: (record) => (
                    <p style={{ margin: 0, color: 'dodgerblue', paddingLeft: '10px' }}>
                        {record.age_limit} | {record.description} | {record.conscious} | {record.price} VND|
                        {record.district} | {record.commune} | {record.address_details} | {record.movie_duration} | {record.age_limit}
                    </p>
                ),
                rowExpandable: (record) => record.name !== 'Not Expandable',
            }}
            dataSource={listCarts.map(item => ({ ...item, key: item.id }))}
            onRow={(record) => ({
                onClick: () => handleRowClick(record),
            })}
            rowClassName={(record) => (record.key === selectedRow ? 'selected-row' : '')}
        />
    );
}
