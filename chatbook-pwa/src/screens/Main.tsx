import { Box, Button, List, ListItem, ListItemButton, ListItemText, Typography } from '@mui/material';
import useFetchProfile from '../hooks/useFetchProfile';
import useFetchLastConversations from '../hooks/useFetchLastConversations';

type ConversationListItem = {
  participantName : string
  lastMessage: string
  ts :string
}
const Main = () => {
  const {profile} : any = useFetchProfile()

  const {conversations} : any = useFetchLastConversations()



  if(profile === null || conversations === null){
    return <Typography>Loading...</Typography>
  }


  const LastConversationsList = ()=>{

    console.log(profile.friendsList)
    const lastConversations = conversations.map((i:any)=>{
      const otherParticipant = profile.friendsList.find((j:any)=> {
        console.log("Other: ",j.id, i.otherParticipantID)
        return j.id === i.otherParticipantID
      } )
     
      const convLI: ConversationListItem = {
        lastMessage:i.lastMessageContent,
        ts: i.lastMessageDeliveredAt,
        participantName: `${otherParticipant.firstName} ${otherParticipant.lastName}`
      }
      return convLI
    })
    return (<List>
          {lastConversations.map((conversation:ConversationListItem,index:number) => (
            <ListItemButton key={`li-${index}`}>
              <ListItemText primary={conversation.participantName}  secondary={conversation.ts}/>
              <ListItemText primary={conversation.lastMessage} />
            </ListItemButton>
          ))}
        </List>
    )
  }
  return (
    <Box display="flex" flexDirection={"column"} height="100vh">
        <Box display={"flex"} width={"100%"}>
            <Typography>{profile.firstName} {profile.lastName}</Typography>
        </Box>
        <Box height={"80%"} display={"flex"} flexDirection={"row"}>
      <Box width="30%" bgcolor="lightgray">
        
       <LastConversationsList/>
      </Box>
      <Box width="70%" bgcolor="coral" textAlign="center">
        <Typography> Select a chat</Typography>
        <Button variant="contained" color="primary" fullWidth>
          Start New Conversation
        </Button>
      </Box>
    </Box>
    </Box>
  );
};

export default Main;
