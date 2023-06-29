import { Box, Stack, TextField, Typography } from "@mui/material"
import { ConversationListItem } from "../../screens/Main"
import useFetchChat from "../../hooks/useFetchChat"

export const Chat = (props: { conversationWithParticipant: ConversationListItem }) => {
    const { chat } = useFetchChat(props.conversationWithParticipant.id)

    console.log(props.conversationWithParticipant)

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
                <TextField placeholder='Type a new message...' fullWidth></TextField>
            </Box>
        </Stack>
    )

}