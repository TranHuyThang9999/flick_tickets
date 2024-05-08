import { Table, Pagination } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import moment from 'moment';
import { RedditCircleFilled } from '@ant-design/icons';
import { showError } from '../../common/log/log';
import './index.css';
export default function PurchaseHistory({ email }) {
  const [orders, setOrders] = useState([]);
  const [total, setTotal] = useState(0);
  const [limit, setLimit] = useState(10);
  const [currentPage, setCurrentPage] = useState(1);

  useEffect(() => {
    fetchData();
  }, [currentPage]); // Gọi fetchData khi trang hiện tại thay đổi

  const fetchData = async () => {
    try {
      const offset = (currentPage - 1) * limit;
      const response = await axios.get('http://localhost:8080/manager/user/order/getlist', {
        params: {
          offset: offset,
          limit: limit,
          email: email,
        }
      });

      if (response.data.result.code === 0) {
        setOrders(response.data.orders);
        setTotal(response.data.total);
      } else if (response.data.result.code === 20) {
        // Xử lý code 20
      } else if (response.data.result.code === 4) {
        showError("Server error");
      } else {
        showError("Unknown error");
      }
    } catch (error) {
      showError("Network error");
    }
  };

  const getStatusText = (status) => {
    switch (status) {
      case 7:
        return 'Khởi tạo đơn hàng';
      case 9:
        return 'Đã thanh toán';
      case 11:
        return 'Đã hủy';
      default:
        return 'Unknown';
    }
  };

  const formatAddressDetails = (addressDetails) => {
    try {
      const parsedAddress = JSON.parse(addressDetails);
      const { cinema_name, description, conscious, district, commune, address_details } = parsedAddress;
      return `${cinema_name}, ${description}, ${conscious}, ${district}, ${commune}, ${address_details}`;
    } catch (error) {
      console.error('Error parsing address details:', error);
      return 'Thông tin địa chỉ không hợp lệ';
    }
  };

  const columns = [
    {
      title: 'Mã đơn hàng',
      dataIndex: 'id',
      key: 'id',
    },
    {
      title: 'Tên phim',
      dataIndex: 'movie_name',
      key: 'movie_name',
    },
    {
      title: 'Tại rạp',
      dataIndex: 'cinema_name',
      key: 'cinema_name',
    },

    {
      title: 'Trạng Thái',
      dataIndex: 'status',
      key: 'status',
      render: (status) => getStatusText(status),
    },
    {
      title: 'Giá',
      dataIndex: 'price',
      key: 'price',
    },

  ];


  if (localStorage.getItem('user_name') === null) {
          return (
              <div>
                  Vui lòng đăng nhập
                  <RedditCircleFilled style={{color:'dodgerblue',fontSize:'30px'}} />
              </div>
          )
      }
  return (
    <div>
      <Table
        columns={columns}
        dataSource={orders.map((order, index) => ({ ...order, key: index }))}
        expandable={{
          expandedRowRender: (record) => (
            <p style={{ margin: 0, color: 'dodgerblue', paddingLeft: '10px' }}>
              {formatAddressDetails(record.addressDetails)} |Thời gian chiếu : 
              {moment.unix(record.movie_time).format('YYYY-MM-DD HH:mm:ss')}|Thời gian phát hành :
              {moment.unix(record.release_date).format('YYYY-MM-DD HH:mm:ss')} |Thời gian mua: {moment.unix(record.created_at).format('YYYY-MM-DD HH:mm:ss')}
            </p>
          ),
          rowExpandable: (record) => record.name !== 'Not Expandable',
        }}
        pagination={false} // Tắt phân trang mặc định của Table
      />
      <Pagination
        total={total}
        defaultPageSize={limit}
        current={currentPage}
        onChange={(page) => setCurrentPage(page)}
      />
    </div>
  );
}
