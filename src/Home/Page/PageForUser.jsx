import React, { useEffect, useState } from 'react';
import './home.css';
import { Avatar, Button, Col, Menu, Row } from 'antd';
import {
  BellFilled, ShoppingCartOutlined, TwitterCircleFilled,
  InteractionFilled, WeiboCircleOutlined,
  AppstoreOutlined, HomeFilled, RightOutlined
} from '@ant-design/icons';
import axios from 'axios';
import FormLogin from '../../dashboard/FormLogin';
import CinemasGetAll from '../../common/cinemas/CinemasGetAll';
import Profile from '../../common/customers/Profile';
import GetTicketOncarousel from '../../common/ViewTicketsForSell/GetTicketOncarousel';

export default function PageForUser() {

  const [islogin, setIslogin] = useState(false);
  const [personalPage, setPersonalPage] = useState(false);
  const username = localStorage.getItem('user_name');
  const [user, setUser] = useState(null);
  const [tickets, setTickets] = useState([]);
  const [statusTicketSale, setStatusTicketSale] = useState(0);
  const [nameCinema, setNameCinema] = useState('');

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
  }, [username]);


  const handlerCheckNextComponent = () => {
    if (localStorage.getItem('user_name') === null) {
      window.location.reload();
      alert.console();
    } else {
      setPersonalPage(true);
    }
  }

  const handleShowingTickets = () => {
    setNameCinema("");
    setStatusTicketSale(15);
  };

  const handleUpcomingTickets = () => {
    setNameCinema("");
    setStatusTicketSale(17);
  };
  const handlerNameCinema = (name) => {
    setStatusTicketSale(0);
    setNameCinema(name);
  }
  const handlerGobackHome = ()=>{
    window.location.reload();
  }
  const conhandlerLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    localStorage.removeItem('user_name');
    window.location.reload();
  }

  const handlerNextLogin = () => {
    setIslogin(true);
  }


  const listRoomCinema = CinemasGetAll();

  useEffect(() => {
    const feedDataTicket = async () => {
      try {
        const respone = await axios.get('http://localhost:8080/manager/customers/ticket')
        if (respone.data.result.code === 0) {
          setTickets(respone.data.list_tickets);
          console.log(respone.data.list_tickets);
          return;
        } else if (respone.data.result.code === 20) {
          return;
        }
      } catch (error) {
        console.log(error);
        return;
      }
    }
    feedDataTicket();
  }, [statusTicketSale]); // Thêm statusTicketSale vào dependency array

  if (personalPage) {
    return (
      <Profile />
    );
  }
  if (islogin) {
    return (
      <FormLogin />
    )
  }

  return (
    <div>
      <div className='layout-header'>
        <div className='layout-header-start'>
          {user && (
            <Avatar src={user.avatar_url} onClick={handlerCheckNextComponent} />
          )}
          <div>
            {!username && (
              <Button onClick={handlerNextLogin}>
                Đăng nhập <InteractionFilled />
              </Button>
            )}
          </div>
        </div>
        <div className='layout-header-center'>
          <div className='layout-header-center-menu-choice-two'>
            <div></div>
            <Menu style={{ backgroundColor: 'blanchedalmond', fontSize: '17px' }} mode="horizontal">
              <Menu.Item>
                <Avatar shape="square" size="large"  src='http://localhost:1234/manager/shader/huythang/638518679.jpeg' onClick={handlerGobackHome}/>
              </Menu.Item>
              <Menu.SubMenu key="SubMenu" icon={<WeiboCircleOutlined />} title={<span>Lịch chiếu</span>}>
                <Menu.Item key="one" icon={<AppstoreOutlined />} onClick={handleShowingTickets}>
                  Đang chiếu
                </Menu.Item>

                <Menu.Item key="two" icon={<AppstoreOutlined />} onClick={handleUpcomingTickets}>
                  Sắp chiếu
                </Menu.Item>
              </Menu.SubMenu>
            </Menu>
          </div>
          <div>
            <Menu style={{ backgroundColor: 'blanchedalmond', fontSize: '17px' }} mode="horizontal">
              <Menu.SubMenu key="SubMenu" title={<span> Rạp chiếu</span>}>
                {listRoomCinema.map((cinema) => (
                  <Menu.Item key={cinema.id} icon={<AppstoreOutlined />}>
                    {cinema.cinema_name}
                  </Menu.Item>
                ))}
              </Menu.SubMenu>
            </Menu>
          </div>
          <div className='layout-header-center-menu-choice-end'>
            <Menu style={{ backgroundColor: 'blanchedalmond', fontSize: '17px' }} mode="horizontal">
              <Menu.SubMenu key="SubMenu" title={<span> Phim chiếu</span>}>
                {tickets.map((ticket) => (
                  <Menu.Item key={ticket.id} icon={<AppstoreOutlined />} onClick={() => handlerNameCinema(ticket.name)}>
                    {ticket.name}
                  </Menu.Item>
                ))}
              </Menu.SubMenu>
            </Menu>
          </div>
        </div>
        <div className='layout-header-end'>

          <div>Thông báo <BellFilled /></div>

          <div>
            <Button> Giỏ hàng <ShoppingCartOutlined /></Button>
          </div>

          <div>
            <Button>Cộng đồng <TwitterCircleFilled /></Button>
          </div>

          <div>
            {username && (
              <Button onClick={conhandlerLogout}>
                Đăng xuất <InteractionFilled />
              </Button>
            )}
          </div>
        </div>
      </div>
      <div className='layout-content'>
        <div className='layout-content-header'><HomeFilled /><RightOutlined /> Cinema</div>
        <Row className='layout-content-body'>
          <Col className='layout-content-descript'>
            <ul>
              <li>Phim đang chiếu 2024</li>
              <li>Lịch phim đang chiếu luôn cập nhật sớm nhất</li>
              <li>Suất phim đang chiếu đầy đủ các rạp</li>
              <li>Đặt lịch phim đang chiếu siêu nhanh</li>
              <li>Đặt vé lịch phim đang chiếu yêu thích mọi nơi</li>
            </ul>

          </Col>
          <Col className='layout-content-image'>
            <img width='600px' height='400px' src='http://localhost:1234/manager/shader/huythang/daidien.png' alt="Avatar" />
          </Col>
        </Row>
      </div>
      <div className='layout-footer'>
        <GetTicketOncarousel status={statusTicketSale} name={nameCinema} />
      </div>
    </div>
  )
}
