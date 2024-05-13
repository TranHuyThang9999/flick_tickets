import { Button, Form, Input, Upload } from 'antd';
import React, { useEffect, useState } from 'react'
import { showError, showSuccess } from '../log/log';
import axios from 'axios';

export default function UpdateProfile() {
    const [form] = Form.useForm();
    const [imageFile, setImageFile] = useState(null);
    const [user, setUser] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get("http://localhost:8080/manager/customer/user/profile", {
                    params: {
                        user_name: localStorage.getItem('user_name')
                    }
                });
                setUser(response.data.customer);
            } catch (error) {
                console.log(error);
            }
        };

        fetchData();
    }, []);

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', localStorage.getItem('user_name'));
            formData.append('address', values.address);
            formData.append('age', values.age);
            formData.append('email', values.email);
            formData.append('phone_number', values.phone_number);
            formData.append('file', imageFile);

            const response = await axios.put(
                'http://localhost:8080/manager/customer/user/update',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );

            if (response.data.result.code === 0) {
                showSuccess('Cập nhật tài khoản thành công');
            } else {
                showError('Lỗi server, vui lòng thử lại');
            }
        } catch (error) {
            console.log(error);
            showError('Lỗi server, vui lòng thử lại');
        }
    };

   

    if (!user) {
        return null; // or loading indicator
    }

    return (
        <div>
            <Form  form={form} className="form-container-update" onFinish={handleFormSubmit}>

                <Form.Item
                    initialValue={user.address}
                    label="Địa chỉ"
                    name="address"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    initialValue={user.age}
                    label="Tuổi"
                    name="age"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    initialValue={user.email}
                    label="Email"
                    name="email"
                    rules={[
                        {
                            type: 'email',
                            message: 'Email không hợp lệ!',
                        },
                    ]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    initialValue={user.phone_number}
                    label="Số điện thoại"
                    name="phone_number"
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Ảnh đại diện"
                    name="image"
                    valuePropName="fileList"
                    getValueFromEvent={(e) => {
                        if (Array.isArray(e)) {
                            return e;
                        }
                        return e && e.fileList;
                    }}
                >
                    <Upload
                        maxCount={1}
                        type=''
                        listType='picture-card'
                        openFileDialogOnClick
                        accept="image/jpeg,image/png"
                        beforeUpload={(file) => {
                            setImageFile(file);
                            return false;
                        }}
                        onRemove={() => {
                            setImageFile(null);
                        }}
                    >
                        +Upload
                    </Upload>
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
