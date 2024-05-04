import React, { useState, useRef } from 'react';
import QrReader from 'react-qr-scanner';
import jsQR from 'jsqr';
import axios from 'axios';
import { Button, Input } from 'antd';
import { showError, showSuccess, showWarning } from '../common/log/log';
import moment from 'moment';

const QRScanner = () => {
  const [resultInforQrCode, setResultInforQrCode] = useState('');
  const [scanEnabled, setScanEnabled] = useState(false);
  const [order, setOrder] = useState(null);
  const [orderIdToQrCode, setOrderIdToQrCode] = useState('');
  const qrReaderRef = useRef(null);
  const [imageSelected, setImageSelected] = useState(false);

  const handleScan = (data) => {
    if (data) {
      setResultInforQrCode(data.text);
    }
  };

  const handleError = (error) => {
    console.error(error);
  };

  const startScan = () => {
    setScanEnabled(true);
  };

  const stopScan = () => {
    setScanEnabled(false);
  };

  const handleFileUpload = async (event) => {
    const file = event.target.files[0];
    if (file) {
      try {
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
              setResultInforQrCode(code.data);
            } else {
              console.error('Failed to decode QR code');
            }
          };
        };
        reader.readAsDataURL(file);
        setImageSelected(true); // Set state to true when image is selected
      } catch (error) {
        console.error('Error while uploading file:', error);
      }
    }
  };
  
  

  const handleButtonClick = async () => {
    try {
      const response = await axios.get('http://localhost:8080/manager/user/verify/aes', {
        params: {
          token: resultInforQrCode
        }
      });
      console.log(response.data);
      if (response.data.result.code === 0) {
        setOrderIdToQrCode(response.data.content);
        handlerDetailOrderById(response.data.content);
        showSuccess("Mời xem thông tin đơn hàng");
      } else if (response.data.result.code === 18) {
        showWarning("QR code hợp lệ");
      } else if (response.data.result.code === 12) {
        showError("Lỗi yêu cầu không hợp lệ");
      } else {
        showError("Lỗi máy chủ");
      }
    } catch (error) {
      console.error(error);
      showError("Lỗi máy chủ");
    }
  };

  const handlerDetailOrderById = async (orderId) => {
    try {
      const response = await axios.get(`http://localhost:8080/manager/customer/look/order/ticket?id=${orderId}`);
      if (response.data.result.code === 0) {
        setOrder(response.data.orders);
      } else if (response.data.result.code === 14) {
        showWarning("Lỗi từ phía máy khách");
      } else if (response.data.result.code === 4) {
        showError("Lỗi máy chủ");
      }
    } catch (error) {
      console.error(error);
      showError("Lỗi máy chủ");
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
        style={{ width: '20%' }}
      />
    );
  } else {
    scannerContent = (
      <div>
        <Input type="file" accept="image/*" onChange={handleFileUpload} />
      </div>
    );
  }

  let scanButton;
  if (scanEnabled) {
    scanButton = <Button onClick={stopScan}>Stop Scan</Button>;
  } else {
    scanButton = <Button onClick={startScan}>Check QR With Camera</Button>;
  }
  const getOrderStatus = (status) => {
    if (status === 9) {
      return 'Thanh toán thành công';
    } else if (status === 11) {
      return 'Đã hủy';
    } else {
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
  //const addressDetailsString = formatAddressDetails(order.addressDetails);

  let orderInfo;
  if (order) {
    const addressDetailsString = order.addressDetails ? formatAddressDetails(order.addressDetails) : '';
    orderInfo = (
      <div style={{padding:'10px'}}>
        <p>ID Đơn hàng: {order.id}</p>
        <p>Giờ chiếu: {order.show_time_id}</p>
        <p>Ngày phát hành: {order.release_date}</p>
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
  <Button onClick={handleButtonClick} disabled={!imageSelected}>Check QRCode</Button>

  return (
    <div style={{width:'600px'}}>
      {scannerContent}
      {scanButton}
      <Button onClick={handleButtonClick}>Check QRCode</Button>
      {orderInfo}
    </div>
  );
};

export default QRScanner;
