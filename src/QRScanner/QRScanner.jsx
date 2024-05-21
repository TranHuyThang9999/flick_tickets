import React, { useState, useRef } from 'react';
import QrReader from 'react-qr-scanner';
import jsQR from 'jsqr';
import axios from 'axios';
import { Button, Input } from 'antd';
import { showError, showWarning } from '../common/log/log';
import moment from 'moment';

const QRScanner = () => {
  const [scanEnabled, setScanEnabled] = useState(false);
  const [order, setOrder] = useState(null);
  const qrReaderRef = useRef(null);

  const handleScan = async (data) => {
    if (data) {
      setScanEnabled(false);
      verifyQRCode(data.text);
    }
  };

  const handleError = (error) => {
    console.error(error);
  };

  const startScan = () => {
    setScanEnabled(true);
    setOrder(null);
  };

  const stopScan = () => {
    setScanEnabled(false);
  };

  const handleFileUpload = (event) => {
    const file = event.target.files[0];
    if (file && (file.type === 'image/png' || file.type === 'image/jpeg' || file.type === 'image/jpg')) {
      setOrder(null);

      const reader = new FileReader();
      reader.onload = (e) => {
        const imageData = e.target.result;
        const image = new Image();
        image.src = imageData;
        image.onload = () => {
          const canvas = document.createElement('canvas');
          canvas.width = image.width;
          canvas.height = image.height;
          const context = canvas.getContext('2d');
          context.drawImage(image, 0, 0);
          const imageData = context.getImageData(0, 0, canvas.width, canvas.height);
          const code = jsQR(imageData.data, imageData.width, imageData.height);
          if (code) {
            verifyQRCode(code.data);
          } else {
            showWarning('Mã QrCode không hợp lệ vui lòng chọn lại');
          }
        };
      };
      reader.readAsDataURL(file);
    } else {
      showWarning('Mã QrCode không hợp lệ vui lòng chọn lại');
    }
  };

  const verifyQRCode = async (token) => {
    try {
      const response = await axios.get('http://localhost:8080/manager/user/verify/aes', {
        params: { token }
      });
      const { result, order } = response.data;
      if (result.code === 0) {
        setOrder(order);
        // showSuccess("Mời xem thông tin đơn hàng");
      } else if (result.code === 18) {
        showWarning("QR code hợp lệ");
      } else if (result.code === 12) {
        showError("Lỗi yêu cầu không hợp lệ");
      } else {
        showError("Lỗi máy chủ");
      }
    } catch (error) {
      console.error(error);
      showError("Lỗi máy chủ");
    }
  };

  const getOrderStatus = (status) => {
    switch (status) {
      case 9:
        return 'Thanh toán thành công';
      case 11:
        return 'Đã hủy';
      default:
        return 'Trạng thái không xác định';
    }
  };

  const formatAddressDetails = (addressDetails) => {
    try {
      const parsedAddress = JSON.parse(addressDetails);
      const { cinema_name, description, conscious, district, commune, address_details } = parsedAddress;
      return `${cinema_name}, ${description}, ${conscious}, ${district}, ${commune}, ${address_details}`;
    } catch (error) {
      console.error('Error parsing address details:', error);
      return 'Thông tin địa chỉ không hợp lệ';
    }
  };

  let scannerContent;
  if (scanEnabled) {
    scannerContent = (
      <QrReader
        ref={qrReaderRef}
        delay={300}
        onError={handleError}
        onScan={handleScan}
        style={{ width: '50%' }}
      />
    );
  } else {
    scannerContent = (
      <div>
        <Input type="file" accept=".png, .jpg, .jpeg" onChange={handleFileUpload} />
      </div>
    );
  }

  let orderInfo;
  if (order) {
    const addressDetailsString = order.address_details ? formatAddressDetails(order.address_details) : '';
    orderInfo = (
      <div style={{ padding: '10px' }}>
        <p>ID Đơn hàng: {order.id}</p>
        <p>Tên phim: {order.movie_name}</p>
        <p>Tên rạp: {order.cinema_name}</p>
        <p>Email: {order.email}</p>
        <p>Mô tả: {order.description}</p>
        <p>Trạng thái: {getOrderStatus(order.status)}</p>
        <p>Giá: {order.price}</p>
        <p>Ghế: {order.seats}</p>
        <p>Giảm giá: {order.sale}</p>
        <p>Thời gian chiếu phim: {moment(order.movie_time * 1000).format('DD/MM/YYYY HH:mm')}</p>
        <p>Địa chỉ: {addressDetailsString}</p>
      </div>
    );
  }

  let scanButton;
  if (scanEnabled) {
    scanButton = <Button onClick={stopScan}>Stop Scan</Button>;
  } else {
    scanButton = <Button onClick={startScan}>Kiểm tra mã QrCode</Button>;
  }

  return (
    <div style={{ width: '600px' }}>
      {scannerContent}
      {scanButton}
      {orderInfo}
    </div>
  );
};

export default QRScanner;
