// DetailedShowSchedule.js
import React, { useState } from 'react';
import { Button } from 'antd';
import TestSelectedSeatsPopup from './test';

export default function TestDetailedShowSchedule({ id }) {
    const [selectedSeats, setSelectedSeats] = useState([]);
    const [modalVisible, setModalVisible] = useState(false);

    const handleCreatePayment = async ({ buyerName, buyerEmail }) => {
        // Logic to handle payment creation using buyerName and buyerEmail
        console.log('Creating payment: 2', buyerName, buyerEmail);
    };

    return (
        <div>
            <TestSelectedSeatsPopup
                selectedSeats={selectedSeats}
                visible={modalVisible}
                onCancel={() => setModalVisible(false)}
                onCreate={handleCreatePayment}
                onSelectChange={setSelectedSeats}
            />
            <h1>
                 {selectedSeats.map((seat, index) => (
            <h1 key={index}>{seat}</h1>
        ))}
            </h1>
            <Button type="primary" onClick={() => setModalVisible(true)}>
                Open Popup
            </Button>
        </div>
    );
}
