// SelectedSeatsPopup.js
import React from 'react';
import { Modal, Input } from 'antd';

function TestSelectedSeatsPopup({ selectedSeats, visible, onCancel, onCreate }) {
  const [buyerName, setBuyerName] = React.useState('');
  const [buyerEmail, setBuyerEmail] = React.useState('');

  const handleNameChange = (e) => {
    setBuyerName(e.target.value);
  };

  const handleEmailChange = (e) => {
    setBuyerEmail(e.target.value);
  };

  const handleCreate = () => {
    onCreate({ buyerName, buyerEmail });
    setBuyerName('');
    setBuyerEmail('');
  };

  return (
    <Modal
      visible={visible}
      title="Selected Seats"
      okText="Create"
      onCancel={onCancel}
      onOk={handleCreate}
    >
      <ul>
        {selectedSeats.map((seat, index) => (
          <li key={index}>{seat}</li>
        ))}
      </ul>
      <Input placeholder="Name" value={buyerName} onChange={handleNameChange} />
      <Input placeholder="Email" value={buyerEmail} onChange={handleEmailChange} />
    </Modal>
  );
}

export default TestSelectedSeatsPopup;
