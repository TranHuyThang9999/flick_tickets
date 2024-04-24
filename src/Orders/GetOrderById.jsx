import { get as axiosGet } from 'axios';
import React, { useState } from 'react';
import { showError, showWarning } from '../common/log/log';

export default function GetOrderById({ id }) {
  const [order, setOrder] = useState('');

  const handleSubmit = async () => {
    try {
      const { data: { code, data, desc } } = await axiosGet(`http://localhost:8080/manager/public/customer/payment/request?id=${id}`);

      if (code === '00') {
        setOrder(data);
      } else if ([ '101', '102', '103' ].includes(code)) {
        showWarning(desc);
        return;
      } else {
        showError("error server");
        return;
      }
    } catch (error) {
      console.log(error);
      return;
    }
  }

  return order;
}