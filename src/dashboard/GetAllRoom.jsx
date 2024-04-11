import React from 'react'
import CinemasGetAll from '../common/cinemas/CinemasGetAll';
import { Table } from 'antd';

export default function GetAllRoom() {
    const cinemas = CinemasGetAll();

    const columns = [
     
        {
            title: 'Tên phòng chiếu',
            dataIndex: 'cinema_name',
            key: 'cinema_name',
        },
      
        {
            title: 'Tỉnh',
            dataIndex: 'conscious',
            key: 'conscious',
        },
        {
            title: 'Huyện',
            dataIndex: 'district',
            key: 'district',
        },
        {
            title: 'Xã',
            dataIndex: 'commune',
            key: 'commune',
        },
        {
            title: 'Mô tả',
            dataIndex: 'description',
            key: 'description',
        },
 

    ];
  return (
    <div>
        <Table dataSource={cinemas} columns={columns}/>
    </div>
  )
}
