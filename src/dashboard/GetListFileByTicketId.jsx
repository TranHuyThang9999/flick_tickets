import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError, showSuccess, showWarning } from '../common/log/log';
import { Image, Table, Button } from 'antd';
import { RetweetOutlined, DeleteOutlined } from '@ant-design/icons';

const dataForNotFoundInDB = [
  {
    "id": 0,
    "url": "http://localhost:1234/manager/shader/huythang/empty_15187.png",
  }
];

export default function GetListFileByTicketId({ id }) {
  const [listFile, setListFile] = useState([]);

  const fetchData = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/manager/user/load?id=${id}`);
      const data = response.data;
      if (data.result.code === 0) {
        setListFile(data.files);
      } else if (data.result.code === 20) {
        setListFile(dataForNotFoundInDB);
      } else if (data.result.code === 4) {
        showError("Lỗi server vui lòng thử lại");
      }
    } catch (error) {
      console.error('Error:', error);
      showError("Lỗi server vui lòng thử lại");
    }
  };

  const handlerDeleteFileById = async (fileId) => {
    try {
      const responseDelete = await axios.delete(`http://localhost:8080/manager/user/delete/file/${fileId}`);
      if (responseDelete.data.result.code === 0) {
        showSuccess('Xóa ảnh thành công');
        setListFile((prevFiles) => prevFiles.filter(file => file.id !== fileId));
      } else {
        showError("Lỗi server vui lòng thử lại");
      }
    } catch (error) {
      console.error('Error:', error);
      showError("Lỗi server vui lòng thử lại");
    }
  };

  useEffect(() => {
    fetchData();
  }, [id]);
  const handleReload =()=>{
    fetchData();
  }
  return (
    <div>
      {listFile.length > 0 && (
        <Table scroll={{ x: 190 }} dataSource={listFile} pagination={false} rowKey="id">
          <Table.Column title="Id" dataIndex="id" key="id" />
          <Table.Column
            title="Ảnh mô tả"
            dataIndex="url"
            key="url"
            render={(url) => <Image width={100} src={url} />}
          />
          <Table.Column
            title={<Button onClick={handleReload}><RetweetOutlined /></Button>}
            render={(text, record) =>
              record.id !== 0 && (
                <Button
                  icon={<DeleteOutlined />}
                  onClick={() => handlerDeleteFileById(record.id)}
                >
                  Xóa
                </Button>
              )
            }
          />
        </Table>
      )}
    </div>
  );
}
