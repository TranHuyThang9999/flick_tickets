import React, { useState } from 'react';
import './index.css';
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  UploadOutlined,
  LineChartOutlined,
  VideoCameraOutlined,
  UsergroupAddOutlined,
  SolutionOutlined,
  LogoutOutlined
} from '@ant-design/icons';
import { Layout, Button, Tabs } from 'antd';
import QRScanner from '../QRScanner/QRScanner';
import CinemasAdd from '../common/cinemas/CinemasAdd';
import CreateAccountStaff from './CreateAccountStaff';
import GetAllStaff from './GetAllStaff';
import AdminUploadTickets from './AdminUploadTickets';
import GetAllTickets from './GetAllTickets';
import MovieAdd from '../common/MovieTypes/MovieUpload';
import PageForUser from '../Home/Page/PageForUser';
import OrderStatistics from '../Orders/OrderStatistics';

const { Header, Sider, Content } = Layout;

const items = [
  {
    key: '1',
    icon: <LineChartOutlined />,
    label: 'Thống kế', // login or logout
    children: <OrderStatistics/>,
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
  },
  {
    key: '7',
    icon: <SolutionOutlined />,
    label: 'Quản lý vé',
    children: <GetAllTickets />
  },
  {
    key: '8',
    icon: <SolutionOutlined />,
    label: 'Thêm loại phim',
    children: <MovieAdd />
  },
];

const HomeAdmin = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [activeTab, setActiveTab] = useState('7');
  const[logout,setLogout] = useState(false);

  const handleTabChange = (key) => {
    setActiveTab(key);
  };

  const handleToggleCollapse = () => {
    setCollapsed(!collapsed);
  };

  const handleLogout = ()=>{
    localStorage.removeItem('token');
    localStorage.removeItem('email');
    localStorage.removeItem('user_name');

    setLogout(true);
  }
  if(logout){
    return(
      <PageForUser/>
    )
  }

  return (
    <Layout>
      <Sider
        style={{
          overflow: 'auto',
          height: '100vh',
          position: 'fixed',
          left: 0,
          top: 0,
          bottom: 0,
        }}
        trigger={null}
        collapsible
        collapsed={collapsed}
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
        <div className="demo-logo-vertical" />

      </Sider>
      <Layout>

        <div style={{ marginLeft: collapsed ? 80 : 196 }}>

          <Header
            style={{
              position: 'sticky',
              top: 0,
              zIndex: 1,
              width: '100%',
              display: 'flex',
              alignItems: 'center',
            }}

          >
            <div className="header-menu">
              <Button
                type="text"
                icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                onClick={handleToggleCollapse}
                style={{
                  fontSize: '70px',
                  width: 70,
                  height: 73,
                  backgroundColor: 'white',
                  marginLeft: '-50px',
                  borderBottomLeftRadius: 0,
                  borderBottomRightRadius: 0,
                }}
              />
            </div>
            <div style={{ display: 'flex' }}>
              <Button
                style={{
                  height: '65px',
                  borderRadius: 0,
                  outline: 'none',
                }}
                onClick={handleLogout}
              >
                Đăng xuất<LogoutOutlined />
              </Button>

            </div>
          </Header>

          <Content
            style={{
              margin: '24px 16px',
              padding: 24,
              minHeight: 280,
              overflow: 'auto'
            }}
          >
            {/* Hiển thị nội dung tương ứng với tab được chọn */}
            {items.map((item) => (
              activeTab === item.key && item.children
            ))}
          </Content>
        </div>
      </Layout>
    </Layout>
  );
};

export default HomeAdmin;