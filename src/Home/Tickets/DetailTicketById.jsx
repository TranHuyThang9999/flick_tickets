import { Col, Row, Button } from 'antd';
import React, { useEffect, useState } from 'react';
import './index.css';
import { showError } from '../../common/log/log';
import axios from 'axios';
import CarouselCustomize from '../../common/CustomizeCarousel/CarouselCustomize';
import DetailedShowSchedule from './DetailedShowSchedule';

export default function GetTicketById({ id }) {
    const [ticket, setTicket] = useState({});
    const [listFile, setListFile] = useState([]);
    const [goback, setGoback] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/user/ticket?id=${id}`);
                const data = response.data;
                setTicket(data.ticket);
            } catch (error) {
                console.error('Error:', error);
            }
        };

        fetchData();
    }, [id]);
    function formatTimestamp(timestamp) {
        const date = new Date(timestamp * 1000); // Nhân 1000 để chuyển đổi từ milliseconds sang seconds

        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');

        return `${year}-${month}-${day} ${hours}:${minutes}`;
    }
    useEffect(() => {
        const fetchFileData = async () => {
            try {
                const result = await fetchDataFile(id);
                if (result.success) {
                    setListFile(result.files);
                } else {
                    showError(result.error);
                }
            } catch (error) {
                console.error('Error:', error);
                showError('Lỗi server vui lòng thử lại');
            }
        };

        fetchFileData();
    }, [id]);
    const handlerGoback = () => {
        setGoback(true);
    }
    if (Object.keys(ticket).length === 0) {
        return <div>Đang tải...</div>;
    }

    if (goback) {
        window.location.reload();
    }

    return (
        <div className="card-ticket-detail">
            <Button onClick={handlerGoback}>Quay lại</Button>
            <Row className="backgroud-card">

                <Col style={{ display: 'flex', padding: '30px' }} flex="300px">
                    <CarouselCustomize images={listFile} />
                </Col>

                <Col className="ticket-info" flex="400px">
                    <ul style={{ listStyle: 'none' }}>
                        <li>
                            <div>

                            </div>
                            <div className='name-film'>{ticket.name}</div>
                            <div className='name-film-header'>
                                <div>{ticket.director}</div>
                                <div>{ticket.actor}</div>
                                <div>{ticket.producer}</div>
                                <div>2024</div>
                                <div>{ticket.movieDuration}</div>
                            </div>
                            <span style={{ fontSize: '16px', color: 'white' }}> Nội dung</span><br />
                            <span
                                style={{
                                    fontSize: '15px',
                                    color: 'white',
                                    fontFamily: 'Tangerine'
                                }}
                            >
                                {ticket.description}
                            </span>
                            <br />

                            <div style={{
                                fontFamily: '-moz-initial',
                                color: '#202020',
                                fontSize: '17px'
                            }}>
                                Ngày phát hành: {formatTimestamp(ticket.release_date)}<br />
                                {ticket.status === 15 && (
                                    <span>Trạng thái: Mở bán</span>
                                )}
                                {ticket.status === 17 && (
                                    <span>Trạng thái: Chưa mở</span>
                                )}
                                <br/>
                                Giới hạn độ tuổi: {ticket.age_limit}<br />
                            </div>

                            {/* Giá: <span style={{ fontSize: '16px' }}><u>{ticket.price} VND</u></span><br /> */}
                        </li>
                    </ul>
                </Col>
                <Col flex='auto'>
                    <DetailedShowSchedule id={ticket.id} statusSaleForTicket={ticket.status} />
                </Col>
            </Row>

        </div>
    );
}

const fetchDataFile = async (ticketId) => {
    try {
        const response = await axios.get(`http://localhost:8080/manager/user/load?id=${ticketId}`);
        const data = response.data;
        if (data.result.code === 0) {
            return { success: true, files: data.files };
        } else if (data.result.code === 20) {
            return { success: false, error: 'Không tìm thấy bản ghi nào' };
        } else if (data.result.code === 4) {
            return { success: false, error: 'Lỗi server vui lòng thử lại' };
        }
    } catch (error) {
        console.error('Error:', error);
        return { success: false, error: 'Lỗi server vui lòng thử lại' };
    }
};