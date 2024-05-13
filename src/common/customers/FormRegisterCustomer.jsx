import { Button, Form, Input, Upload } from 'antd';
import axios from 'axios';
import React, { useState } from 'react';
import { showError, showSuccess, showWarning } from '../log/log';
import { WeiboCircleFilled } from '@ant-design/icons';
import './index.css';
import FormLogin from '../../dashboard/FormLogin';

export default function FormRegisterCustomer() {
    const [form] = Form.useForm();
    const [imageFile, setImageFile] = useState(null);
    const [goBackFormLogin, setGoBackFormLogin] = useState(false);
    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', values.user_name);
            formData.append('password', values.password);
            formData.append('address', values.address);
            formData.append('age', values.age);
            formData.append('email', values.email);
            formData.append('phone_number', values.phone_number);
            formData.append('file', imageFile);

            const response = await axios.post(
                'http://localhost:8080/manager/customer/user/register',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                    // Authorization: token,
                }
            );

            if (response.data.result.code === 0) {
                showSuccess('Tạo tài khoản thành công');
            } else if (response.data.result.code === 22) {
                showWarning('Tên tài khoản đã tồn tại');
            } else {
                showError('Lỗi server, vui lòng thử lại');
            }
        } catch (error) {
            console.log(error);
            showError('Lỗi server, vui lòng thử lại');
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
    const handlerGoback = () => {
        setGoBackFormLogin(true);
    }
    if (goBackFormLogin) {
        return (
            <FormLogin />
        )
    }
    return (
        <div>
            <div><WeiboCircleFilled /></div>
            <Form {...layout} form={form} className="form-container-register-customer" onFinish={handleFormSubmit}>
                <Form.Item
                    label="Tên tài khoản"
                    name="user_name"
                    rules={[
                        {
                            required: true,
                            message: 'Vui lòng nhập tên tài khoản!',
                        },
                    ]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Mật khẩu"
                    name="password"
                    rules={[
                        {
                            required: true,
                            message: 'Vui lòng nhập mật khẩu!',
                        },
                    ]}
                >
                    <Input.Password />
                </Form.Item>

                <Form.Item
                    label="Địa chỉ"
                    name="address"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Tuổi"
                    name="age"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Email"
                    name="email"
                    rules={[
                        {
                            required: true,
                            type: 'email',
                            message: 'Email không hợp lệ!',
                        },
                    ]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Số điện thoại"
                    name="phone_number"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Ảnh đại diện"
                    name="image"
                    valuePropName="fileList"
                    getValueFromEvent={(e) => {
                        if (Array.isArray(e)) {
                            return e;
                        }
                        return e && e.fileList;
                    }}
                >
                    <Upload
                        maxCount={1}
                        type=''
                        listType='picture-card'
                        openFileDialogOnClick
                        accept="image/jpeg,image/png"
                        beforeUpload={(file) => {
                            setImageFile(file);
                            return false;
                        }}
                        onRemove={() => {
                            setImageFile(null);
                        }}
                    >
                        +Upload
                    </Upload>
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button onClick={handlerGoback} style={{ marginRight: '10px' }}>Quay lại</Button>
                    <Button type="primary" htmlType="submit">
                        Đăng ký
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
