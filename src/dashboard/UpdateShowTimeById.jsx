import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showError, showSuccess, showWarning } from '../common/log/log';
import { Button, DatePicker, Drawer, Form, InputNumber, Select } from 'antd';
import moment from 'moment';
import CinemasGetAll from '../common/cinemas/CinemasGetAll';
import './index.css';
export default function UpdateShowTimeById({ show_time_id }) {

    const [showTime, setShowTime] = useState(null);
    const [cinemaName, setCinemaName] = useState('');
    const [form] = Form.useForm();
    const [visible, setVisible] = useState(false); // State để kiểm soát trạng thái mở của Drawer


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

            const releaseDateTimestamp = moment(values.movie_time).unix();

            const formData = new FormData();
            formData.append('id', show_time_id);
            formData.append('cinema_name', cinemaName);
            formData.append('movie_time', releaseDateTimestamp);
            formData.append('quantity', values.quantity);
            formData.append('price', values.price);

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
                return;
            } else if (response.data.result.code === 26) {
                showWarning("Suất chiếu đã tồn tại vui lòng chọn lại");
                return;
            } else {
                showError('Lỗi server, vui lòng thử lại');
                return;
            }
        } catch (error) {
            console.log(error);
            showError('Lỗi server, vui lòng thử lại');
            return;
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
            </Drawer>

        </div>
    )
}
