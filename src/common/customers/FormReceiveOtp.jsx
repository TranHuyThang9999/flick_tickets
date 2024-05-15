import { Button, Form, InputNumber } from 'antd';
import axios from 'axios';
import React from 'react';
import './index.css';
import { GooglePlusCircleFilled } from '@ant-design/icons';
import { showError, showSuccess, showWarning } from '../log/log';

export default function FormReceiveOtp() {
    const [form] = Form.useForm();

    const layout = {
        labelCol: { span: 8 },
        wrapperCol: { span: 16 },
    };

    const handleFormSubmit = async (values) => {
        const email = localStorage.getItem('email');
        try {
            const formData = new FormData();
            formData.append('email', email);
            formData.append('otp', values.otp);

            const response = await axios.post('http://localhost:8080/manager/customer/verify/', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });

            if (response.data.result.code === 0) {
                localStorage.setItem('email',email);
                showSuccess('Đăng nhập thành công');

                const responseAuth = await axios.get(`http://localhost:8080/manager/customer/auth2?email=${email}`)
                if(responseAuth.data.result.code ===0){
                    localStorage.setItem('token',responseAuth.data.jwt_token.refresh_token);
                    window.location.reload();
                    return;
                }else{
                    showWarning('error client');
                    return;
                }
            } else if (response.data.result.code === 22) {
                showWarning('Mã OTP không hợp lệ');
                return;
            }
            else {
                showError('Lỗi từ máy chủ');
                return;
            }
        } catch (error) {
            console.error(error);
            showError('Lỗi kết nối đến máy chủ');
        }
    };

    return (
        <div>
            <Form {...layout} form={form} className="form-container-register-with-otp" onFinish={handleFormSubmit}>
                <GooglePlusCircleFilled style={{ fontSize: '19px' }} />
                <Form.Item
                    label='OTP'
                    name='otp'
                    rules={[{ required: true, message: 'Vui lòng nhập OTP được gửi về địa chỉ email!' }]}
                >
                    <InputNumber style={{ width: '100%' }} />
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Đăng nhập
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
