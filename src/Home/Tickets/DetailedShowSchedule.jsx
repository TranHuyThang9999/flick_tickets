import React, { useState, useEffect } from 'react';
import { Button, Drawer, Table, Spin } from 'antd'; // Thêm Spin từ antd để hiển thị trạng thái loading
import { showWarning, showError } from '../../common/log/log';
import SelectedSeat from '../../common/cinemas/SelectedSeat';
import axios from 'axios';

export default function DetailedShowSchedule({ id }) {
  const [showTimeTicket, setShowTimeTicket] = useState([]);
  const [open, setOpen] = useState(false);
  const [selectedRecord, setSelectedRecord] = useState(null);
  const [fetchingData, setFetchingData] = useState(false);
  const [selectPopChid, setSelectPopChid] = useState([]);
  const [loadingPayment, setLoadingPayment] = useState(false); // Trạng thái loading cho phần thanh toán

  const showDrawer = (record) => {
    setSelectedRecord(record);
    setOpen(true);
  };

  const onClose = () => {
    setSelectedRecord(null);
    setOpen(false);
  };

  useEffect(() => {
    fetchData();
  }, [id]);

  const fetchData = async () => {
    setFetchingData(true);
    try {
      const response = await fetch(`http://localhost:8080/manager/user/getlist/time?id=${id}`);
      const data = await response.json();
      setShowTimeTicket(data.showtimes);
      if (data.result.code === 20) {
        showWarning("Không tìm thấy bản ghi nào");
      } else if (data.result.code === 4) {
        showError("Lỗi server vui lòng thử lại");
      }
    } catch (error) {
      console.error('Error:', error);
      showError("Lỗi server vui lòng thử lại", error);
    } finally {
      setFetchingData(false);
    }
  };

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
      title: 'Code',
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
      title: 'Xã/Phường',
      dataIndex: 'commune',
      key: 'commune',
    },
    {
      title: 'Địa chỉ Chi tiết',
      dataIndex: 'address_details',
      key: 'address_details',
    },
    {
      title: 'Giá tiền',
      dataIndex: 'price',
      key: 'price',
    },
    {
      title: '',
      render: (record) => (
        <div>
          <Button type="primary" onClick={() => showDrawer(record)}>
            Chọn ghế xem phim
          </Button>
        </div>
      ),
    },
  ];

  const pagination = {
    pageSize: 4,
    position: ['bottomLeft'],
  };

  const handleCreatePayment = async () => {
    setLoadingPayment(true); // Bắt đầu loading khi bắt đầu thanh toán
    try {
      const amount = selectedRecord.price * selectPopChid.length;

      const requestData = {
        amount: amount,
        description: `Số ghế đã chọn: ${selectPopChid.join(', ')}`,
        cancelUrl: "http://localhost:3000/",
        returnUrl: "http://localhost:3000/",
        buyerName: "John Doe",
        buyerEmail: "john@example.com",
        buyerPhone: "123456789",
        buyerAddress: "123 Main St"
      };

      const response = await axios.post('http://localhost:8080/manager/public/customer/payment/pay', requestData, {
        headers: {
          'Content-Type': 'application/json'
        }
      });

      const paymentResult = response.data;
      if (paymentResult && paymentResult.checkoutUrl) {
        // Chuyển hướng người dùng đến trang thanh toán
        window.location.href = paymentResult.checkoutUrl;
      }
    } catch (error) {
      console.log(error);
      showError('Error server 1');
    }
    setLoadingPayment(false); // Kết thúc loading sau khi hoàn thành hoặc gặp lỗi
  };

  console.log("list for pop : ", selectPopChid);

  return (
    <div>
      <Table scroll={{ x: 90 }} dataSource={showTimeTicket} columns={columns} pagination={pagination} loading={fetchingData} />
      <Drawer
        title="Phòng"
        width={1000}
        onClose={onClose}
        visible={open}
        bodyStyle={{
          paddingBottom: 80,
        }}
        style={{
          background: 'linear-gradient(90deg, rgba(102, 153, 204, 1) 0%, rgba(102, 204, 102, 1) 100%)',
        }}
      >
        {selectedRecord && (
          <div style={{ padding: '10px 16px' }}>
            <SelectedSeat
              SelectedSeatGetFormApi={selectedRecord.selected_seat}
              heightContainerUseSaveData={selectedRecord.height_container}
              widthContainerUseSavedate={selectedRecord.width_container}
              numSquares={selectedRecord.original_number}
              onCreate={setSelectPopChid}
            />
          </div>
        )}
        <Button type="primary" onClick={handleCreatePayment} loading={loadingPayment}> {/* Sử dụng loading từ antd cho nút thanh toán */}
          {loadingPayment ? 'Đang xử lý...' : 'Mua'}
        </Button>
      </Drawer>
    </div>
  );
}
