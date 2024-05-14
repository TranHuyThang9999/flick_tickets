import { Button, Col, Image, Row } from 'antd'
import React from 'react'

export default function Blogs() {

  const handlerGobackHome = () =>{
    window.location.reload();
  }

  return (
    <div>
      <Button onClick={handlerGobackHome}>Quay lại trang chủ</Button>
      <Row>
        <Col style={{ fontSize: '20px', padding: '10px' }} span={12}>
          Call me Chihiro - Tôi Là Chihiro là kể về sự tái sinh của một cô gái từng làm nghề bán dâm, chuyển thể từ truyện tranh Chihiro-san của tác giả Hiroyuki Yasuda.

          Chihiro (Arimura Kasumi) trong quá khứ là một gái mại dâm, cô tình cờ ghé qua một thị trấn ven biển và thưởng thức một bữa cơm bento. Bữa ăn ngon miệng khiến cô quyết định ở lại thị trấn và làm thu ngân tại một cửa hàng cơm hộp.

          Từ đây, quán ăn nhanh chóng thu hút sự chú ý của cánh đàn ông trong khu vực. Dù luôn tươi cười và thân thiện với khách hàng, nhưng bên trong, Chihiro vẫn cảm thấy trống trải trong tâm hồn.

          Tuy vậy, qua thời gian dài trò chuyện với khách hàng về nhiều vấn đề khác nhau trong cuộc sống, cô dần tìm thấy sự an ủi và ý nghĩa mới. Nhờ vào những cuộc trò chuyện này, cô từ bỏ quá khứ và bắt đầu xây dựng lại cuộc sống mới cho mình.


        </Col>
        <Col span={12}>
          <Image src='https://cdn.tgdd.vn/Files/2024/05/07/1565419/diem-danh-6-bo-phim-moi-nhat-cua-ngoc-nu-arimura-kasumi-202405071658273055.jpg' />
        </Col>
      </Row>
      <Row>
        <Col style={{ fontSize: '20px', padding: '10px' }} span={12}>
          Điểm IMDb: 7.7/10

          Ngày ra mắt: 26/04/2024

          Thể loại phim: Hành động, Tội phạm

          Thời lượng phim: 109 phút

          Quốc gia: Hàn Quốc

          Đạo diễn: Heo Myeong Haeng

          Diễn viên chính: Ma Dong-seok, Kim Mu-yeol, Lee Joo-bin

          Giải thưởng nổi bật: Đang cập nhật

          Link xem phim: Đang cập nhật
        </Col>
        <Col span={12}>
          <Image src='https://cdn.tgdd.vn/Files/2024/04/23/1564920/le-30-4-1-5-nhieu-phim-hay-ra-rap-cho-ban-thuong-thuc-202404231426583396.jpg' />
        </Col>
      </Row>
      <Row>
        <Col style={{ fontSize: '20px', padding: '10px' }} span={12}>
          Điểm IMDb: 6.1/10

          Ngày ra mắt: 26/04/2024

          Thể loại phim: Hoạt hình, Hài, Phiêu lưu, Gia đình

          Thời lượng phim: 88 phút

          Quốc gia: Hoa Kỳ

          Đạo diễn: Christopher Jenkins

          Diễn viên lồng tiếng: Mo Gilligan, Simone Ashley, Sophie Okonedo

          Giải thưởng nổi bật: Đang cập nhật

          Link xem phim: Đang cập nhật
        </Col>
        <Col span={12}>
          <Image src='https://cdn.tgdd.vn/Files/2024/04/23/1564920/le-30-4-1-5-nhieu-phim-hay-ra-rap-cho-ban-thuong-thuc-202404231428542929.jpg' />
        </Col>
      </Row>
    </div>
  )
}
