import { Button, Form, Input, Spin } from 'antd';
import React, { useState } from 'react';
import { GooglePlusSquareFilled } from '@ant-design/icons';
import './index.css';
import FormReceiveOtp from './FormReceiveOtp';
import axios from 'axios';
import { showError } from '../log/log';

export default function LoginWithEmail() {
    const [isNextFormOtp, setIsNextFormOtp] = useState(false);
    const [loading, setLoading] = useState(false);
    const [form] = Form.useForm();

    const layout = {
        labelCol: { span: 8 },
        wrapperCol: { span: 16 },
    };

    const handleFormSubmit = async (values) => {
        setLoading(true);
        try {
            const response = await axios.post(`http://localhost:8080/manager/customer/send/${values.email}`);
            if (response.data.result.code === 0) {
                localStorage.setItem('email', values.email);
                setIsNextFormOtp(true);
            } else {
                showError('error server');
            }
        } catch (error) {
            console.error(error);
            showError('Failed to send request');
        } finally {
            setLoading(false);
        }
    };

    if (isNextFormOtp) {
        return <FormReceiveOtp />;
    }

    return (
        <div>
            <Spin spinning={loading} tip="Vui lòng chờ...">
                <Form {...layout} form={form} className="form-container-register-with-email" onFinish={handleFormSubmit}>
                    <GooglePlusSquareFilled style={{ fontSize: '25px' }} />
                    <Form.Item
                        label='Email'
                        name="email"
                        rules={[{ required: true, type: 'email', message: 'Vui lòng nhập địa chỉ email!' }]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                        <Button type="primary" htmlType="submit">
                            Tiếp tục
                        </Button>
                    </Form.Item>
                </Form>
            </Spin>
        </div>
    );
}
