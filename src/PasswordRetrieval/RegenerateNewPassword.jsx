import { Button, Form, Input } from 'antd';
import React from 'react';
import { showError, showSuccess } from '../common/log/log';
import axios from 'axios';

export default function RegenerateNewPassword() {
    const [form] = Form.useForm();

    const layout = {
        labelCol: { span: 8 },
        wrapperCol: { span: 16 },
    };

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', localStorage.getItem('user_name'));
            formData.append('new_password', values.new_password);

            const response = await axios.put('http://localhost:8080/manager/customer/reset/password', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });
            if (response.data.result.code === 0) {
                showSuccess('Cập nhật password thành công');
                window.location.reload();
                return;
            } else {
                showError('Error server');
                return;
            }
        } catch (error) {
            console.error(error);
            showError('Lỗi kết nối đến máy chủ');
        }
    };

    return (
        <div>
            <Form
                {...layout}
                form={form}
                className="form-container-register-with-otp"
                onFinish={handleFormSubmit}
            >
                <Form.Item
                    name='new_password'
                    label='Tạo lại mật khẩu'
                    rules={[
                        { required: true, message: 'Vui lòng nhập mật khẩu mới!' },
                    ]}
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item
                    name='confirm_password'
                    label='Nhập lại mật khẩu'
                    dependencies={['new_password']}
                    rules={[
                        { required: true, message: 'Vui lòng nhập lại mật khẩu!' },
                        ({ getFieldValue }) => ({
                            validator(_, value) {
                                if (!value || getFieldValue('new_password') === value) {
                                    return Promise.resolve();
                                }
                                return Promise.reject(new Error('Hai mật khẩu không khớp!'));
                            },
                        }),
                    ]}
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Tiếp tục
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
