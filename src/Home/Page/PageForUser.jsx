import React, { useEffect, useState } from 'react';
import './home.css';
import { Avatar, Button, Col, Drawer, Menu, Row } from 'antd';
import {
  BellFilled, ShoppingCartOutlined, TwitterCircleFilled,
  InteractionFilled, WeiboCircleOutlined,
  AppstoreOutlined, HomeFilled, RightOutlined, WeiboSquareFilled,
  MailFilled
} from '@ant-design/icons';
import axios from 'axios';
import FormLogin from '../../dashboard/FormLogin';
import CinemasGetAll from '../../common/cinemas/CinemasGetAll';
import Profile from '../../common/customers/Profile';
import GetTicketOncarousel from '../../common/ViewTicketsForSell/GetTicketOncarousel';
import GetListCart from '../../cart/GetListCart';
import QRScanner from '../../QRScanner/QRScanner';
import PurchaseHistory from '../Tickets/PurchaseHistory';

export default function PageForUser() {

  const [islogin, setIslogin] = useState(false);
  const [personalPage, setPersonalPage] = useState(false);
  const username = localStorage.getItem('user_name');
  const [user, setUser] = useState(null);
  const [tickets, setTickets] = useState([]);
  const [statusTicketSale, setStatusTicketSale] = useState(0);
  const [nameCinema, setNameCinema] = useState('');
  const [movieTheaterName, setMovieTheaterName] = useState('');
  const [openCart, setOpencart] = useState(false);
  const [openCheck, setOpenCheck] = useState(false);
  const [openHistoryOrder,setOpenHistoryOrder] = useState(false);

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
  const handlerMovieTheaterName = (movieTheaterName) => {
    setStatusTicketSale(0);
    setNameCinema("");
    setMovieTheaterName(movieTheaterName);
  }

  const handlerGobackHome = () => {
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

  const showDrawer = () => {
    setOpencart(true);
  };
  const onClose = () => {
    setOpencart(false);
  };
  const showDrawerCheck = () => {
    setOpenCheck(true);
  };
  const onCloseCCheck = () => {
    setOpenCheck(false);
  };
  const showDrawerHistoryOrder = () => {
    setOpenHistoryOrder(true);
  };
  const onCloseCHistoryOrder= () => {
    setOpenHistoryOrder(false);
  };
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
              <Button className='layout-header-start-button-login' onClick={handlerNextLogin}>
                Đăng nhập <InteractionFilled />
              </Button>
            )}
          </div>
        </div>
        <div className='layout-header-center'>
          <div className='layout-header-center-menu-choice-two'>
            <Menu className='layout-header-center-menu-item-title' style={{ backgroundColor: 'blanchedalmond', fontSize: '17px' }} mode="horizontal">
              <Menu.Item>
                <Avatar className='layout-header-center-menu-item-title-avatar' shape="square" size="large"
                  src='http://localhost:1234/manager/shader/huythang/638518679.jpeg' onClick={handlerGobackHome} />
              </Menu.Item>
              <Menu.SubMenu className='layout-header-center-menu-item-title-sub' key="SubMenu" icon={<WeiboCircleOutlined />} title={<span>Lịch chiếu</span>}>
                <Menu.Item key="one" icon={<AppstoreOutlined />} onClick={handleShowingTickets}>
                  Đang chiếu
                </Menu.Item>

                <Menu.Item key="two" icon={<AppstoreOutlined />} onClick={handleUpcomingTickets}>
                  Sắp chiếu
                </Menu.Item>
              </Menu.SubMenu>
            </Menu>
          </div>
          <div className='layout-header-center-menu-choice-select-from-api'>
            <Menu style={{ backgroundColor: 'blanchedalmond', fontSize: '17px' }} mode="horizontal">
              <Menu.SubMenu className='layout-header-center-menu-choice-select-from-api-cinema-name' key="SubMenu" title={<span> Rạp chiếu</span>}>
                {listRoomCinema.map((cinema) => (
                  <Menu.Item key={cinema.id} icon={<AppstoreOutlined />} onClick={() => handlerMovieTheaterName(cinema.cinema_name)}>
                    {cinema.cinema_name}
                  </Menu.Item>
                ))}
              </Menu.SubMenu>
            </Menu>
          </div>
          <div className='layout-header-center-menu-choice-end'>
            <Menu style={{ backgroundColor: 'blanchedalmond', fontSize: '17px' }} mode="horizontal">
              <Menu.SubMenu className='layout-header-center-menu-choice-select-from-api-movie-name' key="SubMenu" title={<span> Phim chiếu</span>}>
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

          <div style={{ paddingRight: '10px' }}>
            <Button onClick={showDrawerCheck}><WeiboSquareFilled /> Kiểm tra vé</Button>
            <WeiboSquareFilled />

            <Drawer
              title="Create a new account"
              width={720}
              onClose={onCloseCCheck}
              open={openCheck}
              styles={{
                body: {
                  paddingBottom: 80,
                },
              }}
            >
              <QRScanner />
            </Drawer>
          </div>

          <div>Thông báo <BellFilled /></div>
          <div>
            <Button onClick={showDrawer}> Giỏ hàng <ShoppingCartOutlined /> </Button>
            <Drawer
              title="Thông tin giỏ hàng"
              width={1200}
              onClose={onClose}
              open={openCart}
              styles={{
                body: {
                  paddingBottom: 80,
                },
              }}
            >
              <GetListCart />
            </Drawer>
          </div>
            <div>
              <Button onClick={showDrawerHistoryOrder}>Lịch sử  mua hàng <MailFilled /></Button>
              <Drawer
              title="Lịch  sử mua hàng"
              width={1200}
              onClose={onCloseCHistoryOrder}
              open={openHistoryOrder}
              styles={{
                body: {
                  paddingBottom: 80,
                },
              }}
            >
              <PurchaseHistory/>
            </Drawer>
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
            <img width='600px' Height='400px' src='http://localhost:1234/manager/shader/huythang/daidien.png' alt="Avatar" />
          </Col>
        </Row>
      </div>
      <div className='layout-footer'>
        <GetTicketOncarousel status={statusTicketSale} name={nameCinema} movie_theater_name={movieTheaterName} />
      </div>
    </div>
  )
}
