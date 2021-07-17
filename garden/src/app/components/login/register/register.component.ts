import { Component, OnInit } from '@angular/core';
import { MessageService } from 'primeng/api'; // 消息弹窗
import { Router } from '@angular/router';
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
        public messageService: MessageService,
        public router: Router
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
            return false;
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
    public async handleRegister() {
        // 校验参数
        if (this.userinfo.nickname == '') {
            alert('请正确输入昵称');
            return false;
        }
        if (this.userinfo.password.length < 6) {
            alert('请正确输入密码,长度大于6位');
            return false;
        }
        if (!/^1[3456789]\d{9}$/.test(this.userinfo.username)) {
            alert('请正确输入手机号');
            return false;
        }
        if (this.userinfo.code == '') {
            alert('请正确输入验证码');
            return false;
        }
        let res = await this.http.post(
            GlobalVariable.base_path + '/api/v1/user/registe',
            { ...this.userinfo }
        );
        // 15051524096
        console.log(res);
        if (res['code == 0']) {
            // 暂存session信息
            // 跳转到login
            sessionStorage.setItem('username', this.userinfo.username);
            this.router.navigate(['login/login']);
        } else {
            alert(res['msg']);
        }
    }

    /**
     * handleAddCaptcha
     * 添加验证码到注册表单中
     */
    public handleAddCaptcha() {}

    // TODO
    // 我想监听 isNeedCaptcha (这个值会有多哥事件改变它的值),如果它一旦大于3,就去触发验证码
    public changedExtraHandler(a) {
        console.log(a);
    }
}
