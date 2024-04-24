import React, { useState, useEffect } from 'react';
import { Col, Row, Space } from 'antd';

function Square({ size = 30, index, onClick, disabled, selected }) {
  const style = {
    width: `${size}px`,
    height: `${size}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
    border: '1px solid brown',
    backgroundColor: selected ? 'red' : disabled ? 'gray' : 'white',
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

export default function SelectedSeat({
  SelectedSeatGetFormApi,
  widthContainerUseSavedate,
  heightContainerUseSaveData,
  numSquares,
  onCreate, // Hàm callback để truyền danh sách ghế đã chọn
}) {
  const [containerWidth, setContainerWidth] = useState(widthContainerUseSavedate || 400);
  const [containerHeight, setContainerHeight] = useState(heightContainerUseSaveData || 400);
  const [disabledSquares, setDisabledSquares] = useState([]);

  useEffect(() => {
    const selectedSeatsArray = SelectedSeatGetFormApi.split(',').map((seat) => parseInt(seat.trim(), 10));
    const disabledSquaresList = Array.from({ length: numSquares }).map((_, index) =>
      selectedSeatsArray.includes(index + 1)
    );
    setDisabledSquares(disabledSquaresList);
  }, [SelectedSeatGetFormApi, numSquares]);

  const handleSquareClick = (index) => {
    setDisabledSquares((prevDisabledSquares) => {
      const updatedDisabledSquares = [...prevDisabledSquares];
      updatedDisabledSquares[index - 1] = !updatedDisabledSquares[index - 1];
      return updatedDisabledSquares;
    });
  };

  useEffect(() => {
    onCreate(
      disabledSquares
        .map((disabled, index) => (disabled ? index + 1 : null))
        .filter((seat) => seat !== null)
        .map((index) => `${index}`)
    );
  }, [disabledSquares, onCreate]);

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
                  selected={false} // Không cần kiểm tra selected nữa vì màu nền được xác định bởi disabled
                />
              ))}
            </SquareContainer>
          </Col>
          <Col style={{ display: 'flex', marginTop: '-18px' }}>
            <div>
              <h2>Selected Squares:</h2>
              <ul>
                {disabledSquares.map((disabled, index) => (
                  disabled && <li key={index}>Square {index + 1}</li>
                ))}
              </ul>
            </div>
          </Col>
        </Row>
      </Space>
    </div>
  );
}
