import React, { useState } from 'react';
import { showError, showSuccess } from '../common/log/log';
import { Form, DatePicker,Upload } from 'antd';
import { Axios } from 'axios';


export default function AdminUploadTickets() {


    const [fileList, setFileList] = useState([]);

    const onChange = ({ fileList }) => {
        setFileList(fileList);
    };



    const handleFormSubmit = async (values) => {

        try {
            const [] = Form.useForm();

            const token = localStorage.getItem('token');

            const formData = new FormData();
            formData.append('name', values.name);
            formData.append('price', values.price);
            formData.append('quantity', values.quantity);
            formData.append('description', values.description);
            formData.append('showtime', values.showtime);
            formData.append('status', values.status ? values.status.value : '');
            formData.append('sales_discount', values.sales_discount);

            fileList.forEach((file) => {
                formData.append('file', file.originFileObj);
            });


            // Send a POST request using Axios
            const response = await Axios.post(
                'http://localhost:8080/manager/user/upload/ticket',
                formData,
                {
                    headers: {
                        Authorization: token,
                        'Content-Type': 'multipart/form-data',
                    },
                }
            );
            if (response.data.result.code == 0) {
                showSuccess("uopload  vé thành công");
            }
        } catch (error) {
            console.log(error);
            showError("lỗi server vui lòng thử lại");
            return;
        }
    };
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
            <Form
                {...layout}
                form={form}
                className='form-container'
                onFinish={handleFormSubmit}
            >

                <Form.Item
                    label='Nhập ảnh mô tả sản phẩm'
                    className='form-row'
                    name='file'
                    rules={[
                        { required: true, message: 'Vui lòng tên sản phẩm của bạn!' },
                    ]}
                >
                    <Upload
                        maxCount={5}
                        fileList={fileList}
                        listType='picture-card'
                        accept='image/jpeg,image/png'
                        onChange={onChange}
                        beforeUpload={() => false} // Prevent auto-upload
                    >
                        {fileList.length < 5 && '+ Upload'}
                    </Upload>
                    {/* {fileList.length > 5 && (
                // <p style={{ color: 'red' }}>Upload tối đa 5 ảnh mô tả sản phẩm</p>
                message.warning('Upload tối đa 5 ảnh mô tả sản phẩm')
            )} */}
                </Form.Item>

                <Form.Item>
                    <DatePicker
                        showTime />
                </Form.Item>

            </Form>
        </div>
    );
}
