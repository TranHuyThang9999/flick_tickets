import React, { useEffect, useState } from 'react';
import { showError, showWarning } from '../common/log/log';
import axios from 'axios';
import { Avatar, Select, Table } from 'antd';

export default function GetAllStaff() {
    const [staff, setStaff] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/customer/staff/getall');
                const data = response.data;
                if (data.result.code === 0) {
                    setStaff(data.customers);
                } else if (data.result.code === 20) {
                    showWarning('Data empty');
                    return;
                } else {
                    console.error('Error retrieving staff');
                    showError('Error retrieving staff');
                    return;
                }
            } catch (error) {
                console.error('Error fetching data:', error);
                showError('Error fetching data');
                return;
            }
        };

        fetchData();
    }, []);

    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            key: 'id',
        },
        {
            title: 'Tên nhân viên',
            dataIndex: 'user_name',
            key: 'user_name',
        },
        {
            title: 'Avatar',
            dataIndex: 'avatar_url',
            key: 'avatar_url',
            render: (avatarUrl) => <Avatar size={64} src={avatarUrl} />,

        },
        {
            title: 'Số điện thoại',
            dataIndex: 'phone_number',
            key: 'phone_number',
        },
        {
            title:'Ca làm'
        },
        {
            title:'Phòng ',
            render: () => <Select style={{width:100}} mode='multiple' />
        }
    ];

    return (
        <div>
            <Table dataSource={staff} columns={columns} />
        </div>
    );
}