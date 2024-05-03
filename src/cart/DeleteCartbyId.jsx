import React, { useState } from 'react';
import axios from 'axios';
import { Button, Popconfirm } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

export default function DeleteCartbyId({ onDelete,cartId }) {

    const handleDelete = async () => {
        try {
            const response = await axios.delete(`http://localhost:8080/manager/cart/delete/${cartId}`);
            console.log(response.data); // In ra dữ liệu phản hồi từ API sau khi xóa thành công
            // Thêm các xử lý khác sau khi xóa thành công nếu cần
            onDelete(cartId);
        } catch (error) {
            console.error('Error:', error);
            // Xử lý lỗi nếu có
        }
    };

    return (
        <div>
            <Popconfirm
                title="Bạn có chắc chắn muốn xóa giỏ hàng này?"
                okText="Yes"
                cancelText="No"
                onConfirm={handleDelete}
            >
                <Button style={{width:'100px',paddingLeft:'10px'}} type="primary" danger><CloseOutlined /></Button>
            </Popconfirm>
        </div>
    );
}
