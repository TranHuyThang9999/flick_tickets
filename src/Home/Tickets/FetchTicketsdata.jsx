import axios from 'axios';

export default async function FetchTicketsdata(values) {
  try {
    const response = await axios.get('http://localhost:8080/manager/customers/ticket', {
      params: values,
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });

    if (response.data.result.code === 0) {
      return {
        code: response.data.result.code,
        data: response.data.list_tickets,
      };
    } else if (response.data.result.code === 20) {
      return {
        code: 20,
        data: null,
      };
    } else {
      throw new Error('Lỗi server');
    }
  } catch (error) {
    throw new Error('Lỗi server');
  }
}