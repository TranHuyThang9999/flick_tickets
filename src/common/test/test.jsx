import { Form } from 'antd';
import React, { useEffect, useState } from 'react';

function Square({ size, borderWidth, borderColor, index, onClick, disabled, selected }) {
  const style = {
    width: `${size}px`,
    height: `${size}px`,
    border: `${borderWidth}px solid ${borderColor}`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    cursor: 'pointer',
    backgroundColor: disabled ? 'gray' : (selected ? 'yellow' : 'white'),
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

function AppSquare({ Number, ContainerWidth, ContainerHeight }) {
  const [numSquares, setNumSquares] = useState(40);
  const [squareSize, setSquareSize] = useState(50);
  const [borderWidth, setBorderWidth] = useState(1);
  const [borderColor, setBorderColor] = useState('black');
  const [containerWidth, setContainerWidth] = useState(400);
  const [containerHeight, setContainerHeight] = useState(400);
  const [disabledSquares, setDisabledSquares] = useState([]);
  const [selectedSquares, setSelectedSquares] = useState([]);
  const [restored, setRestored] = useState(false);

  const handleNumSquaresChange = (e) => {
    const value = parseInt(e.target.value);
    setNumSquares(value);
  };

  const handleSquareSizeChange = (e) => {
    const value = parseInt(e.target.value);
    setSquareSize(value);
  };

  const handleBorderWidthChange = (e) => {
    const value = parseInt(e.target.value);
    setBorderWidth(value);
  };

  const handleBorderColorChange = (e) => {
    const value = e.target.value;
    setBorderColor(value);
  };

  const handleContainerWidthChange = (e) => {
    const value = parseInt(e.target.value);
    setContainerWidth(value);
  };

  const handleContainerHeightChange = (e) => {
    const value = parseInt(e.target.value);
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

  const handleSaveState = () => {
    if (restored) {
      const stateToSave = {
        numSquares,
        squareSize,
        selectedSquares,
        containerHeight,
        containerWidth,
      };

      localStorage.setItem('appState', JSON.stringify(stateToSave));
    }
  };

  useEffect(() => {
    const savedState = localStorage.getItem('appState');

    if (savedState && !restored) {
      const parsedState = JSON.parse(savedState);
      setNumSquares(parsedState.numSquares);
      setSquareSize(parsedState.squareSize);
      setSelectedSquares(parsedState.selectedSquares);
      setContainerHeight(parsedState.containerHeight);
      setContainerWidth(parsedState.containerWidth);
      setRestored(true);
    }
  }, [restored]);

  useEffect(() => {
    handleSaveState();
  }, [numSquares, squareSize, selectedSquares, containerHeight, containerWidth]);

  return (
    <div>
      <Form>
        <Form.Item>
          <label>
            Number of squares:
            <input
              type="number"
              value={numSquares}
              onChange={handleNumSquaresChange}
            />
          </label>
        </Form.Item>
        <Form.Item>
          <label>
            Container width:
            <input
              type="number"
              value={containerWidth}
              onChange={handleContainerWidthChange}
            />
          </label>
        </Form.Item>
        <Form.Item>
          <label>
            Container height:
            <input
              type="number"
              value={containerHeight}
              onChange={handleContainerHeightChange}
            />
          </label>
        </Form.Item>
      </Form>

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

export default AppSquare;