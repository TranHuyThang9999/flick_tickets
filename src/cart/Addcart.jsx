import { Button } from 'antd';
import axios from 'axios';
import React from 'react';
import { PlusCircleFilled } from '@ant-design/icons';
import { showSuccess } from '../common/log/log';

export default function Addcart({ show_time_id, seats_position, price }) {
    const user_name = localStorage.getItem("user_name");

    const handlerAddcart = async () => {
        try {
            const formData = new FormData();
            formData.append('user_name', user_name);
            formData.append('show_time_id', show_time_id);
            formData.append('seats_position', seats_position);
            formData.append('price', price);

            const response = await axios.post('http://localhost:8080/manager/cart/add', formData);
            showSuccess("add cart success");
            console.log(response.data); // In case you want to see the response
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div>
            <Button onClick={handlerAddcart}>Add Cart<PlusCircleFilled /></Button>
        </div>
    )
}
