import axios from 'axios';
const serverUrl = 'https://lipeiheng.fun'
const post = axios.create({
    baseURL: serverUrl,
    timeout: 5000,
    headers : {
        'Content-Type': 'application/json'
    },
    method : "post" ,
})
post.interceptors.response.use(
    (response) => {
        const code = response.data.code
        if ( code != 0 ) {
            return Promise.reject(response.data.msg)
        } 
        else {
            return Promise.resolve(response.data.data)
        }
    } ,
    (err) => {
        return Promise.reject(err)
    }
)
export {
    serverUrl , 
    post , 
}