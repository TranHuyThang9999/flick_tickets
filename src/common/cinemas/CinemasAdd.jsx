import { Button, Col, Form, Input, Row } from 'antd';
import React, { useState } from 'react'
import { showError, showSuccess, showWarning } from '../log/log';
import Axios from 'axios';
import OpenApiAddress from '../OpenApiAddress/OpenApiAddress';
import GetAllRoom from '../../dashboard/GetAllRoom';


export default function CinemasAdd() {
    const [address, setAddress] = useState(null);
    const handleAddressChange = (selectedAddress) => {
        setAddress(selectedAddress);
    };
    const [form] = Form.useForm();
    console.log(address);
    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();

            formData.append('cinema_name', values.cinema_name.trim());
            formData.append('description', values.description.trim());
            formData.append('conscious', address.city);
            formData.append('district', address.district);
            formData.append('commune', address.commune);
            formData.append('address_details', values.address_details);
            // Send a POST request using Axios

            const response = await Axios.post(
                'http://localhost:8080/manager/user/add/cinema',
                formData,
                {
                    headers: {
                        // Authorization: token,
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );
            if (response.data.result.code === 0) {
                showSuccess('Upload phong thành công');
                return;
            } else if (response.data.result.code === 30) {
                showWarning("tên phòng đã tồn tạo");
                return;
            } else {
                showError("error server");
                return;
            }
        } catch (error) {
            console.log(error);
            showError("error server 1");
            return;
        }

    }

    const layout = {
        labelCol: {
            span: 8,
        },
        wrapperCol: {
            span: 16,
        },
    };
    return (
        <div >
            <Row>
                <Col style={{padding:'0 16px'}}>
                    <Form
                        {...layout}
                        form={form}
                        onFinish={handleFormSubmit}
                    >
                        <Form.Item
                            label="Nhập tên phong"
                            className="form-row"
                            name="cinema_name"
                            rules={[{ required: true, message: 'Vui lòng nhập tên phong!' }]}
                        >
                            <Input trim />
                        </Form.Item>

                        <Form.Item
                            label="Mô tả phòng"
                            className="form-row"
                            name="description"
                            rules={[{ required: true, message: 'Vui lòng nhập Mô tả phòng!' }]}
                        >
                            <Input trim />
                        </Form.Item>

                        <Form.Item
                            style={{ display: 'block' }}
                            label="Nhập địa chỉ phòng chiếu"
                            className="form-row"
                        >
                            <OpenApiAddress onAddressChange={handleAddressChange} />
                        </Form.Item>
                        <Form.Item
                            label="Chi tiết địa chỉ phòng chiếu"
                            className="form-row"
                            name="address_details"
                            rules={[{ required: true, message: 'Vui lòng nhập chi tiết địa chỉ phòng chiếu!' }]}

                        >
                            <Input trim />
                        </Form.Item>
                        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                            <Button type="primary" htmlType="submit">
                                Submit
                            </Button>
                        </Form.Item>
                    </Form>
                </Col>
                <Col>
                    <GetAllRoom />
                </Col>
            </Row>

        </div>
    )
}
