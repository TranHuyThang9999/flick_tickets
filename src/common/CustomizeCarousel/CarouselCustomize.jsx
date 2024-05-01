import React, { useState } from 'react';
import { Image } from 'antd';
import { ArrowLeftOutlined, ArrowRightOutlined } from '@ant-design/icons';
import './CarouselCustomize.css';

export default function CarouselCustomize({ images }) {
  const [index, setIndex] = useState(0);
  const hasNext = images.length > 1 && index < images.length - 1;
  const hasPrevious = images.length > 1 && index > 0;

  function handleNextClick() {
    if (hasNext) {
      setIndex(index + 1);
    } else {
      setIndex(0);
    }
  }

  function handlePreviousClick() {
    if (hasPrevious) {
      setIndex(index - 1);
    } else {
      setIndex(images.length - 1);
    }
  }

  const currentImage = images.length > 0 ? images[index] : null;

  return (
    <div className="carousel-container">
      <div className="carousel-image-container">
        {currentImage && (
          <div style={{ maxWidth: '100%', maxHeight: '100%' }}>
            <Image
              style={{ width: '100%', height: '100%', objectFit: 'contain',borderRadius:'10px'}}
              src={currentImage.url}
            />
          </div>
        )}
      </div>
      <div className="carousel-buttons-container">
        <button className="carousel-button-shadow" onClick={handlePreviousClick}>
          <ArrowLeftOutlined style={{ color: 'darkblue', fontSize: '15px' }} />
        </button>
        <button  className="carousel-button-shadow" onClick={handleNextClick}>
          <ArrowRightOutlined style={{ color: 'darkblue', fontSize: '15px' }} />
        </button>
      </div>
    </div>
  );
}
