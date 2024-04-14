import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError, showWarning } from '../common/log/log';
import { Image, Table } from 'antd';

export default function GetListFileByTicketId({ id }) {
  const [listFile, setListFile] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/manager/user/load?id=${id}`);
        const data = response.data;
        setListFile(data.files);
        if (data.result.code === 0) {
          // Handle success condition if needed
        } else if (data.result.code === 20) {
          showWarning("Không tìm thấy bản ghi nào");
          return;
        } else if (data.result.code === 4) {
          showError("Lỗi server vui lòng thử lại");
          return;
        }
      } catch (error) {
        console.error('Error:', error);
        showError("Lỗi server vui lòng thử lại");
        return;
      }
    };

    fetchData();
  }, [id]);

  return (
    <div>
      {listFile.length > 0 && (
        <Table scroll={{ x: 190 }} dataSource={listFile} pagination={false}>
          <Table.Column title="Id" dataIndex="id" key="id" />
          <Table.Column
            title="Ảnh mô tả"
            dataIndex="url"
            key="url"
            render={(url) => <Image width={100} src={url} />}
          />
        </Table>
      )}
    </div>
  );
}