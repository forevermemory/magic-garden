import { Component, OnInit } from '@angular/core';
import { MessageService } from 'primeng/api'; // 消息弹窗

// 自定义
import { GlobalVariable } from '../../../../global/config';
import { HttprequestService } from '../../../services/http-request.service';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
    providers: [MessageService],
})
export class RegisterComponent implements OnInit {
    constructor(
        public http: HttprequestService,
        public messageService: MessageService
    ) {}
    // 是否触发验证码 超过三次触发
    isNeedCaptcha = 0;
    // 用户注册信息
    userinfo: any = {
        nickname: '',
        username: '',
        password: '',
        sex: '女',
        sexList: ['男', '女'],
        code: '',
        captcha: '',
    };
    ngOnInit(): void {}

    /**
     * handleGetSmsCode
     * 发送短信验证码
     */
    public async handleGetSmsCode() {
        if (!/^1[3456789]\d{9}$/.test(this.userinfo.username)) {
            this.isNeedCaptcha++;
            alert('请正确输入手机号');
        }
        // 发送短信验证码
        let res = await this.http.get(
            GlobalVariable.base_path + '/api/v1/user/sendsms',
            {
                params: { phone: this.userinfo.username },
            }
        );
        if (res['code'] == 0) {
            alert('短信验证码发送成功');
        }
    }

    /**
     * handleRegister
     * 注册逻辑
     */
    public handleRegister() {}

    // TODO
    // i want to monitor isNeedCaptcha ,when isNeedCaptcha >3 trigger captcha
}
