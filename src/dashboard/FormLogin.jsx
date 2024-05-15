import { Button, Checkbox, Form, Input, message } from 'antd';
import axios from 'axios';
import React, { useState } from 'react';
import {
    LockOutlined, UserOutlined,
    WeiboCircleFilled, AndroidOutlined, GooglePlusCircleFilled
}
    from '@ant-design/icons';
import './index.css';
import { showError } from '../common/log/log';
import FormRegisterCustomer from '../common/customers/FormRegisterCustomer';
import LoginWithEmail from '../common/customers/LoginWithEmail';
import PasswordRetrieval from '../PasswordRetrieval/PasswordRetrieval';

export default function FormLogin() {

    const [form] = Form.useForm();
    const [isLoginAdmin, setIsloginAdmin] = useState(false);
    const [isLoginCustomer, setIsLoginCustomer] = useState(false);
    const [role, setRole] = useState(13); // Default role is 13
    const [isNextRegister, setIsNextRegister] = useState(false);
    const [isLoginEmail,setIsLoginEmail] = useState(false);
    const [isResetPassword,setIsResetPassword] =useState(false);
    const handleCheckboxChange = (e) => {
        const newRole = e.target.checked ? 1 : 13; // Set role to 1 if checkbox is checked, otherwise 13
        setRole(newRole);
    };

    const handlernextFormRegister = () => {
        setIsNextRegister(true);
    }
    const errorMessage = () => {
        message.error('Lỗi hệ thống vui lòng thử lại');
    };

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', values.user_name);
            formData.append('password', values.password);
            formData.append('role', role); // Truyền role tương ứng

            const response = await axios.post('http://localhost:8080/manager/customer/manager/login', formData);

            if (response.data.result.code === 0) {
                localStorage.setItem('user_name', response.data.user_name);
                localStorage.setItem('email', response.data.email);
                localStorage.setItem('token', response.data.jwt_token.refresh_token);


                if (role === 1) {
                    setIsloginAdmin(true);
                } else {
                    setIsLoginCustomer(true);
                }

            } else if (response.data.result.code === 24) {
                alert('Thông tin tài khoản hoặc mật khẩu không chính xác. Vui lòng thử lại.');
            } else {
                showError("error server");
            }
        } catch (error) {
            console.error(error);
            errorMessage();
            showError("error server");
        }
    };

    if (isLoginCustomer) {
        window.location.reload();
    }

    const handlerLoginWithEmail =()=>{
        setIsLoginEmail(true);
    }
    const handlerResetPassword =()=>{
        setIsResetPassword(true);
    }
    // Trả về HomeAdmin component nếu đăng nhập là admin
    if (isLoginAdmin) {
        // return <HomeAdmin />;
        window.location.reload();
    }
    if(isLoginEmail){
        return(
            <LoginWithEmail/>
        )
    }
    if (isNextRegister) {
        return (
            <FormRegisterCustomer />
        )
    }
    if(isResetPassword){
        return(
            <PasswordRetrieval/>
        )
    }
    // Phần còn lại của mã để hiển thị biểu mẫu đăng nhập
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
                        <h2>Đăng nhập <WeiboCircleFilled /></h2>
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
                        <div className="login-form-forgot" href="/#">
                            <div>
                                <Checkbox onChange={handleCheckboxChange}>Quản trị viên <AndroidOutlined /></Checkbox>
                            </div>
                            <div style={{ display: 'flex', justifyContent: 'space-between' }}>
                                <a href='/##' onClick={handlerResetPassword}>
                                    Quên mật khẩu
                                </a>
                                <a href='/#' onClick={handlernextFormRegister}>
                                    Đăng ký tài khoản
                                </a>
                            </div>
                        </div>
                    </Form.Item>

                    <Form.Item style={{ display: 'flex', justifyContent: 'center' }}>
                        <Button style={{ fontSize: '15px' }} type="primary" htmlType='submit'>
                            Đăng nhập
                        </Button>
                        <a onClick={handlerLoginWithEmail} href='/#' style={{ marginLeft: '10px', color: 'gray' }}>
                            Đăng nhập với Gmail <GooglePlusCircleFilled style={{ color: 'gray' }} />
                        </a>
                    </Form.Item>

                    <Form.Item>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
}
