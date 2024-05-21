import { DatePicker } from 'antd';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, Legend, CartesianGrid } from 'recharts';

const { RangePicker } = DatePicker;

export default function StatisticalBar() {
    const [movieSales, setMovieSales] = useState([]);
    const [dateRange, setDateRange] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            if (!dateRange) return;

            const [startDate, endDate] = dateRange;
            try {
                const response = await axios.get('http://localhost:8080/manager/user/statistical', {
                    params: {
                        start_time: startDate.unix(),
                        end_time: endDate.unix()
                    }
                });
                if (response.data.result.code === 0) {
                    const orders = response.data.orders;
                    const sales = {};

                    // Tính tổng số vé bán và tổng doanh thu mỗi phim
                    orders.forEach(order => {
                        const movieName = order.movie_name;
                        const price = order.price * order.sale; // Giá vé * số lượng vé
                        if (!sales[movieName]) {
                            sales[movieName] = price;
                        } else {
                            sales[movieName] += price;
                        }
                    });

                    // Chuyển dữ liệu sang mảng để sử dụng trong BarChart
                    const salesArray = Object.entries(sales).map(([movieName, totalRevenue]) => ({
                        movieName,
                        totalRevenue
                    }));

                    setMovieSales(salesArray);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };
        fetchData();
    }, [dateRange]);

    return (
        <div>
            <RangePicker
                renderExtraFooter={() => 'extra footer'}
                onChange={(dates) => setDateRange(dates)}
            />
            <div>
                <h2>Tổng doanh thu mỗi phim:</h2>
                <BarChart width={600} height={400} data={movieSales}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="movieName" />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <Bar dataKey="totalRevenue" fill="#8884d8" />
                </BarChart>
            </div>
        </div>
    );
}
