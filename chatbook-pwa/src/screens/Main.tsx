import { Box, Button, List, ListItem, ListItemButton, ListItemText, Typography } from '@mui/material';
import useFetchProfile from '../hooks/useFetchProfile';
import useFetchLastConversations from '../hooks/useFetchLastConversations';

const Main = () => {
  const {profile} : any = useFetchProfile()

  const {conversations} : any = useFetchLastConversations()

  console.log(profile,conversations)
  if(profile === null){
    return <Typography>Loading...</Typography>
  }
  return (
    <Box display="flex" flexDirection={"column"} height="100vh">
        <Box display={"flex"} width={"100%"}>
            <Typography>{profile.firstName} {profile.lastName}</Typography>
        </Box>
        <Box height={"80%"} display={"flex"} flexDirection={"row"}>
      <Box width="30%" bgcolor="lightgray">
        <List>
          {conversations.map((conversation:any) => (
            <ListItemButton key={conversation.id}>
              <ListItemText primary={conversation.name} />
            </ListItemButton>
          ))}
        </List>
       
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
