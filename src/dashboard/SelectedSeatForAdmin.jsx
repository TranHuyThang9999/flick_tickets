import React, { useState, useEffect } from 'react';
import { Col, Row, Space } from 'antd';

function Square({ size = 30, index, onClick, disabled, selected, inSelectedSeatGetFormApi }) {
  const style = {
    width: `${size}px`,
    height: `${size}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
    border: '1px solid brown',
    backgroundColor: selected ? 'red' : inSelectedSeatGetFormApi ? 'yellow' : disabled ? 'gray' : 'white',
  };

  const handleClick = () => {
    if (!disabled) {
      onClick(index);
    }
  };

  return (
    <div style={style} onClick={handleClick} disabled={disabled}>
      {index}
    </div>
  );
}

function SquareContainer({ width, height, children }) {
  const style = {
    width: `${width}px`,
    height: `${height}px`,
    display: 'flex',
    flexWrap: 'wrap',
    border: '1px solid black',
    boxSizing: 'border-box',
  };

  return <div style={style}>{children}</div>;
}

export default function SelectedSeatForAdmin({
  SelectedSeatGetFormApi,
  widthContainerUseSavedate,
  heightContainerUseSaveData,
  numSquares,
  onCreate, // Hàm callback để truyền danh sách ghế đã chọn
}) {
  const [containerWidth, setContainerWidth] = useState(widthContainerUseSavedate || 400);
  const [containerHeight, setContainerHeight] = useState(heightContainerUseSaveData || 400);
  const [disabledSquares, setDisabledSquares] = useState([]);
  const [selectedSeats, setSelectedSeats] = useState([]); // Thêm state mới để lưu trữ vị trí các ghế đã chọn

  useEffect(() => {
    const selectedSeatsArray = SelectedSeatGetFormApi.split(',').map((seat) => parseInt(seat.trim(), 10));
    const disabledSquaresList = Array.from({ length: numSquares }).map((_, index) =>
      selectedSeatsArray.includes(index + 1)
    );
    setDisabledSquares(disabledSquaresList);
  }, [SelectedSeatGetFormApi, numSquares]);

  const handleSquareClick = (index) => {
    if (disabledSquares[index - 1]) {
      return; // Nếu ô đã bị disable thì không xử lý click
    }
  
    setSelectedSeats((prevSelectedSeats) => {
      const selectedIndex = prevSelectedSeats.indexOf(index);
      if (selectedIndex !== -1) {
        // Nếu ô đã được chọn trước đó, loại bỏ nó khỏi danh sách
        return [...prevSelectedSeats.slice(0, selectedIndex), ...prevSelectedSeats.slice(selectedIndex + 1)];
      } else {
        // Nếu ô chưa được chọn, thêm nó vào danh sách
        return [...prevSelectedSeats, index];
      }
    });
  };
  
  useEffect(() => {
    onCreate(selectedSeats.map(index => `${index}`)); // Truyền danh sách vị trí ghế đã chọn thay vì disabledSquares
  }, [selectedSeats, onCreate]);

  return (
    <div className="form-selected-seat">
      <Space direction="vertical" size="middle" style={{ display: 'flex' }}>
        <Row>
          <Col style={{ padding: '0 16px' }}>
            <SquareContainer width={containerWidth} height={containerHeight}>
              {Array.from({ length: numSquares }).map((_, index) => (
                <Square
                  key={index + 1}
                  index={index + 1}
                  onClick={handleSquareClick}
                  disabled={disabledSquares[index]}
                  selected={selectedSeats.includes(index + 1)} // Sử dụng selectedSeats thay vì disabledSquares
                />
              ))}
            </SquareContainer>
          </Col>
          <Col style={{ display: 'flex', marginTop: '-18px' }}>
            <div>
              <h2>Selected Squares:</h2>
              <ul>
                {selectedSeats.map((seat, index) => ( // Hiển thị danh sách các ghế đã chọn
                  <li key={index}>Square {seat}</li>
                ))}
              </ul>
            </div>
          </Col>
        </Row>
      </Space>
    </div>
  );
}
