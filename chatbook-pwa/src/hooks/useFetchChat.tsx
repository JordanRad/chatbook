import { useEffect, useState } from 'react';
import UserService from '../services/UserService';

const useFetchChat = (conversationID:string) => {
    const [chat, setChat] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetch = async () => {
            setError(null);
            const response = await UserService.getChatHistory(conversationID)

            console.log("Resp: ",response.statusText)
            if (response.status === 200) {
                const data = response.data;
                setChat(data.messages);
            };
        }

        fetch();
    }, []);

    return { chat, error };
};

export default useFetchChat;
