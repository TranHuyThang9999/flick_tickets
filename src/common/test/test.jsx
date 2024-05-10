import { Button, Checkbox, Form, Input } from 'antd';
import React, { useState } from 'react';

export default function TestCheckBox() {
  const [role, setRole] = useState(13); // Default role is 13
  const [inputValue, setInputValue] = useState('');
  const [result, setResult] = useState('');

  const handleCheckboxChange = (e) => {
    const newRole = e.target.checked ? 1 : 13; // Set role to 1 if checkbox is checked, otherwise 13
    setRole(newRole);
  };

  const handleInputChange = (e) => {
    setInputValue(e.target.value);
  };

  const handleButtonClick = () => {
    // Perform actions based on the role value and input value
    if (role === 1 && inputValue === 'admin') {
      setResult('Thành công. Vai trò: Admin.');
    } else if (role === 13 && inputValue === 'user') {
      setResult('Thành công. Vai trò: Người dùng.');
    } else {
      setResult('Thất bại.');
    }
  };

  return (
    <div>
      <Form
        className='login-form'
        initialValues={{
          remember: true,
        }}
      >
        <Form.Item>
          <Input onChange={handleInputChange} placeholder="Nhập 'admin' hoặc 'user'" />
        </Form.Item>
        <Form.Item>
          <Checkbox onChange={handleCheckboxChange}>Admin</Checkbox>
        </Form.Item>
        <Form.Item>
          <Button onClick={handleButtonClick}>Check</Button>
        </Form.Item>
      </Form>
      {result && <p>{result}</p>}
    </div>
  );
}
