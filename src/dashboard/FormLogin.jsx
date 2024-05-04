import { Button, Form, Input, message } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { LockOutlined, UserOutlined, LoginOutlined } from '@ant-design/icons';
import './index.css';

import HomeAdmin from '../dashboard/HomeaAdmin';
import PageForUser from '../Home/Page/PageForUser';

export default function FormLogin() {
    const [form] = Form.useForm();
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isLoginRole3, setIsLoginRole3] = useState(false);
    const [loading, setLoading] = useState(true);
    const[nextFromHome,setNextFromHome] = useState(false);

    const errorMessage = () => {
        message.error('Lỗi hệ thống vui lòng thử lại');
    };

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', values.user_name);
            formData.append('password', values.password);

            const response = await axios.post('http://localhost:8080/manager/customer/manager/login', formData);

            if (response.data.result.code === 0) {
                localStorage.setItem('user_name', values.user_name);
                localStorage.setItem('token', response.data.jwt_token.refresh_token);

                // Đăng nhập thành công, cập nhật trạng thái đăng nhập
                setIsLoggedIn(true);
            } else {
                alert('Thông tin tài khoản hoặc mật khẩu không chính xác. Vui lòng thử lại.');
            }
        } catch (error) {
            console.error(error);
            errorMessage();
        }
    };


    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            checkToken(token);
        } else {
            setLoading(false);
        }
    }, []);
    const handleNextHome = ()=>{
        setNextFromHome(true);
    }
    const checkToken = async (token) => {
        try {
            const response = await axios.get(`http://localhost:8080/manager/customer/${token}`);
            if (response.data.Errors === 4) {
                setLoading(false);
            } else {
                localStorage.setItem('user', JSON.stringify(response.data));
                if (response.data.role === 3) {
                    setIsLoggedIn(true);
                    setIsLoginRole3(false);
                    return;
                } else if (response.data.role === 13) {
                    setIsLoggedIn(true);
                    setIsLoginRole3(true);
                    return;
                } else {
                    setIsLoggedIn(false);
                    setIsLoginRole3(false);
                }
                setLoading(false);
            }
        } catch (error) {
            console.error(error);
            setLoading(false);
        }
    }

    // Nếu đăng nhập thành công và là role 3, chuyển sang component GetListTicket
    if (isLoggedIn && isLoginRole3) {
        window.location.reload(); // Reload the page
        return <HomeAdmin />;

    }

    // Nếu đăng nhập thành công và không phải role 3, chuyển sang component HomeAdmin
    if (isLoggedIn) {
        window.location.reload(); // Reload the page
        return <PageForUser />;
    }

    // Nếu đang loading, hiển thị thông báo loading
    if (loading) {
        return <div>Loading...</div>;
    }
    if(nextFromHome){
        return(
            <PageForUser/>
        )
    }
    // Nếu không đăng nhập, hiển thị form đăng nhập
    return (
        <div className="container-login-user">
            <div className='form-login'>
                <Form
                    className='login-form'
                    form={form}
                    onFinish={handleFormSubmit}
                    initialValues={{
                        remember: true,
                    }}
                >
                    <Form.Item
                        labelAlign='right'
                        className='form-login-user-label-header form-row'
                    >
                        <h2>Classy Login Form</h2>
                    </Form.Item>

                    <Form.Item
                        className='form-row'
                        name='user_name'
                        rules={[{ required: true, message: 'Vui lòng nhập tài khoản của bạn!' }]}
                    >
                        <Input
                            className='form-login-input'
                            prefix={<UserOutlined className="site-form-item-icon" />}
                            placeholder="Username"
                        />
                    </Form.Item>

                    <Form.Item
                        className='form-row'
                        name='password'
                        rules={[{ required: true, message: 'Vui lòng nhập mật khẩu của bạn!' }]}
                    >
                        <Input.Password
                            className='form-login-input'
                            prefix={<LockOutlined className="site-form-item-icon" />}
                            placeholder="Password"
                        />
                    </Form.Item>

                    <Form.Item>
                        <div className="login-form-forgot" href="/">
                            <a>
                                Forgot password
                            </a>
                        </div>
                        <div className="login-form-forgot" href="/">
                            <a onClick={handleNextHome}>
                            Quay lại trang chủ
                            </a>
                        </div>
                    </Form.Item>

                    <Form.Item style={{ display: 'flex', justifyContent: 'center' }}>
                        <Button style={{ fontSize: '15px' }} type="primary" htmlType='submit'>
                            Sign in
                            <LoginOutlined />
                        </Button>
                    </Form.Item>

                    <Form.Item>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}
