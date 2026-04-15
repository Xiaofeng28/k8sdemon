// 用户实体类
export interface User {
    userid: number;
    username: string;
    password: string;
}
// 用户信息数组
export type UserArray = User[]