import React, { useState } from 'react';
import './index.css';
import { Button, Form, Input, Spin } from 'antd';
import axios from 'axios';
import { WeiboCircleFilled } from '@ant-design/icons';
import { showError } from '../common/log/log';

import FormCheckOtp from './FormCheckOtp';
import FormLogin from '../dashboard/FormLogin';

export default function PasswordRetrieval() {
  const [loading, setLoading] = useState(false);
  const [isNextNewPassword, setIsNextNewPassword] = useState(false);
  const [isGoback,setIsGoback] = useState(false);

  const [form] = Form.useForm();

  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };

  const handleFormSubmit = async (values) => {
    setLoading(true);
    try {
      const formData = new FormData();
      formData.append('email', values.email);
      formData.append('user_name', values.user_name);
      const response = await axios.post('http://localhost:8080/manager/customer/check', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      if (response.data.result.code === 0) {
        localStorage.setItem('email', values.email);
        localStorage.setItem('user_name', values.user_name);
        setIsNextNewPassword(true);
        return;
      } else {
        showError('error server');
        return;
      }
    } catch (error) {
      console.error(error);
      showError('Failed to send request');
    } finally {
      setLoading(false);
    }
  };
  if (isNextNewPassword) {
    return (
      <FormCheckOtp />
    )
  }
  if(isGoback){
    return(
      <FormLogin/>
    )
  }
  return (
    <div>
      <Form
        {...layout}
        form={form}
        className="form-container-reset-with-email"
        onFinish={handleFormSubmit}
      >
        <WeiboCircleFilled />
        <Form.Item
          name="user_name"
          label="Nhập tài khoản"
          rules={[{ required: true, message: 'Vui lòng nhập tài khoản!' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name="email"
          label="Nhập địa chỉ email đã đăng ký"
          rules={[
            { required: true, message: 'Vui lòng nhập địa chỉ email!' },
            { type: 'email', message: 'Địa chỉ email không hợp lệ!' },
          ]}
        >
          <Input />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Spin spinning={loading} tip="Vui lòng chờ...">
          <Button style={{marginRight:'15px'}} onClick={()=>setIsGoback(true)}>Quay lại</Button>
            <Button type="primary" htmlType="submit">
              Tiếp tục
            </Button>
          </Spin>
        </Form.Item>
      </Form>
    </div>
  );
}
