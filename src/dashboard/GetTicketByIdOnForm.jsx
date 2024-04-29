import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showWarning, showError, showSuccess } from '../common/log/log';
import { Button, DatePicker, Form, Input, InputNumber, Select } from 'antd';
import CinemasGetAll from '../common/cinemas/CinemasGetAll';
import MovieGetAll from '../common/MovieTypes/MovieGetAll';
import moment from 'moment';

export default function GetTicketByIdOnForm({ id }) {

    const [ticket, setTicket] = useState(null);
    const [cinemaName, setCinemaName] = useState('');
    const [movieType, setMovieType] = useState([]);
    const [timestampListAdd, setTimestampListAdd] = useState([]);
    const [timestampsList, setTimestampsList] = useState([]);
    const [cinemaOptions, setCinemaOptions] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(
                    `http://localhost:8080/manager/user/ticket?id=${id}`
                );

                const { result, ticket: fetchedTicket } = response.data;

                if (result.code === 0) {
                    setTicket(fetchedTicket);
                } else if (result.code === 14) {
                    showWarning('error client');
                } else {
                    showError('error server');
                }
            } catch (error) {
                showError('error server');
                console.log(error);
            }
        };

        fetchData();
    }, [id]);

    //
    //date
    const handleDateChange = (date, dateString) => {
        if (date && moment(date).isValid()) {
            setTimestampsList((prevTimestampsList) => [
                ...prevTimestampsList,
                moment(dateString).unix(),
            ]);
        }
    };

    //
    const options = [];
    const cinemas = CinemasGetAll();
    for (let index = 0; index < cinemas.length; index++) {
        options.push({
            label: cinemas[index].cinema_name,
            value: cinemas[index].cinema_name,
        })

    }
    //  const optionsMovieType = [];
    const movieTypes = MovieGetAll();
    const optionsMovieType = movieTypes.map((movieType) => ({
        label: movieType.movieTypeName,
        value: movieType.movieTypeName,
    }));


    const [form] = Form.useForm();

    const handleFormSubmit = async (values) => {
        try {
            const releaseDateTimestamp = moment(values.release_date).unix();

            // const token = localStorage.getItem('token');
            const formData = new FormData();
            formData.append('name', values.name);
            formData.append('price', values.price);
            formData.append('quantity', values.quantity);
            formData.append('description', values.description);
            formData.append('status', values.status ? values.status.value : '');
            formData.append('sale', values.sale);
            formData.append('release_date', releaseDateTimestamp);
            formData.append('cinema_name', cinemaName);// sua lai
            formData.append('movie_time', timestampListAdd);
            //
            formData.append('movie_duration', values.movie_duration);
            formData.append('age_limit', values.age_limit);
            formData.append('director', values.director);
            formData.append('actor', values.actor);
            formData.append('producer', values.producer);
            formData.append('movie_type', movieType);

            console.log(timestampListAdd);

            // Send a POST request using Axios
            const response = await axios.post(
                'http://localhost:8080/manager/user/upload/ticket',
                formData,
                {
                    headers: {
                        // Authorization: token,
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );

            if (response.data.result.code === 0) {
                showSuccess('Upload vé thành công');
                return;
            } else if (response.data.result.code === 26) {
                showWarning('các phòng vé đã tồn tại thời gian chiếu phim  vui lòng chọn phòng khác hoặc thời gian khác');
                return;
            } else if (response.data.result.code === 28) {
                showError('lỗi phía máy khách');
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


    function convertToList(movieTypeStr) {
        // Chia chuỗi thành mảng bằng dấu phẩy và loại bỏ khoảng trắng
        let movieTypeList = movieTypeStr.split(',').map(type => type.trim());
        return movieTypeList;
    }

    useEffect(() => {
        const getRoomName = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/user/getlist/time?id=${id}`)
                if (response.data.result.code === 0) {
                    setCinemaOptions(response.data.showtimes);
                } else {
                    showError("error server 1");
                    console.error('Error fetching cinema options:', response.data.result.message);
                    return;
                }
            } catch (error) {
                console.error('Error fetching cinema options:', error);
                showError("error server");
            }
        }

        if (id) {
            getRoomName();
        }
    }, [id]);



    const defaultValueRoom = Array.from(new Set(cinemaOptions.map((item) => item.cinema_name))).map((cinema_name) => ({
        label: cinema_name,
        value: cinema_name,
    })).map(option => option.value);

    function formatTimestamp(timestamp) {
        const date = new Date(timestamp * 1000); // Nhân 1000 để chuyển đổi từ milliseconds sang seconds

        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}`;
    }
    
    const defaultValueShowTime = Array.from(new Set(cinemaOptions.map((item) => item.movie_time))).map((movie_time) => ({
        label: movie_time,
        value: movie_time,
    })).map(option => formatTimestamp(option.value));

    const optionsGetTimeSelect = timestampsList.map(timestamp => ({
        value: timestamp,
        // label: moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
        label: formatTimestamp(timestamp) // Sử dụng hàm formatTimestamp thay vì moment.unix
    }));



    return (
        <div>
            {ticket && (
                <div>
                    <Form {...layout} form={form} className="form-container" onFinish={handleFormSubmit}>
                        <Form.Item
                            label="Nhập tên vé"
                            className="form-row"
                            name="name"
                            initialValue={ticket.name}
                            rules={[{ required: true, message: 'Vui lòng nhập tên vé!' }]}
                        >
                            <Input />
                        </Form.Item>

                        <Form.Item
                            label="Nhập giá vé"
                            className="form-row"
                            name="price"
                            initialValue={ticket.price}
                            rules={[{ required: true, message: 'Vui lòng nhập giá vé!' }]}
                        >
                            <InputNumber />
                        </Form.Item>

                        <Form.Item
                            label="Nhập số lượng vé trên 1 phòng"
                            className="form-row"
                            name="quantity"
                            initialValue={ticket.quantity}
                            rules={[{ required: true, message: 'Vui lòng nhập số lượng vé!' }]}
                        >
                            <InputNumber />
                        </Form.Item>

                        <Form.Item
                            label="Nhập mô tả vé"
                            className="form-row"
                            name="description"
                            initialValue={ticket.description}
                            rules={[{ required: true, message: 'Vui lòng nhập mô tả vé!' }]}
                        >
                            <Input />
                        </Form.Item>


                        <Form.Item
                            label="Nhập trạng thái"
                            className="form-row"
                            name="status"
                        >
                            <Select
                                labelInValue
                                defaultValue={ticket.status}
                                style={{
                                    width: 120,
                                    height: 42,
                                }}
                                options={[
                                    {
                                        value: '15',
                                        label: 'Mở bán',
                                    },
                                    {
                                        value: '17',
                                        label: 'Đóng bán',
                                    },
                                ]}
                            />
                        </Form.Item>

                        <Form.Item
                            label="Nhập giảm giá"
                            className="form-row"
                            name="sale"
                            initialValue={ticket.sale}
                        >
                            <InputNumber />
                        </Form.Item>

                        <Form.Item
                            label="Ngày phát hành"
                            className="form-row"
                            name="release_date"
                            initialValue={moment(ticket.release_date)}
                        >
                            <DatePicker showTime />
                        </Form.Item>

                        <Form.Item
                            label="Lịch chiếu phim"
                            className="form-row"
                            name="movie_time"
                        // rules={[{ required: true, message: 'Vui lòng nhập thời lượng phim!' }]}
                        >
                            <div className='showTime'>
                                <DatePicker
                                    showTime
                                    onChange={handleDateChange}
                                    picker="datetime"
                                    size="small"
                                />
                                <Select
                                    defaultValue={defaultValueShowTime}
                                    allowClear
                                    mode="multiple"
                                    placeholder="Please select"
                                    options={optionsGetTimeSelect}
                                    onChange={(value) => setTimestampListAdd(value)}
                                />
                            </div>
                        </Form.Item>

                        <Form.Item
                            label="Chọn phòng để chiếu phim"
                            className="form-row"
                            name="cinema_name"
                        >
                            <Select
                                defaultValue={defaultValueRoom}
                                allowClear
                                mode="multiple"
                                options={options}
                                onChange={(value) => setCinemaName(value)}
                            />
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

                        <Form.Item
                            label="Thời lượng phim"
                            name="movie_duration"
                            initialValue={ticket.movieDuration}
                            rules={[{ required: true }]}>
                            <InputNumber />
                        </Form.Item>

                        <Form.Item
                            label="Giới hạn độ tuổi"
                            name="age_limit"
                            initialValue={ticket.age_limit}
                            rules={[{ required: true }]}>
                            <InputNumber />
                        </Form.Item>

                        <Form.Item
                            label="Đạo diễn"
                            name="director"
                            initialValue={ticket.director}
                            rules={[{ required: true }]}>
                            <Input />
                        </Form.Item>

                        <Form.Item
                            label="Diễn viên"
                            name="actor"
                            initialValue={ticket.actor}
                            rules={[{ required: true }]}>
                            <Input />
                        </Form.Item>

                        <Form.Item
                            label="Nhà sản xuất"
                            name="producer"
                            initialValue={ticket.producer}
                            rules={[{ required: true }]}>
                            <Input />
                        </Form.Item>



                        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                            <Button type="primary" htmlType="submit">
                                Submit
                            </Button>
                        </Form.Item>
                    </Form>
                </div>
            )}
        </div>
    );
}