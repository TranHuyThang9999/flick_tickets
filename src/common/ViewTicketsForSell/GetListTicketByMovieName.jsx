import { Button, Col, Row } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import Carousel from 'react-multi-carousel';
import 'react-multi-carousel/lib/styles.css';
import CarouselCustomize from '../CustomizeCarousel/CarouselCustomize';
import { SelectOutlined } from '@ant-design/icons';
import './index.css';
import GetTicketById from '../../Home/Tickets/DetailTicketById';
import { showError } from '../log/log';
import PageForUser from '../../Home/Page/PageForUser';
export default function GetListTicketByMovieName({ movieName }) {

    const [tickets, setTickets] = useState([]);
    const [listFile, setListFile] = useState([]);
    const [isDetail, setIsDetail] = useState(false);
    const [selectedTicketId, setSelectedTicketId] = useState(null); // Thêm state để lưu trữ ticket_id được chọn
    const [goBack, setGoback] = useState(false);
    useEffect(() => {
        const fetchTickets = async () => {
            try {
                const response = await axios.get('http://localhost:8080/manager/customer/ticket/movie', {
                    params: {
                        movie_name: movieName,
                    },
                });
                if (response.data.result.code === 0) {
                    setTickets(response.data.tickets);
                    fetchFiles(response.data.tickets);
                } else if (response.data.result.code === 20) {
                    // Handle error or do something else
                }
            } catch (error) {
                console.log(error);
                showError('error server');
                return;
            }
        };
        fetchTickets();
    }, [movieName]); // Add status to dependency array

    const fetchFiles = async (tickets) => {
        const filesList = [];
        for (const ticket of tickets) {
            const result = await fetchDataFile(ticket.id);
            if (result.success) {
                filesList.push({ ticketId: ticket.id, files: result.files });
            }
        }
        setListFile(filesList);
    };

    const responsiveConfig = {
        desktop: {
            breakpoint: { max: 3000, min: 1024 },
            items: 4,
        },
        tablet: {
            breakpoint: { max: 1024, min: 464 },
            items: 2,
        },
        mobile: {
            breakpoint: { max: 464, min: 0 },
            items: 1,
        },
    };

    const handlerDetail = (ticketId) => {
        setSelectedTicketId(ticketId);
        setIsDetail(true);
    };
    const handlderGobacl = () => {
        setGoback(true);
    }
    if (isDetail) {
        return <GetTicketById id={selectedTicketId} />;
    }
    if(goBack){
        return <PageForUser/>
    }


    return (
        <div className='list-body-find-name'>
            <span><Button onClick={handlderGobacl}>Quay lại</Button></span>
            <h3>Thông tin vé</h3>
            <div className='list-display-carousel'>
                {tickets.length > 0 && (
                    <Carousel responsive={responsiveConfig}>
                        {tickets.map((ticket) => (
                            <div style={{ paddingLeft: '50px' }} key={ticket.id}>
                                <Row>
                                    <Col>
                                        <CarouselCustomize images={listFile.find((item) => item.ticketId === ticket.id)?.files || []} />
                                    </Col>
                                </Row>
                                <Row style={{ paddingTop: '10px', width: '270px', justifyContent: 'center' }}>
                                    <Col span={12} offset={6}>
                                        <p>{ticket.description}</p>
                                        <p>{ticket.name}</p>
                                        <Button onClick={() => handlerDetail(ticket.id)}>
                                            <SelectOutlined /> Chi tiết vé
                                        </Button>
                                    </Col>
                                </Row>
                            </div>
                        ))}
                    </Carousel>
                )}
            </div>
        </div>
    )
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