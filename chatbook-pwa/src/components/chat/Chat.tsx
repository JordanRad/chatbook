import { Box, Stack, TextField, Typography } from "@mui/material"
import { ConversationListItem } from "../../screens/Main"
import useFetchChat from "../../hooks/useFetchChat"
import { useEffect, useRef, useState } from "react"

export const Chat = (props: { conversationWithParticipant: ConversationListItem }) => {
    const { chat,updateWithMessage } = useFetchChat(props.conversationWithParticipant.id)
    const socketRef = useRef<WebSocket | null>(null);
    const [currentMessage, setCurrentMessage] = useState("")
    useEffect(() => {
        // Establish the WebSocket connection
        socketRef.current = new WebSocket('ws://localhost:6001/');

        // WebSocket event handlers
        socketRef.current.onopen = () => {
            console.log('WebSocket connection established');
        };

        socketRef.current.onmessage = (event) => {
            console.log('Received message:', event.data);
            updateWithMessage(event.data)
            
        };

        socketRef.current.onclose = () => {
            console.log('WebSocket connection closed');
        };

        return () => {
            // Clean up the WebSocket connection on component unmount
            socketRef.current?.close();
        };
    }, []);

    console.log(props.conversationWithParticipant)

    const sendMessage = () => {
        if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
            socketRef.current.send(currentMessage);
            console.log('Sent message:', currentMessage);
        }
    };
    const handleKeyUp = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === 'Enter') {
          sendMessage();
        }
      };

      console.log(chat.length)
    const Messages = () => {
        const messagesList = chat.map((i: any) => {

            const messageClass = props.conversationWithParticipant.participantID === i.senderID ? "left-message" : "right-message";
            return (
                <Box key={i.timestamp} className={`chat-message ${messageClass}`}>
                    <Typography variant="body1">{i.content}</Typography>
                </Box>
            )
        })

        return (
            <Stack display={"flex"} flexDirection={"column-reverse"} sx={{ overflowY: "scroll", height: "65vh", scrollBehavior: "smooth" }}>
                {messagesList}
            </Stack>
        )
    }
    return (
        <Stack>
            <Typography>{props.conversationWithParticipant.participantName}</Typography>
            <Messages />
            <Box display={"flex"} flexDirection={"row"} alignContent={"start"} mt={2}>
                <TextField onChange={(e)=>setCurrentMessage(e.target.value)} onKeyUp={handleKeyUp} placeholder='Type a new message...' fullWidth></TextField>
            </Box>
        </Stack>
    )

}