import React, { useState, useRef } from 'react';
import QrReader from 'react-qr-scanner';
import jsQR from 'jsqr';

import CheckQr from './CheckQr';

const QRScanner = () => {

  const [resultInfor, setResult] = useState('');
  const [scanEnabled, setScanEnabled] = useState(false);
  const qrReaderRef = useRef(null);

  const handleScan = (data) => {
    if (data) {
      setResult(data.text);
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

  const handleFileUpload = (event) => {
    const file = event.target.files[0];
    if (file) {
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
            setResult(code.data);
          } else {
            console.error('Failed to decode QR code');
          }
        };
      };
      reader.readAsDataURL(file);
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
        <input type="file" accept="image/*" onChange={handleFileUpload} />
        <p>Click "Check QR" to enable camera, or upload a QR code image</p>
      </div>
    );
  }

  let scanButton;
  if (scanEnabled) {
    scanButton = <button onClick={stopScan}>Stop Scan</button>;
  } else {
    scanButton = <button onClick={startScan}>Check QR With Camera</button>;
  }

  return (
    <div>
      {scannerContent}
      {scanButton}
      <p>{resultInfor}</p>
      <CheckQr token={resultInfor}/>
    </div>
  );
};

export default QRScanner;