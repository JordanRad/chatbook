import { AxiosInstance } from "axios"
import RestClient from "./RestClient"

export type Friend = {
    id: string
    firstName: string
    lastName: string
    email: string
}

export type Profile = {
    id: string
    firstName: string
    lastName: string
    email: string
    friendsList: Friend[]
}

class UserService {
    private client: AxiosInstance

    constructor() {
        this.client = RestClient
    }

    async getProfile() {
        return await this.client.get('http://localhost:8091/api/user-management/v1/users/profile')
    }

    async getLastConversations() {
        return await this.client.get('http://localhost:8092/api/chat/v1/conversations')
    }
    async getChatHistory(conversationID :string) {
        const timestamp = new Date().toISOString().replace("."," ").split(" ")[0].replace("T"," ")
        return await this.client.get(`http://localhost:8092/api/chat/v1/conversations/${conversationID}/history?beforeTimestamp=${timestamp}`)
    }

}

export default new UserService();