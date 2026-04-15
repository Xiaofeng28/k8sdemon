// 导入 axios 依赖
import axios from "axios";

// 利用 axios.create 方法创建一个 axios 实例：可以设置基础路径、超时的时间
const request = axios.create({
    // 请求的基础路径的设置
    baseURL: import.meta.env.VITE_APP_BASE_API,
    // 超时的时间的设置，超出五秒请求就是失败的
    timeout: 5000
});


// 请求拦截器
request.interceptors.request.use((config) => {
    // config:请求拦截器回调注入的对象(配置对象)
    // 配置对象的身上有一个重要的属性：headers 属性
    // 可以通过请求头携带公共参数：token
    return config;
})

// 响应拦截器
request.interceptors.response.use((response) => {
    // 响应拦截器成功的回调，一个会进行简化数据
    // 返回数据部分
    return response.data;
}, (error) => {
    // 根据 Http 状态码处理错误
    const status = error.response.status
    switch (status) {
        case 404:
            // 处理 404 错误
            break;
        case 500 | 501 | 502 |503 | 504 | 505:
            // 服务器错误
            break;
        case 401:
            // 处理 401 错误
            break;
    }
    return Promise.reject(new Error(error.message))
})

// 务必对外暴露我们配置好的 axios
export default request;