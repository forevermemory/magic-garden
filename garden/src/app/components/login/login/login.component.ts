import { Component, OnInit } from '@angular/core';
import { GlobalVariable } from '../../../../global/config';
import { HttprequestService } from '../../../services/http-request.service';

import { MenuItem } from 'primeng/api';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit {
    constructor(private http: HttprequestService) {}
    userInfo: any = {
        username: '',
        password: '',
    };
    public items__: MenuItem[];
    ngOnInit(): void {
        this.items__ = [
            { label: 'planA' },
            { label: 'planB' },
            { label: 'planC' },
        ];
    }
    // 登陆
    public async handleLogin() {
        let res = await this.http.get(
            GlobalVariable.base_path +
                'api/v1/articles/list?page=2&page_size=10'
        );
        if (res['code'] == 0) {
            console.log(res);
        } else {
            //
        }
    }
    handleOnItemClick(event) {
        console.log(event.item.label);
    }
}
