import React, { useState, useEffect } from 'react';
import './index.css';

function Square({ size = 30, index, onClick, disabled, selected, inSelectedSeatGetFormApi }) {
  const style = {
    width: `${size}px`,
    height: `${size}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
    border: '1px solid brown',
    backgroundColor: inSelectedSeatGetFormApi ? 'yellow' : selected ? 'red' : disabled ? 'gray' : 'white',
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
}) {
  const [containerWidth, setContainerWidth] = useState(widthContainerUseSavedate || 400);
  const [containerHeight, setContainerHeight] = useState(heightContainerUseSaveData || 400);
  const [disabledSquares, setDisabledSquares] = useState([]);
  const [selectedSquares, setSelectedSquares] = useState([]);
  const [restored, setRestored] = useState(false);

  useEffect(() => {
    // Lấy danh sách các ô bị vô hiệu hóa từ SelectedSeatGetFormApi
    const disabledSquaresList = Array.from({ length: numSquares }).map((_, index) =>
      SelectedSeatGetFormApi.includes(index + 1)
    );
    setDisabledSquares(disabledSquaresList);

    // Đánh dấu những ô đã được chọn từ SelectedSeatGetFormApi
    const selectedSquaresList = Array.from({ length: numSquares })
      .map((_, index) => index + 1)
      .filter((index) => SelectedSeatGetFormApi.includes(index));
    setSelectedSquares(selectedSquaresList);
  }, [SelectedSeatGetFormApi, numSquares]);

  const handleSquareClick = (index) => {
    const isDisabled = disabledSquares[index - 1];
    const isSelected = selectedSquares.includes(index);

    if (isDisabled) {
      return;
    }

    if (isSelected) {
      const updatedSelectedSquares = selectedSquares.filter((squareIndex) => squareIndex !== index);
      setSelectedSquares(updatedSelectedSquares);
    } else {
      setSelectedSquares([...selectedSquares, index]);
    }
  };
  console.log(selectedSquares
    .filter((index) => !SelectedSeatGetFormApi.includes(index))
    .map((index) => `${index}`));

  return (
    <div>
      <SquareContainer width={containerWidth} height={containerHeight}>
        {Array.from({ length: numSquares }).map((_, index) => (
        <Square
        key={index + 1}
        index={index + 1}
        onClick={handleSquareClick}
        disabled={disabledSquares[index]}
        selected={selectedSquares.includes(index + 1)}
        inSelectedSeatGetFormApi={SelectedSeatGetFormApi.includes(index + 1)}
      />
        ))}
      </SquareContainer>

      <div>
        <h2>Selected Squares:</h2>
        <ul>
          {selectedSquares
            .filter((index) => !SelectedSeatGetFormApi.includes(index))
            .map((index) => (
              <li key={index}>Square {index}</li>
            ))}
        </ul>
      </div>
    </div>
  );
}