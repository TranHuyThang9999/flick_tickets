import React, { useState } from 'react';
import './index.css';

function Square({ size = 30, index, onClick, disabled, selected }) {
  const style = {
    width: `${size}px`,
    height: `${size}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
    border: '1px solid brown',
    backgroundColor: disabled ? 'gray' : selected ? 'red' : 'white', // Thay đổi màu đỏ
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

export default function SelectedSeat({
  SelectedSeatGetFormApi,
  widthContainerUseSavedate,
  heightContainerUseSaveData,
  numSquares,
}) {
  const [containerWidth, setContainerWidth] = useState(widthContainerUseSavedate || 400);
  const [containerHeight, setContainerHeight] = useState(heightContainerUseSaveData || 400);
  const [disabledSquares, setDisabledSquares] = useState([]);
  const [selectedSquares, setSelectedSquares] = useState([]);
  const [restored, setRestored] = useState(false);

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
      <SquareContainer width={containerWidth} height={containerHeight}>
        {Array.from({ length: numSquares }).map((_, index) => (
          <Square
            key={index}
            index={index + 1}
            onClick={() => handleSquareClick(index)}
            disabled={disabledSquares[index]}
            selected={selectedSquares.includes(index) || SelectedSeatGetFormApi.includes(index + 1)} // Thêm điều kiện để hiển thị màu đỏ
          />
        ))}
      </SquareContainer>

      <div>
        <h2>Selected Squares:</h2>
        <ul>
          {selectedSquares.map((index) => (
            <li key={index}>Square {index + 1}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}