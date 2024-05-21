import { Button, Drawer, Table } from 'antd';
import React, { useEffect, useState } from 'react';
import { showError, showWarning } from '../log/log';
import {  UnorderedListOutlined, RetweetOutlined } from '@ant-design/icons';
import SelectedSeatForAdmin from '../../dashboard/SelectedSeatForAdmin';
import DeleteShowTimeById from '../../dashboard/DeleteShowTimeById';
import UpdateShowTimeById from '../../dashboard/UpdateShowTimeById';

export default function DetailedShowSchedule({ id }) {// display for admin ShowTime
  const [showTimeTicket, setShowTimeTicket] = useState([]);
  const [open, setOpen] = useState(false);
  const [selectedRecord, setSelectedRecord] = useState(null);
  // const [selectPopChid, setSelectPopChid] = useState([]);
  // const [selectedRow, setSelectedRow] = useState(null);
  // const [openUpdateShowTime, setopenUpdateShowTime] = useState(false);
  // const [selectedRecordId, setSelectedRecordId] = useState(null);

  const showDrawer = (record) => {
    // const { id } = record; // Extract the id from the record
    setSelectedRecord(record); // Pass the entire record
    setOpen(true);
  };

  const onClose = () => {
    setOpen(false);
  };

  useEffect(() => {
    fetchData();
  }, [id]);

  const fetchData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/manager/user/getlist/time/admin?id=${id}`);
      const data = await response.json();
      setShowTimeTicket(data.showtimes);
      if (data.result.code === 0) {
      } else if (data.result.code === 20) {
        showWarning("Không tìm thấy bản ghi nào");
        return;
      } else if (data.result.code === 4) {
        showError("Lỗi server vui lòng thử lại");
        return;
      }
    } catch (error) {
      console.error('Error:', error);
      showError("Lỗi server vui lòng thử lại", error);
      return;
    }
  };

  const handleDelete = async (show_time_id) => {
    try {
      await fetch(`http://localhost:8080/manager/user/delete/showtime?id=${show_time_id}`, {
        method: 'DELETE',
      });
      fetchData(); // Reload data after deletion
    } catch (error) {
      console.error('Error deleting show time:', error);
      showError("Lỗi xóa suất chiếu, vui lòng thử lại", error);
    }
  };
  const handleReload = () => {
    fetchData(); // Reload data when reload button is clicked
  };

  // const showDrawerUpdateShowTime = (record) => {
  //   setopenUpdateShowTime(true);
  //   setSelectedRecordId(record.id); // Lưu ID của bản ghi được chọn
  // };

  // const onCloseDrawreUpdateShowTime = () => {
  //   setopenUpdateShowTime(false);
  // };

  function formatTimestamp(timestamp) {
    const date = new Date(timestamp * 1000);

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
  }

  const columns = [
    {
      title: 'Mã suất chiếu',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: 'Phòng để chiếu',
      dataIndex: 'cinema_name',
      key: 'cinema_name',
    },
    {
      title: 'Thời gian chiếu phim',
      dataIndex: 'movie_time',
      key: 'movie_time',
      render: (movie_time) => formatTimestamp(movie_time),
    },
    {
      title: 'Mô tả',
      dataIndex: 'description',
      key: 'description',
    },
    {
      title: <span><Button onClick={handleReload}><RetweetOutlined /></Button> </span>, // Use RetweetOutlined icon as title
      render: (record) => (
        <div style={{ display: 'flex' }}>
          <Button type="primary" onClick={() => showDrawer(record)}>
            <UnorderedListOutlined />
          </Button>
          <DeleteShowTimeById onDelete={() => handleDelete(record.id)} showTimeId={record.id} />
          {/* <Button onClick={() => showDrawerUpdateShowTime(record)}><EditFilled /></Button> */}
          <UpdateShowTimeById show_time_id={record.id}/>

          <Drawer
            title="Phòng"
            width={1000}
            onClose={onClose}
            visible={open}
            bodyStyle={{
              paddingBottom: 80,
            }}
          >
            {selectedRecord && (
              <div style={{ padding: '10px 16px' }}>
                <SelectedSeatForAdmin
                  SelectedSeatGetFormApi={selectedRecord.selected_seat}
                  heightContainerUseSaveData={selectedRecord.height_container}
                  widthContainerUseSavedate={selectedRecord.width_container}
                  numSquares={selectedRecord.original_number}
                  // onCreate={setSelectPopChid}
                />
              </div>
            )}
          </Drawer>
        </div>
      ),
    },

  ];
  // const handleRowClick = (record) => {
  //   if (selectedRow === record.key) {
  //     setSelectedRow(null);
  //   } else {
  //     setSelectedRow(record.key);
  //   }
  // };
  return (
    <div>
      <Table
        scroll={{ x: 120 }}
        dataSource={showTimeTicket}
        columns={columns}
        expandable={{
          expandedRowRender: (record) => (
            <p style={{ margin: 0, color: 'dodgerblue', paddingLeft: '10px' }}>
              |{record.price} VND | {record.description} | {record.conscious} |
              {record.district} | {record.commune} | {record.address_details} |Số lượng ghế: {record.original_number}
            </p>
          ),
          rowExpandable: (record) => record.name !== 'Not Expandable',
        }}
        // onRow={(record) => ({
        //   onClick: () => handleRowClick(record),
        // })}
        // rowClassName={(record) => (record.key === selectedRow ? 'selected-row' : '')}
      />
    </div>
  );
}
