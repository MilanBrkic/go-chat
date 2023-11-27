# go-chat

Users Collection:

This collection stores information about the users who are registered on your chat app.
Each user document might contain fields such as:
\_id: A unique identifier for the user.
username: The username of the user.
email: The email address of the user.
password: The hashed password for user authentication.
Other user-related information (e.g., profile picture, display name).

Conversations Collection:

This collection represents the chat conversations between users, including one-on-one and group chats.
Each conversation document might contain fields such as:
\_id: A unique identifier for the conversation.
participants: An array of user IDs involved in the conversation.
type: A field indicating whether it's a one-on-one or group chat.
Other conversation-related metadata (e.g., chat title, creation timestamp).

Messages Collection:

This collection stores the chat messages exchanged between users in conversations.
Each message document might contain fields such as:
\_id: A unique identifier for the message.
conversation_id: The ID of the conversation to which the message belongs.
sender_id: The ID of the user who sent the message.
content: The text or content of the message.
timestamp: The timestamp indicating when the message was sent.
Other message-related metadata (e.g., message type, attachments).
