import { Button, Space, AutoComplete } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import GetListTicketByMovieName from '../ViewTicketsForSell/GetListTicketByMovieName';

export default function TestSelect() {
  const [tickets, setTickets] = useState([]);
  const [statusFindMovieName, setStatusMovieName] = useState(false);
  const [selectedMovie, setSelectedMovie] = useState('');
  const [searchInput, setSearchInput] = useState('');

  useEffect(() => {
    const fetchTickets = async () => {
      try {
        const response = await axios.get('http://localhost:8080/manager/customers/ticket');
        if (response.data.result.code === 0) {
          setTickets(response.data.list_tickets);
        } else if (response.data.result.code === 20) {
          console.log('No tickets available.');
        }
      } catch (error) {
        console.error('Error fetching tickets:', error);
      }
    };

    fetchTickets();
  }, []);

  const handleSearch = () => {
    setStatusMovieName(true);
    setSelectedMovie(searchInput);
  };

  const dataSource = tickets.map((ticket) => ({
    value: ticket.name,
  }));

  if (statusFindMovieName) {
    return <GetListTicketByMovieName movieName={selectedMovie} />;
  }

  return (
    <div>
      <Space direction="vertical" size="middle">
        <Space.Compact>
          <AutoComplete
            options={dataSource}
            onSelect={(value) => setSearchInput(value)}
            onChange={(value) => setSearchInput(value)}
            placeholder="Nhập chữ S để tìm kiếm"
            style={{ width: 200 }}
          />
          <Button onClick={handleSearch}>Tìm kiếm</Button>
        </Space.Compact>
      </Space>
    </div>
  );
}
