import { Layout, Modal } from 'antd';
import React, { useEffect, useState } from 'react';
import './home.css';
import FormLogin from '../../dashboard/FormLogin';

const { Header, Footer, Content } = Layout;

export default function PageForUser() {
  const [loginVisible, setLoginVisible] = useState(false);

  const showLoginModal = () => {
    setLoginVisible(true);
  };

  const handleLoginCancel = () => {
    setLoginVisible(false);
  };

  return (
    <div>
      <Layout className=''>
        <Header className='header'>
          <div onClick={showLoginModal}>Đăng nhập</div>
          <div style={{ display: 'flex' }}>
            <div>Tạo tài khoản</div>
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
