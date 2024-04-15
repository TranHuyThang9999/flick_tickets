import { Layout, Flex, Tabs } from 'antd';
import React from 'react';
import './home.css';

const { Header, Footer, Sider, Content } = Layout;
const items = [
  {
    key: '1',
    label: 'Tab 1',
    children: 'Content of Tab Pane 1',
  },
  {
    key: '2',
    label: 'Tab 2',
    children: 'Content of Tab Pane 2',
  },
  {
    key: '3',
    label: 'Tab 3',
    children: 'Content of Tab Pane 3',
  },
];

export default function PageForUser() {

  return (
    <div >
      <Layout className=''>
        <Header
          style={{
            backgroundColor: 'beige',
            display: 'flex',
            alignItems: 'center',
          }}
        >
          header
          <Tabs defaultActiveKey='1' items={items} />

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
    </div>
  )
}
