import { Layout, Modal, Avatar, Popover, Button } from 'antd';
import React, { useEffect, useState } from 'react';
import './home.css';
import FormLogin from '../../dashboard/FormLogin';
import axios from 'axios';

const { Header, Footer, Content } = Layout;

export default function PageForUser() {
  const [user, setUser] = useState(null);
  const username = localStorage.getItem('user_name');
  const [loginVisible, setLoginVisible] = useState(false);

  const showLoginModal = () => {
    setLoginVisible(true);
  };

  const handleLoginCancel = () => {
    setLoginVisible(false);
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:8080/manager/customer/user/profile", {
          params: {
            user_name: username
          }
        });
        setUser(response.data.customer);
      } catch (error) {
        console.log(error);
      }
    };

    fetchData();
  }, []);

  // Hàm renderUser khi user đã được tải thành công từ API
  const renderUser = () => {
    if (user) {
      return (
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <Popover content={user.user_name} trigger="hover">
            <Avatar src={user.avatar_url} />
          </Popover>
          <div style={{ marginLeft: '10px' }}>{user.name}</div>
        </div>
      );
    } else {
      return <div onClick={showLoginModal}>
        <Button style={{background:'beige'}}>
        Đăng nhập
        </Button>
        
      </div>;
    }
  };

  return (
    <div>
      <Layout className=''>
        <Header className='header'>
          {renderUser()}
          <div style={{ display: 'flex' }}>
            <div>Blog</div>
          </div>
        </Header>
        <Content
          style={{
            padding: '0 48px',
            backgroundColor: 'red'
          }}
        >
          qwacd
        </Content>
        <Footer
          style={{
            textAlign: 'center',
          }}
        >
          cds
        </Footer>
      </Layout>
      <Modal
        title="Đăng nhập"
        visible={loginVisible}
        onCancel={handleLoginCancel}
        footer={null}
      >
        <FormLogin />
      </Modal>
    </div>
  )
}
