import { Button, Form, Input } from 'antd';
import React from 'react'
import { showError, showSuccess, showWarning } from '../log/log';
import  Axios  from 'axios';


export default function CinemasAdd() {

    const [form] = Form.useForm();

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();

            formData.append('cinema_name', values.cinema_name);
            formData.append('description', values.description);
            // Send a POST request using Axios

            const response = await Axios.post(
                'http://localhost:8080/manager/user/add/cinema',
                formData,
                {
                    headers: {
                        // Authorization: token,
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );
            if (response.data.result.code === 0) {
                showSuccess('Upload phong thành công');
                return;
            } else if (response.data.result.code === 30) {
                showWarning("tên phòng đã tồn tạo");
                return;
            } else {
                showError("error server");
                return;
            }
        } catch (error) {
            console.log(error);
            showError("error server 1");
            return;
        }

    }

    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };
    return (
        <div>
            <Form
                {...layout}
                form={form}
                onFinish={handleFormSubmit}
            >
                <Form.Item
                    label="Nhập tên phong"
                    className="form-row"
                    name="cinema_name"
                    rules={[{ required: true, message: 'Vui lòng nhập tên phong!' }]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Mô tả phòng"
                    className="form-row"
                    name="description"
                    rules={[{ required: true, message: 'Vui lòng nhập Mô tả phòng!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        </div>
    )
}
