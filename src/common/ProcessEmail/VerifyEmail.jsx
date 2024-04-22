import { Form, Input, Button, InputNumber } from 'antd';
import axios from 'axios';
import React, { useState } from 'react';
import { showError, showSuccess, showWarning } from '../log/log';

export default function VerifyEmail({emaiInput}) {
    const [otp, setOtp] = useState(0);

    const handleSubmit = async () => {
        try {
            const response = await axios.post('http://localhost:8080/manager/customer/verify/', {
                email: emaiInput,
                otp: otp
            });
            console.log('Verification successful:', response.data);
            // Handle success, maybe redirect user or show a success message
            if (response.data.result.code === 0) {
                showSuccess("Xác minh thành công vui lòng thanh toán");
                return;
            } else if (response.data.result.code === 22) {
                showWarning("Mã xác thực không chính xác vui lòng nhập lại");
                return;
            } else {
                showError("Lỗi server vui lòng thử lại");
                return;
            }
        } catch (error) {
            console.error('Verification failed:', error);
            showError("Lỗi server vui lòng thử lại");
            return;
        }
    };

    const handleOtpChange = (value) => {
        setOtp(value);
    };
    

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
            <Form {...layout} onFinish={handleSubmit}>
        
                <Form.Item label="OTP" name="otp" rules={[{ required: true, message: 'Please enter OTP' }]}>
                    <InputNumber onChange={handleOtpChange} />
                </Form.Item>
                <Form.Item  wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Verify
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
