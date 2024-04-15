import { Button, Form, InputNumber } from 'antd';
import axios from 'axios';
import React from 'react';
import { showSuccess, showError } from '../log/log';

export default function UpdateSizeRoom({ ticket_id }) {

    const [form] = Form.useForm();

    const handlerUpdateSizeRoom = async (values) => {
        try {
            const formData = new FormData();
            formData.append('ticket_id', ticket_id);
            formData.append('width_container', values.width_container);
            formData.append('height_container', values.height_container);

            const response = await axios.put(
                'http://localhost:8080/manager/user/update/size/room',
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );
            console.log(response.data);
            if (response.data.result.code === 0) {
                showSuccess('Cập nhật chiều dài, chiều rộng thành công');
            } else {
                showError('Lỗi máy chủ');
            }
        } catch (error) {
            console.error(error);
            showError('Lỗi máy chủ');
        }
    };

    return (
        <div>
            <Form form={form} onFinish={handlerUpdateSizeRoom}>
                <Form.Item name="width_container" label="Chiều dài">
                    <InputNumber />
                </Form.Item>
                <Form.Item name="height_container" label="Chiều rộng">
                    <InputNumber />
                </Form.Item>
                <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        Cập nhật phòng
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}