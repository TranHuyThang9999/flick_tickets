import { Space, Table } from 'antd';
import React, { useEffect, useState } from 'react';
import DetailedShowSchedule from '../common/Showtimes/DetailedShowSchedule';
import GetListFileByTicketId from './GetListFileByTicketId';


export default function GetTicketById({ id }) {

    const [ticket, setTicket] = useState([]);


    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/manager/user/ticket?id=${id}`);
                const data = await response.json();
                setTicket(data.ticket);
            } catch (error) {
                console.error('Error:', error);
            }
        };

        fetchData();
    }, [id]);


    if (!ticket.id) {
        return <div>Đang tải...</div>;
    }
    console.log("tikcet : ",ticket.selected_seat);
    return (
        <div>

            {ticket && (
                <Table scroll={{ x: 190 }} dataSource={[ticket]} pagination={false}>
                    <Table.Column title="Tên" dataIndex="name" key="name" />
                    <Table.Column title="Giá" dataIndex="price" key="price" />
                    <Table.Column title="Số vé tối đa" dataIndex="max_ticket" key="max_ticket" />
                    <Table.Column title="Số lượng" dataIndex="quantity" key="quantity" />
                    <Table.Column title="Mô tả" dataIndex="description" key="description" />
                    <Table.Column title="Giảm giá" dataIndex="sale" key="sale" />
                    <Table.Column title="Ngày phát hành" dataIndex="release_date" key="release_date" />
                    <Table.Column title="Trạng thái" dataIndex="status" key="status" />
                    <Table.Column title="Ghế đã chọn" dataIndex="selected_seat" key="selected_seat" />
                    <Table.Column title="Thời lượng phim" dataIndex="movieDuration" key="movieDuration" />
                    <Table.Column title="Giới hạn độ tuổi" dataIndex="age_limit" key="age_limit" />
                    <Table.Column title="Đạo diễn" dataIndex="director" key="director" />
                    <Table.Column title="Diễn viên" dataIndex="actor" key="actor" />
                    <Table.Column title="Nhà sản xuất" dataIndex="producer" key="producer" />
                    {/* <Table.Column title="Ngày tạo" dataIndex="created_at" key="created_at" render={formatTimestamp} />
              <Table.Column title="Ngày cập nhật" dataIndex="updated_at" key="updated_at" render={formatTimestamp} /> */}
                </Table>

            )
            }
            <DetailedShowSchedule id={ticket.id} numSquares={ticket.quantity} selectedSeatGetFormApi={ticket.selected_seat} />
            <GetListFileByTicketId id={ticket.id} />
            
            <Space direction='vertical' style={{ marginLeft: '20px' }}>

                <Space>
                    {/* <UpdateSizeRoom ticket_id={ticket.id} /> */}
                </Space>

                <p>Mô tả phòng</p>
            </Space>

        </div>
    );
}