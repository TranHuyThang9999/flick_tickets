import React, { useState } from 'react';
import axios from 'axios';
import { showError } from '../../common/log/log';

export default function PaymentOrder({dataReq}) {
  const [payment, setPayment] = useState('');

  try {
    const items = dataReq.items.map(item => ({
      name: item.name,
    //   quantity: item.quantity,
      price: item.price
    }));

    const requestData = {
      amount: dataReq.amount,
      description: dataReq.description,
      cancelUrl: "http://localhost:3000/",
      returnUrl: "http://localhost:3000/",
      items: items,
      buyerName: dataReq.buyerName,
      buyerEmail: dataReq.buyerEmail,
      buyerPhone: dataReq.buyerPhone,
      buyerAddress: dataReq.buyerAddress
    };

    axios
      .post('http://localhost:8080/manager/public/customer/payment/pay', requestData, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then(response => {
        setPayment(response.data);
      })
      .catch(error => {
        console.log(error);
        showError('Error server 1');
      });
  } catch (error) {
    console.log(error);
    showError('Error server 1');
    return;
  }

  return payment;
}