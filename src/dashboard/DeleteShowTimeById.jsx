import { Button, Popconfirm } from 'antd'
import React from 'react'
import { CloseCircleFilled } from '@ant-design/icons';
import { showError, showSuccess } from '../common/log/log';
import axios from 'axios';

export default function DeleteShowTimeById({ onDelete, showTimeId }) {

    const handlerDelete = async () => {
        try {
            const response = await axios.delete(`http://localhost:8080/manager/use/delete/byid?id=${showTimeId}`)
            if (response.data.result.code === 0) {
                showSuccess("Xóa thành công");
                onDelete(showTimeId);
                return;
            } else if(response.data.result.code === 4){
                showError("error server");
                return;
            }else{
                showError("error server");
                return;  
            }
        } catch (error) {
            console.log(error)
            showError("error server");
            return;
        }
    }

    return (
        <div>
            <Popconfirm
                title="Bạn có chắc chắn xóa xuất chiếu này?"
                okText="Yes"
                cancelText="No"
                onConfirm={handlerDelete}
            >
                <Button style={{ width: '45px', paddingLeft: '10px' }} type="primary" danger><CloseCircleFilled  /></Button>
            </Popconfirm>
        </div>
    )
}
