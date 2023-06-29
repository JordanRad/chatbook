import { AxiosInstance } from "axios"
import RestClient from "./RestClient"


export type RegisterDetails = {
    firstName: string,
    lastName: string,
    email: string,
    password: string,
    confirmedPassword: string
}



class AuthService {
    private client: AxiosInstance

    constructor() {
        this.client = RestClient
    }

    async registerUser(details: RegisterDetails) {
        console.log("Method")
        return await this.client.post('/users/register', details)
    }

    async loginUser(email: string, password: string) {
        return await this.client.post('http://localhost:8091/api/user-management/v1/auth/login', {
            email: email,
            password: password
        })
    }

}

export default new AuthService();