import React, { useEffect, useState } from 'react';

export default function HTMLContent() {
  const [content, setContent] = useState('');

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8080/load');
      const result = await response.text();
      setContent(result);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div dangerouslySetInnerHTML={{ __html: content }} />
  );
}