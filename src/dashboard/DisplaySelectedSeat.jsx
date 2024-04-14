import { Form, InputNumber } from 'antd';
import React, { useState } from 'react';

function Square({ size, borderWidth, borderColor, index, onClick, disabled, selected }) {
  const style = {
    width: `${size}px`,
    height: `${size}px`,
    border: `${borderWidth}px solid ${borderColor}`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
    backgroundColor: disabled ? 'gray' : selected ? 'yellow' : 'white',
  };

  return (
    <div style={style} onClick={onClick} disabled={disabled}>
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

export default function DisplaySelectedSeat() {
  const [numSquares, setNumSquares] = useState(40);
  const [squareSize, setSquareSize] = useState(50);
  const [borderWidth, setBorderWidth] = useState(1);
  const [borderColor, setBorderColor] = useState('black');
  const [containerWidth, setContainerWidth] = useState(400);
  const [containerHeight, setContainerHeight] = useState(400);
  const [disabledSquares, setDisabledSquares] = useState([]);
  const [selectedSquares, setSelectedSquares] = useState([]);

  const renderSelectedSquares = () => {
    return selectedSquares.map((squareIndex) => (
      <div key={squareIndex} style={{ marginBottom: '5px' }}>
        Ô button {squareIndex + 1}
      </div>
    ));
  };

  const handleNumSquaresChange = (value) => {
    setNumSquares(value);
  };

  const handleContainerWidthChange = (value) => {
    setContainerWidth(value);
  };

  const handleContainerHeightChange = (value) => {
    setContainerHeight(value);
  };

  const handleSquareClick = (index) => {
    const updatedDisabledSquares = [...disabledSquares];
    updatedDisabledSquares[index] = !updatedDisabledSquares[index];
    setDisabledSquares(updatedDisabledSquares);

    const selectedSquareIndex = selectedSquares.indexOf(index);
    if (selectedSquareIndex !== -1) {
      const updatedSelectedSquares = [...selectedSquares];
      updatedSelectedSquares.splice(selectedSquareIndex, 1);
      setSelectedSquares(updatedSelectedSquares);
    } else {
      setSelectedSquares([...selectedSquares, index]);
    }
  };

  return (
    <div>
      <Form>
        <Form.Item label="Nhập độ dài của phòng dạp" className="form-row" name="ContainerHeight">
          <InputNumber value={containerHeight} onChange={handleContainerHeightChange} />
        </Form.Item>
        <Form.Item label="Nhập độ rộng của phòng dạp" className="form-row" name="ContainerWidth">
          <InputNumber value={containerWidth} onChange={handleContainerWidthChange} />
        </Form.Item>
        <Form.Item label="Số lượng hình vuông" className="form-row" name="NumSquares">
          <InputNumber value={numSquares} onChange={handleNumSquaresChange} />
        </Form.Item>
        <SquareContainer width={containerWidth} height={containerHeight}>
          {Array.from({ length: numSquares }).map((_, index) => (
            <Square
              key={index}
              size={squareSize}
              borderWidth={borderWidth}
              borderColor={borderColor}
              index={index + 1}
              onClick={() => handleSquareClick(index)}
              disabled={disabledSquares[index]}
              selected={selectedSquares.includes(index)}
            />
          ))}
        </SquareContainer>
        <div style={{ marginTop: '20px' }}>
          <h3>Các ô button đã chọn:</h3>
          {renderSelectedSquares()}
        </div>
      </Form>
    </div>
  );
}