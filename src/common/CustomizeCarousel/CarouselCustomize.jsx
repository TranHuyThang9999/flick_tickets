import React, { useState } from 'react';
import { Image } from 'antd';
import { ArrowLeftOutlined, ArrowRightOutlined } from '@ant-design/icons';

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
    <div style={{ height: '350px', border: 'solid 1px', width: '300px', display: 'flex', flexDirection: 'column', backgroundColor: 'beige' }}>
      <div style={{ flex: 1, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        {currentImage && (
          <div style={{ maxWidth: '100%', maxHeight: '100%' }}>
            <Image
              style={{ width: '100%', height: '100%', objectFit: 'contain' }}
              src={currentImage.url}
            />
          </div>
        )}
      </div>
      <div style={{ display: 'flex', justifyContent: 'center', marginBottom: '10px' }}>
        <button
          style={{ backgroundColor: 'beige', marginRight: '5px', padding: '1px 20px', border: 'none', color: 'white' }}
          onClick={handlePreviousClick}
        >
          <ArrowLeftOutlined style={{ color: 'darkblue', fontSize: '15px' }} />
        </button>
        <button
          style={{ backgroundColor: 'beige', marginLeft: '5px', padding: '1px 20px', border: 'none', color: 'white' }}
          onClick={handleNextClick}
        >
          <ArrowRightOutlined style={{ color: 'darkblue', fontSize: '15px' }} />
        </button>
      </div>
    </div>
  );
}