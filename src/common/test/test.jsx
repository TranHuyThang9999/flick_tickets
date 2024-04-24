import { Button, Input, Modal, Space } from 'antd';
import React, { useState } from 'react'

export default function DemoSend() {

  const [open, setOpen] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);

  const showModal = () => {
    setOpen(true);
  };

  const handleOk = () => {
    setConfirmLoading(true);
    setTimeout(() => {
      setOpen(false);
      setConfirmLoading(false);
    }, 2000);
  };

  const handleCancel = () => {
    setOpen(false);
  };
  
  return (
    <>
      <Button type="primary" onClick={showModal}>
        Open Modal with async logic
      </Button>
      <Modal
        title="Title"
        open={open}
        onOk={handleOk}
        confirmLoading={confirmLoading}
        onCancel={handleCancel}
      >
       <div style={{padding:'0 10px'}}>
          <label>Nhập số điện thoại</label>
          <Input/>
          <label>Nhập email để nhận vé</label>
          <Input/>
       </div>
       
      </Modal>
    </>
  );
}