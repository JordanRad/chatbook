import { useEffect, useState } from 'react';
import UserService from '../services/UserService';

const useFetchLastConversations = () => {
    const [conversations, setConversations] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchProfile = async () => {
            setError(null);
            const response = await UserService.getLastConversations()

            console.log("Resp: ",response.statusText)
            if (response.status === 200) {
                const data = response.data;
                setConversations(data.resources);
            };
        }

        fetchProfile();
    }, []);

    return { conversations, error };
};

export default useFetchLastConversations;
