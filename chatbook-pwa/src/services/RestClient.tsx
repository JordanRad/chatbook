import axios, { InternalAxiosRequestConfig } from "axios"

type ProtectedRoute = {
    method: string,
    path: string
}

const protectedRoutes: ProtectedRoute[] = [
    {
        method: "get",
        path: "/users/profile"
    },
    {
        method: "get",
        path: "/conversations"
    },
]

const isProtectedRoute = (method: string, url: string): boolean => {
    let result = false
    protectedRoutes.forEach(i => {
        if (i.method === method && url.includes(i.path)) result = true
    })
    return result
}


const onRequest = async (config: InternalAxiosRequestConfig) => {
    if (config.method && config.url) {
        if (isProtectedRoute(config.method, config.url)) {
            let user: any = JSON.parse(sessionStorage.getItem("token") || "")
            if (config.headers) {
                config.headers["Authorization"] = `Bearer ${user.token}`
                config.headers["Access-Control-Allow-Origin"] = "*"
            }
        }
    }
    return config
}

const instance = axios.create({
    baseURL: "",
})

instance.interceptors.request.use(
    (config) => onRequest(config),
    (error) => {
        return Promise.reject(error)
    }
)

export default instance