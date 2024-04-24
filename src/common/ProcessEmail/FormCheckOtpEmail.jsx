import { Button, Col, Row, Steps, message } from 'antd';
import React, { useState, useRef } from 'react';
import SendEmail from './SendEmail';
import VerifyEmail from './VerifyEmail';
import axios from 'axios';
import { showError, showSuccess } from '../log/log';
import './index.css';
const { Step } = Steps;

export default function FormCheckOtpEmail() {
    
    const [current, setCurrent] = useState(0);
    const [emailInput, setEmailInput] = useState('');
    const sendEmailRef = useRef(null);

    const next = () => {
        setCurrent(current + 1);
    };

    const prev = () => {
        setCurrent(current - 1);
    };

    const handleSendEmail = async () => {
        try {
            const response = await axios.post(`http://localhost:8080/manager/customer/send/${emailInput}`);
            if (response.data.result.code === 0) {
                showSuccess("Vui lòng kiểm tra email");
                next(); // Chuyển đến bước tiếp theo
            } else {
                showError("Lỗi server");
            }
        } catch (error) {
            console.error('Error:', error);
            showError("Lỗi server vui lòng thử lại");
        }
    };

    const handleEmailChange = (value) => {
        setEmailInput(value);
    };

    const handleFormSendEmail = () => {
        if (emailInput) {
            handleSendEmail();
        } else {
            showError("Vui lòng nhập địa chỉ email");
        }
    };

    const steps = [
        {
            title: 'Nhập địa chỉ email',
            content: <SendEmail onEmailChange={handleEmailChange} />,
        },
        {
            title: 'Nhập mã OTP',
            content: <VerifyEmail emaiInput={emailInput} />,
        },
    ];

    return (
        <>
            <div style={{margin:'10px'}} className='form-check-otp'>
                <Steps current={current}>
                    {steps.map(item => (
                        <Step key={item.title} title={item.title} />
                    ))}
                </Steps>

                <div style={{ marginTop: 16 }}>
                    {React.cloneElement(steps[current].content, { ref: sendEmailRef })}
                </div>

                <div style={{ marginTop: 24 }}>
                    {current > 0 && (
                        <Button style={{ margin: '0 8px' }} onClick={prev}>
                            Quay lại
                        </Button>
                    )}
                    {current < steps.length - 1 && (
                        <Button style={{ margin: '10px' }} type="primary" onClick={handleFormSendEmail}>
                            Tiếp tục
                        </Button>
                    )}
                    {current === steps.length - 1 && (
                        <Button type="primary" onClick={() => message.success('Hoàn tất!')}>
                            Hoàn tất
                        </Button>
                    )}
                </div>

            </div>

        </>
    );
}
