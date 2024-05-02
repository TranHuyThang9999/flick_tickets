import React, { useState, useEffect } from 'react';
import { Button, Drawer, Table, Modal, Input, Popconfirm } from 'antd'; // Thêm Modal và Input từ antd để hiển thị form và ô input
import { showWarning, showError } from '../../common/log/log';
import SelectedSeat from '../../common/cinemas/SelectedSeat';
import axios from 'axios';
import Cookies from 'js-cookie'; // Import thư viện js-cookie

export default function DetailedShowSchedule({ id,statusSaleForTicket }) {
  const [showTimeTicket, setShowTimeTicket] = useState([]);
  const [open, setOpen] = useState(false);
  const [selectedRecord, setSelectedRecord] = useState(null);
  const [fetchingData, setFetchingData] = useState(false);
  const [selectPopChid, setSelectPopChid] = useState([]);
  const [loadingPayment, setLoadingPayment] = useState(false); // Trạng thái loading cho phần thanh toán
  const [showModal, setShowModal] = useState(false); // Trạng thái hiển thị modal
  const [phoneNumber, setPhoneNumber] = useState(''); // Trạng thái để lưu số điện thoại nhập vào
  const [email, setEmail] = useState(''); // Trạng thái để lưu email nhập vào

  const showDrawer = (record) => {
    const { show_time_id } = record; // Extract the show_time_id from the record
    setSelectedRecord({ ...record, show_time_id }); // Pass the show_time_id along with the record
    setOpen(true);
  };

  const onClose = () => {
    setSelectedRecord(null);
    setOpen(false);
  };

  const fetchData = async () => {
    setFetchingData(true);
    try {
      const response = await axios.get(`http://localhost:8080/manager/user/getlist/time?id=${id}`);// lay ve theo id
      const data = response.data; // Truy cập dữ liệu từ response.data
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


  useEffect(() => {
    fetchData();
  }, [id]);

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
  console.log(selectedRecord);
  const pagination = {
    pageSize: 4,
    position: ['bottomLeft'],
  };

  const item = selectPopChid
  const items = item.map((item) => ({
    name: "Vị trí ghế : " + item,
    quantity: 1,
    price: selectedRecord && selectedRecord.price
  }));

  const handleCreatePayment = async () => {
    setLoadingPayment(true); // Bắt đầu loading khi bắt đầu thanh toán
    try {
      const amount = selectedRecord.price * selectPopChid.length;

      const requestData = {
        amount: amount,
        description: 'Xin Cam on',
        items: items,
        ShowTimeId: selectedRecord.id,
        seats: selectPopChid.join(","),
        cancelUrl: "http://localhost:8080/manager/public/customer/payment/calcel",
        returnUrl: "http://localhost:8080/manager/public/customer/payment/return",
        buyerName: "John Doe",
        buyerEmail: email, // Sử dụng email đã nhập vào
        buyerPhone: phoneNumber, // Sử dụng số điện thoại đã nhập vào
        buyerAddress: "123 Main St"
      };

      const response = await axios.post('http://localhost:8080/manager/public/customer/payment/pay', requestData, {
        headers: {
          'Content-Type': 'application/json'
        }
      });

      localStorage.setItem("order_id", response.data.orderCode);
      Cookies.set("order_id", response.data.orderCode, { expires: 30 }); // Đặt cookie với thời gian sống là 1 tháng (30 ngày)

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
              statusSale={statusSaleForTicket}
            />
          </div>
        )}

        <Button type="primary" onClick={() => setShowModal(true)} disabled={selectPopChid.length === 0 || loadingPayment}> {/* Sử dụng điều kiện để vô hiệu hóa nút khi selectPopChid rỗng hoặc loadingPayment đang true */}
          {loadingPayment ? 'Đang xử lý...' : 'Mua'}
        </Button>
      </Drawer>
      {/* Modal */}
      <Modal
        title="Thông tin thanh toán"
        visible={showModal} // Sử dụng trạng thái showModal để điều khiển sự hiển thị của modal
        onCancel={() => setShowModal(false)}
        footer={[
          <Button key="back" onClick={() => setShowModal(false)}>
            Quay lại
          </Button>,
          //
          <Popconfirm
            title="Xác nhận thanh toán?"
            okText="Xác nhận"
            cancelText="Hủy"
            onConfirm={handleCreatePayment}
          >
            <Button key="submit" type="primary" loading={loadingPayment}>
              {loadingPayment ? 'Đang xử lý...' : 'Thanh toán'}
            </Button>
          </Popconfirm>,
        ]}
      >
        <div style={{ padding: '0 10px' }}>
          <label>Nhập số điện thoại</label>
          <Input onChange={(e) => setPhoneNumber(e.target.value)} />
          <label>Nhập email để nhận vé</label>
          <Input type='email' onChange={(e) => setEmail(e.target.value)} />
        </div>
      </Modal>
    </div>
  );
}
