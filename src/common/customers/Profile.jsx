import { Avatar, Button, Drawer, Image, Space } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import './index.css';
import UpdateProfile from './UpdateProfile';
import PageForUser from '../../Home/Page/PageForUser';
import {
     ShoppingCartOutlined,HomeFilled,
    InteractionFilled,
  } from '@ant-design/icons';
export default function Profile() {
    const [user, setUser] = useState(null);
    const username = localStorage.getItem('user_name');
    const [drawerVisible, setDrawerVisible] = useState(false);
    const [isGoback,setIsGoback]= useState(false);
    const showDrawer = () => {
        setDrawerVisible(true);
    };

    const onCloseDrawer = () => {
        setDrawerVisible(false);
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

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        localStorage.removeItem('user_name');
        window.location.reload(); // Reload the page
    };
    const handleGoback = ()=>{
        setIsGoback(true);
    }
    if(isGoback){
        return <PageForUser/>
    }
    return (
        <div style={{ background: '' }}>
            {user && (
                <div className='header'>
                    <div style={{ display: 'flex' }}>
                        <Avatar src={user.avatar_url} />
                        <div style={{ paddingLeft: '10px',paddingTop:'10px' }}> {user.user_name}</div>
                    </div>
                    <div>
                        <Space>
                            <Space.Compact>
                                <Button  onClick={handleLogout}><InteractionFilled/> Đăng xuất</Button>
                            </Space.Compact>
                            <Space.Compact>
                                <Button onClick={handleGoback}><HomeFilled /> Quay về trang chủ</Button>
                            </Space.Compact>
                            <Space.Compact>
                                <Button><ShoppingCartOutlined/> Giỏ hàng</Button>
                            </Space.Compact>
                        </Space>
                    </div>
                </div>
            )}
            <div className='footer'>
                {user && (
                    <div style={{ padding: '10px', borderRadius: '8px' }}>
                        <Image width={60} src={user.avatar_url} />
                        <div style={{ marginTop: '10px' }}>
                            <p style={{ color: '#1890ff' }}>Địa chỉ: {user.address}</p>
                            <p style={{ color: '#1890ff' }}>Tuổi: {user.age}</p>
                            <p style={{ color: '#1890ff' }}>Email: {user.email}</p>
                            <p style={{ color: '#1890ff' }}>Số điện thoại: {user.phone_number}</p>
                        </div>
                    </div>
                )}
            </div>
            <Drawer
                title="Cập nhật thông tin"
                width={600}
                onClose={onCloseDrawer}
                visible={drawerVisible}
                placement="right"
                destroyOnClose={true}
            >
                <UpdateProfile />
            </Drawer>
            <Button type="primary" onClick={showDrawer}>Cập nhật</Button>
        </div>
    );
}
