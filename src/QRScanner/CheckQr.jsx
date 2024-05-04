import React from 'react';
import axios from 'axios';
import { showError, showSuccess, showWarning } from '../common/log/log';
import { Button } from 'antd';
// ko dung
export default function CheckQr({ token }) {

  const handleButtonClick = async () => {
    try {
      const response = await axios.get('http://localhost:8080/manager/user/verify/aes', {
        params: {
          token: token
        }
      });
      console.log(response.data);
      // Xử lý dữ liệu nhận được từ API tại đây
      if (response.data.result.code === 0) {
        showSuccess("ok");
        return;
      } else if (response.data.result.code === 18) {
        showWarning("QRcode valid");
        return;
      }else if(response.data.result.code === 12){
        showError("bad req");
        return;
      } 
      else {
        showError("error server");
        return;
      }
    } catch (error) {
      console.error(error);
      showError("error server");
      return;
    }
  };

  return (
    <div>
      <Button onClick={handleButtonClick} >
        Submit
      </Button>
    </div>
  );
}
