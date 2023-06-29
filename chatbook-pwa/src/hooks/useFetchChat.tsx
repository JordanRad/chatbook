import { useEffect, useState } from 'react';
import UserService from '../services/UserService';

const useFetchChat = (conversationID:string) => {
    const [chat, setChat] = useState<any[]>([]);
    const [error, setError] = useState(null);

    const updateWithMessage = (message:any)=>{
        const m = {
            timestamp: Date.now(),
            senderID: "vjfskvjsf",
            content: message
        }
        setChat((prevChat)=>[m,...prevChat])
    }
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

    return { chat, error, updateWithMessage };
};

export default useFetchChat;
