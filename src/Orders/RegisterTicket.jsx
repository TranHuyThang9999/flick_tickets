// File: RegisterTicket.js

import { useEffect } from 'react';
import axios from 'axios';
import { showError } from '../common/log/log';

const RegisterTicket = ({ show_time_id, email, seats }) => {

    const registerTicket = async () => {
        try {
            const response = await axios.post(
                'http://localhost:8080/manager/customer/order/ticket',
                {
                    show_time_id: show_time_id,
                    email: email,
                    seats: seats
                },
                {
                    withCredentials: true
                }
            );
            if (response.data.result.code === 0) {
                return;
            }
        } catch (error) {
            showError("error server");
            console.error(error);
            return;
        }
    };

    useEffect(() => {
        registerTicket();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return null;
};

export default RegisterTicket;
