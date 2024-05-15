import { Button, Form, InputNumber } from 'antd';
import axios from 'axios';
import React, { useState } from 'react';
import './index.css';
import { GooglePlusCircleFilled } from '@ant-design/icons';
import { showError, showWarning } from '../common/log/log';
import RegenerateNewPassword from './RegenerateNewPassword';

export default function FormCheckOtp() {
    const [form] = Form.useForm();
    const [isNext, setIsNext] = useState(false);

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
                setIsNext(true);
                return;
            } else if (response.data.result.code === 22) {
                showWarning('Vui lòng nhập lại OTP');
                return;
            } else {
                showError('error server');
                return;
            }


        } catch (error) {
            console.error(error);
            showError('Lỗi kết nối đến máy chủ');
            return;
        }
    };
    if(isNext){
        return(
            <RegenerateNewPassword/>
        )
    }
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
                        Tiếp theo
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
