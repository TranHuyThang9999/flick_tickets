import { Button, Col, Form, Input, Row, Table } from 'antd';
import React, { useEffect, useState } from 'react'
import { showError, showSuccess, showWarning } from '../log/log';
import axios from 'axios';
import MovieGetAll from './MovieGetAll';

export default function MovieAdd() {

    const listMovieType = MovieGetAll();

    const [form] = Form.useForm();
    const handleFormSubmit = async (values) => {
        try {
            const formData = new FormData();

            formData.append('movieTypeName', values.movieTypeName);

            // Send a POST request using Axios

            const response = await axios.post(
                'http://localhost:8080/manager/user/movie/add',
                formData,
                {
                    headers: {
                        // Authorization: token,
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );
            if (response.data.result.code === 0) {
                showSuccess('Tạo thành công');
                return;
            } else if (response.data.result.code === 40) {
                showWarning("Tên loại phim đã tồn tại");
                return;
            } else {
                showError("error server");
                return;
            }
        } catch (error) {
            console.log(error);
            showError("error server");
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
        <div>
            <Row>
                <Col style={{ padding: '0 16px' }}>
                    <Form
                        {...layout}
                        form={form}
                        onFinish={handleFormSubmit}
                    >
                        <Form.Item
                            label="Nhập  loại phim"
                            className="form-row"
                            name="movieTypeName"
                            rules={[{ required: true, message: 'Vui lòng nhập loại phim!' }]}
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                            <Button type="primary" htmlType="submit">
                                Submit
                            </Button>
                        </Form.Item>
                    </Form>
                </Col>
                <Col>

                    <Table dataSource={listMovieType} rowKey="id">

                        <div>
                            <Table.Column title="Thể loại phim" dataIndex="movieTypeName" key="id" />
                        </div>

                    </Table>
                </Col>
            </Row>
        </div>
    )
}
