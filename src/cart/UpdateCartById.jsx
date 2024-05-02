import React, { useState, useEffect } from 'react';
import { Form, Input, Button } from 'antd';
import axios from 'axios';

export default function UpdateCartById({ cartId }) {
    const [formData, setFormData] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/cart/getlist?id=${cartId}`);
                if (response.data.result.code === 0 && response.data.carts.length > 0) {
                    const cartData = response.data.carts;
                    setFormData(cartData);
                }
            } catch (error) {
                console.error('Error:', error);
                // Xử lý lỗi nếu có
            }
        };
        fetchData();
    }, [cartId]);

    return (
        <div>
            <Form
                labelCol={{ span: 4 }}
                wrapperCol={{ span: 12 }}
                layout="horizontal"
            >
                <Form.Item initialValue={formData.id} label="ID" name="id">
                    <Input disabled />
                </Form.Item>

                <Form.Item label="Vị trí ghế" name="seats_position">
                    <Input />
                </Form.Item>
                <Form.Item wrapperCol={{ offset: 4, span: 12 }}>
                    <Button type="primary" htmlType="submit">
                        Update
                    </Button>
                </Form.Item>
            </Form>
        </div>
    );
}
