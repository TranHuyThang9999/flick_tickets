import React, { useState } from 'react'
import CinemasGetAll from '../common/cinemas/CinemasGetAll';
import { Button, DatePicker, Form, InputNumber, Select } from 'antd';
import { showError, showSuccess, showWarning } from '../common/log/log';
import axios from 'axios';
import moment from 'moment';

export default function AddShowTime({ ticketId }) {
  const [cinemaName, setCinemaName] = useState('');
  const [timestampsList, setTimestampsList] = useState([]);
  const [timestampListAdd, setTimestampListAdd] = useState([]);


  const [form] = Form.useForm();

  const handleFormSubmit = async (values) => {
    if (cinemaName.length === 0) {
      showWarning('Vui lòng chọn phòng để chiếu phim!');
      return;
    }

    if (timestampListAdd.length === 0) {
      showWarning('Vui lòng chọn lịch chiếu phim!');
      return;
    }
    try {

      const formData = new FormData();
      formData.append('ticket_id', ticketId);
      formData.append('quantity', values.quantity);
      formData.append('cinema_name', cinemaName);// sua lai
      formData.append('movie_time', timestampListAdd);
      formData.append('quantity', values.quantity);
      formData.append('price', values.price);


      // Send a POST request using Axios
      const response = await axios.post(
        'http://localhost:8080/manager/use/add/time',
        formData,
        {
          headers: {
            // Authorization: token,
            'Content-Type': 'multipart/form-data',
          },
        }
      );

      if (response.data.result.code === 0) {
        showSuccess('Upload vé thành công');
        return;
      } else if (response.data.result.code === 26) {
        showWarning('các phòng vé đã tồn tại thời gian chiếu phim  vui lòng chọn phòng khác hoặc thời gian khác');
        return;
      } else if (response.data.result.code === 28) {
        showError('lỗi phía máy khách');
        return;
      } else {
        showError('Lỗi server, vui lòng thử lại');
        return;
      }
    } catch (error) {
      console.log(error);
      showError('Lỗi server, vui lòng thử lại');
      return;
    }
  };

  //
  const options = [];
  const cinemas = CinemasGetAll();
  for (let index = 0; index < cinemas.length; index++) {
    options.push({
      label: cinemas[index].cinema_name,
      value: cinemas[index].cinema_name,
    })
  }

  const layout = {
    labelCol: {
      span: 8,
    },
    wrapperCol: {
      span: 16,
    },
  };

  const optionsGetTimeSelect = timestampsList.map(timestamp => ({
    value: timestamp,
    label: moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
  }));
  //date
  const handleDateChange = (date, dateString) => {
    if (date && moment(date).isValid()) {
      setTimestampsList((prevTimestampsList) => [
        ...prevTimestampsList,
        moment(dateString).unix(),
      ]);
    }
  };

  return (
    <div>
      <Form {...layout} style={{ width: '' }} form={form} className="form-container-upload-ticket" onFinish={handleFormSubmit}>

        <Form.Item
          label="Nhập giá vé"
          className="form-row"
          name="price"
          rules={[{ required: true, message: 'Vui lòng nhập giá vé!' }]}
        >
          <InputNumber />
        </Form.Item>

        <Form.Item
          label="Nhập số lượng vé trên 1 phòng"
          className="form-row"
          name="quantity"
          rules={[{ required: true, message: 'Vui lòng nhập số lượng vé!' }]}
        >
          <InputNumber />
        </Form.Item>


        <Form.Item
          label="Lịch chiếu phim"
          className="form-row"
          name="movie_time"
        >
          <div className='showTime'>
            <DatePicker
              showTime
              onChange={handleDateChange}
              picker="datetime"
              size="small"
            />
            <Select
              allowClear
              mode="multiple"
              placeholder="Please select"
              options={optionsGetTimeSelect}
              onChange={(value) => setTimestampListAdd(value)}
            />
          </div>
        </Form.Item>

        <Form.Item
          label="Chọn phòng để chiếu phim"
          className="form-row"
          name="cinema_name"
        >
          <Select
            allowClear
            mode='multiple'
            options={options}
            onChange={(value) => setCinemaName(value)}

          />

        </Form.Item>


        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit">
            Submit
          </Button>
        </Form.Item>
      </Form>

    </div>
  )
}
