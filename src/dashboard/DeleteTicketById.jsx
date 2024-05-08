import axios from 'axios'
import React from 'react'
import { showError } from '../common/log/log';
import { Button, Popconfirm } from 'antd';
export default function DeleteTicketById({ onDelete, ticketId }) {// ko dung

    const handleDeleteById = async () => {
        try {
            const response = await axios.delete(`http://localhost:8080/manager/user/delete/ticket/${ticketId}`)
            if (response.data.result.code !== 0) {
                showError("error server");
                return;
            }
            onDelete(ticketId)
        } catch (error) {
            showError("error server");
            return;
        }
    }

    return (
        <div>
            <Popconfirm
                title="Bạn có chắc chắn xóa vé này đi cùng các suất chiếu ?"
                okText="Yes"
                cancelText="No"
                onConfirm={handleDeleteById}
            >
                <Button style={{ width: '100px', paddingLeft: '10px' }} type="primary">Xóa</Button>
            </Popconfirm>
        </div>
    )
}
