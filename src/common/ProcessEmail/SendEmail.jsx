import {  Form, Input } from 'antd';
import React, { useState } from 'react';


export default function SendEmail({ onEmailChange }) {
    const [emailInput, setEmailInput] = useState('');

    const handleEmailChange = (e) => {
        const { value } = e.target;
        setEmailInput(value);
        onEmailChange(value); // Gọi hàm callback để truyền email lên component cha
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
            <Form {...layout}>
                <Form.Item
                    label="Email"
                    name="email"
                    rules={[
                        {
                            type: 'email',
                            message: 'Email không hợp lệ!',
                        },
                        {
                            required: true,
                            message: 'Vui lòng nhập email!',
                        },
                    ]}
                >
                    <Input value={emailInput} onChange={handleEmailChange} />
                </Form.Item>
            </Form>
        </div>
    );
}