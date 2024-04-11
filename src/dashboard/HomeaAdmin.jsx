import React, { useEffect, useState } from 'react';
import './index.css';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  LineChartOutlined,
  VideoCameraOutlined,
  UsergroupAddOutlined,
  SolutionOutlined
} from '@ant-design/icons';
import { Layout, Button, Tabs } from 'antd';
import AdminUploadTickets from './AdminUploadTickets';
import QRScanner from '../QRScanner/QRScanner';
import CinemasAdd from '../common/cinemas/CinemasAdd';
import CreateAccountStaff from './CreateAccountStaff';
import GetAllStaff from './GetAllStaff';

const { Header, Sider, Content } = Layout;

const items = [
  {
    key: '1',
    icon: <LineChartOutlined />,
    label: 'Thống kế',
    children: null,
  },
  {
    key: '2',
    icon: <VideoCameraOutlined />,
    label: ' Kiểm tra vé',
    children: <QRScanner />,
  },
  {
    key: '3',
    icon: <UploadOutlined />,
    label: 'Thêm phòng chiếu vé',
    children: <CinemasAdd />,
  },
  {
    key: '4',
    icon: <UploadOutlined />,
    label: 'Tạo vé',
    children: <AdminUploadTickets />,
  },
  {
    key: '5',
    icon: <UsergroupAddOutlined />,
    label: 'Thêm nhân viên ',
    children: <CreateAccountStaff />
  },
  {
    key: '6',
    icon: <SolutionOutlined />,
    label: ' Quản lý nhân viên ',
    children: <GetAllStaff />
  }
];

const HomeAdmin = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [activeTab, setActiveTab] = useState('4');

  const handleTabChange = (key) => {
    setActiveTab(key);
  };

  const handleToggleCollapse = () => {
    setCollapsed(!collapsed);
  };




  return (
    <Layout hasSider>
      <Sider
        trigger={null}
        collapsible
        collapsed={collapsed}
        width={200}
      >
        <div />

        <Tabs
          theme="dark"
          style={{ backgroundColor: 'beige', height: 'auto' }}
          tabPosition={collapsed ? 'top' : 'left'}
          activeKey={activeTab}
          onChange={handleTabChange}
        >
          {items.map((item) => (
            <Tabs.TabPane key={item.key} tab={<span>{item.icon}{!collapsed && item.label}</span>} />
          ))}
        </Tabs>
      </Sider>

      <Layout>
        <Header

      
        >
          <div className="header-menu">
            <Button
              type="text"
              icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              onClick={handleToggleCollapse}
              style={{
                fontSize: '16px',
                width: 64,
                height: 64,
                backgroundColor: 'red'
              }}
            />
            <h2 className="header-menu-login">logout</h2>
          </div>
        </Header>

        <Content
          style={{
            margin: '24px 16px',
            padding: 24,
            minHeight: 280,
            overflow:'auto'
          }}
        >
          {/* Hiển thị nội dung tương ứng với tab được chọn */}
          {items.map((item) => (
            activeTab === item.key && item.children
          ))}
        </Content>
      </Layout>
    </Layout>
  );
};

export default HomeAdmin;