import { Image } from 'antd';
import React from 'react';
import Carousel from 'react-multi-carousel';
import 'react-multi-carousel/lib/styles.css';

export default function Slider({ data,withImage,deviceType }) {
  const responsive = {
    desktop: {
      breakpoint: { max: 3000, min: 1024 },
      items: 3,
      slidesToSlide: 3,
    },
    tablet: {
      breakpoint: { max: 1024, min: 464 },
      items: 2,
      slidesToSlide: 2,
    },
    mobile: {
      breakpoint: { max: 464, min: 0 },
      items: 1,
      slidesToSlide: 1,
    },
  };

  return (
    <div>
      <Carousel
        swipeable={false}
        draggable={false}
        showDots={true}
        responsive={responsive}
        ssr={true}
        infinite={true}
        autoPlay={deviceType !== 'mobile'}
        autoPlaySpeed={1000}
        keyBoardControl={true}
        customTransition="all .5"
        transitionDuration={500}
        containerClass="carousel-container"
        removeArrowOnDeviceType={['tablet', 'mobile']}
        deviceType={deviceType}
        dotListClass="custom-dot-list-style"
        itemClass="carousel-item-padding-40-px"
      >
        <div>
            <Image width={100} src='http://localhost:1234/manager/shader/huythang/559888234.jpeg'/>
        </div>
        <div>Item 2
        <Image width={100} src='http://localhost:1234/manager/shader/huythang/174814364.png'/>
        </div>
        <div>Item 3
        <Image width={100} src='http://localhost:1234/manager/shader/huythang/532981219.jpg'/>
        </div>
        <div>Item 4
        <Image width={100} src='http://localhost:1234/manager/shader/huythang/857678238.jpeg'/>
        </div>
        <div>Item 5
        <Image width={100} src='http://localhost:1234/manager/shader/huythang/999919283.png'/>
        </div>
        <div>
            {data.map((item)=>(
                <div key={item.id}>
                    <Image src={data.url} width={withImage}/>
                </div>
            ))}
        </div>
      </Carousel>
    </div>
  );
}