import { Button, Form, Upload } from 'antd';
import axios from 'axios';
import React, { useState } from 'react'
import { showError, showSuccess } from '../common/log/log';

export default function UpSertFileByTicketId({ ticketId }) {

  const [form] = Form.useForm();
  const [fileList, setFileList] = useState([]);

  const onChange = ({ fileList }) => {
    setFileList(fileList);
  };
  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };

  const handlerSubmitUpsertFile = async () => {
    try {
      const formData = new FormData();
      formData.append('ticket_id', ticketId);
      fileList.forEach((file) => {
        formData.append('file', file.originFileObj);
      });
      const response = await axios.put('http://localhost:8080/manager/user/upload', formData, {
        headers: {
          // Authorization: token,
          'Content-Type': 'multipart/form-data',
        },
      })
      if (response.data.result.code === 0) {
        showSuccess('Upload ảnh thành công');
        return;
      } else {
        showError('error server');
        return;
      }
    } catch (error) {
      console.log(error);
      showError('error server');
      return;
    }
  }

  return (
    <div>
      <Form {...layout} form={form} className="form-container-upsert-with-file" onFinish={handlerSubmitUpsertFile}>

        <Form.Item
          label="Nhập ảnh mô tả vé"
          className="form-row"
          name="file"
          rules={[{ required: true, message: 'Vui lòng chọn ảnh mô tả vé!' }]}
        >
          <Upload
            fileList={fileList}
            listType="picture-card"
            accept="image/jpeg,image/png"
            onChange={onChange}
            beforeUpload={() => false} // Prevent auto-upload
          >
            {'+ Upload'}
          </Upload>
        </Form.Item>

        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button type="primary" htmlType="submit">
          Thêm ảnh mô tả
          </Button>
        </Form.Item>
      </Form>
    </div>
  )
}
