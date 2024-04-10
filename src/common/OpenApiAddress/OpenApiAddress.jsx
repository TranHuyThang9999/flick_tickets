import { Select } from 'antd';
import React, { useState, useEffect } from 'react';

export default function OpenApiAddress() {
  const [cities, setCities] = useState([]);
  const [districts, setDistricts] = useState([]);
  const [communes, setCommunes] = useState([]);

  useEffect(() => {
    // Gọi API để lấy danh sách tỉnh/thành phố
    fetch('http://localhost:8080/manager/public/customer/cities')
      .then(response => response.json())
      .then(data => setCities(data.cities.map(city => city['Tỉnh Thành Phố'])))
      .catch(error => console.error(error));
  }, []);

  const handleCityChange = (selectedCity) => {
    // Gọi API để lấy danh sách huyện/quận dựa trên tỉnh/thành phố đã chọn
    fetch(`http://localhost:8080/manager/public/customer/districts?name=${encodeURIComponent(selectedCity)}`)
      .then(response => response.json())
      .then(data => setDistricts(data.districts.map(district => district['Quận Huyện'])))
      .catch(error => console.error(error));
  };

  const handleDistrictChange = (selectedDistrict) => {
    // Gọi API để lấy danh sách xã/phường dựa trên huyện/quận đã chọn
    fetch(`http://localhost:8080/manager/public/customer/communes?name=${encodeURIComponent(selectedDistrict)}`)
      .then(response => response.json())
      .then(data => setCommunes(data.communes.map(commune => commune['Phường Xã'])))
      .catch(error => console.error(error));
  };

  return (
    <div style={{ display: 'flex' }}>
      <p>Tỉnh</p>
      <Select
        style={{ width: 200 }}
        onChange={handleCityChange}
      >
        {cities.map(city => (
          <Select.Option key={city} value={city}>{city}</Select.Option>
        ))}
      </Select>
      <p>Huyện</p>
      <Select
        style={{ width: 200 }}
        onChange={handleDistrictChange}
      >
        {districts.map(district => (
          <Select.Option key={district} value={district}>{district}</Select.Option>
        ))}
      </Select>
      <p>Xã</p>
      <Select
        style={{ width: 200 }}
      >
        {communes.map(commune => (
          <Select.Option key={commune} value={commune}>{commune}</Select.Option>
        ))}
      </Select>
    </div>
  );
}