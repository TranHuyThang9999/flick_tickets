import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { showError, showSuccess } from '../common/log/log';
import { Form, Input, Button, DatePicker, InputNumber, Select } from 'antd';
import moment from 'moment';
import './index.css';
import MovieGetAll from '../common/MovieTypes/MovieGetAll';

export default function UpdateTicketById({ ticketId }) {

    const [ticket, setTicket] = useState(null);
    const [form] = Form.useForm();
    const [movieType, setMovieType] = useState([]);

    function convertToList(movieTypeStr) {
        // Chia chuỗi thành mảng bằng dấu phẩy và loại bỏ khoảng trắng
        let movieTypeList = movieTypeStr.split(',').map(type => type.trim());
        return movieTypeList;
    }
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/user/ticket?id=${ticketId}`)
                if (response.data.result.code === 0) {
                    setTicket(response.data.ticket);
                    return;
                } else {
                    showError('error server');
                    return;
                }
            } catch (error) {
                console.log(error);
                return;
            }
        }
        fetchData();
    }, [ticketId]);

    //  const optionsMovieType = [];
    const movieTypes = MovieGetAll();
    const optionsMovieType = movieTypes.map((movieType) => ({
        label: movieType.movieTypeName,
        value: movieType.movieTypeName,
    }));
    const handleFormSubmit = async (values) => {
        try {

            const formData = new FormData();

            formData.append('id', ticketId);
            formData.append('name', values.name);
            formData.append('price', values.price);
            formData.append('description', values.description);
            formData.append('sale', values.sale);
            formData.append('release_date', moment(values.release_date).unix());
            formData.append('status', values.status ? values.status.value : '');
            formData.append('movieDuration', values.movieDuration);
            formData.append('age_limit', values.age_limit);
            formData.append('director', values.director);
            formData.append('actor', values.actor);
            formData.append('producer', values.producer);
            formData.append('movie_type', movieType);


            // Send a POST request using Axios
            const response = await axios.put(
                'http://localhost:8080/manager/use/ticket/updates',
                formData,
                {
                    headers: {
                        // Authorization: token,
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );

            if (response.data.result.code === 0) {
                showSuccess('Cập nhật thành công');
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
    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };

    if (!ticket) {
        return null; // Or you can render a loading indicator
    }
    console.log(ticket);
    return (
        <div>
            <Form {...layout} form={form} onFinish={handleFormSubmit}>
                <Form.Item
                    label="Tên phim"
                    name="name"
                    initialValue={ticket.name}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Giá vé"
                    name="price"
                    initialValue={ticket.price}
                >
                    <InputNumber />
                </Form.Item>

                <Form.Item
                    label="Mô tả"
                    name="description"
                    initialValue={ticket.description}

                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Sale"
                    name="sale"
                    initialValue={ticket.sale}
                >
                    <InputNumber />
                </Form.Item>

                <Form.Item
                    label="Ngày phát hành"
                    name="release_date"
                // initialValue={moment.unix(ticket.release_date).format('YYYY-MM-DD HH:mm:ss')}
                >
                    <DatePicker
                        defaultValue={moment.unix(ticket.release_date)}
                        showTime={{ format: 'HH:mm:ss' }}
                    />
                </Form.Item>
                <Form.Item
                    label="Nhập trạng thái"
                    className="form-row"
                    name="status"
                >
                    <Select
                        defaultValue={ticket.status}
                        labelInValue

                        style={{
                            width: 120,
                            height: 42,
                        }}
                        options={[
                            {
                                value: 15,
                                label: 'Mở bán',
                            },
                            {
                                value: 17,
                                label: 'Đóng bán',
                            },
                        ]}
                    />
                </Form.Item>
                <Form.Item label="Thời lượng phim" name="movieDuration" initialValue={ticket.movieDuration}>
                    <InputNumber />
                </Form.Item>

                <Form.Item label="Giới hạn độ tuổi" name="age_limit" initialValue={ticket.age_limit}>
                    <InputNumber />
                </Form.Item>

                <Form.Item label="Đạo diễn" name="director" initialValue={ticket.director}>
                    <Input />
                </Form.Item>

                <Form.Item label="Diễn viên" name="actor" initialValue={ticket.actor}>
                    <Input />
                </Form.Item>

                <Form.Item label="Nhà sản xuất" name="producer" initialValue={ticket.producer}>
                    <Input />
                </Form.Item>
                <Form.Item
                    label="Thuộc thể loại phim"
                    className="form-row"
                    name="movie_type"
                >
                    <Select
                        defaultValue={convertToList(ticket.movie_type)}
                        allowClear
                        mode="multiple"
                        options={optionsMovieType} // Ensure correct variable is used here
                        onChange={(value) => setMovieType(value)}
                    />

                </Form.Item>
                <Form.Item >
                    <Button {...layout} type="primary" htmlType="submit">
                        Cập nhật
                    </Button>
                </Form.Item>
            </Form>
        </div>
    )
}
