import React, { useEffect, useState } from 'react';
import { showError, showSuccess, showWarning } from '../common/log/log';
import { Form, DatePicker, Upload, Input, Button, Select, InputNumber } from 'antd';
import Axios from 'axios';
import CinemasGetAll from '../common/cinemas/CinemasGetAll';
import "./index.css";
import moment from "moment";

export default function AdminUploadTickets() {

    const [cinemaName, setCinemaName] = useState('');
    const [timestampsList, setTimestampsList] = useState([]);

    const [form] = Form.useForm();
    const [fileList, setFileList] = useState([]);

    const onChange = ({ fileList }) => {
        setFileList(fileList);
    };
    //date
    const handleDateChange = (date, dateString) => {
        if (date && moment(date).isValid()) {
            setTimestampsList((prevTimestampsList) => [
                ...prevTimestampsList,
                moment(dateString).unix(),
            ]);
        }
    };

    console.log("time : ", timestampsList);

    //
    const options = [];
    const cinemas = CinemasGetAll();
    for (let index = 0; index < cinemas.length; index++) {
        options.push({
            label: cinemas[index].cinema_name,
            value: cinemas[index].cinema_name,
        })

    }

    console.log("cinema:", cinemaName);
    const optionsGetTime = timestampsList.map((timestamp) => ({
        value: timestamp,
        label: moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }));
    const handleFormSubmit = async (values) => {
        try {
            const releaseDateTimestamp = moment(values.release_date).unix();

            // const token = localStorage.getItem('token');
            const formData = new FormData();
            formData.append('name', values.name);
            formData.append('price', values.price);
            formData.append('quantity', values.quantity);
            formData.append('description', values.description);
            formData.append('showtime', values.showtime);
            formData.append('status', values.status ? values.status.value : '');
            formData.append('sale', values.sale);
            formData.append('release_date', releaseDateTimestamp);
            formData.append('cinema_name', cinemaName);// sua lai
            formData.append('movie_time', timestampsList);

            fileList.forEach((file) => {
                formData.append('file', file.originFileObj);
            });

            // Send a POST request using Axios
            const response = await Axios.post(
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
            } else if (response.data.result.code == 26) {
                showWarning('các phòng vé đã tồn tại thời gian chiếu phim  vui lòng chọn phòng khác hoặc thời gian khác');
            } else {
                showError('Lỗi server, vui lòng thử lại');
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

    const optionsGetTimeSelect = timestampsList.map(timestamp => ({
        value: timestamp,
        label: moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
    }));

    return (
        <div>
            <Form {...layout} form={form} className="form-container" onFinish={handleFormSubmit}>
                <Form.Item
                    label="Nhập tên sản phẩm"
                    className="form-row"
                    name="name"
                    rules={[{ required: true, message: 'Vui lòng nhập tên sản phẩm!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Nhập giá sản phẩm"
                    className="form-row"
                    name="price"
                    rules={[{ required: true, message: 'Vui lòng nhập giá sản phẩm!' }]}
                >
                    <InputNumber />
                </Form.Item>

                <Form.Item
                    label="Nhập số lượng sản phẩm"
                    className="form-row"
                    name="quantity"
                    rules={[{ required: true, message: 'Vui lòng nhập số lượng sản phẩm!' }]}
                >
                    <InputNumber />
                </Form.Item>

                <Form.Item
                    label="Nhập mô tả sản phẩm"
                    className="form-row"
                    name="description"
                    rules={[{ required: true, message: 'Vui lòng nhập mô tả sản phẩm!' }]}
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

                        style={{
                            width: 120,
                            height: 42,
                        }}
                        options={[
                            {
                                value: '13',
                                label: 'Mở bán',
                            },
                            {
                                value: '15',
                                label: 'Đóng bán',
                            },
                        ]}
                    />
                </Form.Item>

                <Form.Item
                    label="Nhập giảm giá"
                    className="form-row"
                    name="sale"
                >
                    <InputNumber />
                </Form.Item>

                <Form.Item
                    label="Ngày phát hành"
                    className="form-row"
                    name="release_date"
                >
                    <DatePicker showTime />
                </Form.Item>

                <Form.Item
                    label="Thời gian chiếu phim"
                    className="form-row"
                    name="movie_time"
                    rules={[{ required: true, message: 'Vui lòng nhập thời lượng phim!' }]}
                >
                    <div className='showTime'>
                        <DatePicker
                            showTime
                            onChange={handleDateChange}
                            picker="datetime"
                            size="small"
                        />
                        <Select
                            allowClear
                            mode="multiple"
                            placeholder="Please select"
                            options={optionsGetTimeSelect}
                        />
                    </div>
                </Form.Item>

                <Form.Item
                    label="Nhập tên rạp"
                    className="form-row"
                    name="cinema_name"
                // rules={[{ required: true, message: 'Vui lòng nhập tên rạp!' }]}
                >
                    <Select
                        allowClear
                        mode='multiple'
                        options={options}
                        onChange={(value) => setCinemaName(value)}

                    />

                </Form.Item>

                <Form.Item
                    label="Nhập ảnh mô tả sản phẩm"
                    className="form-row"
                    name="file"
                    rules={[{ required: true, message: 'Vui lòng chọn ảnh mô tả sản phẩm!' }]}
                >
                    <Upload
                        fileList={fileList}
                        listType="picture-card"
                        accept="image/jpeg,image/png"
                        onChange={onChange}
                        beforeUpload={() => false} // Prevent auto-upload
                    >
                        {'+ Upload'}
                    </Upload>
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}