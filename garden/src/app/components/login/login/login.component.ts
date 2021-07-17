import { Component, OnInit } from '@angular/core';
import { GlobalVariable } from '../../../../global/config';
import { HttprequestService } from '../../../services/http-request.service';
import { Router } from '@angular/router';

import { MenuItem } from 'primeng/api';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
    constructor(private http: HttprequestService, private router: Router) {}

    userinfo: any = {
        username: '',
        password: '',
        code: '',
        captcha: '',
    };
    // 封装个菜单组件 TODO
    public items__: MenuItem[];
    ngOnInit(): void {
        this.items__ = [
            { label: 'planA' },
            { label: 'planB' },
            { label: 'planC' },
        ];
        this.userinfo.username = sessionStorage.getItem('username');
        // console.log(this.userinfo);
    }
    // 登陆
    public async handleLogin() {
        let res = await this.http.post(
            GlobalVariable.base_path + '/api/v1/user/login',
            { ...this.userinfo }
        );
        if (res['code'] == 0) {
            // 跳转
            // home/my_home
            localStorage.setItem('loginUser', JSON.stringify(res['data']));
            this.router.navigate(['home/my_home']);
        } else {
            alert(res['msg']);
        }
    }
    handleOnItemClick(event) {
        console.log(event.item.label);
    }
}
