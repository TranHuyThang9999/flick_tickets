import { Button, DatePicker, Select, Space } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { showError } from '../common/log/log';

const { RangePicker } = DatePicker;

export default function RevenueOrder() {
    const [orders, setOrders] = useState([]);
    const [sum, setSum] = useState(0);
    const [selectedMovie, setSelectedMovie] = useState(null);
    const [selectedCinema, setSelectedCinema] = useState(null);
    const [dateRange, setDateRange] = useState(null);
    const [distinctCinemas, setDistinctCinemas] = useState([]);
    const [distinctMovieNames, setDistinctMovieNames] = useState([]);

    useEffect(() => {
        axios.get('http://localhost:8080/manager/user/order/getlist', {
            params: {
                status: 9
            }
        })
            .then(response => {
                if (response.data.result.code === 0) {
                    setOrders(response.data.orders);
                }
            })
            .catch(error => {
                console.error('There was an error fetching the order list!', error);
            });
    }, []);

    const fetchDistinctCinemas = async () => {
        try {
            const response = await axios.get('http://localhost:8080/manager/user/history');
            if (response.data.result.code === 0) {
                setDistinctCinemas(response.data.orders);
            } else {
                showError("Error server");
            }
        } catch (error) {
            console.error('There was an error fetching the cinema names!', error);
        }
    };

    useEffect(() => {
        fetchDistinctCinemas();
    }, []);

    useEffect(() => {
        if (selectedCinema) {
            const fetchMovieNames = async () => {
                try {
                    const response = await axios.get(`http://localhost:8080/manager/user/history/movie/name?cinema_name=${selectedCinema}`);
                    if (response.data.result.code === 0) {
                        setDistinctMovieNames(response.data.orders);
                    } else {
                        showError("Error server");
                    }
                } catch (error) {
                    console.error('There was an error fetching the movie names!', error);
                }
            };

            fetchMovieNames();
        }
    }, [selectedCinema]);

    const handleCalculate = () => {
        if (selectedMovie && selectedCinema && dateRange) {
            const [startDate, endDate] = dateRange;

            axios.get('http://localhost:8080/manager/user/order/revenue', {
                params: {
                    cinema_name: selectedCinema,
                    movie_name: selectedMovie,
                    time_distance_start: startDate.unix(),
                    time_distance_end: endDate.unix()
                }
            })
                .then(response => {
                    if (response.data.result.code === 0) {
                        setSum(response.data.sum);
                    }
                })
                .catch(error => {
                    console.error('There was an error fetching the revenue data!', error);
                });
        } else {
            alert('Please select a movie, cinema, and date range.');
        }
    };

    const optionsMovie = distinctMovieNames.map((movie) => ({
        value: movie.movie_name,
        label: movie.movie_name
    }));

    const optionsCinema = distinctCinemas.map((cinema) => ({
        value: cinema.cinema_name,
        label: cinema.cinema_name
    }));

    return (
        <div>
            <Space>
                <Space.Compact>
                    <Select
                        placeholder='Nhập tên phòng'
                        showSearch
                        style={{ width: '200px' }}
                        options={optionsCinema}
                        onChange={setSelectedCinema}
                    />
                    <Select
                        placeholder='Nhập tên phim'
                        showSearch
                        style={{ width: '200px' }}
                        options={optionsMovie}
                        onChange={setSelectedMovie}
                    />
                    <RangePicker
                        renderExtraFooter={() => 'extra footer'}
                        onChange={(dates) => setDateRange(dates)}
                    />
                </Space.Compact>
            </Space>
            <Button onClick={handleCalculate}>Tính</Button>
            <div>
                <h3>Doanh thu: {sum} VND</h3>
            </div>
        </div>
    );
}
