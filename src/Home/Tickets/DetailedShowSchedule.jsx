import React, { useState, useEffect } from 'react';
import { Button, Drawer, Table, Modal, Input, Popconfirm, Form } from 'antd';
import { showWarning, showError } from '../../common/log/log';
import SelectedSeat from '../../common/cinemas/SelectedSeat';
import axios from 'axios';
import Cookies from 'js-cookie';
import Addcart from '../../cart/Addcart';

export default function DetailedShowSchedule({ id, statusSaleForTicket }) {
  const [showTimeTicket, setShowTimeTicket] = useState([]);
  const [open, setOpen] = useState(false);
  const [selectedRecord, setSelectedRecord] = useState(null);
  const [fetchingData, setFetchingData] = useState(false);
  const [selectPopChid, setSelectPopChid] = useState([]);
  const [loadingPayment, setLoadingPayment] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [phoneNumber, setPhoneNumber] = useState('');
  const [email, setEmail] = useState('');
  const [selectedRow, setSelectedRow] = useState(null);

  const [form] = Form.useForm();

  const showDrawer = (record) => {
    const { show_time_id } = record;
    setSelectedRecord({ ...record, show_time_id });
    setOpen(true);
  };

  const onClose = () => {
    setSelectedRecord(null);
    setOpen(false);
  };

  const fetchData = async () => {
    setFetchingData(true);
    try {
      const response = await axios.get(`http://localhost:8080/manager/user/getlist/time?id=${id}`);
      const data = response.data;
      setShowTimeTicket(data.showtimes);
      if (data.result.code === 20) {
        showWarning("Hiện tạo chưa có xuất chiếu nào");
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
  const handleRowClick = (record) => {
    if (selectedRow === record.key) {
        setSelectedRow(null);
    } else {
        setSelectedRow(record.key);
    }
};

  const columns = [
    { title: 'Code', dataIndex: 'id', key: 'id' },
    { title: 'Chiếu tại rạp', dataIndex: 'cinema_name', key: 'cinema_name' },
    { title: 'Thời gian chiếu phim', dataIndex: 'movie_time', key: 'movie_time', render: (movie_time) => formatTimestamp(movie_time) },
    { title: 'Mô tả', dataIndex: 'description', key: 'description' },
    { title: 'Địa chỉ Chi tiết', dataIndex: 'address_details', key: 'address_details' },
    { title: 'Giá tiền', dataIndex: 'price', key: 'price' },
    { title: 'Giảm giá %', dataIndex: 'discount', key: 'discount' },

    {
      title: '',
      render: (record) => (
        <div>
          <Button type="primary" onClick={() => showDrawer(record)}>Chọn ghế xem phim</Button>
        </div>
      ),
    },
  ];

  const pagination = { pageSize: 4, position: ['bottomLeft'] };

  const item = selectPopChid;
  const items = item.map((item) => ({
    name: "Vị trí ghế : " + item,
    quantity: 1,
    price: selectedRecord && selectedRecord.price
  }));
  const amountTitile = selectedRecord &&
    (selectedRecord.price * selectPopChid.length) - (selectedRecord.price * selectPopChid.length) * (selectedRecord.discount / 100) + 'VND';

  const handleCreatePayment = async () => {
    setLoadingPayment(true);
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
        buyerEmail: email,
        buyerPhone: phoneNumber,
        buyerAddress: "123 Main St"
      };

      const response = await axios.post('http://localhost:8080/manager/public/customer/payment/pay', requestData, {
        headers: {
          'Content-Type': 'application/json'
        }
      });

      localStorage.setItem("order_id", response.data.orderCode);
      Cookies.set("order_id", response.data.orderCode, { expires: 30 });

      const paymentResult = response.data;
      if (paymentResult.resp_order === 44) {
        setLoadingPayment(false);
        showWarning("Ghế này đã được người mua trước vui lòng chọn lại");
        return;
      }
      if (paymentResult && paymentResult.checkoutUrl) {
        window.location.href = paymentResult.checkoutUrl;
      }
    } catch (error) {
      console.log(error);
      showError('Error server 1');
    }
    setLoadingPayment(false);
  };

  return (
    <div>
      <Table
        scroll={{ x: 90 }}
        dataSource={showTimeTicket}
        columns={columns}
        pagination={pagination}
        loading={fetchingData}
        expandable={{
          expandedRowRender: (record) => (
            <div>
              {record.conscious} || {record.district} ||{record.commune}
            </div>
          ),
        }}
        onRow={(record) => ({
          onClick: () => handleRowClick(record),
      })}
      rowClassName={(record) => (record.key === selectedRow ? 'selected-row' : '')}
      />
      <Drawer
        title="Phòng"
        width={1000}
        onClose={onClose}
        visible={open}
        bodyStyle={{ paddingBottom: 80 }}
        style={{ background: 'linear-gradient(90deg, rgba(102, 153, 204, 1) 0%, rgba(102, 204, 102, 1) 100%)' }}
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
        <div style={{ display: 'flex', padding: '10px', marginRight: '10px', marginTop: '100px' }}>
          <Button style={{ width: '120px', marginRight: '10px' }} type="primary" onClick={() => setShowModal(true)} disabled={selectPopChid.length === 0 || loadingPayment}>
            {loadingPayment ? 'Đang xử lý...' : 'Mua'}
          </Button>
          {selectPopChid.length > 0 && (
            <Addcart
              show_time_id={selectedRecord ? selectedRecord.id : null}
              seats_position={item}
              price={selectedRecord ? selectedRecord.price * item.length : 0}
            />
          )}
        </div>
      </Drawer>
      <Modal
        title="Nhập thông tin để nhận vé"
        visible={showModal}
        onCancel={() => setShowModal(false)}
        footer={[
          <Button key="back" onClick={() => setShowModal(false)}>Quay lại</Button>,
          <Popconfirm
            title="Xác nhận thanh toán?"
            okText="Xác nhận"
            cancelText="Hủy"
            onConfirm={form.submit}
          >
            <Button key="submit" type="primary" loading={loadingPayment}>{loadingPayment ? 'Đang xử lý...' : 'Thanh toán'}</Button>
          </Popconfirm>,
        ]}
      >
        <Form form={form} onFinish={handleCreatePayment}>
          <Form.Item label='Tổng số tiền'>
            <Input disabled value={amountTitile} />
          </Form.Item>
          <Form.Item
            label='Nhập số điện thoại'
            name='phoneNumber'
            rules={[{ required: true, message: 'Vui lòng nhập số điện thoại!' }]}
          >
            <Input onChange={(e) => setPhoneNumber(e.target.value)} />
          </Form.Item>
          <Form.Item
            label='Nhập email để nhận vé'
            name='email'
            rules={[
              { required: true, message: 'Vui lòng nhập email để nhận vé!' },
              { type: 'email', message: 'Vui lòng nhập email hợp lệ!' }
            ]}
          >
            <Input type='email' onChange={(e) => setEmail(e.target.value)} />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}
