// 统一管理模块接口
import type { User, UserArray } from "@/type/user";
import request from "@/utils/request";

// 通过枚举管理模块的接口地址
enum API {
    // 定义数据接口地址
    USER_Info = '/user/info',
    USER_Add = '/user/add',
    USER_Delet = '/user/delete',
    USER_Update = '/user/update'
}

// 1.获取用户信息
export const reqUserInfo = () => request.get<any, UserArray>(API.USER_Info)

// 2.添加一个用户
export const reqAddOneUser = (username:string, password:string) => request.post<any, number>(API.USER_Add, {username,password})

// 3.删除一个用户
export const reqDeleteOneUser = (userid:number) => request.post<any, number>(API.USER_Delet, userid)

// 4.更新一个用户的信息
export const reqUpdateOneUserInfo = (user:User) => request.post<any, number>(API.USER_Update,user)