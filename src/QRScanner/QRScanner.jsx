import React, { useState, useRef } from 'react';
import QrReader from 'react-qr-scanner';

const QRScanner = () => {
  
  const [result, setResult] = useState('');
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
    scannerContent = <p>Click "Check QR" to enable camera</p>;
  }
  let scanButton;
  if (scanEnabled) {
    scanButton = <button onClick={stopScan}>Stop Scan</button>;
  } else {
    scanButton = <button onClick={startScan}>Check QR</button>;
  }
  return (
    <div>
      {scannerContent}
      <p>{result}</p>
     {scanButton}
    </div>
  );
};

export default QRScanner;