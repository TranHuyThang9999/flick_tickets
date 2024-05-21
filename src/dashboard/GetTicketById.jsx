import { Button, Modal, Space, Table } from 'antd';
import React, { useEffect, useState } from 'react';
import { PlusCircleFilled } from '@ant-design/icons';
import DetailedShowSchedule from '../common/Showtimes/DetailedShowSchedule';
import GetListFileByTicketId from './GetListFileByTicketId';
import AddShowTime from './AddShowTime';
import UpSertFileByTicketId from './UpSertFileByTicketId';


export default function GetTicketById({ id }) {

    const [ticket, setTicket] = useState([]);
    const [showAddShowTimeModal, setShowAddShowTimeModal] = useState(false);
    const [showAddFileModal, setShowAddFileModal] = useState(false);

    const toggleAddShowTimeModal = () => {
        setShowAddShowTimeModal(!showAddShowTimeModal);
    };

    const toggleAddFileModal = () => {
        setShowAddFileModal(!showAddFileModal);
    };
    const fetchData = async () => {
        try {
            const response = await fetch(`http://localhost:8080/manager/user/ticket?id=${id}`);
            const data = await response.json();
            setTicket(data.ticket);
        } catch (error) {
            console.error('Error:', error);
        }
    };

    useEffect(() => {
      
        fetchData();
    }, [id]);

    if (!ticket.id) {
        return <div>Đang tải...</div>;
    }
    function formatTimestamp(timestamp) {
        const date = new Date(timestamp * 1000); // Nhân 1000 để chuyển đổi từ milliseconds sang seconds

        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}`;
    }

    return (
        <div>
            <Button onClick={toggleAddShowTimeModal}>
                Thêm xuất chiếu <PlusCircleFilled />
            </Button>
            <Button onClick={toggleAddFileModal}>Thêm ảnh mô tả <PlusCircleFilled /></Button>
            <Modal
                width={800}
                visible={showAddShowTimeModal}
                onCancel={toggleAddShowTimeModal}
                footer={null}
            >
                <AddShowTime ticketId={ticket.id} />
            </Modal>
            <Modal
                width={800}
                visible={showAddFileModal}
                onCancel={toggleAddFileModal}
                footer={null}
            >
                <UpSertFileByTicketId ticketId={id} />
            </Modal>
            {ticket && (
                <Table scroll={{ x: 190 }} dataSource={[ticket]} pagination={false}>
                    <Table.Column title="Tên phim" dataIndex="name" key="name" />
                    <Table.Column title="Mô tả" dataIndex="description" key="description" />
                    <Table.Column title="Giảm giá" dataIndex="sale" key="sale" />
                    <Table.Column title="Ngày phát hành" dataIndex="release_date" key="release_date" render={formatTimestamp} />
                    <Table.Column title="Trạng thái" dataIndex="status" key="status" />
                    <Table.Column title="Ghế đã chọn" dataIndex="selected_seat" key="selected_seat" />
                    <Table.Column title="Thời lượng phim" dataIndex="movieDuration" key="movieDuration" />
                    <Table.Column title="Giới hạn độ tuổi" dataIndex="age_limit" key="age_limit" />
                    <Table.Column title="Đạo diễn" dataIndex="director" key="director" />
                    <Table.Column title="Diễn viên" dataIndex="actor" key="actor" />
                    <Table.Column title="Nhà sản xuất" dataIndex="producer" key="producer" />
                    <Table.Column title="Ngày tạo" dataIndex="created_at" key="created_at" render={formatTimestamp} />

                </Table>
            )
            }
            <DetailedShowSchedule
                id={ticket.id}
            />{/*ci tiet phong*/}

            <GetListFileByTicketId id={ticket.id} />{/*listfile*/}

            <Space direction='vertical' style={{ marginLeft: '20px' }}>

                <Space>
                    {/* <UpdateSizeRoom ticket_id={ticket.id} /> */}
                </Space>

            </Space>

        </div>
    );
}