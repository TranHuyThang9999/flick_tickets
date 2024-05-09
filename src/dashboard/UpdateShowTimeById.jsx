import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showError, showSuccess } from '../common/log/log';
import { Button, DatePicker, Form, InputNumber, Select } from 'antd';
import moment from 'moment';
import CinemasGetAll from '../common/cinemas/CinemasGetAll';

export default function UpdateShowTimeById({ show_time_id }) {

    const [showTime, setShowTime] = useState(null);
    const [cinemaName, setCinemaName] = useState('');
    const [form] = Form.useForm();

    useEffect(() => {
        const fetchdata = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/use/showtime?id=${show_time_id}`);
                if (response.data.result.code === 0) {
                    setShowTime(response.data.show_time);
                    return;
                } else {
                    showError("error server");
                    return;
                }
            } catch (error) {
                console.log(error);
                showError("error server");
                return;
            }
        }
        fetchdata();
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
            formData.append('price', values.price);
            formData.append('movie_time', values.movie_time);
            formData.append('quantity', values.quantity);

            const response = await axios.put(
                'http://localhost:8080/manager/use/ticket/updates',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );

            if (response.data.result.code === 0) {
                showSuccess('Cập nhật thành công');
            } else {
                showError('Lỗi server, vui lòng thử lại');
            }
        } catch (error) {
            console.log(error);
            showError('Lỗi server, vui lòng thử lại');
        }
    };
    
    const options = [];
    const cinemas = CinemasGetAll();
    for (let index = 0; index < cinemas.length; index++) {
        options.push({
            label: cinemas[index].cinema_name,
            value: cinemas[index].cinema_name,
        });
    }

    if (!showTime) {
        return null;
    }

    return (
        <div>
            <Form {...layout} form={form} onFinish={handleFormSubmit}>
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
                        mode='multiple'
                        options={options}
                        onChange={(value) => setCinemaName(value)}
                        maxCount={1}
                    />
                </Form.Item>
                <Form.Item
                    label="Thời gian chiếu"
                    name="movie_time"
                >
                    <DatePicker
                        allowClear
                        defaultValue={moment.unix(showTime.movie_time)}
                        showTime
                    />
                </Form.Item>
                <Button type="primary" htmlType="submit">Cập nhật lại suất chiếu</Button>
            </Form>
        </div>
    )
}
