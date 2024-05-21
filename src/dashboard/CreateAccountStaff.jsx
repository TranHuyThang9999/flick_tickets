import { Button, Form, Input, Select, Upload } from 'antd'
import axios from 'axios'
import React, { useState } from 'react'
import { showError, showSuccess, showWarning } from '../common/log/log';
import CinemasGetAll from '../common/cinemas/CinemasGetAll';

// curl --location 'localhost:8080/manager/customer/staff/register' \
// --form 'user_name="go 111"' \
// --form 'file=@"/home/huythang/9656856.png"' \
// --form 'address="Thai Binh"' \
// --form 'age="12"' \
// --form 'email="tranhuythang9999@gmail.com"' \
// --form 'phone_number="111"'

export default function CreateAccountStaff() {

    const [form] = Form.useForm();
    const [imageFile, setImageFile] = useState(null);

    

    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();
            formData.append('user_name', values.user_name);
            formData.append('file', imageFile);
            formData.append('age', values.age);
            formData.append('email', values.email);
            formData.append('phone_number', values.phone_number);


            const response = await axios.post('http://localhost:8080/manager/customer/staff/register',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',

                    },
                }
            );
            if (response.data.result.code === 0) {
                showSuccess('Tạo thành công');
                return;
            } else if (response.data.result.code === 22) {
                showWarning("Tên tài khoản đã tồn tạo");
                return;
            } else {
                showError("error server");
                return;
            }
        } catch (error) {

        }
    }
    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };
    const options = [];
    const cinemas = CinemasGetAll();
    for (let index = 0; index < cinemas.length; index++) {
      options.push({
        label: cinemas[index].cinema_name,
        value: cinemas[index].cinema_name,
      })
    }

    return (
        <div>
            <Form {...layout} form={form} className="form-container" onFinish={handleFormSubmit}>
                <Form.Item
                    label="Tên nhân viên"
                    name="user_name"
                    rules={[
                        {
                            required: true,
                            message: 'Vui lòng nhập tên người dùng',
                        },
                    ]}
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

                <Form.Item
                    label="Tuổi"
                    name="age"
                    rules={[
                        {
                            required: true,
                            message: 'Vui lòng nhập tuổi',
                        },
                    ]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Email"
                    name="email"
                    rules={[
                        {
                            type: 'email',
                            message: 'Email không hợp lệ',
                        },
                        {
                            required: true,
                            message: 'Vui lòng nhập email',
                        },
                    ]}
                >
                    <Input />
                </Form.Item>

                <Form.Item
                    label="Số điện thoại"
                    name="phone_number"
                    rules={[
                        {
                            required: true,
                            message: 'Vui lòng nhập số điện thoại',
                        },
                    ]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    label='Ca làm việc'
                >
                

                </Form.Item>
                <Form.Item
                    label='Làm tại rạp'
                >
                <Select
                    allowClear
                    options={options}
                    mode='multiple'
                />
                </Form.Item>
                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                </Form.Item>
            </Form>
        </div>
    )
}
