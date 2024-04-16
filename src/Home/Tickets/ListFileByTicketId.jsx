import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError, showWarning } from '../../common/log/log';
import {  Image } from 'antd';
import {
    ArrowLeftOutlined,
    ArrowRightOutlined,

} from '@ant-design/icons';


export default function ListFileByTicketId({ id }) {
    const [listFile, setListFile] = useState([]);
    const [index, setIndex] = useState(0);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/manager/user/load?id=${id}`);
                const data = response.data;
                if (data.result.code === 0) {
                    setListFile(data.files);
                } else if (data.result.code === 20) {
                    showWarning("Không tìm thấy bản ghi nào");
                } else if (data.result.code === 4) {
                    showError("Lỗi server vui lòng thử lại");
                }
            } catch (error) {
                console.error('Error:', error);
                showError("Lỗi server vui lòng thử lại");
            }
        };

        fetchData();
    }, [id]);

    console.log(listFile);

    const hasNext = listFile.length > 1 && index < listFile.length - 1;
    const hasPrevious = listFile.length > 1 && index > 0;

    function handleNextClick() {
        if (hasNext) {
            setIndex(index + 1);
        } else {
            setIndex(0);
        }
    }

    function handlePreviousClick() {
        if (hasPrevious) {
            setIndex(index - 1);
        } else {
            setIndex(listFile.length - 1);
        }
    }

    let sculpture = listFile.length > 0 ? listFile[index] : null;

    return (
        <div style={{ height: '350px', border: 'solid 1px', width: '300px', display: 'flex', flexDirection: 'column', backgroundColor: 'beige' }}>
            <div style={{ flex: 1, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
                {sculpture && (
                    <div style={{ maxWidth: '100%', maxHeight: '100%' }}>
                        <Image
                            style={{ width: '100%', height: '100%', objectFit: 'contain' }}
                            src={sculpture.url}
                        />
                    </div>
                )}
            </div>
            <div style={{ display: 'flex', justifyContent: 'center', marginBottom: '10px' }}>
                <button
                    style={{ backgroundColor: 'beige', marginRight: '5px', padding: '1px 20px', border: 'none', color: 'white' }}
                    onClick={handlePreviousClick}
                >
                     <ArrowLeftOutlined style={{
                        color: 'darkblue', fontSize: '15px'
                    }} />
                </button>
                <button
                    style={{ backgroundColor: 'beige', marginLeft: '5px', padding: '1px 20px', border: 'none', color: 'white' }}
                    onClick={handleNextClick}
                >
                    <span>
                        
                    </span>
                    <ArrowRightOutlined style={{
                        color: 'darkblue', fontSize: '15px'
                    }} />
                </button>
            </div>

        </div>
    );
}