import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

@Injectable({
    providedIn: 'root',
})
export class HttprequestService {
    constructor(private http: HttpClient, private router: Router) {}
    // get请求
    get(url, options = {}) {
        return new Promise((resolve, reject) => {
            this.http.get(url, options).subscribe((response) => {
                // handle success
                resolve(response);
            });
        });
    }

    // post请求
    post(url, data = {}, options = {}) {
        const httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
        };
        return new Promise((resolve, reject) => {
            this.http.post(url, data, options).subscribe((response) => {
                resolve(response);
            });
        });
    }

    // // 将token存入  localStorage
    // saveToken(token,user){
    //   localStorage.setItem('token',token)
    //   localStorage.setItem('user',user)
    // }

    // // 从localStorage 获取token
    // getToken(){
    //   return localStorage.getItem('token')
    // }
    // // 从localStorage 获取token
    // getUser(){
    //   return localStorage.getItem('user')
    // }

    // // 清空 token
    // removeToken(){
    //   localStorage.removeItem('token')
    //   localStorage.removeItem('user')
    // }

    // // 是否显示 用户信息 或者 登录注册
    //   isLogin(){
    //     if(this.getToken()!=null){
    //       // 登录
    //       return true
    //     }

    //       return false
    //   }

    // 刷新页面
    refreshPage() {
        this.router.navigate(['/home']); //动态路由
    }
}
