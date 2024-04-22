import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { showError, showWarning } from '../../common/log/log';
import CarouselCustomize from '../../common/CustomizeCarousel/CarouselCustomize';

export default function ListFileByTicketId({ id }) {
  const [listFile, setListFile] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/manager/user/load?id=${id}`);
        const data = response.data;
        if (data.result.code === 0) {
          setListFile(data.files);
        } else if (data.result.code === 20) {
          showWarning('Không tìm thấy bản ghi nào');
        } else if (data.result.code === 4) {
          showError('Lỗi server vui lòng thử lại');
        }
      } catch (error) {
        console.error('Error:', error);
        showError('Lỗi server vui lòng thử lại');
      }
    };

    fetchData();
  }, [id]);

  return (
    <div>
      <CarouselCustomize images={listFile} />
    </div>
  );
}