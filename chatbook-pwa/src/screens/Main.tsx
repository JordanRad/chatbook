import { Box, Button, List, ListItemButton, ListItemText, Typography } from '@mui/material';
import useFetchProfile from '../hooks/useFetchProfile';
import useFetchLastConversations from '../hooks/useFetchLastConversations';
import { useEffect, useRef, useState } from 'react';
import { Chat } from '../components/chat/Chat';

export type ConversationListItem = {
  id: string
  myID: string
  participantName: string
  participantID: string
  lastMessage: string
  ts: string
  online?: boolean
}
const Main = () => {
  const { profile }: any = useFetchProfile()

  const { conversations }: any = useFetchLastConversations()
  const [selectedConversation, setSelectedConversation] = useState<any>(null)
  const [activeFriendsIDs, setActiveFriendsIDs] = useState<any>([])

  const onListItemClickHandler = (conversationID: string) => {
    console.log(conversations.find((i: any) => i.ID === conversationID))
    const conversation = conversations.find((i: any) => i.ID === conversationID)
    if (conversation === null) {
      console.error("Cannot set null conversation value")
    }
    setSelectedConversation(conversation)
  }

  const updateActiveFriendsIDs = (ids: string[]) => {
    setActiveFriendsIDs(ids);
  }
  const socketRef = useRef<WebSocket | null>(null);
  useEffect(() => {
    if (profile === null) {
      return
    }
    // Establish the WebSocket connection
    socketRef.current = new WebSocket(`ws://localhost:6001/friends/realtime-status?id=${profile.id}`);

    // WebSocket event handlers
    socketRef.current.onopen = () => {
      console.log('WebSocket connection established');
    };

    socketRef.current.onmessage = (event) => {
      console.log('Received message:', event.data);
      const friendsIDs = JSON.parse(event.data)
      updateActiveFriendsIDs(friendsIDs)
    };

    socketRef.current.onclose = () => {
      console.log('WebSocket connection closed');
    };

    return () => {
      // Clean up the WebSocket connection on component unmount
      socketRef.current?.close();
    };
  }, [profile]);


  if (profile === null || conversations === null) {
    return <Typography>Loading...</Typography>
  }

  const lastConversationsWithParticipants = conversations.map((i: any) => {
    const otherParticipant = profile.friendsList.find((j: any) => j.id === i.otherParticipantID)
    
    const convLI: ConversationListItem = {
      id: i.ID,
      myID: profile.id,
      lastMessage: i.lastMessageContent,
      ts: i.lastMessageDeliveredAt,
      participantName: `${otherParticipant.firstName} ${otherParticipant.lastName}`,
      participantID: otherParticipant.id,
      online: Array.from(new Set(activeFriendsIDs.activeFriendsIDs.filter((id: string) => id !== ''))).includes(otherParticipant.id)
    }
    return convLI
  })


  const LastConversationsList = () => {
    return (<List>
      {lastConversationsWithParticipants.map((conversation: ConversationListItem) => (
        <ListItemButton onClick={(e) => onListItemClickHandler(conversation.id)} key={`li-${conversation.id}}`}>
          <Box sx={{ backgroundColor: conversation.online ? "green" : "gray" }} className="circle"></Box>
          <ListItemText primary={conversation.participantName} secondary={conversation.ts} />
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

          <LastConversationsList />
        </Box>
        <Box width="70%" bgcolor={selectedConversation !== null ? "lightblue" : "coral"} textAlign="center">
          {selectedConversation === null ?
            (
              <>
                <Typography> Select a chat</Typography>
                <Button variant="contained" color="primary" fullWidth>
                  Start New Conversation
                </Button></>)
            :
            (
              <Chat conversationWithParticipant={lastConversationsWithParticipants.find((i: ConversationListItem) => i.id === selectedConversation.ID)} />
            )}
        </Box>
      </Box>
    </Box>
  );
};

export default Main;
