import { Button, DatePicker, Form, Select } from 'antd';
import moment from 'moment';
import React, { useState } from 'react';

export default function TestGetAllSelect() {
  const [timestampsList, setTimestampsList] = useState([]);
  const [timestampListAdd, setTimestampListAdd] = useState([]); //timestampListAdd add value
  const options = [];



  const handleChange = (value) => {
    console.log(`Selected: ${value}`);
  };
  //date
  const handleDateChange = (date, dateString) => {
    if (date && moment(date).isValid()) {
      setTimestampsList((prevTimestampsList) => [
        ...prevTimestampsList,
        moment(dateString).unix(),
      ]);
    }
  };
  const optionsGetTimeSelect = timestampsList.map(timestamp => ({
    value: timestamp,
    label: moment.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
  }));

  // Hàm xử lý khi nhấn nút "Clear"
  const handleClear = () => {
    setTimestampsList([]); // Xóa toàn bộ danh sách thời gian đã chọn
  };


  return (
    <div>
      <Form>
        <Form.Item>
          <Select
            mode='multiple'
            allowClear
            style={{ width: '100%' }}
            defaultValue={['a10', 'c12']}
            onChange={handleChange}
            options={options}
          />
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
            <Button onChange={handleClear}>Clear time</Button>
          </div>
        </Form.Item>
      </Form>
    </div>
  );
}
