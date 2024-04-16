import { Card, Carousel, Col, Image, Row, Table } from 'antd';
import React, { useEffect, useState } from 'react'
import './index.css';
import { showError } from '../../common/log/log';
import axios from 'axios';


export default function GetTicketById({ id }) {
    const [ticket, setTicket] = useState([]);
    const [listFile, setListFile] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/manager/user/ticket?id=${id}`);
                const data = await response.json();
                setTicket(data.ticket);
            } catch (error) {
                console.error('Error:', error);
            }
        };

        fetchData();
    }, [id]);
    console.log(ticket.id);

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
                showError("Lỗi server vui lòng thử lại");
            }
        };

        fetchFileData();
    }, [id]);

    console.log(listFile);
    const onChange = (currentSlide) => {
        console.log(currentSlide);
    };

    if (!ticket.id) {
        return <div>Đang tải...</div>;
    }

    return (
        <div className='card-ticket-detail'>
            <p>
                {listFile.map((file) => (
                    <div key={file.id}>
                        <image width={200} src={file.url} alt='File' />
                    </div>
                ))}
            </p>

            <Row className='backgroud-card'>
                <Col style={{
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center'
                }} flex="400px">
                    ;k
                    {/* <Carousel afterChange={onChange}> */}
                    {listFile.map((file) => (
                        <div key={file.id}>
                            <image width={200} src={file.url} alt='File' />
                        </div>
                    ))}
                    {/* </Carousel> */}
                </Col>
                <Col
                    style={{
                        display: 'flex',
                        paddingTop: '100px'
                    }}
                    flex="auto">
                    Fill Rest
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
            return { success: false, error: "Không tìm thấy bản ghi nào" };
        } else if (data.result.code === 4) {
            return { success: false, error: "Lỗi server vui lòng thử lại" };
        }
    } catch (error) {
        console.error('Error:', error);
        return { success: false, error: "Lỗi server vui lòng thử lại" };
    }
};