import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showError, showSuccess, showWarning } from '../common/log/log';
import { Button, DatePicker, Drawer, Form, InputNumber, Select } from 'antd';
import CinemasGetAll from '../common/cinemas/CinemasGetAll';
import './index.css';

export default function UpdateShowTimeById({ show_time_id }) {

    const [showTime, setShowTime] = useState(null);
    const [cinemaName, setCinemaName] = useState('');
    const [form] = Form.useForm();
    const [visible, setVisible] = useState(false); // State để kiểm soát trạng thái mở của Drawer

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/use/showtime?id=${show_time_id}`);
                if (response.data.result.code === 0) {
                    setShowTime(response.data.show_time);
                } else {
                    showError("error server");
                }
            } catch (error) {
                console.log(error);
                showError("error server");
            }
        }
        fetchData();
    }, [show_time_id]);

    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('id', show_time_id);
            formData.append('cinema_name', cinemaName);
            formData.append('movie_time', values.movie_time.unix());
            formData.append('quantity', values.quantity);
            formData.append('price', values.price);
            formData.append('discount', values.discount);

            const response = await axios.put(
                'http://localhost:8080/manager/user/showtime/update',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );

            if (response.data.result.code === 0) {
                showSuccess('Cập nhật thành công');
            } else if (response.data.result.code === 26) {
                showWarning("Suất chiếu đã tồn tại vui lòng chọn lại");
            } else {
                showError('Lỗi server, vui lòng thử lại');
            }
        } catch (error) {
            console.log(error);
            showError('Lỗi server, vui lòng thử lại');
        }
    };

    const options = CinemasGetAll().map(cinema => ({
        label: cinema.cinema_name,
        value: cinema.cinema_name,
    }));

    if (!showTime) {
        return null;
    }

    const handleUpdateClick = () => {
        setVisible(true); // Mở Drawer khi nhấn nút cập nhật
    };

    const handleCloseDrawer = () => {
        setVisible(false); // Đóng Drawer khi cần
    };

    return (
        <div>
            <Button onClick={handleUpdateClick}>Cập nhật</Button>
            <Drawer
                title="Cập nhật suất chiếu"
                width={500}
                onClose={handleCloseDrawer}
                visible={visible} // Trạng thái mở của Drawer được kết nối với state
                bodyStyle={{ paddingBottom: 80 }}
            >
                <Form style={{ width: '600px' }} {...layout} form={form} className="form-container-update-show-time" onFinish={handleFormSubmit}>
                    <Form.Item
                        initialValue={showTime.price}
                        label="Nhập giá vé"
                        className="form-row"
                        name="price"
                    >
                        <InputNumber />
                    </Form.Item>

                    <Form.Item
                        initialValue={showTime.quantity}
                        label="Nhập số lượng vé trên 1 phòng"
                        className="form-row"
                        name="quantity"
                    >
                        <InputNumber />
                    </Form.Item>
                    <Form.Item
                        label="Chọn phòng để chiếu phim"
                        className="form-row"
                        name="cinema_name"
                    >
                        <Select
                            defaultValue={showTime.cinema_name}
                            allowClear
                            options={options}
                            onChange={(value) => setCinemaName(value)}
                        />
                    </Form.Item>
                    <Form.Item
                        label='Giảm giá vé'
                        initialValue={showTime.discount}
                        name="discount"
                    >
                        <InputNumber />
                    </Form.Item>
                    <Form.Item
                        label="Thời gian chiếu"
                        name="movie_time"
                    >
                        <DatePicker
                            
                            // defaultValue={showTime.movie_time ? moment.unix(showTime.movie_time) : null}
                            showTime
                        />
                    </Form.Item>
                    <Button type="primary" htmlType="submit">Cập nhật lại suất chiếu</Button>
                </Form>
            </Drawer>

        </div>
    )
}
